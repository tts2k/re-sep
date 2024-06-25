import type { Font } from "$lib/stylePresets";
import { writable } from "svelte/store";

type PreviewConfig = {
	font: Font;
	fontSize: number;
	justify: boolean;
	margin: {
		left: number;
		right: number;
	};
};

const defaultPreviewConfig: PreviewConfig = {
	font: "serif",
	fontSize: 3,
	justify: false,
	margin: {
		left: 1,
		right: 1,
	},
};

export const previewConfig = writable<PreviewConfig>(defaultPreviewConfig);
