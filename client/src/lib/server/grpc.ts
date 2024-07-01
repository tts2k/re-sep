import { credentials } from "@grpc/grpc-js";
import { env } from "$env/dynamic/private";
import { ContentClient } from "@/proto/content";
import { building } from "$app/environment";
import { AuthClient } from "@/proto/auth";

const AUTH_URL = env.AUTH_URL;
const CONTENT_URL = env.CONTENT_URL;
const NODE_ENV = env.NODE_ENV;

let contentClient: ContentClient;
let authClient: AuthClient;

const initGRPC = () => {
	if (!AUTH_URL) {
		throw new Error("No auth URL");
	}

	if (!CONTENT_URL) {
		throw new Error("No content URL");
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
