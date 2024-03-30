<script>
	import Header from "@/components/header/Header.svelte";
	import "../app.pcss";
	import { ModeWatcher } from "mode-watcher";
</script>

<ModeWatcher></ModeWatcher>
<Header></Header>
<slot></slot>
