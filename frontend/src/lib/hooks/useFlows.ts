import { createQuery } from '@tanstack/svelte-query';
import { listFlows, getFlow, type FlowFilters } from '$lib/api/flows';

export function useFlows(filters: () => FlowFilters) {
	return createQuery(() => ({
		queryKey: ['flows', filters()],
		queryFn: () => listFlows(filters()),
	}));
}

export function useFlow(id: () => number) {
	return createQuery(() => ({
		queryKey: ['flow', id()],
		queryFn: () => getFlow(id()),
	}));
}
