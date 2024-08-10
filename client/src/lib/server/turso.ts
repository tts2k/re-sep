import { createClient, type Client, type ResultSet } from "@libsql/client";
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
	let articleRes: ResultSet;

	if (entryName !== "") {
		articleRes = await contentClient.execute({
			sql: "SELECT * FROM articles WHERE entry_name = ?",
			args: [entryName],
		});
	} else {
		articleRes = await contentClient.execute(
			"SELECT * FROM articles ORDER BY RANDOM() LIMIT 1",
		);
	}

	if (articleRes.rows.length === 0) {
		throw new NotFoundError("Article not found");
	}

	let userConfig: UserConfig;
	if (email) {
		const ucRes = await configClient.execute({
			sql: "SELECT json(config) as config FROM config WHERE email = ?",
			args: [email],
		});

		if (ucRes.rows.length === 0) {
			userConfig = defaultConfig;
		} else {
			console.log(ucRes.rows[0]);
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
	email: string,
): Promise<UserConfig> => {
	const config = await configClient.execute({
		sql: `SELECT email FROM config WHERE email = ?`,
		args: [email],
	});
	console.log(config);

	let ucRes: ResultSet;
	if (config.rows.length === 0) {
		ucRes = await configClient.execute({
			sql: `INSERT INTO config (
				email, config
			) VALUES (
				?, json(?)
			)`,
			args: [email, JSON.stringify(uc)],
		});
	} else {
		ucRes = await configClient.execute({
			sql: `
				UPDATE config AS c
				SET config = json(?)
				FROM config
				WHERE c.email = ?
			`,
			args: [JSON.stringify(uc), email],
		});
	}

	if (ucRes.rowsAffected === 0) {
		throw new MutationFailed("Failed to upsert user config");
	}

	return uc;
};

// Init
init();
