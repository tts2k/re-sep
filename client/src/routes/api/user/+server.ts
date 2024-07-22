import { json, type RequestHandler } from "@sveltejs/kit";
import authService from "@/server/authService";
import { promiseResult } from "@/server/utils";
import { logger } from "@/server/logger";

export const POST: RequestHandler = async ({ cookies, locals, request }) => {
	const user = locals.user;
	const token = cookies.get("token");

	// unauthorized
	if (!user?.sub || !token) {
		logger.error("Unauthorized POST call to request handler");
		return json(
			{ message: "unauthorized" },
			{ status: 401, statusText: "unauthorized" },
		);
	}

	const data = await request.json();

	// Validate is on the go service side
	const res = await promiseResult(authService.updateUserConfig(token, data));
	if (res.isErr()) {
		logger.error("Error updating user config", res.error.message);
		return json(
			{ message: "Error updating user config" },
			{ status: 403, statusText: "Bad request" },
		);
	}

	return json({ message: "success" });
};
