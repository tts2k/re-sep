import type { ArticleService } from "./type";
import type { Article } from "@/proto/content";
import * as turso from "../turso";

export const getArticle = async (
	entryName: string,
	email: string,
): Promise<Article> => {
	return await turso.getArticle(entryName, email);
};

const service: ArticleService = {
	getArticle,
};

export default service;
