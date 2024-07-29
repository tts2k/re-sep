import adapterNode from "@sveltejs/adapter-node";
import adapterVercel from "@sveltejs/adapter-vercel";
import adapterAuto from "@sveltejs/adapter-auto";
import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://kit.svelte.dev/docs/integrations#preprocessors
	// for more information about preprocessors
	preprocess: vitePreprocess(),

	kit: {
		alias: {
			"@/*": "./src/lib/*",
		},
	},
};

if (process.env.PLATFORM_SELF) {
	config.kit.adapter = adapterNode();
} else {
	config.kit.adapter = adapterAuto();
}

export default config;
