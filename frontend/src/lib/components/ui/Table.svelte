<script lang="ts">
	import { cn } from '$lib/utils';
	interface Props {
		class?: string;
		children?: import('svelte').Snippet;
		[key: string]: any;
	}
	let { class: className = '', children, ...rest }: Props = $props();
</script>

<div class="relative w-full overflow-auto">
	<table class={cn('w-full caption-bottom text-sm', className)} {...rest}>
		{#if children}
			{@render children()}
		{/if}
	</table>
</div>
