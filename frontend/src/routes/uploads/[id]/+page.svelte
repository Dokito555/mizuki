<script lang="ts">
	import { page } from '$app/stores';
	import { useUpload, useAnalyzeUpload, useAIAnalyzeUpload } from '$lib/hooks/useUploads';
	import { useFlows } from '$lib/hooks/useFlows';
	import { Button, Card, Progress } from '$lib/components/ui';
	import { Play, Brain, AlertTriangle, CheckCircle, Clock } from 'lucide-svelte';
	import { goto } from '$app/navigation';
	import { showToast } from '$lib/stores/toast.svelte';

	const rawId = $derived(Number($page.params.id));
	const uploadId = $derived(isNaN(rawId) ? null : rawId);
	const uploadQuery = useUpload(() => uploadId!, { enabled: uploadId !== null });
	const flowsQuery = useFlows(() => ({ upload_id: uploadId!, page_size: 50 }), { enabled: uploadId !== null });
	const analyzeMutation = useAnalyzeUpload();
	const aiAnalyzeMutation = useAIAnalyzeUpload();

	async function handleAnalyze() {
		if (uploadId === null) return;
		try {
			await analyzeMutation.mutateAsync(uploadId);
			showToast('Analysis started', 'success');
		} catch { /* error toast already shown by Axios interceptor */ }
	}

	async function handleAIAnalyze() {
		if (uploadId === null) return;
		try {
			await aiAnalyzeMutation.mutateAsync(uploadId);
			showToast('Batch AI analysis queued', 'success');
		} catch { /* error toast already shown by Axios interceptor */ }
	}

	const meta = $derived(uploadQuery.data);
	const progress = $derived(meta ? meta.progress_pct : 0);
	const status = $derived(meta ? meta.status : 'loading');
	const StatIcon = $derived(status === 'done' ? CheckCircle : status === 'error' ? AlertTriangle : Clock);
	const statColor = $derived(status === 'done' ? 'text-green-500' : status === 'error' ? 'text-red-500' : 'text-blue-500');
</script>

<div class="space-y-6">
	{#if uploadQuery.isPending}
		<p class="text-muted-foreground">Loading...</p>
	{:else if uploadQuery.isError}
		<Card class="p-4 border-red-300 dark:border-red-700 bg-red-50 dark:bg-red-950/30">
			<p class="text-sm text-red-600 dark:text-red-400">Failed to load upload: {uploadQuery.error?.message}</p>
		</Card>
	{:else if uploadQuery.data}
		<div class="flex items-center justify-between">
			<div>
				<h1 class="text-2xl font-bold tracking-tight">{meta.filename}</h1>
				<p class="text-sm text-muted-foreground mt-1">Uploaded {new Date(meta.created_at).toLocaleString()}</p>
			</div>
			<div class="flex items-center gap-2">
				<Button onclick={handleAnalyze} variant="outline" size="sm" loading={analyzeMutation.isPending}>
					<Play class="h-4 w-4 mr-1" /> Analyze
				</Button>
				<Button onclick={handleAIAnalyze} size="sm" loading={aiAnalyzeMutation.isPending}>
					<Brain class="h-4 w-4 mr-1" /> AI Analyze
				</Button>
			</div>
		</div>

		<div class="grid gap-4 md:grid-cols-4">
			<Card class="p-4">
				<p class="text-sm text-muted-foreground">Status</p>
				<div class="flex items-center gap-2 mt-1">
					<StatIcon class={"h-4 w-4 " + statColor} />
					<span class="font-medium capitalize">{status}</span>
				</div>
			</Card>
			<Card class="p-4">
				<p class="text-sm text-muted-foreground">Packets</p>
				<p class="text-xl font-bold mt-1">{meta.packets_processed.toLocaleString()}</p>
			</Card>
			<Card class="p-4">
				<p class="text-sm text-muted-foreground">Flows</p>
				<p class="text-xl font-bold mt-1">{meta.flow_count}</p>
			</Card>
			<Card class="p-4">
				<p class="text-sm text-muted-foreground">File Size</p>
				<p class="text-xl font-bold mt-1">{(meta.file_size / 1024 / 1024).toFixed(1)} MB</p>
			</Card>
		</div>

		{#if status === 'queued' || status === 'parsing' || status === 'inserting'}
			<Card class="p-4">
				<p class="text-sm font-medium mb-2">Processing...</p>
				<Progress value={progress} max={100} />
				<p class="text-xs text-muted-foreground mt-1">{meta.packets_processed.toLocaleString()} packets processed</p>
			</Card>
		{/if}

		{#if meta.error}
			<Card class="p-4 border-red-300 dark:border-red-700 bg-red-50 dark:bg-red-950/30">
				<p class="text-sm font-medium text-red-600 dark:text-red-400">Error: {meta.error}</p>
			</Card>
		{/if}

		<div class="space-y-3">
			<h2 class="text-lg font-semibold">Flows</h2>
			{#if flowsQuery.isPending}
				<p class="text-sm text-muted-foreground">Loading flows...</p>
			{:else if flowsQuery.isError}
				<Card class="p-4 border-red-300 dark:border-red-700 bg-red-50 dark:bg-red-950/30">
					<p class="text-sm text-red-600 dark:text-red-400">Failed to load flows: {flowsQuery.error?.message}</p>
				</Card>
			{:else if flowsQuery.data?.data?.length}
				<div class="overflow-x-auto">
					<table class="w-full text-sm">
						<thead>
							<tr class="border-b border-gray-200 dark:border-gray-800">
								<th class="text-left py-2 px-3 font-medium">Src IP:Port</th>
								<th class="text-left py-2 px-3 font-medium">Dst IP:Port</th>
								<th class="text-left py-2 px-3 font-medium">Proto</th>
								<th class="text-right py-2 px-3 font-medium">Score</th>
								<th class="text-left py-2 px-3 font-medium">Threats</th>
								<th class="text-right py-2 px-3 font-medium">Packets</th>
							</tr>
						</thead>
						<tbody>
							{#each flowsQuery.data.data as flow}
								<tr class="border-b border-gray-100 dark:border-gray-900 hover:bg-gray-50 dark:hover:bg-gray-900/50 cursor-pointer" onclick={() => goto(`/flows/${flow.id}`)}>
									<td class="py-2 px-3 font-mono text-xs">{flow.src_ip}:{flow.src_port}</td>
									<td class="py-2 px-3 font-mono text-xs">{flow.dst_ip}:{flow.dst_port}</td>
									<td class="py-2 px-3">{flow.protocol}</td>
									<td class="py-2 px-3 text-right font-mono">{flow.score.toFixed(0)}</td>
									<td class="py-2 px-3 max-w-[200px] truncate">{flow.threats?.join(', ') ?? '-'}</td>
									<td class="py-2 px-3 text-right">{flow.packet_count}</td>
								</tr>
							{/each}
						</tbody>
					</table>
				</div>
			{:else}
				<p class="text-sm text-muted-foreground">No flows extracted yet.</p>
			{/if}
		</div>
	{:else}
		<p class="text-muted-foreground">Upload not found.</p>
	{/if}
</div>

