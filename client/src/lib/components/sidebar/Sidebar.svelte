<script lang="ts">
	import { Button } from "$lib/components/ui/button";
	import Toc from "./Toc.svelte";
	import UnpinnedIcon from "~icons/radix-icons/drawing-pin";
	import PinnedIcon from "~icons/radix-icons/drawing-pin-filled";
	import { sidebarStatus } from "./stores/sidebarStore";
	import { clickOutside } from "$lib/actions/clickOutside";
	import { slide } from "svelte/transition";
	import { metadata } from "@/stores/articleMetadata";
	import { currentTocItem } from "./stores/tocStore";

	let tocRoot: HTMLElement;

	const onPinClick = () => {
		$sidebarStatus.pin = !$sidebarStatus.pin;
	};

	const closeSidebar = () => {
		if ($sidebarStatus.pin) {
			return;
		}
		$sidebarStatus.open = false;
	};

	$: pinIcon = $sidebarStatus.pin ? PinnedIcon : UnpinnedIcon;

	$: if (tocRoot && $currentTocItem !== "") {
		const active = tocRoot.querySelector(
			`a[href="${"#" + $currentTocItem}"]`,
		);
		active?.classList.add("active");
	}
</script>

{#if $sidebarStatus.open}
	<div
		use:clickOutside={closeSidebar}
		transition:slide={{ axis: "x" }}
		class="bg-border h-screen top-0 fixed z-20
		shadow-black shadow-2xl"
		style="width: 400px;"
	>
		<div
			class="space-y-4 py-4 overflow-hidden whitespace-nowrap text-ellipsis"
		>
			<div class="flex justify-end pr-3">
				<Button variant="ghost" size="icon" on:click={onPinClick}>
					<svelte:component
						this={pinIcon}
						font-size="16"
						class="-rotate-45"
					/>
				</Button>
			</div>
			<div id="toc" class="px-10 py-5" bind:this={tocRoot}>
				<Toc items={$metadata.toc} />
			</div>
		</div>
	</div>
{/if}