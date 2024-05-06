import { credentials } from "@grpc/grpc-js";
import { env } from "$env/dynamic/private";
import { ContentClient } from "@/proto/content";

const GRPC_URL = env.GRPC_URL;

if (!GRPC_URL) {
	throw new Error("No gRPC URL");
}

const contentClient = new ContentClient(GRPC_URL, credentials.createInsecure());

const deadline = new Date();
deadline.setSeconds(deadline.getSeconds() + 5);

contentClient.waitForReady(deadline, (error?: Error) => {
	if (error) {
		console.log(`GRPC content client connect error: ${error.message}`);
	}
});

export { contentClient };
