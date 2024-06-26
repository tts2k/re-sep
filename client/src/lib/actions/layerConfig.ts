import type { Action } from "svelte/action";
import type { UserConfig } from "@/stores/userConfig";
import { type Writable } from "svelte/store";

let currentConfig: UserConfig;
let rootElement: HTMLElement;

/*
 * Not needed anymore. Comment in case it can be used in the future.
 * */
//type DiffMap = {
//	[key: string]: {
//		oldValue: number | boolean | string;
//		newValue: number | boolean | string;
//	};
//};
// const diffObjects = (
// 	diffMap: DiffMap,
// 	oldObject: Record<string, any>,
// 	newObject: Record<string, any>,
// 	parentKey?: string,
// ) => {
// 	for (const [key, value] of Object.entries(oldObject)) {
// 		// Gonna break on array but I'd prefer to not think about that right now
// 		if (typeof value === "object") {
// 			diffObjects(diffMap, oldObject[key], newObject[key], key);
// 		}
//
// 		// Skip unsupported types
// 		if (!["number", "string", "boolean"].includes(typeof value)) {
// 			continue;
// 		}
//
// 		// Compare values
// 		if (value !== newObject[key]) {
// 			const diffMapKey = parentKey ? parentKey + "." + key : key;
// 			diffMap[diffMapKey] = {
// 				oldValue: value,
// 				newValue: newObject[key],
// 			};
// 		}
// 	}
// };
//
// const diffConfig = (userConfig: UserConfig) => {
// 	const diffMap: DiffMap = {};
//
// 	diffObjects(diffMap, currentConfig, userConfig);
//
// 	return diffMap;
// };

const replaceElementClass = (
	query: string,
	oldClass: string,
	newClass: string,
) => {
	const nodes = rootElement.querySelectorAll(query);

	for (const element of nodes.values()) {
		const ok = element.classList.replace(oldClass, newClass);
		if (!ok) {
			throw new Error(`Replacing class failed`);
		}
	}
};

const userConfigSubscribe = (value: UserConfig) => {
	if (!currentConfig) {
		currentConfig = value;
		return;
	}

	if (!value.layered) {
		return;
	}
};

// Config changes will be layered on exisiting config for better UX and server
// resources
// Normally they would be pre-applied from server side
export const layerConfig: Action<HTMLElement, Writable<UserConfig>> = (
	element: HTMLElement,
	userConfig: Writable<UserConfig>,
) => {
	rootElement = element;
	userConfig.subscribe(userConfigSubscribe);
};
