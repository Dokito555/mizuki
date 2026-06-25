<script lang="ts">
	import { cn } from '$lib/utils';
	import { getToasts, dismiss, showToast } from '$lib/stores/toast.svelte';
	import { X } from 'lucide-svelte';

	export { showToast };

	let toasts = $derived(getToasts());

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
</script>

<div class="fixed bottom-4 right-4 z-50 flex flex-col gap-2" style="pointer-events: none">
	{#each toasts as toast}
		<div class={toastClass(toast.type)} style="pointer-events: auto">
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
