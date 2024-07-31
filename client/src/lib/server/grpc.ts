import { Metadata, credentials } from "@grpc/grpc-js";
import { env } from "$env/dynamic/private";
import { ContentClient } from "@/proto/main";
import { building } from "$app/environment";
import { AuthClient } from "@/proto/main";
import { logger } from "./logger";
import { createJWTToken } from "./utils";

const AUTH_URL = env.AUTH_URL;
const CONTENT_URL = env.CONTENT_URL;
const NODE_ENV = env.NODE_ENV;
const JWT_SECRET = env.JWT_SECRET;

let contentClient: ContentClient;
let authClient: AuthClient;

if (!JWT_SECRET) {
	throw new Error("No JWT secret");
}

/**
 * Create a GRPC Metadata object with the correct authorization headers
 * Short lived token only for getting the data
 */
export const createMetadata = async (id: string): Promise<Metadata> => {
	const metadata = new Metadata();

	const oAuthToken = await createJWTToken(id, JWT_SECRET);

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
			logger.error(`GRPC content client connect error: ${error.message}`);
		}
	});

	authClient.waitForReady(deadline, (error?: Error) => {
		if (error) {
			logger.error(`GRPC auth client connect error: ${error.message}`);
		}
	});
};

if (!building) {
	initGRPC();
}

export { contentClient, authClient };
