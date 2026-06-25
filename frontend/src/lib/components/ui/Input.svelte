<script lang="ts">
	import { cn } from '$lib/utils';
	interface Props {
		type?: string;
		placeholder?: string;
		disabled?: boolean;
		value?: string;
		class?: string;
		[key: string]: any;
	}
	let { type = 'text', placeholder = '', disabled = false, value = $bindable(''), class: className = '', ...rest }: Props = $props();
</script>

<input
	{type}
	class={cn('flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500 focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50', className)}
	{placeholder}
	{disabled}
	bind:value
	{...rest}
/>
