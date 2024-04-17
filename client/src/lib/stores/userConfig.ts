import { browser } from "$app/environment";
import { writable } from "svelte/store";

const availableFonts = ["serif", "sans-serif", "Open Dyslexic"] as const;

type Font = (typeof availableFonts)[number];

export type UserConfig = {
	layered: boolean;
	font: Font;
	margin: {
		left: number;
		right: number;
	};
	lineHeight: number;
};

const defaultConfig: UserConfig = {
	layered: false,
	font: "serif",
	margin: {
		left: 300,
		right: 300,
	},
	lineHeight: 1.5,
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
