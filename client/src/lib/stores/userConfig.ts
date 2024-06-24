import { browser } from "$app/environment";
import type { Font } from "@/stylePresets";
import { writable } from "svelte/store";

export type UserConfig = {
	layered: boolean;
	font: Font;
	fontSize: number;
	margin: {
		left: number;
		right: number;
	};
};

const defaultConfig: UserConfig = {
	layered: false,
	font: "serif",
	fontSize: 3,
	margin: {
		left: 300,
		right: 300,
	},
};

let stored: UserConfig = defaultConfig;
export const userConfig = writable<UserConfig>(stored);

if (browser) {
	const localConfig = localStorage.getItem("userConfig");
	stored = localConfig ? JSON.parse(localConfig) : stored;

	userConfig.subscribe((value) => {
		localStorage.setItem("userConfig", JSON.stringify(value));
	});
}
