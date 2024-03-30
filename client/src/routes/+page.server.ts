import type { PageServerLoad } from "./$types";
import { mockData } from "@/server/loadMockData";

export const load: PageServerLoad = () => {
	return mockData;
};
