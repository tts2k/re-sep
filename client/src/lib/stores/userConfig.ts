import { browser } from "$app/environment";
import { defaultConfig, type UserConfigLayer } from "@/defaultConfig";
import { writable } from "svelte/store";

let stored: UserConfigLayer = defaultConfig;
export const userConfig = writable<UserConfigLayer>(stored);

if (browser) {
	const localConfig = localStorage.getItem("userConfig");
	stored = localConfig ? JSON.parse(localConfig) : stored;

	userConfig.subscribe((value) => {
		localStorage.setItem("userConfig", JSON.stringify(value));
	});
}
