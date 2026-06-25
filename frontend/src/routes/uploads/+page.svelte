<script lang="ts">
	import { useUploads, useUploadFile } from '$lib/hooks/useUploads';
	import { browser } from '$app/environment';
	import { goto } from '$app/navigation';
	import { Button, Card, Input } from '$lib/components/ui';
	import { Upload, Search } from 'lucide-svelte';

	let page = $state(1);
	let pageSize = 20;
	const uploadsQuery = useUploads(() => page);

	let isDragging = $state(false);
	let dragCounter = $state(0);
	let fileInput = $state<HTMLInputElement | undefined>(undefined);

	const uploadMutation = useUploadFile();

	function handleDragOver(e: DragEvent) {
		e.preventDefault();
		e.stopPropagation();
	}
	function handleDragEnter(e: DragEvent) {
		e.preventDefault();
		dragCounter++;
		isDragging = true;
	}
	function handleDragLeave(e: DragEvent) {
		e.preventDefault();
		dragCounter--;
		if (dragCounter === 0) isDragging = false;
	}
	function handleDrop(e: DragEvent) {
		e.preventDefault();
		isDragging = false;
		dragCounter = 0;
		if (e.dataTransfer?.files?.[0]) {
			doUpload(e.dataTransfer.files[0]);
		}
	}

	async function doUpload(file: File) {
		if (file.size === 0) return;
		const result = await uploadMutation.mutateAsync({ file });
		goto(`/uploads/${result.id}`);
	}

	function handleFileSelect() {
		if (fileInput?.files?.[0]) {
			doUpload(fileInput.files[0]);
			fileInput.value = '';
		}
	}
</script>

<div class="space-y-6">
	<h1 class="text-2xl font-bold tracking-tight">Uploads</h1>

	<div
		class="border-2 border-dashed rounded-lg p-8 text-center transition-colors {isDragging ? 'border-primary-500 bg-primary-50 dark:bg-primary-950/30' : 'border-gray-300 dark:border-gray-700'}"
		role="button"
		tabindex="0"
		ondragover={handleDragOver}
		ondragenter={handleDragEnter}
		ondragleave={handleDragLeave}
		ondrop={handleDrop}
		onclick={() => fileInput?.click()}
		onkeydown={(e) => e.key === 'Enter' && fileInput?.click()}
	>
		<Upload class="h-8 w-8 mx-auto text-muted-foreground" />
		<p class="mt-2 text-sm font-medium">Drop a PCAP/PCAPNG file here or click to browse</p>
		<p class="text-xs text-muted-foreground mt-1">Max file size: 500 MB</p>
		<input type="file" accept=".pcap,.pcapng,application/vnd.tcpdump.pcap" class="hidden" bind:this={fileInput} onchange={handleFileSelect} />
	</div>

	<div class="space-y-3">
		<h2 class="text-lg font-semibold">All Uploads</h2>
		{#if uploadsQuery.isPending}
			<p class="text-sm text-muted-foreground">Loading...</p>
		{:else if uploadsQuery.data?.data?.length}
			<div class="grid gap-3 md:grid-cols-2 lg:grid-cols-3">
				{#each uploadsQuery.data.data as upload}
					<a href="/uploads/{upload.id}">
						<Card class="p-4 hover:border-primary-500 transition-colors cursor-pointer">
							<p class="font-medium text-sm truncate">{upload.filename}</p>
							<div class="flex items-center gap-2 mt-1">
								<span class="text-xs px-2 py-0.5 rounded-full bg-primary-100 dark:bg-primary-900 text-primary-700 dark:text-primary-300">
									{upload.status}
								</span>
								<span class="text-xs text-muted-foreground">{upload.packets_processed.toLocaleString()} packets</span>
							</div>
							<p class="text-xs text-muted-foreground mt-1">{new Date(upload.created_at).toLocaleString()}</p>
						</Card>
					</a>
				{/each}
			</div>
			<div class="flex items-center gap-2 mt-4">
				<Button disabled={page <= 1} onclick={() => page--} variant="outline" size="sm">Previous</Button>
				<span class="text-sm text-muted-foreground">Page {page} of {uploadsQuery.data?.meta.total_pages ?? 1}</span>
				<Button disabled={page >= (uploadsQuery.data?.meta.total_pages ?? 1)} onclick={() => page++} variant="outline" size="sm">Next</Button>
			</div>
		{:else}
			<p class="text-sm text-muted-foreground">No uploads yet.</p>
		{/if}
	</div>
</div>
