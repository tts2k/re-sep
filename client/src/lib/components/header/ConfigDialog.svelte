<script lang="ts">
	import * as Dialog from "$lib/components/ui/dialog";
	import { Slider } from "../ui/slider";

	export let open = false;
</script>

<Dialog.Root bind:open>
	<Dialog.Content>

		<Dialog.Header>
			<Dialog.Title>Font and Layout Settings</Dialog.Title>
		</Dialog.Header>

		<Slider/>

	</Dialog.Content>
</Dialog.Root>
