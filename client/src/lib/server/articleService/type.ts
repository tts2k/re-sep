export type TocItem = {
	label: string;
	id: string;
	subItems: TocItem[];
};

export type Article = {
	title: string;
	entryName: string;
	author: string[];
	toc: TocItem[];
	htmlText: string;
	issued: string;
	modified: string;
};

export interface ArticleService {
	getArticle(entryName: string): Article;
}
