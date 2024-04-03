import { writable } from "svelte/store";

const initStatus = {
	pin: false,
	open: false,
};

export const sidebarStatus = writable(initStatus);
