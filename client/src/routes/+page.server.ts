import type { PageServerLoad } from "./$types";
import articleService from "$lib/server/articleService";
import { promiseResult } from "@/server/utils";
import { error } from "@sveltejs/kit";
import { NotFoundError } from "@/server/error";

export const prerender = false;

export const load: PageServerLoad = async ({ cookies }) => {
	const token = cookies.get("token");

	const article = await promiseResult(
		articleService.getArticle("action", token),
	);

	if (article.isErr()) {
		console.log(article.error);
		if (article.error instanceof NotFoundError) {
			return error(404, { message: "Not found" });
		}

		return error(500, {
			message: article.error.message,
		});
	}

	return article.value;
};
