<script lang="ts">
	import type { TocItem } from "@/stores/articleMetadata";
	import { currentTocItem } from "./stores/tocStore";

	export let items: TocItem[];
	export let indent = 0;
	export let parentPrefix = "";

	const setCurrentTocItem = (id: string) => {
		$currentTocItem = id;
	};
</script>

<ul>
	{#each items as item, index}
		<li style="margin-left: {indent}px;">
			<a href="#{item.id}" on:click={() => setCurrentTocItem(item.id)}>
				{parentPrefix}{index + 1}. {item.label}
			</a>
		</li>
		{#if item.subItems.length > 0}
			<svelte:self
				items={item.subItems}
				parentPrefix={parentPrefix + (index + 1) + "."}
				indent={indent + 24}
			/>
		{/if}
	{/each}
</ul>

<style lang="postcss">
	a:global(.active) {
		@apply bg-foreground text-background;
	}
</style>
