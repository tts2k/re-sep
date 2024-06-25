/* Font */
export const AvailableFonts = ["serif", "sans-serif", "Open Dyslexic"] as const;
export type Font = (typeof AvailableFonts)[number];

export const FontPreset: Record<Font, string> = {
	serif: "font-serif",
	"sans-serif": "font-sans",
	"Open Dyslexic": "font-open-dyslexic",
} as const;

/* Font size */
export type FontSizePreset = {
	h1: string;
	h2: string;
	h3: string;
	h4: string;
	text: string;
};

const fontSizePresets = {
	h1: ["text-5xl", "text-6xl", "text-7xl", "text-8xl", "text-9xl"],
	h2: ["text-2xl", "text-3xl", "text-4xl", "text-5xl", "text-6xl"],
	h3: ["text-xl", "text-2xl", "text-3xl", "text-4xl", "text-5xl"],
	h4: ["text-lg", "text-xl", "text-2xl", "text-3xl", "text-4xl"],
	text: ["text-base", "text-lg", "text-xl", "text-3xl", "text-4xl"],
};

export const getFontSizePreset = (preset: number): FontSizePreset => {
	return {
		h1: fontSizePresets.h1[preset],
		h2: fontSizePresets.h2[preset],
		h3: fontSizePresets.h3[preset],
		h4: fontSizePresets.h4[preset],
		text: fontSizePresets.text[preset],
	};
};

/* Margin */
export const MarginPresets = {
	left: ["ml-0", "ml-1", "ml-2", "ml-3", "ml-4"],
	right: ["mr-0", "mr-1", "mr-2", "mr-3", "mr-4"],
} as const;
