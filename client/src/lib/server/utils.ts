import { SignJWT, type JWTPayload } from "jose";
import { Err, Ok, type Result } from "ts-results-es";

export const promiseResult = async <T>(
	promise: Promise<T>,
): Promise<Result<T, Error>> => {
	let result: Result<T, Error>;

	try {
		const res = await promise;
		result = Ok(res);
	} catch (error: unknown) {
		if (!(error instanceof Error)) {
			throw Error("Cannot handle non-Error type");
		}
		// Assuming error is always Error type, which is not always
		result = Err(error);
	}

	return result;
};

export const createJWTToken = async (id: string, secret: string) => {
	const tokenPayload: JWTPayload = { sub: id };

	const secretEnc = new TextEncoder().encode(secret);

	// Generate and sign the token
	const oAuthToken = await new SignJWT(tokenPayload)
		.setProtectedHeader({ alg: "HS256", typ: "JWT" })
		.setIssuedAt()
		.setExpirationTime("1h")
		.sign(secretEnc);

	return oAuthToken;
};
