import type { ArticleService } from "./type";
import { contentClient } from "../turso";
import type { Article, TOCItem } from "@/proto/content";
import zlib from "node:zlib";
import { promisify } from "node:util";
import { NotFoundError } from "../error";
import mustache from "mustache";
import { defaultConfig } from "@/defaultConfig";
import { getFontSizeMap } from "@/stylePresets";

const doGunzip = promisify(zlib.gunzip);

export const getArticle = async (_: string): Promise<Article> => {
	const result = await contentClient.execute("Select * from articles;");

	if (result.rows.length === 0) {
		throw new NotFoundError("Article not found");
	}

	const [row] = result.rows;

	const htmlTextBuffer = await doGunzip(row.html_text as ArrayBuffer);
	const htmlText = htmlTextBuffer.toString();
	const fszMap = getFontSizeMap(defaultConfig.fontSize);
	const rendered = mustache.render(htmlText, fszMap);

	return {
		title: row.title as string,
		entryName: row.entry_name as string,
		issued: new Date(row.issued as string),
		modified: new Date(row.modified as string),
		authors: JSON.parse(row.author as string) as string[],
		htmlText: rendered,
		toc: JSON.parse(row.toc as string) as TOCItem[],
	};
};

const service: ArticleService = {
	getArticle,
};

export default service;
