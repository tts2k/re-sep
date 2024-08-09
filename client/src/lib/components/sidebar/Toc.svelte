<script lang="ts">
	import type { TOCItem } from "@/proto/content";

	export let items: TOCItem[];
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
		{#if item.subItems && item.subItems.length > 0}
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
