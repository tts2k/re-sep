<script lang="ts">
	import { Label } from "@/components/ui/label";
	import * as Select from "@/components/ui/select";
	import { type Font, AvailableFonts } from "@/stores/userConfig";
	import { Slider } from "@/components/ui/slider";
	import { userConfig } from "@/stores/userConfig";
	import type { Selected } from "bits-ui";
	import { Button } from "@/components/ui/button";

	type FormValues = {
		font: Selected<Font>;
		fontSize: number;
	};

	const formValues: FormValues = {
		font: {
			label: $userConfig.font,
			value: $userConfig.font,
		},
		fontSize: $userConfig.fontSize
	};

	const onSliderValueChange = (value: number[]) => {
		formValues.fontSize = value[0]
	}

	const saveConfig = () => {
		$userConfig.fontSize = formValues.fontSize
		$userConfig.font = formValues.font.value
	}
</script>

<div class="w-[30%] flex flex-col gap-4">
	<div class="border border-border p-8 rounded-md">
		<Label for="font" class="text-md font-bold">Font</Label>
		<Select.Root bind:selected={formValues.font}>
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
			value={[formValues.fontSize]}
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
		<Button on:click={saveConfig}> Save </Button>
	</div>
</div>
