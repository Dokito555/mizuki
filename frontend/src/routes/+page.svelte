<script lang="ts">
	import { useUploads } from '$lib/hooks/useUploads';
	import { useStats } from '$lib/hooks/useStats';
	import Card from '$lib/components/ui/Card.svelte';
	import { BarChart3, Upload, GitBranch, AlertTriangle } from 'lucide-svelte';

	let currentPage = $state(1);
	const uploadsQuery = useUploads(() => currentPage);
	const statsQuery = useStats();
	$effect(() => {
		if (uploadsQuery.data?.meta && currentPage > uploadsQuery.data.meta.total_pages) {
			currentPage = 1;
		}
	});
</script>

<div class="space-y-6">
	<h1 class="text-2xl font-bold tracking-tight">Dashboard</h1>

	<div class="grid gap-4 md:grid-cols-4">
		<Card class="p-4">
			<div class="flex items-center gap-3">
				<Upload class="h-5 w-5 text-primary-600" />
				<div>
					<p class="text-sm text-muted-foreground">Uploads</p>
					<p class="text-2xl font-bold">{uploadsQuery.data?.meta.total ?? 0}</p>
				</div>
			</div>
		</Card>
		<Card class="p-4">
			<div class="flex items-center gap-3">
				<GitBranch class="h-5 w-5 text-primary-600" />
				<div>
					<p class="text-sm text-muted-foreground">Flows</p>
					<p class="text-2xl font-bold">{statsQuery.data?.total_flows ?? 0}</p>
				</div>
			</div>
		</Card>
		<Card class="p-4">
			<div class="flex items-center gap-3">
				<AlertTriangle class="h-5 w-5 text-red-500" />
				<div>
					<p class="text-sm text-muted-foreground">Threats</p>
					<p class="text-2xl font-bold">{statsQuery.data?.total_threats ?? 0}</p>
				</div>
			</div>
		</Card>
		<Card class="p-4">
			<div class="flex items-center gap-3">
				<BarChart3 class="h-5 w-5 text-primary-600" />
				<div>
					<p class="text-sm text-muted-foreground">Progress</p>
					<p class="text-2xl font-bold">-</p>
				</div>
			</div>
		</Card>
	</div>

	<div class="space-y-4">
		<h2 class="text-lg font-semibold">Recent Uploads</h2>
		{#if uploadsQuery.isPending}
			<p class="text-sm text-muted-foreground">Loading...</p>
		{:else if uploadsQuery.isError}
			<Card class="p-4 border-red-300 dark:border-red-700 bg-red-50 dark:bg-red-950/30">
				<p class="text-sm text-red-600 dark:text-red-400">Failed to load uploads: {uploadsQuery.error?.message}</p>
			</Card>
		{:else if uploadsQuery.data?.data?.length}
			<div class="grid gap-3 md:grid-cols-2 lg:grid-cols-3">
				{#each uploadsQuery.data.data as upload}
					<a href="/uploads/{upload.id}" class="block">
						<Card class="p-4 hover:border-primary-500 transition-colors">
							<p class="font-medium text-sm truncate">{upload.filename}</p>
							<p class="text-xs text-muted-foreground mt-1">{upload.status}</p>
							<p class="text-xs text-muted-foreground">{new Date(upload.created_at).toLocaleString()}</p>
						</Card>
					</a>
				{/each}
			</div>
		{:else}
			<p class="text-sm text-muted-foreground">No uploads yet. Upload a PCAP file to get started.</p>
		{/if}
	</div>
</div>
