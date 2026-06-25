import { createQuery } from '@tanstack/svelte-query';
import { listFlows, getFlow, type FlowFilters } from '$lib/api/flows';

export function useFlows(filters: () => FlowFilters, opts?: { enabled?: boolean }) {
	return createQuery(() => ({
		queryKey: ['flows', filters()],
		queryFn: () => listFlows(filters()),
		enabled: opts?.enabled,
	}));
}

export function useFlow(id: () => number, opts?: { enabled?: boolean }) {
	return createQuery(() => ({
		queryKey: ['flow', id()],
		queryFn: () => getFlow(id()),
		enabled: opts?.enabled,
	}));
}
