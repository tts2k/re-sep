import { userConfig } from "@/stores/userConfig";
import type { AuthService } from "./type";
import { get } from "svelte/store";
import type { UserConfig } from "@/proto/user_config";
import type { AuthResponse } from "@/proto/auth";

const mockUser = {
	sub: "demo",
	name: "demo",
};

const mockAuthResponse = {
	token: "demo",
	user: mockUser,
};

const auth = async (): Promise<AuthResponse> => {
	return mockAuthResponse;
};

const updateUserConfig = async (): Promise<UserConfig> => {
	return get(userConfig);
};

const service: AuthService = {
	auth,
	updateUserConfig,
};

export default service;
