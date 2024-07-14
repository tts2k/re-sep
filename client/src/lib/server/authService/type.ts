import type { UserConfig } from "@/stores/userConfig";

export type User = {
	id: string;
	name: string;
};

export type AuthResponse = {
	token: string;
	user: User;
};

/* Gonna need to do something about auth on different platforms.
 * But that's something I'm gonna think about later on
 */
export interface AuthService {
	auth(token: string): Promise<AuthResponse>;
	updateUserConfig(token: string): Promise<UserConfig>;
}
