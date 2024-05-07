import type { PageServerLoad } from "./$types";
import articleService from "$lib/server/articleService";
import { error } from "@sveltejs/kit";

export const prerender = false;

export const load: PageServerLoad = async () => {
	try {
		const article = await articleService.getArticle("blame");
		return article;
	} catch (err) {
		console.error(err);
		error(500, "Internal server error");
	}
};
