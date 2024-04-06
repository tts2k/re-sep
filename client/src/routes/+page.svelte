<script lang="ts">
	import type { PageServerData } from "./$types";
	import { toc } from "$lib/actions/toc";
	import { currentTocItem } from "$lib/components/sidebar/stores/tocStore";
	import { metadata } from "@/stores/articleMetadata";
	import { ScrollArea } from "$lib/components/ui/scroll-area";

	export let data: PageServerData;

	$metadata = {
		title: data.title,
		authors: data.author,
		toc: data.toc,
	};
</script>

<!-- class="mt-24 mb-24 ml-10 mr-10 font-serif lg:ml-72 lg:mr-72" -->

<ScrollArea orientation="vertical" class="rounded-md border p-4 h-screen">
	<article
		use:toc={{ store: currentTocItem }}
		class="mt-24 mb-24 ml-10 mr-10 font-serif lg:ml-72 lg:mr-72"
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
