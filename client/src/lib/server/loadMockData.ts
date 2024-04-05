import mockHtml from "./blame.html?raw";

type TocItem = {
	label: string;
	id: string;
	subItems: TocItem[];
};

type MockData = {
	title: string;
	author: string[];
	toc: TocItem[];
	content: string;
};

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

// The mock data html is raw and not sanitized. Ideally it should be sanitized from the golang
// backend before saving into the database.
// All mock data will be processed from the golang backend. Sveltekit server only job is to handle SSR
export const mockData: MockData = {
	title: "Blame",
	author: ["Tognazzini, Neal", "Coates, D. Justin"],
	toc: toc,
	content: mockHtml,
};
