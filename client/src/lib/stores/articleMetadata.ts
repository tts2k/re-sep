import type { TocItem } from "@/server/articleService/type";
import { writable } from "svelte/store";

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
