import type { PageServerLoad } from "./$types";
import articleService from "$lib/server/articleService";
import { promiseResult } from "@/server/utils";
import { error } from "@sveltejs/kit";
import { NotFoundError } from "@/server/error";
import type { Article } from "@/proto/content";
import type { UserConfig } from "@/proto/user_config";

export const prerender = false;

type LoadResponse = {
	article: Article;
	userConfig: UserConfig;
};

export const load: PageServerLoad = async ({
	locals,
}): Promise<LoadResponse> => {
	const session = await locals.auth();

	const article = await promiseResult(
		articleService.getArticle("", session?.user?.email),
	);

	if (article.isErr()) {
		if (article.error instanceof NotFoundError) {
			return error(404, { message: "Not found" });
		}

		return error(500, {
			message: article.error.message,
		});
	}

	if (!article.value.article) {
		return error(500, {
			message: "",
		});
	}

	if (!article.value.userConfig) {
		return error(500, {
			message: "",
		});
	}

	return {
		article: article.value.article,
		userConfig: article.value.userConfig,
	};
};
