import { createClient } from "@libsql/client";
import { env } from "$env/dynamic/private";
import { NotFoundError } from "./error";
import { promisify } from "util";
import zlib from "node:zlib";
import type { Article, TOCItem } from "@/proto/content";
import { getFontSizeMap } from "@/stylePresets";
import { defaultConfig } from "@/defaultConfig";
import mustache from "mustache";
import type { UserConfig } from "@/proto/user_config";

const contentClient = createClient({
	url: env.CONTENT_DATABASE_URL || "",
	authToken: env.CONTENT_AUTH_TOKEN || "",
});

const userClient = createClient({
	url: env.USER_DATABASE_URL || "",
	authToken: env.USER_AUTH_TOKEN || "",
});

const doGunzip = promisify(zlib.gunzip);

export const getArticle = async (entryName: string): Promise<Article> => {
	const articlePromise = contentClient.execute({
		sql: "SELECT * FROM articles WHERE entry_name = ?",
		args: [entryName],
	});

	const userPromise = userClient.execute({
		sql: "SELECT config FROM users LIMIT 1",
		args: [],
	});

	const [articleRes, userConfigRes] = await Promise.all([
		articlePromise,
		userPromise,
	]);

	if (articleRes.rows.length === 0) {
		throw new NotFoundError("Article not found");
	}

	let userConfig: UserConfig;
	if (userConfigRes.rows.length === 0 || !userConfigRes.rows[0].config) {
		userConfig = defaultConfig;
	} else {
		userConfig = JSON.parse(
			userConfigRes.rows[0].config as string,
		) as UserConfig;
	}

	const [articleRow] = articleRes.rows;

	const htmlTextBuffer = await doGunzip(articleRow.html_text as ArrayBuffer);
	const htmlText = htmlTextBuffer.toString();
	const fszMap = getFontSizeMap(defaultConfig.fontSize);

	return {
		title: articleRow.title as string,
		entryName: articleRow.entry_name as string,
		issued: new Date(articleRow.issued as string),
		modified: new Date(articleRow.modified as string),
		authors: JSON.parse(articleRow.author as string) as string[],
		htmlText: mustache.render(htmlText, fszMap),
		toc: JSON.parse(articleRow.toc as string) as TOCItem[],
	};
};

export { contentClient };
