import type { Article } from "@/proto/content";
import { contentClient, createMetadata } from "../grpc";
import type { ArticleService } from "./type";

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

					resolve(response);
				},
			);
		};
	} else {
		promiseExecutor = (resolve, reject) => {
			contentClient.getArticle({ entryName }, (error, response) => {
				if (error !== null) {
					reject(error);
				}

				resolve(response);
			});
		};
	}

	return new Promise<Article>(promiseExecutor);
};

const service: ArticleService = {
	getArticle,
};

export default service;
