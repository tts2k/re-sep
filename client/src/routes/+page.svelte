<script lang="ts">
	import Reader from "@/components/reader/Reader.svelte";
	import type { PageServerData } from "./$types";

	export let data: PageServerData;
</script>

<svelte:head>
	<title>{data.article?.title}</title>
</svelte:head>

<Reader article={data.article} uc={data.userConfig} />
