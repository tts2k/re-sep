import type { PageServerLoad } from "./$types";
import articleService from "$lib/server/articleService";
import { error } from "@sveltejs/kit";

export const load: PageServerLoad = () => {
	try {
		const article = articleService.getArticle("");
		return article;
	} catch (err) {
		console.error(err);
		error(500, "Internal server error");
	}
};
