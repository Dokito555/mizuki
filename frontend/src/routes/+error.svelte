<script lang="ts">
	import { Button, Card } from '$lib/components/ui';
	import { ShieldAlert } from 'lucide-svelte';

	let { status, error }: { status: number; error: Error } = $props();
</script>

<div class="flex items-center justify-center min-h-screen bg-white dark:bg-gray-950 p-4">
	<Card class="p-8 text-center max-w-md">
		<ShieldAlert class="h-12 w-12 mx-auto text-red-500" />
		<h1 class="text-2xl font-bold mt-4 text-gray-900 dark:text-gray-100">{status}</h1>
		<p class="text-muted-foreground mt-2">
			{error?.message ?? 'An unexpected error occurred.'}
		</p>
		<a href="/" class="mt-6 inline-block">
			<Button>Back to Dashboard</Button>
		</a>
	</Card>
</div>
