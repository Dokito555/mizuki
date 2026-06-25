import { api } from './client';
import type { AIAnalysis } from '$lib/types';

export async function getAIAnalysis(flowId: number) {
	const res = await api.get<AIAnalysis>(`/flows/${flowId}/ai`);
	return res.data;
}

export async function analyzeFlow(flowId: number) {
	const res = await api.post<{
		flow_id: number;
		model: string;
		is_fallback: boolean;
		analysis: AIAnalysis['analysis'];
	}>(`/flows/${flowId}/ai-analyze`);
	return res.data;
}
