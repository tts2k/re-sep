import { browser } from "$app/environment";
import type { UserConfig } from "@/proto/user_config";
import { writable } from "svelte/store";

export type UserConfigLayer = UserConfig & { layered?: boolean };

const defaultConfig: UserConfigLayer = {
	layered: false,
	font: "serif",
	fontSize: 3,
	justify: false,
	margin: {
		left: 3,
		right: 3,
	},
};

let stored: UserConfigLayer = defaultConfig;
export const userConfig = writable<UserConfigLayer>(stored);

if (browser) {
	const localConfig = localStorage.getItem("userConfig");
	stored = localConfig ? JSON.parse(localConfig) : stored;

	userConfig.subscribe((value) => {
		localStorage.setItem("userConfig", JSON.stringify(value));
	});
}
