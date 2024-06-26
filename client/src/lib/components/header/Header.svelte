<script lang="ts">
	import { Button } from "$lib/components/ui/button";
	import * as DropdownMenu from "$lib/components/ui/dropdown-menu";
	import SunIcon from "~icons/radix-icons/sun";
	import MoonIcon from "~icons/radix-icons/moon";
	import VDotsIcon from "~icons/radix-icons/dots-vertical";
	import TextLeftIcon from "~icons/radix-icons/text-align-left";
	import EnterFullScreenIcon from "~icons/radix-icons/enter-full-screen";
	import ExitFullScreenIcon from "~icons/radix-icons/exit-full-screen";
	import { mode, toggleMode } from "mode-watcher";
	import ConfigDialog from "./ConfigDialog.svelte";
	import { sidebarStatus } from "$lib/components/sidebar/stores/sidebarStore";
	import { metadata } from "@/stores/articleMetadata";
	import UserHoverCard from "./UserHoverCard.svelte";

	let fullscreenElement: Document["fullscreenElement"];
	let dropdownOpen: boolean;
	let userHoverOpen: boolean;
	let configDialogOpen: boolean;
	let showHeader = false;
	let hovering = false;

	const onMouseMove = () => {
		hovering = true;
		showHeader = true;
	};

	const onMouseLeave = () => {
		hovering = false;
		if (dropdownOpen || userHoverOpen) {
			return;
		}

		showHeader = false;
	};

	const onDropdownOpenChange = (open: boolean) => {
		if (!open && !hovering) {
			showHeader = false;
		}
	};

	const enterFullScreen = () => {
		document.documentElement.requestFullscreen();
	};

	const exitFullScreen = () => {
		document.exitFullscreen();
	};

	const openSidebar = (e: Event) => {
		e.stopPropagation();
		$sidebarStatus.open = true;
	};

	const onHoverCardOpenChange = (value: boolean) => {
		if (value) {
			return;
		}

		if (!hovering) {
			showHeader = false;
		}
	};

	$: toggleFullScreen = !fullscreenElement ? enterFullScreen : exitFullScreen;

	$: headerVisibility = showHeader ? "opacity-100" : "opacity-0";
</script>

<svelte:document bind:fullscreenElement />

<header
	tabindex="0"
	role="menubar"
	on:mousemove={onMouseMove}
	on:mouseleave={onMouseLeave}
	class="{headerVisibility} fixed w-full top-0 flex flex-grow h-12 items-center justify-center gap-4 border-b
	bg-background px-4 md:px-6 transition-opacity duration-500 shadow-md
	shadow-foreground/5 z-10"
>
	<div class="flex flex-1 justify-start gap-4">
		<Button variant="ghost" size="icon" on:click={openSidebar}>
			<TextLeftIcon font-size="24" />
		</Button>
		<UserHoverCard
			bind:open={userHoverOpen}
			onOpenChange={onHoverCardOpenChange}
		/>
	</div>

	<div class="flex flex-1 justify-center">
		<h1 class="font-bold">{$metadata.title}</h1>
	</div>

	<div class="flex flex-1 justify-end gap-4">
		<!-- Fullscreen button -->
		<Button variant="outline" size="icon" on:click={toggleFullScreen}>
			{#if !fullscreenElement}
				<EnterFullScreenIcon />
			{:else}
				<ExitFullScreenIcon />
			{/if}
		</Button>

		<!-- Settings dropdown -->
		<DropdownMenu.Root
			bind:open={dropdownOpen}
			onOpenChange={onDropdownOpenChange}
		>
			<DropdownMenu.Trigger asChild let:builder>
				<Button variant="outline" size="icon" builders={[builder]}>
					<VDotsIcon />
				</Button>
			</DropdownMenu.Trigger>
			<DropdownMenu.Content
				on:mouseover={onMouseMove}
				on:mouseleave={onMouseLeave}
				on:focus={onMouseMove}
			>
				<DropdownMenu.Item on:click={() => (configDialogOpen = true)}>
					Font & Layout Settings
				</DropdownMenu.Item>
				<DropdownMenu.Separator />
				<DropdownMenu.CheckboxItem>
					Autohide Cursor
				</DropdownMenu.CheckboxItem>
			</DropdownMenu.Content>
		</DropdownMenu.Root>

		<!-- Theme mode toggle -->
		<Button variant="outline" size="icon" on:click={toggleMode}>
			{#if $mode === "light"}
				<SunIcon />
			{:else}
				<MoonIcon />
			{/if}
		</Button>
	</div>
</header>

<ConfigDialog bind:open={configDialogOpen} />
