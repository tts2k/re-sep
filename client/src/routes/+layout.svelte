<script>
	import Header from "@/components/header/Header.svelte";
	import "../app.css";
	import { ModeWatcher } from "mode-watcher";
</script>

<ModeWatcher />
<Header />
<slot />
