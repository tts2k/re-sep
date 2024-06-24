<script lang="ts">
	import { ScrollArea } from "@/components/ui/scroll-area";
	import PreviewContent from "./PreviewContent.svelte";
	import { Button } from "@/components/ui/button";
	import PlusIcon from "~icons/radix-icons/plus";
	import MinusIcon from "~icons/radix-icons/minus";
	import BorderAllIcon from "~icons/radix-icons/border-all";
	import TooltipWrapper from "./TooltipWrapper.svelte";

	let showBorder = false;

	const zoomLvls = [50, 75, 100];
	let curZoomLvl = 1;

	$: zoomlvl = zoomLvls[curZoomLvl];

	const zoomIn = () => {
		if (curZoomLvl + 1 < zoomLvls.length) {
			curZoomLvl++;
		}
	};

	const zoomOut = () => {
		if (curZoomLvl - 1 >= 0) {
			curZoomLvl--;
		}
	};

	const toggleBorder = () => {
		showBorder = !showBorder
	}
</script>

<div class="w-[70%] h-[85vh] font-serif">
	<ScrollArea
		orientation="both"
		class="border p-4 h-full flex-col
		rounded-md relative"
	>
		<PreviewContent scale={zoomlvl} showBorder={showBorder}/>

		<div class="absolute bottom-4 right-4 flex flex-col gap-4">
			<TooltipWrapper text="Zoom in">
				<Button variant="outline" size="icon" on:click={zoomIn}>
					<PlusIcon />
				</Button>
			</TooltipWrapper>

			<TooltipWrapper text="Zoom out">
				<Button variant="outline" size="icon" on:click={zoomOut}>
					<MinusIcon />
				</Button>
			</TooltipWrapper>

			<TooltipWrapper text="Toggle border">
				<Button variant="outline" size="icon" on:click={toggleBorder}>
					<BorderAllIcon />
				</Button>
			</TooltipWrapper>
		</div>
	</ScrollArea>
</div>
