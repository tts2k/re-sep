import type { AuthResponse } from "@/proto/auth";
import type { UserConfig } from "@/proto/user_config";

/* Gonna need to do something about auth on different platforms.
 * But that's something I'm gonna think about later on
 */
export interface AuthService {
	getUser(provider: string, id: string): Promise<AuthResponse>;
	updateUserConfig(token: string, uc: UserConfig): Promise<UserConfig>;
}
