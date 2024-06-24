import type { Font } from "$lib/stylePresets";
import { userConfig } from "@/stores/userConfig";
import { writable } from "svelte/store";

type PreviewConfig = {
	font: Font;
	fontSize: number;
};

const defaultPreviewConfig: PreviewConfig = {
	font: "serif",
	fontSize: 3,
};

export const previewConfig = writable<PreviewConfig>(defaultPreviewConfig);
