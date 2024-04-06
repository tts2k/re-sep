<script lang="ts">
	import type { TocItem } from "@/stores/articleMetadata";

	export let items: TocItem[];
	export let indent = 0;
	export let parentPrefix = "";
</script>

<ul>
	{#each items as item, index}
		<li style="padding-left: {indent}px;" class="truncate">
			<a href="#{item.id}">
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
