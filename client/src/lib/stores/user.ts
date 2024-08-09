import { writable } from "svelte/store";

type User = {
	loggedIn: boolean;
	name?: string;
};

const defaultConfig: User = {
	loggedIn: false,
};

export const user = writable<User>(defaultConfig);
