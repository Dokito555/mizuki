<script lang="ts">
	import { cn } from '$lib/utils';
	interface Props {
		value?: number;
		max?: number;
		class?: string;
		[key: string]: any;
	}
	let { value = 0, max = 100, class: className = '', ...rest }: Props = $props();
</script>

<div class={cn('relative h-2 w-full overflow-hidden rounded-full bg-secondary', className)} {...rest}>
	<div
		class="h-full bg-primary-600 transition-all duration-300 ease-out"
		style="width: {(value / max) * 100}%"
		role="progressbar"
		aria-valuenow={value}
		aria-valuemin={0}
		aria-valuemax={max}
	/>
</div>
