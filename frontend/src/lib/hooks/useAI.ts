import { createQuery, createMutation } from '@tanstack/svelte-query';
import { getAIAnalysis, analyzeFlow } from '$lib/api/ai';

export function useAIAnalysis(flowId: () => number) {
	return createQuery(() => ({
		queryKey: ['ai-analysis', flowId()],
		queryFn: () => getAIAnalysis(flowId()),
	}));
}

export function useAnalyzeFlow() {
	return createMutation(() => ({
		mutationFn: (flowId: number) => analyzeFlow(flowId),
	}));
}
