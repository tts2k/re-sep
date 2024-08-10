import type { TOCItem } from "@/proto/content";
import { writable } from "svelte/store";

export type ArticleMetadata = {
	title: string;
	authors: string[];
	entryName: string;
	toc: TOCItem[];
};

const initialMetadata: ArticleMetadata = {
	title: "",
	entryName: "",
	authors: [],
	toc: [],
};

export const metadata = writable(initialMetadata);
