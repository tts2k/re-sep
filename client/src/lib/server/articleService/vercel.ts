import type { ArticleService } from "./type";
import type { Article } from "@/proto/content";
import * as turso from "../turso";

export const getArticle = async (entryName: string): Promise<Article> => {
	return await turso.getArticle(entryName);
};

const service: ArticleService = {
	getArticle,
};

export default service;
