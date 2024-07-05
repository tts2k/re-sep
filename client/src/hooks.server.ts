import { building } from "$app/environment";
import { logger } from "@/server/logger";
import { redirect, type Handle } from "@sveltejs/kit";
import { promiseResult } from "@/server/utils";
import authService from "$lib/server/authService";

export const handle: Handle = async ({ event, resolve }) => {
	logger.info(`path: ${event.url.pathname}`);
	if (building) {
		return await resolve(event);
	}

	delete event.locals.user;

	// Token check
	const token = event.url.searchParams.get("token");
	if (!token) {
		return await resolve(event);
	}

	// Reset cookies
	event.cookies.set("token", "", {
		path: "/",
		maxAge: 0,
	});

	const auth = await promiseResult(authService.auth(token));
	if (auth.isErr()) {
		logger.error(`Error during auth: ${auth.error}`);
		throw redirect(302, "/?error=unauthorized");
	}

	event.locals.user = auth.value.user;

	if (!event.locals.user?.id) {
		logger.error("No user found");
		throw redirect(302, "/?error=unauthorized");
	}

	event.cookies.set("token", auth.value.token, {
		path: "/",
		maxAge: 604800, // 7 days
		sameSite: "lax",
		secure: true,
		httpOnly: true,
	});

	return await resolve(event);
};
