import { env } from "process";
import { writable } from "svelte/store";

type User = {
	loggedIn: boolean;
	name?: string;
};

const defaultConfig: User = {
	loggedIn: false,
};

export const user = writable<User>(defaultConfig);

export const login = async (provider: string) => {
	const response = await fetch(`${env.PUBLIC_AUTH_URL}/health`);
	if (response.status !== 200) {
		console.error(response);
		throw new Error("Error: Server is not running");
	}

	window.location.href = `${env.PUBLIC_AUTH_URL}/oauth/${provider}/login`;
};
