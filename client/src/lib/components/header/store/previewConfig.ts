import type { UserConfig } from "@/proto/user_config";
import { writable } from "svelte/store";

const defaultPreviewConfig: UserConfig = {
	font: "serif",
	fontSize: 3,
	justify: false,
	margin: {
		left: 3,
		right: 3,
	},
};

export const previewConfig = writable<UserConfig>(defaultPreviewConfig);
