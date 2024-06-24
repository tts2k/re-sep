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

const h1SizePresets = [
	"text-5xl",
	"text-6xl",
	"text-7xl",
	"text-8xl",
	"text-9xl",
] as const;

const h2SizePresets = [
	"text-2xl",
	"text-3xl",
	"text-4xl",
	"text-5xl",
	"text-6xl",
] as const;

const h3SizePresets = [
	"text-xl",
	"text-2xl",
	"text-3xl",
	"text-4xl",
	"text-5xl",
] as const;

const h4SizePresets = [
	"text-lg",
	"text-xl",
	"text-2xl",
	"text-3xl",
	"text-4xl",
] as const;

const textSizePresets = [
	"text-base",
	"text-lg",
	"text-xl",
	"text-3xl",
	"text-4xl",
] as const;

export const getFontSizePreset = (preset: number): FontSizePreset => {
	return {
		h1: h1SizePresets[preset],
		h2: h2SizePresets[preset],
		h3: h3SizePresets[preset],
		h4: h4SizePresets[preset],
		text: textSizePresets[preset],
	};
};
