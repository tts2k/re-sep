import { browser } from "$app/environment";
import type { UserConfig } from "@/proto/user_config";
import { writable } from "svelte/store";

const defaultConfig: UserConfig & { layered: boolean } = {
	layered: false,
	font: "serif",
	fontSize: 3,
	justify: false,
	margin: {
		left: 3,
		right: 3,
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
