<script lang="ts">
	import type { PageServerData } from "./$types";
	import { ScrollArea } from "$lib/components/ui/scroll-area";
	import { toc } from "$lib/actions/toc";
	import { layerConfig } from "$lib/actions/layerConfig";
	import { currentTocItem } from "$lib/components/sidebar/stores/tocStore";
	import { metadata } from "@/stores/articleMetadata";
	import { userConfig } from "@/stores/userConfig";

	export let data: PageServerData;

	$metadata = {
		title: data.title,
		authors: data.author,
		toc: data.toc,
	};
</script>

<svelte:head>
	<title>{data.title}</title>
</svelte:head>

<ScrollArea
	orientation="vertical"
	class="border p-4 h-screen
	flex-col"
>
	<article
		use:toc={{ store: currentTocItem }}
		use:layerConfig={userConfig}
		class="mt-24 mb-24 font-serif h-screen flex flex-col"
		style="margin-left: 300px; margin-right: 300px;"
	>
		{@html data.content}
	</article>
</ScrollArea>

<!-- shadcn typography -->
<style lang="postcss">
	article :global(#pubinfo) {
		@apply mb-2;
	}

	article :global(h1) {
		@apply scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-7xl
		text-center mb-32;
	}

	article :global(h2) {
		@apply scroll-m-20 border-b pb-2 text-4xl font-semibold tracking-tight
		transition-colors mt-10;
	}

	article :global(h3) {
		@apply scroll-m-20 text-3xl font-semibold tracking-tight mt-10;
	}

	article :global(h4) {
		@apply scroll-m-20 text-2xl font-semibold tracking-tight mt-10;
	}

	article :global(p) {
		@apply leading-7 [&:not(:first-child)]:mt-6 text-lg lg:text-xl;
	}

	article :global(em) {
		@apply text-lg lg:text-xl;
	}

	article :global(ul) {
		@apply my-6 ml-6 list-disc text-lg lg:text-xl;
	}

	article :global(li) {
		@apply mt-2;
	}

	article :global(blockquote) {
		@apply mt-6 border-l-2 pl-6 italic text-lg;
	}
</style>
