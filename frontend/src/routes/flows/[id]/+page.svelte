<script lang="ts">
	import { page } from '$app/stores';
	import { useFlow } from '$lib/hooks/useFlows';
	import { useAIAnalysis, useAnalyzeFlow } from '$lib/hooks/useAI';
	import { Button, Card } from '$lib/components/ui';
	import { showToast } from '$lib/stores/toast.svelte';
	import { Brain, Activity, Clock, GitBranch, Wifi, AlertTriangle, Tag } from 'lucide-svelte';

	const rawId = $derived(Number($page.params.id));
	const flowId = $derived(isNaN(rawId) ? null : rawId);
	const flowQuery = useFlow(() => flowId!, { enabled: flowId !== null });
	const aiQuery = useAIAnalysis(() => flowId!, { enabled: flowId !== null });
	const analyzeFlowMutation = useAnalyzeFlow();

	function scoreColor(score: number): string {
		if (score >= 70) return 'text-red-500 font-bold';
		if (score >= 40) return 'text-yellow-500 font-medium';
		if (score > 0) return 'text-blue-500';
		return 'text-muted-foreground';
	}

	async function runAI() {
		if (flowId === null) return;
		try {
			await analyzeFlowMutation.mutateAsync(flowId);
			aiQuery.refetch();
			showToast('AI analysis requested', 'success');
		} catch { /* error toast already shown by Axios interceptor */ }
	}
</script>

<div class="space-y-6">
	{#if flowQuery.isPending}
		<p class="text-muted-foreground">Loading...</p>
	{:else if flowQuery.isError}
		<Card class="p-4 border-red-300 dark:border-red-700 bg-red-50 dark:bg-red-950/30">
			<p class="text-sm text-red-600 dark:text-red-400">Failed to load flow: {flowQuery.error?.message}</p>
		</Card>
	{:else if flowQuery.data}
		{@const flow = flowQuery.data}
		<div class="flex items-center justify-between">
			<h1 class="text-2xl font-bold tracking-tight">Flow #{flow.id}</h1>
			<div class="flex items-center gap-2">
				<Button onclick={runAI} loading={analyzeFlowMutation.isPending}>
					<Brain class="h-4 w-4 mr-1" /> Analyze with AI
				</Button>
			</div>
		</div>

		<div class="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
			<Card class="p-4">
				<div class="flex items-center gap-2 text-muted-foreground text-xs font-medium uppercase tracking-wider mb-2">
					<Wifi class="h-4 w-4" /> Source
				</div>
				<p class="font-mono text-sm">{flow.src_ip}:{flow.src_port}</p>
				<p class="text-xs text-muted-foreground mt-1">{flow.src_mac ?? '-'}</p>
			</Card>
			<Card class="p-4">
				<div class="flex items-center gap-2 text-muted-foreground text-xs font-medium uppercase tracking-wider mb-2">
					<Wifi class="h-4 w-4" /> Destination
				</div>
				<p class="font-mono text-sm">{flow.dst_ip}:{flow.dst_port}</p>
				<p class="text-xs text-muted-foreground mt-1">{flow.dst_mac ?? '-'}</p>
			</Card>
			<Card class="p-4">
				<div class="flex items-center gap-2 text-muted-foreground text-xs font-medium uppercase tracking-wider mb-2">
					<Activity class="h-4 w-4" /> Protocol
				</div>
				<p class="font-mono text-sm">{flow.protocol}</p>
				<p class="text-xs text-muted-foreground mt-1">{flow.app_protocol ?? '-'}</p>
			</Card>
			<Card class="p-4">
				<div class="flex items-center gap-2 text-muted-foreground text-xs font-medium uppercase tracking-wider mb-2">
					<AlertTriangle class="h-4 w-4" /> Score
				</div>
				<p class="font-mono text-2xl font-bold {scoreColor(flow.score)}">{flow.score.toFixed(0)}</p>
			</Card>
		</div>

		<div class="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
			<Card class="p-4">
				<p class="text-xs text-muted-foreground">Packets</p>
				<p class="font-mono text-lg font-medium mt-1">{flow.packet_count.toLocaleString()}</p>
			</Card>
			<Card class="p-4">
				<p class="text-xs text-muted-foreground">Bytes</p>
				<p class="font-mono text-lg font-medium mt-1">{(flow.byte_count / 1024 / 1024).toFixed(2)} MB</p>
			</Card>
			<Card class="p-4">
				<p class="text-xs text-muted-foreground">IAT Avg</p>
				<p class="font-mono text-lg font-medium mt-1">{flow.iat_avg_ms.toFixed(1)} ms</p>
			</Card>
			<Card class="p-4">
				<p class="text-xs text-muted-foreground">IAT StdDev</p>
				<p class="font-mono text-lg font-medium mt-1">{flow.iat_std_dev_ms.toFixed(1)} ms</p>
			</Card>
		</div>

		{#if flow.threats?.length}
			<Card class="p-4 border-yellow-300 dark:border-yellow-700 bg-yellow-50 dark:bg-yellow-950/30">
				<div class="flex items-center gap-2 mb-2">
					<AlertTriangle class="h-5 w-5 text-yellow-500" />
					<span class="font-semibold">Detected Threats</span>
				</div>
				<ul class="space-y-1">
					{#each flow.threats as threat}
						<li class="flex items-center gap-2 text-sm">
							<Tag class="h-3 w-3 text-yellow-500" />
							{threat}
						</li>
					{/each}
				</ul>
			</Card>
		{/if}

		{#if flow.tls_sni || flow.tls_version}
			<Card class="p-4">
				<h3 class="font-semibold text-sm mb-2">TLS / DNS</h3>
				<div class="grid grid-cols-2 gap-4 text-sm">
					{#if flow.tls_version}
						<div>
							<span class="text-muted-foreground">Version:</span> {flow.tls_version}
						</div>
					{/if}
					{#if flow.tls_sni}
						<div>
							<span class="text-muted-foreground">SNI:</span> {flow.tls_sni}
						</div>
					{/if}
					{#if flow.dns_queries?.length}
						<div class="col-span-2">
							<span class="text-muted-foreground">DNS Queries:</span>
							<ul class="list-disc list-inside mt-1">
								{#each flow.dns_queries as q}
									<li class="font-mono text-xs">{q}</li>
								{/each}
							</ul>
						</div>
					{/if}
				</div>
			</Card>
		{/if}

		{#if aiQuery.data}
			{@const ai = aiQuery.data}
			{#if ai.status === 'analyzed' && ai.analysis}
				<Card class="p-4 border-primary-300 dark:border-primary-700">
					<div class="flex items-center justify-between mb-3">
						<div class="flex items-center gap-2">
							<Brain class="h-5 w-5 text-primary-600" />
							<h3 class="font-semibold">AI Analysis</h3>
						</div>
						<span class="text-xs text-muted-foreground">Model: {ai.model} | Confidence: {(ai.analysis.confidence * 100).toFixed(0)}%</span>
					</div>
					<p class="text-sm leading-relaxed">{ai.analysis.narrative}</p>
					{@const isFallback = ai.analysis.is_fallback}
					{#if isFallback}
						<p class="text-xs text-yellow-500 mt-2">Heuristic fallback (AI provider unavailable)</p>
					{/if}
					{#if ai.analysis.mitre_ids?.length}
						<div class="mt-3">
							<p class="text-xs font-medium text-muted-foreground mb-1">MITRE ATT&CK</p>
							<div class="flex flex-wrap gap-1">
								{#each ai.analysis.mitre_ids as id}
									<span class="text-xs px-2 py-0.5 rounded-full bg-primary-100 dark:bg-primary-900 text-primary-700 dark:text-primary-300 font-mono">{id}</span>
								{/each}
							</div>
						</div>
					{/if}
					{#if ai.analysis.attribution}
						<p class="text-xs text-muted-foreground mt-2">
							<span class="font-medium">Attribution:</span> {ai.analysis.attribution}
						</p>
					{/if}
					{#if ai.analysis.remediation?.length}
						<div class="mt-3">
							<p class="text-xs font-medium text-muted-foreground mb-1">Remediation</p>
							<ul class="space-y-1">
								{#each ai.analysis.remediation as r}
									<li class="text-sm flex items-start gap-2">
										<span class="text-green-500 mt-0.5">&#8226;</span>
										{r}
									</li>
								{/each}
							</ul>
						</div>
					{/if}
					{#if ai.analysis.correlations?.length}
						<div class="mt-3">
							<p class="text-xs font-medium text-muted-foreground mb-1">Correlations</p>
							{#each ai.analysis.correlations as corr}
								<div class="text-sm p-2 rounded bg-gray-50 dark:bg-gray-900 mb-1">
									<p class="font-medium">{corr.pattern}</p>
									<p class="text-xs text-muted-foreground">{corr.description}</p>
								</div>
							{/each}
						</div>
					{/if}
				</Card>
			{:else}
				<Card class="p-4">
					<div class="flex items-center gap-2 text-muted-foreground">
						<Brain class="h-5 w-5" />
						<p class="text-sm">No AI analysis yet. Click "Analyze with AI" to run.</p>
					</div>
				</Card>
			{/if}
		{:else if !aiQuery.isPending}
			<Card class="p-4">
				<div class="flex items-center gap-2 text-muted-foreground">
					<Brain class="h-5 w-5" />
					<p class="text-sm">No AI analysis yet. Click "Analyze with AI" to run.</p>
				</div>
			</Card>
		{/if}

		{#if flow.packet_samples?.length}
			<Card class="p-4">
				<h3 class="font-semibold text-sm mb-3">Packet Samples ({flow.packet_samples.length})</h3>
				<div class="overflow-x-auto">
					<table class="w-full text-xs">
						<thead>
							<tr class="border-b border-gray-200 dark:border-gray-800">
								<th class="text-left py-1 px-2 font-medium">#</th>
								<th class="text-left py-1 px-2 font-medium">Timestamp</th>
								<th class="text-right py-1 px-2 font-medium">Size</th>
							</tr>
						</thead>
						<tbody>
							{#each flow.packet_samples as sample, i}
								<tr class="border-b border-gray-100 dark:border-gray-900">
									<td class="py-1 px-2">{i + 1}</td>
									<td class="py-1 px-2 font-mono">{new Date(sample.ts).toISOString()}</td>
									<td class="py-1 px-2 text-right">{sample.size} B</td>
								</tr>
							{/each}
						</tbody>
					</table>
				</div>
			</Card>
		{/if}
	{:else}
		<p class="text-muted-foreground">Flow not found.</p>
	{/if}
</div>
