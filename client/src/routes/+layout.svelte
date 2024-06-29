<script>
	import Header from "@/components/header/Header.svelte";
	import Sidebar from "@/components/sidebar/Sidebar.svelte";
	import "../app.pcss";
	import { ModeWatcher } from "mode-watcher";
	import { sidebarStatus } from "@/components/sidebar/stores/sidebarStore";
	import { fade } from "svelte/transition";
	import { Toaster } from "svelte-sonner";
</script>

<ModeWatcher />
<Toaster position="top-center" />
<!--Dim the background on sidebar open-->
{#if $sidebarStatus.open && !$sidebarStatus.pin}
	<div
		role="none"
		class="fixed w-full h-full bg-black opacity-55 left-0 top-0 z-20"
		transition:fade
	></div>
{/if}
<Sidebar />
<Header />
<slot />
