import { writable } from "svelte/store";

export type TocItem = {
	label: string;
	id: string;
	subItems: TocItem[];
};

export type ArticleMetadata = {
	title: string;
	authors: string[];
	toc: TocItem[];
};

const initialMetadata: ArticleMetadata = {
	title: "",
	authors: [],
	toc: [],
};

export const metadata = writable(initialMetadata);
