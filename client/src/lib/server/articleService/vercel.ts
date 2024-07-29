import type { ArticleService } from "./type";
import { contentClient } from "../turso";
import type { Article, TOCItem } from "@/proto/content";
import zlib from "node:zlib";
import { promisify } from "node:util";

const doGunzip = promisify(zlib.gunzip);

export const getArticle = async (entryName: string): Promise<Article> => {
	const result = await contentClient.execute({
		sql: "SELECT * FROM articles where entry_name = ?",
		args: [entryName],
	});

	if (result.rows.length === 0) {
		throw new NotFoundError("Article not found");
	}

	const [row] = result.rows;

	const htmlTextBuffer = await doGunzip(row.html_text as ArrayBuffer);

	return {
		title: row.title as string,
		entryName: row.entry_name as string,
		issued: new Date(row.issued as string),
		modified: new Date(row.modified as string),
		authors: JSON.parse(row.author as string) as string[],
		htmlText: htmlTextBuffer.toString(),
		toc: JSON.parse(row.toc as string) as TOCItem[],
	};
};

const service: ArticleService = {
	getArticle,
};

export default service;
