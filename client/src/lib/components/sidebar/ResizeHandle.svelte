<script lang="ts">
	import { createEventDispatcher } from "svelte";

	export let resizing = false;

	const dispatch = createEventDispatcher();

	let startX = 0;

	const onMouseMove = (e: MouseEvent) => {
		if (!resizing) {
			return;
		}

		document.body.classList.add("cursor-col-resize");

		const deltaX = e.clientX - startX;
		dispatch("resize", deltaX);
	};

	const onMouseDown = (e: MouseEvent) => {
		resizing = true;

		startX = e.clientX;

		window.addEventListener("mouseup", onMouseUp);
		window.addEventListener("mousemove", onMouseMove);
	};

	const onMouseUp = () => {
		// Move the code block to the end of the event queue,
		// so that Sidebar's clickOutside can know and not run after resizing
		document.body.classList.remove("cursor-col-resize");
		setTimeout(() => {
			startX = 0;
			resizing = false;
			window.removeEventListener("mouseup", onMouseUp);
			window.removeEventListener("mousemove", onMouseMove);
		});
	};
</script>

<div
	tabindex="0"
	role="button"
	on:mousedown|stopPropagation|preventDefault={onMouseDown}
	class="absolute h-full w-2 top-0 right-0 cursor-col-resize"
></div>
