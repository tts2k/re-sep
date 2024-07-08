import { contentClient, createMetadata } from "../grpc";
import type { Article, ArticleService } from "./type";

const getArticle = async (entryName: string, token?: string) => {
	let promiseExecutor: ConstructorParameters<typeof Promise<Article>>[0];

	if (token) {
		const metadata = await createMetadata(token);
		promiseExecutor = (resolve, reject) => {
			contentClient.getArticle(
				{ entryName },
				metadata,
				(error, response) => {
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
				},
			);
		};
	} else {
		promiseExecutor = (resolve, reject) => {
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
		};
	}

	return new Promise<Article>(promiseExecutor);
};

const service: ArticleService = {
	getArticle,
};

export default service;
