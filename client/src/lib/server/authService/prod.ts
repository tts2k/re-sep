import { authClient, createMetadata } from "../grpc";
import type { AuthService } from "./type";
import type { UserConfig } from "@/proto/user_config";
import type { AuthResponse } from "@/proto/auth";

const auth = async (token: string): Promise<AuthResponse> => {
	const metadata = await createMetadata(token);

	return new Promise<AuthResponse>((resolve, reject) => {
		authClient.auth({}, metadata, (error, response) => {
			if (error !== null) {
				reject(error);
				return;
			}

			if (!response.user || !response.token) {
				reject(new Error("Error during auth"));
				return;
			}

			if (!response.user.sub) {
				reject(new Error("No user found"));
			}

			const authResponse: AuthResponse = {
				token: response.token,
				user: {
					sub: response.user.sub,
					name: response.user.name,
				},
			};

			resolve(authResponse);
		});
	});
};

const updateUserConfig = async (
	token: string,
	uc: UserConfig,
): Promise<UserConfig> => {
	const metadata = await createMetadata(token);

	return new Promise<UserConfig>((resolve, reject) => {
		authClient.updateUserConfig(uc, metadata, (error, response) => {
			if (error !== null) {
				reject(error);
				return;
			}

			if (!response || !response.margin) {
				reject(new Error("Error getting user config data"));
				return;
			}

			const uc: UserConfig = {
				fontSize: response.fontSize,
				justify: response.justify,
				font: response.font as UserConfig["font"],
				margin: response.margin,
			};

			resolve(uc);
		});
	});
};

const service: AuthService = {
	auth,
	updateUserConfig,
};

export default service;
