import { createClient, type Client } from "@libsql/client";
import { env } from "$env/dynamic/private";
import { MutationFailed, NotFoundError } from "./error";
import { promisify } from "node:util";
import zlib from "node:zlib";
import type { Article, TOCItem } from "@/proto/content";
import { getFontSizeMap } from "@/stylePresets";
import { defaultConfig } from "@/defaultConfig";
import mustache from "mustache";
import type { UserConfig } from "@/proto/user_config";

let contentClient: Client;
let configClient: Client;

const init = () => {
	contentClient = createClient({
		url: env.CONTENT_DATABASE_URL || "",
		authToken: env.CONTENT_AUTH_TOKEN || "",
	});
	configClient = createClient({
		url: env.CONFIG_DATABASE_URL || "",
		authToken: env.CONFIG_AUTH_TOKEN || "",
	});
};

const doGunzip = promisify(zlib.gunzip);

export const getArticle = async (
	entryName: string,
	email: string,
): Promise<Article> => {
	const articleRes = await contentClient.execute({
		sql: "SELECT * FROM articles WHERE entry_name = ?",
		args: [entryName],
	});

	if (articleRes.rows.length === 0) {
		throw new NotFoundError("Article not found");
	}

	let userConfig: UserConfig;
	if (email) {
		const ucRes = await configClient.execute({
			sql: "SELECT json(config) FROM config WHERE email = ?",
			args: [email],
		});

		if (ucRes.rows.length === 0) {
			userConfig = defaultConfig;
		} else {
			userConfig = JSON.parse(
				ucRes.rows[0].config as string,
			) as UserConfig;
		}
	} else {
		userConfig = defaultConfig;
	}

	const [articleRow] = articleRes.rows;

	const htmlTextBuffer = await doGunzip(articleRow.html_text as ArrayBuffer);
	const htmlText = htmlTextBuffer.toString();
	const fszMap = getFontSizeMap(defaultConfig.fontSize - 1);

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

export const updateUserConfig = async (
	uc: UserConfig,
	userId: string,
): Promise<UserConfig> => {
	const ucRes = await configClient.execute({
		sql: `INSERT INTO config (
			email, config
		) VALUES (
			?, json(?)
		)
		ON CONFLICT (entry_name) DO UPDATE SET
			config=json(excluded.title),
		`,
		args: [userId, JSON.stringify(uc)],
	});

	if (ucRes.rows.length === 0) {
		throw new MutationFailed("Failed to upsert user config");
	}

	return JSON.parse(ucRes.rows[0].config as string) as UserConfig;
};

// Init
init();
