import type { PageServerLoad } from "./$types";
import articleService from "$lib/server/articleService";
import { promiseResult } from "@/server/utils";
import { error } from "@sveltejs/kit";

export const prerender = false;

export const load: PageServerLoad = async ({ cookies }) => {
	const token = cookies.get("token");

	const article = await promiseResult(
		articleService.getArticle("action", token),
	);

	if (article.isErr()) {
		if (article.error instanceof NotFoundError) {
			throw error(404, "Not found");
		}

		throw error(500, article.error);
	}

	console.log(article.value.toc[3].subItems);

	return article.value;
};
