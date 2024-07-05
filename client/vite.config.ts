import { sveltekit } from "@sveltejs/kit/vite";
import { defineConfig } from "vite";
import Icons from "unplugin-icons/vite";
import ConditionalCompile from "vite-plugin-conditional-compiler";

export default defineConfig({
	plugins: [
		sveltekit(),
		ConditionalCompile(),
		Icons({
			compiler: "svelte",
		}),
	],
});
