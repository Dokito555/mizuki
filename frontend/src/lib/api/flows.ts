import { api } from './client';
import type { Flow, FlowDetail, PaginatedResponse } from '$lib/types';

export interface FlowFilters {
	src_ip?: string;
	dst_ip?: string;
	protocol?: string;
	min_score?: number;
	upload_id?: number;
	since?: string;
	until?: string;
	page?: number;
	page_size?: number;
	sort_by?: string;
	sort_desc?: boolean;
}

export async function listFlows(filters: FlowFilters = {}) {
	const params: Record<string, string | number | boolean | undefined> = {
		page: filters.page ?? 1,
		page_size: filters.page_size ?? 20,
	};
	if (filters.src_ip) params.src_ip = filters.src_ip;
	if (filters.dst_ip) params.dst_ip = filters.dst_ip;
	if (filters.protocol) params.protocol = filters.protocol;
	if (filters.min_score !== undefined) params.min_score = filters.min_score;
	if (filters.upload_id) params.upload_id = filters.upload_id;
	if (filters.since) params.since = filters.since;
	if (filters.until) params.until = filters.until;
	if (filters.sort_by) params.sort_by = filters.sort_by;
	if (filters.sort_desc !== undefined) params.sort_desc = filters.sort_desc;

	const res = await api.get<PaginatedResponse<Flow>>('/flows', { params });
	return res.data;
}

export async function getFlow(id: number) {
	const res = await api.get<{ data: FlowDetail }>(`/flows/${id}`);
	return res.data.data;
}
