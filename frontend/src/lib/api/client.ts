import axios from 'axios';

export const api = axios.create({
	baseURL: '/api',
	headers: { 'Content-Type': 'application/json' },
});

api.interceptors.response.use(
	(res) => res,
	(err) => {
		const msg = err.response?.data?.error ?? err.message;
		return Promise.reject(new Error(msg));
	}
);
