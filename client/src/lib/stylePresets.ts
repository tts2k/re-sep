/* Font */
export const AvailableFonts = ["serif", "sans-serif", "Open Dyslexic"];
export type Font = (typeof AvailableFonts)[number];

export const FontPreset: Record<Font, string> = {
	serif: "font-serif",
	"sans-serif": "font-sans",
	"Open Dyslexic": "font-open-dyslexic",
};

/* biome-ignore lint: lint/suspiciouos/noConstEnum: for limited use and
 * experiment only, as const enum can break stuff
 * https://www.typescriptlang.org/docs/handbook/enums.html#const-enums
 */
export const enum FontSizeTag {
	H1 = 0,
	H2 = 1,
	H3 = 2,
	H4 = 3,
	TEXT = 4,
}

export const FontSizePresets = [
	["text-9xl", "text-6xl", "text-5xl", "text-4xl", "text-3xl"],
	["text-8xl", "text-5xl", "text-4xl", "text-3xl", "text-2xl"],
	["text-7xl", "text-4xl", "text-3xl", "text-2xl", "text-xl"],
	["text-6xl", "text-3xl", "text-2xl", "text-xl", "text-lg"],
	["text-5xl", "text-2xl", "text-xl", "text-lg", "text-base"],
];

/* Margin */
export const MarginPresets = {
	left: ["ml-0", "ml-5p", "ml-10p", "ml-20p", "ml-30p", "ml-40p"],
	right: ["mr-0", "mr-5p", "mr-10p", "mr-20p", "mr-30p", "ml-40p"],
} as const;

export const getFontSizeMap = (size: number) => {
	const m: Record<string, string> = {};

	m.h1 = FontSizePresets[FontSizeTag.H1][size];
	m.h2 = FontSizePresets[FontSizeTag.H2][size];
	m.h3 = FontSizePresets[FontSizeTag.H3][size];
	m.h4 = FontSizePresets[FontSizeTag.H4][size];
	m.text = FontSizePresets[FontSizeTag.TEXT][size];

	return m;
};
