package pcap

import (
	"fmt"
	"net"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type ExtractedLayers struct {
	SrcIP     net.IP
	DstIP     net.IP
	SrcPort   int
	DstPort   int
	Protocol  string
	SrcMAC    string
	DstMAC    string
	TLSVersion string
	SNI       string
	DNSQuery  string
	AppProto  string
	Payload   []byte
}

func ExtractAll(p gopacket.Packet) (*ExtractedLayers, error) {
	el := &ExtractedLayers{}

	extractMACs(p, el)
	extractNetwork(p, el)
	extractTransport(p, el)
	extractTLS(p, el)
	extractDNS(p, el)
	detectAppProtocol(p, el, el.DstPort)

	app := p.ApplicationLayer()
	if app != nil {
		payload := app.Payload()
		if len(payload) > 64 {
			payload = payload[:64]
		}
		el.Payload = payload
	}

	if el.SrcIP == nil || el.DstIP == nil {
		return nil, fmt.Errorf("no network layer found")
	}

	return el, nil
}

func extractMACs(p gopacket.Packet, el *ExtractedLayers) {
	eth := p.Layer(layers.LayerTypeEthernet)
	if eth != nil {
		e, ok := eth.(*layers.Ethernet)
		if ok {
			el.SrcMAC = e.SrcMAC.String()
			el.DstMAC = e.DstMAC.String()
		}
	}
}

func extractNetwork(p gopacket.Packet, el *ExtractedLayers) {
	netLayer := p.NetworkLayer()
	if netLayer == nil {
		return
	}

	switch v := netLayer.(type) {
	case *layers.IPv4:
		el.SrcIP = v.SrcIP
		el.DstIP = v.DstIP
		el.Protocol = ipProtoToStr(int(v.Protocol))
	case *layers.IPv6:
		el.SrcIP = v.SrcIP
		el.DstIP = v.DstIP
		el.Protocol = ipProtoToStr(int(v.NextHeader))
	}
}

func extractTransport(p gopacket.Packet, el *ExtractedLayers) {
	tcp := p.Layer(layers.LayerTypeTCP)
	if tcp != nil {
		t, ok := tcp.(*layers.TCP)
		if ok {
			el.SrcPort = int(t.SrcPort)
			el.DstPort = int(t.DstPort)
			return
		}
	}

	udp := p.Layer(layers.LayerTypeUDP)
	if udp != nil {
		u, ok := udp.(*layers.UDP)
		if ok {
			el.SrcPort = int(u.SrcPort)
			el.DstPort = int(u.DstPort)
		}
	}
}

func extractTLS(p gopacket.Packet, el *ExtractedLayers) {
	app := p.ApplicationLayer()
	if app == nil {
		return
	}

	payload := app.Payload()
	if len(payload) < 1 || payload[0] != 0x16 {
		return
	}

	tlsVersion, sni := parseClientHello(payload)
	el.TLSVersion = tlsVersion
	el.SNI = sni
}

func parseClientHello(data []byte) (version, sni string) {
	if len(data) < 5 {
		return "", ""
	}

	if int(data[1])+2 > len(data) {
		return "", ""
	}

	switch data[2] {
	case 0x03:
		switch data[3] {
		case 0x01:
			version = "TLSv1.0"
		case 0x02:
			version = "TLSv1.1"
		case 0x03:
			version = "TLSv1.2"
		case 0x04:
			version = "TLSv1.3"
		default:
			version = "TLSv?" + fmt.Sprintf("%02x", data[3])
		}
	}

	sni = extractSNI(data)
	return version, sni
}

func extractSNI(data []byte) string {
	if len(data) < 43 {
		return ""
	}

	sessionIDLen := int(data[43])
	pos := 44 + sessionIDLen
	if pos+2 > len(data) {
		return ""
	}

	cipherLen := int(data[pos])<<8 | int(data[pos+1])
	pos += 2 + cipherLen
	if pos+2 > len(data) {
		return ""
	}

	compLen := int(data[pos]) + 1
	pos += compLen
	if pos+4 > len(data) {
		return ""
	}

	extLen := int(data[pos])<<8 | int(data[pos+1])
	pos += 2
	extEnd := pos + extLen
	if extEnd > len(data) {
		extEnd = len(data)
	}

	for pos+4 <= extEnd {
		extType := int(data[pos])<<8 | int(data[pos+1])
		extLen := int(data[pos+2])<<8 | int(data[pos+3])
		pos += 4

		if extType == 0x0000 && extLen > 5 && pos+extLen <= extEnd {
			sniListLen := int(data[pos+2])<<8 | int(data[pos+3])
			if sniListLen > 0 && pos+4+sniListLen <= extEnd {
				nameLen := int(data[pos+4])<<8 | int(data[pos+5])
				if nameLen > 0 && pos+6+nameLen <= extEnd {
					sniName := string(data[pos+6 : pos+6+nameLen])
					if strings.HasSuffix(sniName, ".") {
						sniName = sniName[:len(sniName)-1]
					}
					return sniName
				}
			}
		}
		pos += extLen
	}

	return ""
}

func extractDNS(p gopacket.Packet, el *ExtractedLayers) {
	dnsLayer := p.Layer(layers.LayerTypeDNS)
	if dnsLayer == nil {
		return
	}

	dns, ok := dnsLayer.(*layers.DNS)
	if !ok || len(dns.Questions) == 0 {
		return
	}

	for _, q := range dns.Questions {
		if len(q.Name) > 0 {
			el.DNSQuery = string(q.Name)
			return
		}
	}
}

func detectAppProtocol(p gopacket.Packet, el *ExtractedLayers, port int) {
	if port == 80 || port == 8080 {
		if isHTTP(p) {
			el.AppProto = "HTTP"
			return
		}
	}

	if port == 443 {
		if el.TLSVersion != "" {
			el.AppProto = "HTTPS"
			return
		}
	}

	if port == 25 || port == 587 {
		if isSMTP(p) {
			el.AppProto = "SMTP"
			return
		}
	}

	if port == 21 {
		el.AppProto = "FTP"
		return
	}

	if port == 22 {
		el.AppProto = "SSH"
		return
	}

	if port == 53 {
		el.AppProto = "DNS"
		return
	}
}

func isHTTP(p gopacket.Packet) bool {
	app := p.ApplicationLayer()
	if app == nil {
		return false
	}
	payload := string(app.Payload())
	return strings.HasPrefix(payload, "GET ") ||
		strings.HasPrefix(payload, "POST ") ||
		strings.HasPrefix(payload, "PUT ") ||
		strings.HasPrefix(payload, "DELETE ") ||
		strings.HasPrefix(payload, "HEAD ") ||
		strings.HasPrefix(payload, "HTTP/")
}

func isSMTP(p gopacket.Packet) bool {
	app := p.ApplicationLayer()
	if app == nil {
		return false
	}
	payload := string(app.Payload())
	return strings.HasPrefix(payload, "EHLO ") ||
		strings.HasPrefix(payload, "HELO ") ||
		strings.HasPrefix(payload, "MAIL ") ||
		strings.HasPrefix(payload, "RCPT ") ||
		strings.HasPrefix(payload, "DATA\r\n") ||
		strings.Contains(payload, "SMTP")
}

func ipProtoToStr(proto int) string {
	switch proto {
	case 6:
		return "TCP"
	case 17:
		return "UDP"
	case 1:
		return "ICMP"
	case 58:
		return "ICMPv6"
	default:
		return fmt.Sprintf("IP%d", proto)
	}
}
