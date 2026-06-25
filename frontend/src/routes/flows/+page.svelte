<script lang="ts">
	import { useFlows } from '$lib/hooks/useFlows';
	import { Input, Select, Button } from '$lib/components/ui';
	import { goto } from '$app/navigation';
	import { Search, ArrowUpDown } from 'lucide-svelte';

	let filters = $state({
		src_ip: '',
		dst_ip: '',
		protocol: '',
		min_score: undefined as number | undefined,
		page: 1,
		page_size: 20,
		sort_by: 'score',
		sort_desc: true,
	});

	const flowsQuery = useFlows(() => filters);

	const protocolOptions = [
		{ value: '', label: 'All Protocols' },
		{ value: 'TCP', label: 'TCP' },
		{ value: 'UDP', label: 'UDP' },
		{ value: 'ICMP', label: 'ICMP' },
	];

	function applyFilter() {
		filters = { ...filters, page: 1 };
	}

	function resetFilters() {
		filters = { src_ip: '', dst_ip: '', protocol: '', min_score: undefined, page: 1, page_size: 20, sort_by: 'score', sort_desc: true };
	}

	function scoreColor(score: number): string {
		if (score >= 70) return 'text-red-500 font-bold';
		if (score >= 40) return 'text-yellow-500 font-medium';
		if (score > 0) return 'text-blue-500';
		return 'text-muted-foreground';
	}
</script>

<div class="space-y-6">
	<h1 class="text-2xl font-bold tracking-tight">Flows</h1>

	<div class="flex flex-wrap items-end gap-3 p-4 border rounded-lg bg-gray-50 dark:bg-gray-900/50">
		<div class="space-y-1">
			<label class="text-xs font-medium" for="src-ip">Source IP</label>
			<Input id="src-ip" placeholder="e.g. 10.0.0.1" bind:value={filters.src_ip} oninput={applyFilter} />
		</div>
		<div class="space-y-1">
			<label class="text-xs font-medium" for="dst-ip">Dest IP</label>
			<Input id="dst-ip" placeholder="e.g. 192.168.1.1" bind:value={filters.dst_ip} oninput={applyFilter} />
		</div>
		<div class="space-y-1">
			<label class="text-xs font-medium" for="protocol">Protocol</label>
			<Select options={protocolOptions} placeholder="All Protocols" bind:value={filters.protocol} onchange={applyFilter} />
		</div>
		<div class="space-y-1">
			<label class="text-xs font-medium" for="min-score">Min Score</label>
			<Input type="number" id="min-score" placeholder="0" bind:value={filters.min_score} oninput={applyFilter} />
		</div>
		<Button variant="outline" size="sm" onclick={resetFilters}>Reset</Button>
	</div>

	{#if flowsQuery.isPending}
		<p class="text-sm text-muted-foreground">Loading...</p>
	{:else if flowsQuery.data?.data?.length}
		<div class="overflow-x-auto border rounded-lg">
			<table class="w-full text-sm">
				<thead class="bg-gray-50 dark:bg-gray-900">
					<tr>
						<th class="text-left py-3 px-3 font-medium cursor-pointer hover:text-primary-600" onclick={() => { filters = { ...filters, sort_by: 'src_ip' }; }}>
							<div class="flex items-center gap-1">Src IP:Port <ArrowUpDown class="h-3 w-3" /></div>
						</th>
						<th class="text-left py-3 px-3 font-medium">Dst IP:Port</th>
						<th class="text-left py-3 px-3 font-medium">Proto</th>
						<th class="text-right py-3 px-3 font-medium cursor-pointer hover:text-primary-600" onclick={() => { filters = { ...filters, sort_by: 'score' }; }}>
							<div class="flex items-center justify-end gap-1">Score <ArrowUpDown class="h-3 w-3" /></div>
						</th>
						<th class="text-left py-3 px-3 font-medium">Threats</th>
						<th class="text-right py-3 px-3 font-medium">Packets</th>
						<th class="text-right py-3 px-3 font-medium">Bytes</th>
					</tr>
				</thead>
				<tbody>
					{#each flowsQuery.data.data as flow}
						<tr class="border-t border-gray-100 dark:border-gray-900 hover:bg-gray-50 dark:hover:bg-gray-900/50 cursor-pointer" onclick={() => goto(`/flows/${flow.id}`)}>
							<td class="py-2.5 px-3 font-mono text-xs">{flow.src_ip}:{flow.src_port}</td>
							<td class="py-2.5 px-3 font-mono text-xs">{flow.dst_ip}:{flow.dst_port}</td>
							<td class="py-2.5 px-3">{flow.protocol}</td>
							<td class="py-2.5 px-3 text-right font-mono {scoreColor(flow.score)}">{flow.score.toFixed(0)}</td>
							<td class="py-2.5 px-3 max-w-[200px] truncate text-xs">{flow.threats?.join(', ') ?? '-'}</td>
							<td class="py-2.5 px-3 text-right">{flow.packet_count.toLocaleString()}</td>
							<td class="py-2.5 px-3 text-right">{(flow.byte_count / 1024).toFixed(0)} KB</td>
						</tr>
					{/each}
				</tbody>
			</table>
		</div>

		<div class="flex items-center justify-between mt-4">
			<span class="text-sm text-muted-foreground">Total: {flowsQuery.data.meta.total} flows</span>
			<div class="flex items-center gap-2">
				<Button disabled={filters.page <= 1} onclick={() => { filters = { ...filters, page: filters.page - 1 }}} variant="outline" size="sm">Previous</Button>
				<span class="text-sm text-muted-foreground">Page {flowsQuery.data.meta.page} of {flowsQuery.data.meta.total_pages}</span>
				<Button disabled={filters.page >= flowsQuery.data.meta.total_pages} onclick={() => { filters = { ...filters, page: filters.page + 1 }}} variant="outline" size="sm">Next</Button>
			</div>
		</div>
	{:else}
		<div class="text-center py-12">
			<p class="text-muted-foreground">No flows found matching your filters.</p>
		</div>
	{/if}
</div>
