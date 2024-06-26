<script lang="ts">
	import { Label } from "@/components/ui/label";
	import * as Select from "@/components/ui/select";
	import { AvailableFonts, type Font } from "@/stylePresets";
	import { Slider } from "@/components/ui/slider";
	import { Button } from "@/components/ui/button";
	import { Checkbox } from "@/components/ui/checkbox";
	import { userConfig } from "@/stores/userConfig";
	import { previewConfig } from "../store/previewConfig";
	import { getContext, onMount } from "svelte";
	import type { Selected } from "bits-ui";
	import type { ConfigDialogContext } from "../ConfigDialog.svelte";

	const configDialog = getContext<ConfigDialogContext>("config-dialog");

	onMount(() => {
		// Staging preview config before applying to global config
		$previewConfig.font = $userConfig.font;
		$previewConfig.fontSize = $userConfig.fontSize;
	});

	// Mapping values for Select component(s)
	$: selectedFont = {
		label: $previewConfig.font,
		value: $previewConfig.font,
	};

	const onFontSizeChange = (value: number[]) => {
		$previewConfig.fontSize = value[0];
	};

	const onFontSelectChange = (selected: Selected<Font> | undefined) => {
		if (!selected) {
			return;
		}

		$previewConfig.font = selected.value;
	};

	const saveConfig = () => {
		$userConfig.layered = true;
		$userConfig.fontSize = $previewConfig.fontSize;
		$userConfig.font = $previewConfig.font;

		configDialog.closeDialog();
	};

	const createOnMarginChange = (direction: "left" | "right") => {
		return (value: number[]) => {
			$previewConfig.margin[direction] = value[0];
		};
	};
</script>

<div class="w-[20%] flex flex-col gap-4">
	<!-- Font family -->
	<div class="border border-border p-8 rounded-md">
		<Label for="font" class="text-md font-bold">Font</Label>
		<Select.Root
			selected={selectedFont}
			onSelectedChange={onFontSelectChange}
		>
			<Select.Trigger id="font" class="mt-8">
				<Select.Value />
			</Select.Trigger>
			<Select.Content>
				{#each AvailableFonts as font}
					<Select.Item value={font}>{font}</Select.Item>
				{/each}
			</Select.Content>
		</Select.Root>
	</div>

	<!-- Font size -->
	<div class="border border-border p-8 rounded-md">
		<Label for="font-size" class="text-md font-bold">Font Size</Label>
		<Slider
			id="font-size"
			class="mt-8"
			min={1}
			max={5}
			step={1}
			value={[$previewConfig.fontSize]}
			onValueChange={onFontSizeChange}
		/>
	</div>

	<!-- Text alignment -->
	<div class="border border-border p-8 rounded-md flex flex-col">
		<Label for="" class="text-md font-bold">Text Alignment</Label>
		<div class="flex flex-row items-center mt-8">
			<Checkbox
				class="mr-4"
				bind:checked={$previewConfig.justify}
				name="justify"
			/>
			<Label for="justify" class="text-md">Justified</Label>
		</div>
	</div>

	<!-- Margin -->
	<div class="border border-border p-8 rounded-md">
		<Label for="" class="text-md font-bold">Margin (percentage)</Label>

		<div class="flex flex-row mt-8">
			<Label for="margin-left" class="min-w-14">Left</Label>
			<Slider
				id="margin-left"
				min={1}
				max={5}
				step={1}
				value={[$previewConfig.margin.left]}
				onValueChange={createOnMarginChange("left")}
			/>
		</div>

		<div class="flex flex-row mt-8">
			<Label for="margin-left" class="min-w-14">Right</Label>
			<Slider
				id="margin-left"
				min={1}
				max={5}
				step={1}
				value={[$previewConfig.margin.right]}
				onValueChange={createOnMarginChange("right")}
			/>
		</div>
	</div>

	<div class="mt-8">
		<Button on:click={saveConfig}>Save</Button>
	</div>
</div>
