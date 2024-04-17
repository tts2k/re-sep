import type { Action } from "svelte/action";
import type { UserConfig } from "@/stores/userConfig";
import { get, type Writable } from "svelte/store";

// Config changes will be layered on exisiting config for better UX and server
// resources
// Normally they would be pre-applied from server side
export const layerConfig: Action<HTMLElement, Writable<UserConfig>> = (
	element: HTMLElement,
	configStore: Writable<UserConfig>,
) => {
	const config = get(configStore);
	if (!config.layered) {
		return;
	}
};
