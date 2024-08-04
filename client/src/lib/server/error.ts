export class NotFoundError extends Error {
	constructor(m: string) {
		super(m);

		Object.setPrototypeOf(this, Error.prototype);
	}
}

export class MutationFailed extends Error {
	constructor(m: string) {
		super(m);

		Object.setPrototypeOf(this, Error.prototype);
	}
}
