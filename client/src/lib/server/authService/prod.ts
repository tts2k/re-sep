import { authClient, createMetadata } from "../grpc";
import type { AuthService, AuthResponse } from "./type";

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

const service: AuthService = {
	auth,
};

export default service;
