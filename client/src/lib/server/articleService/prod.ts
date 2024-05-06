import { contentClient } from "../grpc";
import type { Article, ArticleService } from "./type";

export const getArticle = async (entryName: string) => {
	return new Promise<Article>((resolve, reject) => {
		contentClient.getArticle({ entryName }, (error, response) => {
			if (error !== null) {
				reject(error);
			}

			const article: Article = {
				title: response.title,
				entryName: response.entryName,
				issued: response.issued?.toLocaleString() || "",
				modified: response.modified?.toLocaleString() || "",
				author: response.authors,
				toc: response.toc,
				htmlText: response.htmlText,
			};

			resolve(article);
		});
	});
};

const service: ArticleService = {
	getArticle,
};

export default service;
