import mockHtml from "../assets/blame.html?raw";
import type { TocItem, Article, ArticleService } from "./type";

const toc: TocItem[] = [
	{
		label: "What is Blame?",
		id: "WhaBla",
		subItems: [
			{
				label: "Cognitive Theories of Blame",
				id: "CogTheBla",
				subItems: [],
			},
			{
				label: "Emotional Theories of Blame",
				id: "EmoTheBla",
				subItems: [],
			},
			{
				label: "Conative Theories of Blame",
				id: "ConTheBla",
				subItems: [
					{
						label: "Dispositions Around a Belief-Desire Pair",
						id: "DisAroBelDesPai",
						subItems: [],
					},
					{
						label: "Attitude Adjustment in Response to Impairment",
						id: "AttAdjResImp",
						subItems: [],
					},
				],
			},
			{
				label: "Functional Theories of Blame",
				id: "FunTheBla",
				subItems: [],
			},
		],
	},
	{
		label: "When is Blame Appropriate?",
		id: "WheBlaApp",
		subItems: [
			{
				label: "Facts about the Person Being Blamed",
				id: "FacAboPers",
				subItems: [
					{
						label: "Moral Agency",
						id: "MorAge",
						subItems: [],
					},
					{
						label: "Freedom and Responsibility",
						id: "Fre",
						subItems: [],
					},
				],
			},
			{
				label: "Facts about the Blaming Interaction",
				id: "FacAboInter",
				subItems: [
					{
						label: "Proportionality",
						id: "Prop",
						subItems: [],
					},
					{
						label: "Epistemic Consideration",
						id: "Epis",
						subItems: [],
					},
				],
			},
			{
				label: "Facts about the Blamer",
				id: "FacAboBlam",
				subItems: [
					{
						id: "Stand",
						label: "Standing",
						subItems: [],
					},
					{
						label: "Hypocrisy",
						id: "Hypo",
						subItems: [],
					},
					{
						label: "Other Blamer-Based Worries",
						id: "OthBlamWor",
						subItems: [],
					},
				],
			},
			{
				label: "Varieties of Blame",
				id: "VarBla",
				subItems: [],
			},
		],
	},
	{
		label: "Bibliography",
		id: "Bib",
		subItems: [],
	},
	{
		label: "Other Internet Resoources",
		id: "Oth",
		subItems: [],
	},
	{
		label: "Related Entires",
		id: "Rel",
		subItems: [],
	},
	{
		label: "Acknowledgements",
		id: "acknowledgments",
		subItems: [],
	},
];

const mockArticle: Article = {
	title: "Blame",
	entryName: "blame",
	author: ["Tognazzini, Neal", "Coates, D. Justin"],
	toc: toc,
	htmlText: mockHtml,
	issued: new Date().toISOString(),
	modified: new Date().toISOString(),
};

export const getArticle = async () => {
	return mockArticle;
};

const service: ArticleService = {
	getArticle,
};

export default service;
