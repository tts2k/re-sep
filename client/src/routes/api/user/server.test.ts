import type { UserConfig } from "@/proto/user_config";
import { test, describe, expect } from "bun:test";
import { POST } from "./+server";
import { json, type RequestEvent } from "@sveltejs/kit";
import { mock } from "bun:test";
import type { Session } from "@auth/sveltekit";

describe("test POST", () => {
	type TestCase = {
		name: string,
		request: Request,
		session: Session,
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
			name: "no session",
			request: request.clone(),
			session: {
				user: undefined,
				expires: ""
			},
			res: unauthorizedResponse.clone()
		},
		{
			name: "authorized, grpc error",
			request: request.clone(),
			session: {
				user: {
					name: "user",
					id: "user"
				},
				expires: ""
			},
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
			session: {
				user: {
					name: "user",
					id: "user"
				},
				expires: ""
			},
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
				async auth() { return t.session },
			} as App.Locals

			const res = await POST({ locals, request: t.request } as RequestEvent)
			expect(res.status).toEqual(t.res.status)
			expect(res.statusText).toEqual(t.res.statusText)

			const expectBody = await t.res.json()
			const resBody = await res.json()
			expect(expectBody).toEqual(resBody)
		})
	}
});
