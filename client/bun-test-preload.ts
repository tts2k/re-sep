import { plugin } from "bun";

plugin({
	name: "sveltekit-env",
	setup(build) {
		build.module("$env/dynamic/private", () => {
			return {
				exports: {
					env: {
						JWT_SECRET: "hTZ66AYJgxylQJmiTXstSdhTuq2D3DUw",
						AUTH_URL: "localhost",
						CONTENT_URL: "localhost",
					},
				},
				loader: "object",
			};
		});
		build.module("$app/environment", () => {
			return {
				exports: {
					building: false,
				},
				loader: "object",
			};
		});
	},
});
