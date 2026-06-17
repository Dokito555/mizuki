package pcap

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcapgo"
)

type ProgressFunc func(packetsProcessed int64, pct int)

type ParseParams struct {
	MergeBidirectional bool
	SamplePayloadLen   int
	OnProgress         ProgressFunc
}

type ParseResult struct {
	FlowStats    []*FlowStats
	TotalPackets int64
	Duration     time.Duration
}

type Engine struct {
	mergeBidi bool
	sampleLen int
}

func NewEngine(mergeBidirectional bool, samplePacketLen int) *Engine {
	return &Engine{
		mergeBidi: mergeBidirectional,
		sampleLen: samplePacketLen,
	}
}

func (e *Engine) Parse(ctx context.Context, r io.Reader, params ParseParams) (*ParseResult, error) {
	start := time.Now()
	builder := NewFlowBuilder(params.MergeBidirectional, e.sampleLen)

	raw, ok := r.(io.ReadSeeker)
	if !ok {
		return nil, fmt.Errorf("reader must implement io.ReadSeeker for PCAP detection")
	}

	fileType := detectFileType(raw)
	if fileType == "" {
		return nil, fmt.Errorf("unknown or invalid pcap file format")
	}

	var packetSource *gopacket.PacketSource

	switch fileType {
	case "pcap":
		pcapReader, err := pcapgo.NewReader(raw)
		if err != nil {
			return nil, fmt.Errorf("open pcap: %w", err)
		}
		packetSource = gopacket.NewPacketSource(pcapReader, pcapReader.LinkType())
	case "pcapng":
		ngReader, err := pcapgo.NewNgReader(raw, pcapgo.NgReaderOptions{})
		if err != nil {
			return nil, fmt.Errorf("open pcapng: %w", err)
		}
		packetSource = gopacket.NewPacketSource(ngReader, ngReader.LinkType())
	default:
		return nil, fmt.Errorf("unsupported file type: %s", fileType)
	}

	packetSource.DecodeOptions = gopacket.DecodeOptions{
		Lazy:   true,
		NoCopy: true,
	}

	var totalPackets int64

	for packet := range packetSource.Packets() {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		if packet == nil {
			continue
		}

		totalPackets++

		el, err := ExtractAll(packet)
		if err != nil {
			continue
		}

		ts := packet.Metadata().Timestamp
		if ts.IsZero() {
			ts = time.Now()
		}

		builder.AddPacket(&PacketInfo{
			SrcIP:      el.SrcIP.String(),
			DstIP:      el.DstIP.String(),
			SrcPort:    el.SrcPort,
			DstPort:    el.DstPort,
			Protocol:   el.Protocol,
			Length:     int64(packet.Metadata().CaptureLength),
			Timestamp:  ts,
			Payload:    el.Payload,
			SrcMAC:     el.SrcMAC,
			DstMAC:     el.DstMAC,
			TLSVersion: el.TLSVersion,
			SNI:        el.SNI,
			DNSQuery:   el.DNSQuery,
			AppProto:   el.AppProto,
		})

		if totalPackets%1000 == 0 && params.OnProgress != nil {
			params.OnProgress(totalPackets, 0)
		}
	}

	flows, err := builder.Flush()
	if err != nil {
		return nil, fmt.Errorf("flush flows: %w", err)
	}

	return &ParseResult{
		FlowStats:    flows,
		TotalPackets: totalPackets,
		Duration:     time.Since(start),
	}, nil
}

func detectFileType(r io.ReadSeeker) string {
	magic := make([]byte, 8)
	_, err := io.ReadFull(r, magic)
	if err != nil {
		return ""
	}

	r.Seek(0, io.SeekStart)

	if len(magic) >= 4 {
		if (magic[0] == 0xd4 && magic[1] == 0xc3 && magic[2] == 0xb2 && magic[3] == 0xa1) ||
			(magic[0] == 0xa1 && magic[1] == 0xb2 && magic[2] == 0xc3 && magic[3] == 0xd4) {
			return "pcap"
		}
		if magic[0] == 0x0a && magic[1] == 0x0d && magic[2] == 0x0d && magic[3] == 0x0a {
			return "pcapng"
		}
	}

	return ""
}
