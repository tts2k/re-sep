import type { AuthService } from "./type";
import type { UserConfig } from "@/proto/user_config";
import * as turso from "../turso";

const updateUserConfig = async (
	uc: UserConfig,
	userId: string,
): Promise<UserConfig> => {
	return await turso.updateUserConfig(uc, userId);
};

const service: AuthService = {
	updateUserConfig,
};

export default service;
