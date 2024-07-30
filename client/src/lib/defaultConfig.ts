import type { UserConfig } from "./proto/user_config";

export type UserConfigLayer = UserConfig & { layered?: boolean };

export const defaultConfig: UserConfigLayer = {
	layered: false,
	font: "serif",
	fontSize: 3,
	justify: false,
	margin: {
		left: 3,
		right: 3,
	},
};
