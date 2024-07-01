export type User = {
	id: string;
	name: string;
};

export type AuthResponse = {
	token: string;
	user: User;
};

/* Gonna need to do something about auth on platforms.
 * But that's something I'm gonna think about later on
 */
export interface AuthService {
	auth(token: string): Promise<AuthResponse>;
}
