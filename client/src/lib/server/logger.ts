import { env } from "$env/dynamic/private";
import pino from "pino";

const getPinoConfig = () => {
	if (env.TARGET === "development") {
		return {
			transport: {
				target: "pino-pretty",
				options: {
					colorize: true,
				},
			},
		};
	} else {
		return {
			level: "info",
		};
	}
};

export const logger = pino(getPinoConfig());
