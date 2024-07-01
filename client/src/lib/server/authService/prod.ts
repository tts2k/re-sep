import { Metadata } from "@grpc/grpc-js";
import { SignJWT } from "jose";
import { env } from "$env/dynamic/private";
import { authClient } from "../grpc";
import type { AuthService, AuthResponse } from "./type";

const JWT_SECRET = env.JWT_SECRET;
if (!JWT_SECRET) {
	throw new Error("JWT Secret is not defined");
}

/**
 * Create a GRPC Metadata object with the correct authorization headers
 * Short lived token only for getting the data
 */
const createMetadata = async (id: string): Promise<Metadata> => {
	const metadata = new Metadata();

	const tokenPayload = {
		id: id,
	};

	const secret = new TextEncoder().encode(JWT_SECRET);

	// Generate and sign the token
	const oAuthToken = await new SignJWT(tokenPayload)
		.setProtectedHeader({ alg: "HS256" })
		.setIssuedAt()
		.setExpirationTime("1h")
		.sign(secret);

	metadata.set("x-authorization", `bearer ${oAuthToken}`);
	return metadata;
};

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
