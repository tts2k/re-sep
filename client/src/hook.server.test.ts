import type { AuthResponse } from "@/proto/auth";
import { describe, mock, test, expect } from "bun:test";
import { handle } from "./hooks.server";
import { type Cookies, type RequestEvent } from "@sveltejs/kit";

class MockCookies implements Cookies {
	cookies: Record<string, any> = {};

	constructor(cookies?: Record<string, any>) {
		if (cookies) {
			this.cookies = cookies;
		}
	}

	set(key: string, value: string, opts: any) {
		this.cookies[key] = {
			value,
			opts,
		};
	}

	get(_: string, _2: any): any {
		return;
	}

	getAll(_2: any): any {
		return;
	}

	serialize(_: string, _2: string, _3: any): string {
		return "";
	}

	delete(_: string, _2: any): void {}
}

describe("Hook handle test", () => {
	type TestCase = {
		name: string;
		mockAuthService?: any;
		throw?: any;
		requestEvent: RequestEvent;
		expect: RequestEvent;
	};

	const testCases: TestCase[] = [
		{
			name: "path name isn't root",
			requestEvent: {
				url: new URL("user", "http://example.com"),
			} as RequestEvent,
			expect: {
				url: new URL("user", "http://example.com"),
			} as RequestEvent,
		},
		{
			name: "path name is root but no token",
			requestEvent: {
				url: new URL("", "http://example.com"),
				locals: {
					user: {
						sub: "user",
						name: "user",
					},
				},
			} as RequestEvent,
			expect: {
				url: new URL("", "http://example.com"),
				locals: {
					user: undefined,
				},
			} as RequestEvent,
		},
		{
			name: "auth but user not found",
			requestEvent: {
				url: new URL("?token=token", "http://example.com"),
				locals: {
					user: {
						sub: "user",
						name: "user",
					},
				},
				cookies: new MockCookies() as Cookies,
			} as RequestEvent<any, any>,
			expect: {
				url: new URL("?token=token", "http://example.com"),
				locals: {
					user: {
						name: "user",
						sub: "sub",
					},
				},
				cookies: new MockCookies({
					token: {
						value: "token",
						opts: {
							httpOnly: true,
							maxAge: 604800,
							path: "/",
							sameSite: "lax",
							secure: true,
						},
					},
				}) as Cookies,
			} as RequestEvent,
			mockAuthService: {
				auth: async (_: string): Promise<AuthResponse> => {
					throw new Error("no user");
				},
			},
			throw: {
				status: 302,
				location: "/?error=unauthorized",
			},
		},
		{
			name: "success run with auth",
			requestEvent: {
				url: new URL("?token=token", "http://example.com"),
				locals: {
					user: {
						sub: "user",
						name: "user",
					},
				},
				cookies: new MockCookies() as Cookies,
			} as RequestEvent<any, any>,
			expect: {
				url: new URL("?token=token", "http://example.com"),
				locals: {
					user: {
						name: "user",
						sub: "sub",
					},
				},
				cookies: new MockCookies({
					token: {
						value: "token",
						opts: {
							httpOnly: true,
							maxAge: 604800,
							path: "/",
							sameSite: "lax",
							secure: true,
						},
					},
				}) as Cookies,
			} as RequestEvent,
			mockAuthService: {
				auth: async (token: string): Promise<AuthResponse> => {
					return {
						token: token,
						user: {
							sub: "sub",
							name: "user",
						},
					};
				},
			},
		},
	];

	for (const t of testCases) {
		test(t.name, async () => {
			mock.module("@/server/authService", () => {
				return {
					default: t.mockAuthService,
				};
			});

			try {
				const resolve = mock((input: any) => input);
				await handle({
					event: t.requestEvent,
					resolve,
				});

				expect(resolve.mock.lastCall?.[0]).toEqual(t.expect);
			} catch (error: unknown) {
				expect(error).toEqual(t.throw);
			}
		});
	}
});
