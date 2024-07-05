import { describe, test } from "bun:test";
import { createMetadata } from "./grpc";
import { jwtVerify } from "jose";
import { expect } from "vitest";
import { env } from "$env/dynamic/private";

describe("test createMetadata", () => {
	test("can generate a token", async () => {
		const metadata = await createMetadata("token");

		const authHeader = metadata.get("x-authorization");
		expect(authHeader.length).toBeGreaterThan(0);

		const bearer = authHeader.at(0);
		expect(bearer?.toString()).toBeTruthy();

		// type narrow
		if (!bearer) {
			return;
		}

		const token = bearer.toString().split(" ")[1];

		const secret = new TextEncoder().encode(env.JWT_SECRET);
		const { payload } = await jwtVerify(token, secret);

		expect(payload.sub).toEqual("token");
	});
});
