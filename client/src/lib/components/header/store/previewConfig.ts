import type { Font } from "$lib/stylePresets";
import { writable } from "svelte/store";

type PreviewConfig = {
	font: Font;
	fontSize: number;
	justify: boolean;
};

const defaultPreviewConfig: PreviewConfig = {
	font: "serif",
	fontSize: 3,
	justify: false,
};

export const previewConfig = writable<PreviewConfig>(defaultPreviewConfig);
