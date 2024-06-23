import type { Action } from "svelte/action";

export const clickOutside: Action<HTMLElement, (e: Event) => void> = (
	element: HTMLElement,
	handler: (e: Event) => void,
) => {
	const handleClick = (e: Event) => {
		if (!element.contains(e.target as Node)) {
			handler(e);
		}
	};

	window.addEventListener("click", handleClick);

	return {
		destroy() {
			window.removeEventListener("click", handleClick);
		},
	};
};
