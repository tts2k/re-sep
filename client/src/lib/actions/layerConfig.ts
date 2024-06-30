import type { Action } from "svelte/action";
import type { UserConfig } from "@/stores/userConfig";
import type { Writable } from "svelte/store";
import { FontSizePresets, FontSizeTag as Tag } from "@/stylePresets";

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

const replaceClass = (query: string, oldClass: string, newClass: string) => {
	const nodes = rootElement.querySelectorAll(query);

	for (const element of nodes.values()) {
		const ok = element.classList.replace(oldClass, newClass);
		if (!ok) {
			element.classList.add(newClass);
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

	if (value.fontSize === currentConfig.fontSize) {
		return;
	}

	const oldFontSize = FontSizePresets[currentConfig.fontSize - 1];
	const newFontSize = FontSizePresets[value.fontSize - 1];

	for (let i = 0; i < oldFontSize.length; i++) {
		switch (i) {
			case Tag.H1:
				replaceClass("h1", oldFontSize[i], newFontSize[i]);
				break;
			case Tag.H2:
				replaceClass("h2", oldFontSize[i], newFontSize[i]);
				break;
			case Tag.H3:
				replaceClass("h3", oldFontSize[i], newFontSize[i]);
				break;
			case Tag.H4:
				replaceClass("h4", oldFontSize[i], newFontSize[i]);
				break;
			case Tag.TEXT:
				replaceClass("p, li, em", oldFontSize[i], newFontSize[i]);
				break;
		}
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
