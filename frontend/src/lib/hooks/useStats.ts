import { createQuery } from '@tanstack/svelte-query';
import { getStats } from '$lib/api/stats';

export function useStats() {
	return createQuery(() => ({
		queryKey: ['stats'],
		queryFn: () => getStats(),
		staleTime: 60_000,
	}));
}
