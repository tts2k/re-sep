import type { PageServerLoad } from "./$types";
import articleService from "$lib/server/articleService";
import { error } from "@sveltejs/kit";
import { promiseResult } from "@/server/utils";

export const prerender = false;

export const load: PageServerLoad = async () => {
	const article = await promiseResult(articleService.getArticle("blame"));
	if (article.isErr()) {
		throw error(500, article.error);
	}
	return article.value;
};
