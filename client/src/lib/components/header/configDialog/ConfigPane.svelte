<script lang="ts">
	import { Label } from "@/components/ui/label";
	import * as Select from "@/components/ui/select";
	import { AvailableFonts, type Font } from "@/stores/userConfig";
	import { Slider } from "@/components/ui/slider";
	import { userConfig } from "@/stores/userConfig";
	import { previewConfig } from "../store/previewConfig";
	import { Button } from "@/components/ui/button";
	import { onMount } from "svelte";
	import type { Selected } from "bits-ui";

	let selectedFont: Selected<Font>;

	onMount(() => {
		// Staging preview config before applying to global config
		$previewConfig.font = $userConfig.font;
		$previewConfig.fontSize = $userConfig.fontSize;
	});

	$: selectedFont = {
		label: $previewConfig.font,
		value: $previewConfig.font,
	};

	const onSliderValueChange = (value: number[]) => {
		$previewConfig.fontSize = value[0];
	};

	const onFontSelectChange = (selected: Selected<Font> | undefined) => {
		if (!selected) {
			return;
		}

		$previewConfig.font = selected.value;
	};

	const saveConfig = () => {
		$userConfig.layered = true
		$userConfig.fontSize = $previewConfig.fontSize;
		$userConfig.font = $previewConfig.font;
	};
</script>

<div class="w-[30%] flex flex-col gap-4">
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

	<div class="border border-border p-8 rounded-md">
		<Label for="font-size" class="text-md font-bold">Font size</Label>
		<Slider
			id="font-size"
			class="mt-8"
			min={1}
			max={5}
			step={1}
			value={[$previewConfig.fontSize]}
			onValueChange={onSliderValueChange}
		/>
	</div>

	<div class="border border-border p-8 rounded-md">
		<Label for="" class="text-md font-bold">Margin</Label>

		<div class="flex flex-row mt-8">
			<Label for="margin-left" class="min-w-14">Left</Label>
			<Slider id="margin-left" min={1} max={5} step={1} />
		</div>

		<div class="flex flex-row mt-8">
			<Label for="margin-left" class="min-w-14">Right</Label>
			<Slider id="margin-left" min={1} max={5} step={1} />
		</div>
	</div>

	<div class="mt-8">
		<Button on:click={saveConfig}>Save</Button>
	</div>
</div>
