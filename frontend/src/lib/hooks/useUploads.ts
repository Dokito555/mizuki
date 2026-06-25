import { createQuery, createMutation } from '@tanstack/svelte-query';
import { listUploads, getUpload, uploadFile, analyzeUpload, aiAnalyzeUpload } from '$lib/api/uploads';
import type { Upload } from '$lib/types';

export function useUploads(page: () => number) {
	return createQuery(() => ({
		queryKey: ['uploads', page()],
		queryFn: () => listUploads(page()),
	}));
}

export function useUpload(id: () => number) {
	return createQuery(() => ({
		queryKey: ['upload', id()],
		queryFn: () => getUpload(id()),
		refetchInterval: (q) => {
			const status = q.state.data?.status;
			if (status === 'queued' || status === 'parsing' || status === 'inserting') {
				return 2000;
			}
			return false;
		},
	}));
}

export function useUploadFile() {
	return createMutation(() => ({
		mutationFn: ({ file, force }: { file: File; force?: boolean }) => uploadFile(file, force ?? false),
	}));
}

export function useAnalyzeUpload() {
	return createMutation(() => ({
		mutationFn: (id: number) => analyzeUpload(id),
	}));
}

export function useAIAnalyzeUpload() {
	return createMutation(() => ({
		mutationFn: (id: number) => aiAnalyzeUpload(id),
	}));
}
