package pcap

import (
	"sync"
	"time"

	"github.com/Dokito555/mizuki/internal/entities"
)

type FlowBuilder struct {
	mu          sync.Mutex
	flows       map[string]*FlowStats
	mergeBidi   bool
	sampleLen   int
}

func NewFlowBuilder(mergeBidirectional bool, samplePacketLen int) *FlowBuilder {
	return &FlowBuilder{
		flows:     make(map[string]*FlowStats),
		mergeBidi: mergeBidirectional,
		sampleLen: samplePacketLen,
	}
}

type PacketInfo struct {
	SrcIP     string
	DstIP     string
	SrcPort   int
	DstPort   int
	Protocol  string
	Length    int64
	Timestamp time.Time
	Payload   []byte
	SrcMAC    string
	DstMAC    string
	TLSVersion string
	SNI       string
	DNSQuery  string
	AppProto  string
}

func (fb *FlowBuilder) AddPacket(p *PacketInfo) {
	fb.mu.Lock()
	defer fb.mu.Unlock()

	key := FlowKey{
		SrcIP:    p.SrcIP,
		DstIP:    p.DstIP,
		SrcPort:  p.SrcPort,
		DstPort:  p.DstPort,
		Protocol: p.Protocol,
	}

	if fb.mergeBidi {
		key = key.Canonical()
	}

	keyStr := key.String()
	stats, ok := fb.flows[keyStr]
	if !ok {
		stats = &FlowStats{
			Key:       key,
			SrcMAC:    p.SrcMAC,
			DstMAC:    p.DstMAC,
			sampleLen: fb.sampleLen,
		}
		fb.flows[keyStr] = stats
	}

	stats.addPacket(p.Timestamp, int(p.Length))

	if stats.TLSVersion == "" && p.TLSVersion != "" {
		stats.TLSVersion = p.TLSVersion
		stats.TLSSNI = p.SNI
	}
	if p.DNSQuery != "" {
		found := false
		for _, q := range stats.DNSQueries {
			if q == p.DNSQuery {
				found = true
				break
			}
		}
		if !found {
			stats.DNSQueries = append(stats.DNSQueries, p.DNSQuery)
		}
	}
	if stats.AppProtocol == "" && p.AppProto != "" {
		stats.AppProtocol = p.AppProto
	}
}

func (fb *FlowBuilder) Flush() ([]*FlowStats, error) {
	fb.mu.Lock()
	defer fb.mu.Unlock()

	result := make([]*FlowStats, 0, len(fb.flows))
	for _, stats := range fb.flows {
		result = append(result, stats)
	}

	fb.flows = make(map[string]*FlowStats)
	return result, nil
}

func (fb *FlowBuilder) Len() int {
	fb.mu.Lock()
	defer fb.mu.Unlock()
	return len(fb.flows)
}

func ToFlowEntity(stats *FlowStats, rawFileID uint, payloadSample []byte) *entities.Flow {
	iatAvg, iatMin, iatMax, iatStd := stats.calculateIAT()

	return &entities.Flow{
		SrcIP:      stats.Key.SrcIP,
		DstIP:      stats.Key.DstIP,
		SrcPort:    stats.Key.SrcPort,
		DstPort:    stats.Key.DstPort,
		Protocol:   stats.Key.Protocol,
		FirstSeen:  stats.FirstSeen,
		LastSeen:   stats.LastSeen,
		PacketCount: stats.PacketCount,
		ByteCount:  stats.ByteCount,
		SrcMAC:     stats.SrcMAC,
		DstMAC:     stats.DstMAC,
		TLSVersion: stats.TLSVersion,
		TLSSNI:     stats.TLSSNI,
		DNSQueries: stats.DNSQueries,
		AppProtocol: stats.AppProtocol,
		PayloadSample: payloadSample,
		IATAvgMs:    iatAvg,
		IATMinMs:    iatMin,
		IATMaxMs:    iatMax,
		IATStdDevMs: iatStd,
		RawFileID:   rawFileID,
	}
}

func ToPacketSampleEntities(flowID uint, stats *FlowStats) []entities.FlowPacketSample {
	n := len(stats.timestamps)
	if n == 0 {
		return nil
	}
	samples := make([]entities.FlowPacketSample, n)
	for i := 0; i < n; i++ {
		samples[i] = entities.FlowPacketSample{
			FlowID:    flowID,
			Timestamp: stats.timestamps[i],
			Size:      stats.packetSizes[i],
		}
	}
	return samples
}
