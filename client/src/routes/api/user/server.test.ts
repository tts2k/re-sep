import type { User } from "@/proto/auth";
import type { UserConfig } from "@/proto/user_config";
import { test, describe, expect } from "bun:test";
import { POST } from "./+server";
import { json, type Cookies, type RequestEvent } from "@sveltejs/kit";
import { mock } from "bun:test";

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

	get(key: string): any {
		return this.cookies[key]?.value;
	}

	getAll(_2: any): any {
		return;
	}

	serialize(_: string, _2: string, _3: any): string {
		return "";
	}

	delete(_: string, _2: any): void {}
}

describe("test POST", () => {
	type TestCase = {
		name: string,
		user?: User,
		request: Request,
		cookies: Cookies,
		res: Response,
		mockAuthService?: any;
	};

	const uc: UserConfig = {
		font: "serif",
		fontSize: 3,
		justify: false,
		margin: {
			left: 3,
			right: 3,
		},
	}

	const request = new Request("https://example.com/api/user", {
		method: "POST",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify(uc)
	})

	const unauthorizedResponse = json(
		{ message: "unauthorized" },
		{ status: 401, statusText: "unauthorized" }
	)

	const testCases: TestCase[] = [
		{
			name: "no user",
			request: request.clone(),
			cookies: new MockCookies(),
			res: unauthorizedResponse.clone()
		},
		{
			name: "no cookie",
			request: request.clone(),
			user: {
				sub: "sub",
				name: "user"
			},
			cookies: new MockCookies(),
			res: unauthorizedResponse.clone()
		},
		{
			name: "authorized, grpc error",
			request: request.clone(),
			user: {
				sub: "sub",
				name: "user"
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
			res: json(
				{ message: "Error updating user config" },
				{ status: 403, statusText: "Bad request" }
			).clone(),
			mockAuthService: {
				updateUserConfig: async (_: string): Promise<UserConfig> => {
					throw new Error("error")
				},
			},
		},
		{
			name: "authorized, success",
			request: request.clone(),
			user: {
				sub: "sub",
				name: "user"
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
			res: json({ message: "success" }).clone(),
			mockAuthService: {
				updateUserConfig: async (_: string): Promise<UserConfig> => {
					return uc
				},
			},
		},
	]

	for (const t of testCases) {
		test(t.name, async () => {
			if (t.mockAuthService) {
				mock.module("@/server/authService", () => {
					return {
						default: t.mockAuthService,
					};
				});
			}

			const locals = {
				user: t.user
			}

			const res = await POST({ cookies: t.cookies, locals, request: t.request } as RequestEvent)
			expect(res.status).toEqual(t.res.status)
			expect(res.statusText).toEqual(t.res.statusText)

			const expectBody = await t.res.json()
			const resBody = await res.json()
			expect(expectBody).toEqual(resBody)
		})
	}
});
