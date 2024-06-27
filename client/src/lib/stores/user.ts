import { browser } from "$app/environment";
import { writable } from "svelte/store";

type User = {
	loggedIn: boolean;
	name?: string;
	token?: string;
};

const defaultConfig: User = {
	loggedIn: false,
};

let stored: User = defaultConfig;
const user = writable<User>(defaultConfig);

if (browser) {
	const localUser = localStorage.getItem("userConfig");
	stored = localUser ? JSON.parse(localUser) : stored;

	user.subscribe((value) => {
		localStorage.setItem("user", JSON.stringify(value));
	});
}
