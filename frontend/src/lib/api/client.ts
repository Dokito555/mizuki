import axios from 'axios';
import { showToast } from '$lib/stores/toast.svelte';

export const api = axios.create({
	baseURL: '/api',
	headers: { 'Content-Type': 'application/json' },
});

api.interceptors.response.use(
	(res) => res,
	(err) => {
		const msg = err.response?.data?.error ?? err.message;
		showToast(msg, 'error');
		return Promise.reject(new Error(msg));
	}
);
