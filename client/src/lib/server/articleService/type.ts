import type { ArticleResponse } from "@/proto/main";

export interface ArticleService {
	// Get article. Return a random article if entry name is empty
	getArticle(
		entryName: string,
		token?: string | null,
	): Promise<ArticleResponse>;
}
