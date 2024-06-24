<script lang="ts" context="module">
	export type ConfigDialogContext = {
		closeDialog: () => void
	}
</script>

<script lang="ts">
	import * as Dialog from "$lib/components/ui/dialog";
	import { setContext } from "svelte";
	import ConfigPane from "./configDialog/ConfigPane.svelte";
	import PreviewPane from "./configDialog/PreviewPane.svelte";

	export let open = false;

	const closeDialog = () => open = false

	setContext<ConfigDialogContext>("config-dialog", { closeDialog })
</script>

<Dialog.Root bind:open>
	<Dialog.Content class="sm:max-w-screen-2xl">
		<Dialog.Header>
			<Dialog.Title>Font and Layout Settings</Dialog.Title>
		</Dialog.Header>
		<div class="flex flex-row gap-4">
			<ConfigPane />
			<PreviewPane />
		</div>
	</Dialog.Content>
</Dialog.Root>
