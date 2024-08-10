import type { ArticleService } from "./type";
import * as turso from "../turso";
import type { ArticleResponse } from "@/proto/main";

export const getArticle = async (
	entryName: string,
	email: string,
): Promise<ArticleResponse> => {
	return await turso.getArticle(entryName, email);
};

const service: ArticleService = {
	getArticle,
};

export default service;
