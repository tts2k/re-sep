<script>
	import { ScrollArea } from "@/components/ui/scroll-area";
	import PreviewContent from "./PreviewContent.svelte";
	import { Button } from "@/components/ui/button";
	import PlusIcon from "~icons/radix-icons/plus";
	import MinusIcon from "~icons/radix-icons/minus";

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
</script>

<div class="w-[70%] h-[85vh] font-serif">
	<ScrollArea
		orientation="both"
		class="border p-4 h-full flex-col
		rounded-md relative"
	>
		<PreviewContent scale="scale-{zoomlvl}" />
		<div class="absolute bottom-4 right-4 flex flex-col gap-4">
			<Button variant="outline" size="icon" on:click={zoomIn}>
				<PlusIcon />
			</Button>
			<Button variant="outline" size="icon" on:click={zoomOut}>
				<MinusIcon />
			</Button>
		</div>
	</ScrollArea>
</div>
