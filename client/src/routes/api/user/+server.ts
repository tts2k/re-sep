import { json, type RequestHandler } from "@sveltejs/kit";
import authService from "@/server/authService";
import { promiseResult } from "@/server/utils";
import { logger } from "@/server/logger";

export const POST: RequestHandler = async ({ locals, request }) => {
	const session = await locals.auth();

	// unauthorized
	if (!session?.user?.email) {
		logger.error("Unauthorized POST call to user api");
		return json(
			{ message: "unauthorized" },
			{ status: 401, statusText: "unauthorized" },
		);
	}

	const data = await request.json();

	// Validate is on the go service side
	const res = await promiseResult(
		authService.updateUserConfig(data, session.user.email),
	);
	if (res.isErr()) {
		console.error(res.error);
		logger.error("Error updating user config", res.error.message);
		return json(
			{ message: "Error updating user config" },
			{ status: 403, statusText: "Bad request" },
		);
	}

	return json({ message: "success" });
};
