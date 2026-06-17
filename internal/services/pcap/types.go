package pcap

import (
	"fmt"
	"math"
	"time"
)

type FlowKey struct {
	SrcIP    string
	DstIP    string
	SrcPort  int
	DstPort  int
	Protocol string
}

func (k FlowKey) String() string {
	return fmt.Sprintf("%s:%d-%s:%d-%s", k.SrcIP, k.SrcPort, k.DstIP, k.DstPort, k.Protocol)
}

func (k FlowKey) Reversed() FlowKey {
	return FlowKey{
		SrcIP:    k.DstIP,
		DstIP:    k.SrcIP,
		SrcPort:  k.DstPort,
		DstPort:  k.SrcPort,
		Protocol: k.Protocol,
	}
}

func (k FlowKey) Canonical() FlowKey {
	a := fmt.Sprintf("%s:%d", k.SrcIP, k.SrcPort)
	b := fmt.Sprintf("%s:%d", k.DstIP, k.DstPort)
	if a <= b {
		return k
	}
	return k.Reversed()
}

type FlowStats struct {
	Key        FlowKey
	FirstSeen  time.Time
	LastSeen   time.Time
	PacketCount int64
	ByteCount  int64
	SrcMAC     string
	DstMAC     string
	TLSVersion string
	TLSSNI     string
	DNSQueries []string
	AppProtocol string

	timestamps  []time.Time
	packetSizes []int
	sampleLen   int
}

func (fs *FlowStats) addPacket(ts time.Time, size int) {
	if fs.FirstSeen.IsZero() || ts.Before(fs.FirstSeen) {
		fs.FirstSeen = ts
	}
	if ts.After(fs.LastSeen) {
		fs.LastSeen = ts
	}
	fs.PacketCount++
	fs.ByteCount += int64(size)

	if len(fs.timestamps) < fs.sampleLen {
		fs.timestamps = append(fs.timestamps, ts)
		fs.packetSizes = append(fs.packetSizes, size)
	}
}

func (fs *FlowStats) calculateIAT() (avgMs, minMs, maxMs, stdDevMs float64) {
	if len(fs.timestamps) < 2 {
		return 0, 0, 0, 0
	}

	var sum, sumSq float64
	minMs = -1

	for i := 1; i < len(fs.timestamps); i++ {
		iat := fs.timestamps[i].Sub(fs.timestamps[i-1]).Seconds() * 1000
		sum += iat
		sumSq += iat * iat
		if minMs < 0 || iat < minMs {
			minMs = iat
		}
		if iat > maxMs {
			maxMs = iat
		}
	}

	n := float64(len(fs.timestamps) - 1)
	avgMs = sum / n
	variance := (sumSq / n) - (avgMs * avgMs)
	if variance > 0 {
		stdDevMs = math.Sqrt(variance)
	}

	return avgMs, minMs, maxMs, stdDevMs
}


