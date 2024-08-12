<script lang="ts">
	import { ScrollArea } from "$lib/components/ui/scroll-area";
	import { toc } from "$lib/actions/toc";
	import { layerConfig } from "$lib/actions/layerConfig";
	import { currentTocItem } from "$lib/components/sidebar/stores/tocStore";
	import { metadata } from "@/stores/articleMetadata";
	import { FontPreset, MarginPresets } from "@/stylePresets";
	import { page } from "$app/stores";
	import { toast } from "svelte-sonner";
	import { userConfig } from "@/stores/userConfig";
	import type { Article } from "@/proto/content";
	import type { UserConfig } from "@/proto/user_config";

	export let article: Article;
	export let uc: UserConfig;

	const error = $page.url.searchParams.get("error");

	$: if (error) {
		toast.error(error);
	}

	$metadata = {
		title: article.title,
		authors: article.authors,
		entryName: article.entryName,
		toc: article.toc,
	};

	$userConfig = {
		layered: false,
		...uc,
	};

	$: font = $userConfig.font ? FontPreset[$userConfig.font] : "font-serif";
	$: justified = $userConfig.justify ? "text-justify" : "";
	$: marginLeft = MarginPresets.left[($userConfig.margin?.left || 1) - 1];
	$: marginRight = MarginPresets.right[($userConfig.margin?.right || 1) - 1];
</script>

<ScrollArea
	orientation="vertical"
	class="border p-4 h-screen
	flex-col"
>
	<article
		use:toc={{ store: currentTocItem }}
		use:layerConfig={userConfig}
		class="mt-24 mb-24 h-screen flex flex-col {justified} {marginLeft}
		{marginRight} {font}"
	>
		{@html article.htmlText}
	</article>
</ScrollArea>

<!-- shadcn typography -->
<style lang="postcss">
	article :global(#pubinfo) {
		@apply mb-2;
	}

	article :global(h1) {
		@apply scroll-m-20 font-extrabold tracking-tight text-center mb-32;
	}

	article :global(h2) {
		@apply scroll-m-20 border-b pb-2 font-semibold tracking-tight
		transition-colors mt-10;
	}

	article :global(h3) {
		@apply scroll-m-20 font-semibold tracking-tight mt-10;
	}

	article :global(h4) {
		@apply scroll-m-20 font-semibold tracking-tight mt-10;
	}

	article :global(p) {
		@apply [&:not(:first-child)]:mt-6;
	}

	article :global(ul) {
		@apply my-6 ml-6 list-disc;
	}

	article :global(li) {
		@apply mt-2;
	}

	article :global(blockquote) {
		@apply mt-6 border-l-2 pl-6 italic;
	}
</style>
