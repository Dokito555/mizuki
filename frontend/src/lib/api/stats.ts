import { api } from './client';
import type { Stats } from '$lib/types';

export async function getStats() {
	const res = await api.get<{ data: Stats }>('/stats');
	return res.data.data;
}
