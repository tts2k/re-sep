import adapter from "@sveltejs/adapter-auto";
import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://kit.svelte.dev/docs/integrations#preprocessors
	// for more information about preprocessors
	preprocess: vitePreprocess(),

	kit: {
		// Auto adapter is temporary. Mostly for demo-ing the FE
		// The final product is probably gonna run on the Node adapter
		adapter: adapter(),
		alias: {
			"@/*": "./src/lib/*",
		},
	},
};

export default config;
