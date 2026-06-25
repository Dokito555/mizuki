<script lang="ts">
	import { cn } from '$lib/utils';
	import { X } from 'lucide-svelte';

	let toasts = $state<{ id: number; message: string; type: 'success' | 'error' | 'info' }[]>([]);
	let idCounter = 0;

	function showToast(message: string, type: 'success' | 'error' | 'info' = 'info') {
		const id = ++idCounter;
		toasts = [...toasts, { id, message, type }];
		setTimeout(() => dismiss(id), 5000);
	}

	function dismiss(id: number) {
		toasts = toasts.filter((t) => t.id !== id);
	}

	function toastClass(type: string) {
		return cn(
			'flex items-center gap-2 px-4 py-3 rounded-lg shadow-lg text-sm font-medium animate-slide-in',
			{
				'bg-green-600 text-white': type === 'success',
				'bg-red-600 text-white': type === 'error',
				'bg-blue-600 text-white': type === 'info',
			}
		);
	}

	export { showToast };
</script>

<div class="fixed bottom-4 right-4 z-50 flex flex-col gap-2 pointer-events-none">
	{#each toasts as toast}
		<div class={toastClass(toast.type)} pointer-events-auto>
			<span>{toast.message}</span>
			<button class="ml-2 hover:opacity-70" onclick={() => dismiss(toast.id)}>
				<X class="h-4 w-4" />
			</button>
		</div>
	{/each}
</div>

<style>
	@keyframes slide-in {
		from { transform: translateX(100%); opacity: 0; }
		to { transform: translateX(0); opacity: 1; }
	}
	.animate-slide-in { animation: slide-in 0.3s ease-out; }
</style>
