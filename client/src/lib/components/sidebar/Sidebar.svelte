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
	import ResizeHandle from "./ResizeHandle.svelte";

	let tocRoot: HTMLElement;
	let resizing: boolean;
	let originalWidth = 500;
	let addedWidth = 0;

	const onPinClick = () => {
		$sidebarStatus.pin = !$sidebarStatus.pin;
	};

	const closeSidebar = () => {
		if (resizing) {
			return;
		}

		if ($sidebarStatus.pin) {
			return;
		}
		$sidebarStatus.open = false;
	};

	const onResize = (e: CustomEvent) => {
		addedWidth = e.detail;
	};

	$: width = originalWidth + addedWidth;

	$: if (!resizing) {
		originalWidth = originalWidth + addedWidth;
		addedWidth = 0;
	}

	$: pinIcon = $sidebarStatus.pin ? PinnedIcon : UnpinnedIcon;

	$: if (tocRoot && $currentTocItem !== "") {
		// Reset active class
		let active = tocRoot.querySelector(".active");
		active?.classList.remove("active");

		// Add active class to new toc item
		active = tocRoot.querySelector(`a[href="${"#" + $currentTocItem}"]`);
		active?.classList.add("active");
	}
</script>

{#if $sidebarStatus.open}
	<div
		use:clickOutside={closeSidebar}
		transition:slide={{ axis: "x" }}
		class="bg-border h-screen top-0 fixed z-20
		shadow-black shadow-2xl"
		style="width: {width}px"
	>
		<div class="space-y-4 py-4 mr-4">
			<div class="flex justify-end pr-3">
				<Button variant="ghost" size="icon" on:click={onPinClick}>
					<svelte:component
						this={pinIcon}
						font-size="16"
						class="-rotate-45"
					/>
				</Button>
			</div>
			<div id="toc" class="mx-10 my-5" bind:this={tocRoot}>
				<Toc items={$metadata.toc} />
			</div>
		</div>
		<div></div>
		<ResizeHandle bind:resizing on:resize={onResize} />
	</div>
{/if}
