import { api } from './client';
import type { Upload, PaginatedResponse } from '$lib/types';

export async function listUploads(page = 1, pageSize = 20) {
	const res = await api.get<PaginatedResponse<Upload>>('/uploads', {
		params: { page, page_size: pageSize },
	});
	return res.data;
}

export async function getUpload(id: number) {
	const res = await api.get<{ data: Upload }>(`/uploads/${id}`);
	return res.data.data;
}

export async function uploadFile(file: File, force = false) {
	const form = new FormData();
	form.append('file', file);
	const res = await api.post<{ data: Upload }>(`/pcap/upload?force=${force}`, form, {
		headers: { 'Content-Type': 'multipart/form-data' },
	});
	return res.data.data;
}

export async function analyzeUpload(id: number) {
	await api.post(`/uploads/${id}/analyze`);
}

export async function aiAnalyzeUpload(id: number) {
	await api.post(`/uploads/${id}/ai-analyze`);
}
