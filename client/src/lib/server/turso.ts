import { createClient } from "@libsql/client";
import { env } from "$env/dynamic/private";

const contentClient = createClient({
	url: env.CONTENT_DATABASE_URL || "",
	authToken: env.CONTENT_AUTH_TOKEN || "",
});

export { contentClient };
