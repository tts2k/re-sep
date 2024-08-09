// import { building } from "$app/environment";
// import { logger } from "@/server/logger";
// import { redirect, type Handle, type RequestEvent } from "@sveltejs/kit";
// import { promiseResult } from "@/server/utils";
// import authService from "@/server/authService";

export { handle } from "@/server/auth";

// const doAuth = async (
// 	event: RequestEvent,
// 	token: string,
// ): Promise<RequestEvent> => {
// 	// Reset cookies
// 	event.cookies.set("token", "", {
// 		path: "/",
// 		maxAge: 0,
// 	});
//
// 	const auth = await promiseResult(authService.auth(token));
// 	if (auth.isErr()) {
// 		logger.error(`Error during auth: ${auth.error}`);
// 		throw redirect(302, "/?error=unauthorized");
// 	}
// 	event.locals.user = auth.value.user;
//
// 	if (!event.locals.user?.sub) {
// 		logger.error("No user found");
// 		throw redirect(302, "/?error=unauthorized");
// 	}
//
// 	event.cookies.set("token", auth.value.token, {
// 		path: "/",
// 		maxAge: 604800, // 7 days
// 		sameSite: "lax",
// 		secure: true,
// 		httpOnly: true,
// 	});
// 	return event;
// };
//
// export const handle: Handle = async ({ event, resolve }) => {
// 	logger.info(`path: ${event.url.pathname}`);
// 	if (building) {
// 		return await resolve(event);
// 	}
//
// 	if (event.url.pathname.startsWith("/api")) {
// 		return await resolve(event);
// 	}
//
// 	// Token check
// 	// If there is no token in search params but there is one in the cookie,
// 	// then do auth to reset the token
// 	const token = event.url.searchParams.get("token");
// 	if (!token) {
// 		const cToken = event.cookies.get("token");
// 		if (cToken) {
// 			const e = await doAuth(event, cToken);
// 			return await resolve(e);
// 		}
//
// 		// reset the user if exists since no cookie
// 		event.locals.user = undefined;
// 		return await resolve(event);
// 	}
//
// 	// If token is in the param, do auth normally
// 	event.locals.user = undefined;
//
// 	event = await doAuth(event, token);
//
// 	return await resolve(event);
// };
