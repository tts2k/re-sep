import { Err, Ok, type Result } from "ts-results-es";

export const promiseResult = async <T>(
	promise: Promise<T>,
): Promise<Result<T, Error>> => {
	let result: Result<T, Error>;

	try {
		const res = await promise;
		result = Ok(res);
	} catch (error: any) {
		if (!(error instanceof Error)) {
			throw Error("Cannot handle non-Error type");
		}
		// Assuming error is always Error type, which is not always
		result = Err(error);
	}

	return result;
};
