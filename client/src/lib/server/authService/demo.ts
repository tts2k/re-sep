import { userConfig } from "@/stores/userConfig";
import type { AuthService } from "./type";
import { get } from "svelte/store";
import type { UserConfig } from "@/proto/user_config";

const updateUserConfig = async (): Promise<UserConfig> => {
	return get(userConfig);
};

const service: AuthService = {
	updateUserConfig,
};

export default service;
