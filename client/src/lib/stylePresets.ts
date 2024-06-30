/* Font */
export const AvailableFonts = ["serif", "sans-serif", "Open Dyslexic"] as const;
export type Font = (typeof AvailableFonts)[number];

export const FontPreset: Record<Font, string> = {
	serif: "font-serif",
	"sans-serif": "font-sans",
	"Open Dyslexic": "font-open-dyslexic",
} as const;

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
	["text-5xl", "text-2xl", "text-xl", "text-lg", "text-base"],
	["text-6xl", "text-3xl", "text-2xl", "text-xl", "text-lg"],
	["text-7xl", "text-4xl", "text-3xl", "text-2xl", "text-xl"],
	["text-8xl", "text-5xl", "text-4xl", "text-3xl", "text-2xl"],
	["text-9xl", "text-6xl", "text-5xl", "text-4xl", "text-3xl"],
];

/* Margin */
export const MarginPresets = {
	left: ["ml-0", "ml-5p", "ml-10p", "ml-20p", "ml-30p", "ml-40p"],
	right: ["mr-0", "mr-5p", "mr-10p", "mr-20p", "mr-30p", "ml-40p"],
} as const;
