import mockHtml from "./blame.html?raw";

type TocItem = {
	name: string;
	subItem: TocItem[];
};

type MockData = {
	toc: TocItem[];
	content: string;
};

export const mockData: MockData = {
	toc: [],
	content: mockHtml,
};
