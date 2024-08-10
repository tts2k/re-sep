<script lang="ts">
	import { metadata } from "@/stores/articleMetadata";
	import GlobeIcon from "~icons/radix-icons/globe";

	$: authors = $metadata.authors.join(" - ");
	$: articleLink = `https://plato.stanford.edu/entries/${$metadata.entryName}/`; // Hardcoded for now
</script>

<div
	class="flex justify-between items-center px-10 py-4 ml-4 border border-accent rounded-xl"
>
	<div class="flex flex-col truncate">
		<span class="font-bold">{$metadata.title} </span>
		{authors}
	</div>

	<div>
		<a href={articleLink} target="_blank">
			<GlobeIcon font-size="20" />
		</a>
	</div>
</div>
