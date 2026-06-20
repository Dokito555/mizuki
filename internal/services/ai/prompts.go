package ai

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Dokito555/mizuki/internal/entities"
)

const (
	MaxContextLen   = 3000
	MaxRelatedFlows = 20
)

type promptFunc func(flow *entities.Flow, related []entities.Flow) (string, string)

func buildFlowContext(flow *entities.Flow) map[string]interface{} {
	return map[string]interface{}{
		"src_ip":        flow.SrcIP,
		"dst_ip":        flow.DstIP,
		"src_port":      flow.SrcPort,
		"dst_port":      flow.DstPort,
		"protocol":      flow.Protocol,
		"packet_count":  flow.PacketCount,
		"byte_count":    flow.ByteCount,
		"iat_avg_ms":    flow.IATAvgMs,
		"iat_std_dev_ms": flow.IATStdDevMs,
		"tls_sni":       flow.TLSSNI,
		"dns_queries":   flow.DNSQueries,
		"app_protocol":  flow.AppProtocol,
		"score":         flow.Score,
		"threats":       flow.Threats,
	}
}

func buildRelatedContext(related []entities.Flow) []map[string]interface{} {
	if len(related) > MaxRelatedFlows {
		related = related[:MaxRelatedFlows]
	}
	ctx := make([]map[string]interface{}, len(related))
	for i, f := range related {
		ctx[i] = buildFlowContext(&f)
	}
	return ctx
}

func promptBorderline(flow *entities.Flow, _ []entities.Flow) (system, user string) {
	system = "You are a network security analyst. Given a network flow with heuristic scores, determine if it is truly malicious or a false positive. Respond with JSON only."

	fc := buildFlowContext(flow)
	data, _ := json.Marshal(fc)

	user = fmt.Sprintf(`Analyze this flow and classify it:
%s

Respond with JSON:
{"is_malicious": bool, "confidence": 0.0-1.0, "reasoning": "string"}`, string(data))

	return system, user
}

func promptNarrative(flow *entities.Flow, _ []entities.Flow) (system, user string) {
	system = "You are a threat intelligence analyst. Generate a detailed threat narrative with MITRE ATT&CK mapping and threat actor attribution. Respond with JSON only."

	fc := buildFlowContext(flow)
	data, _ := json.Marshal(fc)

	user = fmt.Sprintf(`Given this confirmed malicious flow, generate analysis:
%s

Respond with JSON:
{"narrative": "string", "mitre_ids": ["TXXXX.XXX"], "attribution": "string or null"}`, string(data))

	return system, user
}

func promptCorrelation(_ *entities.Flow, related []entities.Flow) (system, user string) {
	system = "You are a network forensics analyst. Analyze multiple flows from the same host for lateral movement, post-exploitation, or C2 patterns. Respond with JSON only."

	ctx := buildRelatedContext(related)
	data, _ := json.Marshal(ctx)

	user = fmt.Sprintf(`Analyze these related flows for cross-flow correlation patterns:
%s

Respond with JSON:
{"patterns": [{"description": "string", "flows": [0,1,...], "confidence": 0.0-1.0}]}`, string(data))

	return system, user
}

func promptRemediation(flow *entities.Flow, _ []entities.Flow) (system, user string) {
	system = "You are a security incident responder. Recommend concrete remediation actions. Respond with JSON only."

	fc := buildFlowContext(flow)
	data, _ := json.Marshal(fc)

	user = fmt.Sprintf(`Given this confirmed threat, recommend remediation actions:
%s

Respond with JSON:
{"remediation": ["action 1", "action 2"], "priority": "high|medium|low"}`, string(data))

	return system, user
}

var PromptTemplates = map[string]promptFunc{
	"borderline": promptBorderline,
	"narrative":  promptNarrative,
	"correlation": promptCorrelation,
	"remediation": promptRemediation,
}

func EstimateTokens(text string) int {
	return len(strings.Fields(text)) + len(text)/4
}
