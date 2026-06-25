<script lang="ts">
import '../app.css';
import { browser } from '$app/environment';
import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query';
import { Toast } from '$lib/components/ui';
import { Sidebar, Shield, ShieldHalf, Upload, GitBranch } from 'lucide-svelte';

	let { children }: { children: import('svelte').Snippet } = $props();

	let theme = $state<'light' | 'dark'>('light');

	if (browser) {
		theme = (localStorage.getItem('theme') as 'light' | 'dark') ?? 'light';
	}

	let applied = $state(false);
	$effect(() => {
		if (browser && !applied) {
			document.documentElement.classList.toggle('dark', theme === 'dark');
			applied = true;
		}
	});

	function toggleTheme() {
		const next = theme === 'light' ? 'dark' : 'light';
		theme = next;
		localStorage.setItem('theme', next);
		document.documentElement.classList.toggle('dark', next === 'dark');
	}

	const queryClient = new QueryClient({
		defaultOptions: {
			queries: { staleTime: 30_000, retry: 1 },
		},
	});

	let sidebarOpen = $state(true);
</script>

<QueryClientProvider client={queryClient}>
	<Toast />
	<div class="flex h-screen overflow-hidden bg-white dark:bg-gray-950 text-gray-900 dark:text-gray-100 transition-colors">
		<aside
			class="flex flex-col border-r border-gray-200 dark:border-gray-800 bg-gray-50 dark:bg-gray-900 transition-all duration-200 {sidebarOpen ? 'w-56' : 'w-0 overflow-hidden'}"
		>
			<div class="flex items-center gap-2 px-4 h-14 border-b border-gray-200 dark:border-gray-800 shrink-0">
				<ShieldHalf class="h-6 w-6 text-primary-600" />
				{#if sidebarOpen}
					<span class="font-semibold text-base">Mizuki</span>
				{/if}
			</div>
			<nav class="flex-1 p-2 space-y-1">
				<a href="/" class="flex items-center gap-3 px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-200 dark:hover:bg-gray-800 transition-colors">
					<Shield class="h-4 w-4" />
					{#if sidebarOpen}Dashboard{/if}
				</a>
				<a href="/uploads" class="flex items-center gap-3 px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-200 dark:hover:bg-gray-800 transition-colors">
					<Upload class="h-4 w-4" />
					{#if sidebarOpen}Uploads{/if}
				</a>
				<a href="/flows" class="flex items-center gap-3 px-3 py-2 rounded-md text-sm font-medium hover:bg-gray-200 dark:hover:bg-gray-800 transition-colors">
					<GitBranch class="h-4 w-4" />
					{#if sidebarOpen}Flows{/if}
				</a>
			</nav>
		</aside>

		<div class="flex flex-1 flex-col min-w-0">
			<header class="flex items-center justify-between h-14 px-4 border-b border-gray-200 dark:border-gray-800 bg-white dark:bg-gray-950 shrink-0">
				<button onclick={() => (sidebarOpen = !sidebarOpen)} class="p-1 rounded-md hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors">
					<Sidebar class="h-5 w-5" />
				</button>
				<div class="flex items-center gap-2">
					<button onclick={toggleTheme} class="p-2 rounded-md hover:bg-gray-100 dark:hover:bg-gray-800 transition-colors text-sm">
						{theme === 'light' ? '🌙' : '☀️'}
					</button>
				</div>
			</header>

			<main class="flex-1 overflow-auto p-6">
				{@render children()}
			</main>
		</div>
	</div>
</QueryClientProvider>
