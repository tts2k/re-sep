import { env } from "$env/dynamic/public";
import { writable } from "svelte/store";

type User = {
	loggedIn: boolean;
	name?: string;
};

const defaultConfig: User = {
	loggedIn: false,
};

export const user = writable<User>(defaultConfig);

// #v-ifdef DEMO
// @ts-ignore
export const login = async (_: string) => {
	user.set({
		loggedIn: true,
		name: "demo",
	});
};

// @ts-ignore
export const logout = async () => {
	user.set({
		loggedIn: false,
	});
};

// #v-else
// @ts-ignore
// biome-ignore lint/suspicious/noRedeclare: conditional compile
export const login = async (provider: string) => {
	const response = await fetch(`${env.PUBLIC_AUTH_URL}/health`);
	if (response.status !== 200) {
		console.error(response);
		throw new Error("Error: Server is not running");
	}

	window.location.href = `${env.PUBLIC_AUTH_URL}/oauth/${provider}/login`;
};

// @ts-ignore
// biome-ignore lint/suspicious/noRedeclare: conditional compile
export const logout = async () => {};
// #v-endif