import { userConfig, type UserConfig } from "@/stores/userConfig";
import { authClient, createMetadata } from "../grpc";
import type { AuthService, AuthResponse } from "./type";
import type { UserConfig as PbUserConfig } from "@/proto/auth";
import { get } from "svelte/store";

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
					id: response.user.sub,
					name: response.user.name,
				},
			};

			resolve(authResponse);
		});
	});
};

const updateUserConfig = async (token: string): Promise<UserConfig> => {
	const metadata = await createMetadata(token);
	const uc = get(userConfig);

	const pbUc: PbUserConfig = {
		fontSize: uc.fontSize,
		font: uc.font,
		justify: uc.justify,
		margin: uc.margin,
	};

	return new Promise<UserConfig>((resolve, reject) => {
		authClient.updateUserConfig(pbUc, metadata, (error, response) => {
			if (error !== null) {
				reject(error);
				return;
			}

			if (!response || !response.margin) {
				reject(new Error("Error getting user config data"));
				return;
			}

			const uc: UserConfig = {
				layered: false,
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
