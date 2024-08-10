import type { PageServerLoad } from "./$types";
import articleService from "$lib/server/articleService";
import { promiseResult } from "@/server/utils";
import { error } from "@sveltejs/kit";
import { NotFoundError } from "@/server/error";

export const prerender = false;

export const load: PageServerLoad = async ({ locals, params }) => {
	const session = await locals.auth();

	const article = await promiseResult(
		articleService.getArticle(params.slug, session?.user?.email),
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
