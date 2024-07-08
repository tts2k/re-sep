import type { AuthResponse, AuthService } from "./type";

const mockUser = {
	id: "demo",
	name: "demo",
};

const mockAuthResponse = {
	token: "demo",
	user: mockUser,
};

const auth = async (): Promise<AuthResponse> => {
	return mockAuthResponse;
};

const service: AuthService = {
	auth,
};

export default service;
