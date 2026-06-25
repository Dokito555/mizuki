<script lang="ts">
	import { cn } from '$lib/utils';
	interface Props {
		placeholder?: string;
		disabled?: boolean;
		value?: string;
		options?: { value: string; label: string }[];
		class?: string;
		[key: string]: any;
	}
	let { placeholder = '', disabled = false, value = $bindable(undefined), options = [], class: className = '', ...rest }: Props = $props();
</script>

<div class="relative">
	<select
		class={cn('flex h-10 w-full appearance-none rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-primary-500 focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50', className)}
		bind:value
		{disabled}
		{...rest}
	>
		{#if placeholder}
			<option value="" disabled selected>{placeholder}</option>
		{/if}
		{#each options as opt}
			<option value={opt.value}>{opt.label}</option>
		{/each}
	</select>
</div>
