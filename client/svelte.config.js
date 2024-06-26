import adapterAuto from "@sveltejs/adapter-auto";
import adapterNode from "@sveltejs/adapter-node";
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

if (process.env.DEMO === "true") {
	config.kit.adapter = adapterAuto();
} else {
	config.kit.adapter = adapterNode();
}

export default config;
