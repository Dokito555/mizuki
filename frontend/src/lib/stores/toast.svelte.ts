type ToastItem = { id: number; message: string; type: 'success' | 'error' | 'info' };

let toasts = $state<ToastItem[]>([]);
let idCounter = 0;

export function showToast(message: string, type: 'success' | 'error' | 'info' = 'info') {
	const id = ++idCounter;
	toasts = [...toasts, { id, message, type }];
	setTimeout(() => dismiss(id), 5000);
}

export function dismiss(id: number) {
	toasts = toasts.filter((t) => t.id !== id);
}

export function getToasts() {
	return toasts;
}
