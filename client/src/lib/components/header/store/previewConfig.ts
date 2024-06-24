import { userConfig, type Font } from "@/stores/userConfig";
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

userConfig.subscribe((value) => {
	previewConfig.set({
		fontSize: value.fontSize,
		font: value.font,
	});
});
