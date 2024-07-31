import type { Article } from "@/proto/content";

export interface ArticleService {
	getArticle(entryName: string, token?: string | null): Promise<Article>;
}
