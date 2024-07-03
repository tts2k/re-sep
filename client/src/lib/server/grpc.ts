import { Metadata, credentials } from "@grpc/grpc-js";
import { env } from "$env/dynamic/private";
import { ContentClient } from "@/proto/content";
import { building } from "$app/environment";
import { AuthClient } from "@/proto/auth";
import { SignJWT } from "jose";

const AUTH_URL = env.AUTH_URL;
const CONTENT_URL = env.CONTENT_URL;
const NODE_ENV = env.NODE_ENV;
const JWT_SECRET = env.JWT_SECRET;

let contentClient: ContentClient;
let authClient: AuthClient;

/**
 * Create a GRPC Metadata object with the correct authorization headers
 * Short lived token only for getting the data
 */
export const createMetadata = async (id: string): Promise<Metadata> => {
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

const initGRPC = () => {
	if (!AUTH_URL) {
		throw new Error("No auth URL");
	}

	if (!CONTENT_URL) {
		throw new Error("No content URL");
	}

	if (!JWT_SECRET) {
		throw new Error("No JWT secret");
	}

	contentClient = new ContentClient(
		CONTENT_URL || "",
		NODE_ENV === "production"
			? credentials.createSsl()
			: credentials.createInsecure(),
	);

	authClient = new AuthClient(
		AUTH_URL || "",
		NODE_ENV === "production"
			? credentials.createSsl()
			: credentials.createInsecure(),
	);

	const deadline = new Date();
	deadline.setSeconds(deadline.getSeconds() + 5);

	contentClient.waitForReady(deadline, (error?: Error) => {
		if (error) {
			console.log(`GRPC content client connect error: ${error.message}`);
		}
	});

	authClient.waitForReady(deadline, (error?: Error) => {
		if (error) {
			console.log(`GRPC auth client connect error: ${error.message}`);
		}
	});
};

if (!building) {
	initGRPC();
}

export { contentClient, authClient };
