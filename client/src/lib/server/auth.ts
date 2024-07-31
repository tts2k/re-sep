import { SvelteKitAuth } from "@auth/sveltekit";
import Google from "@auth/sveltekit/providers/google";

// TODO
// import Mastodon from "@auth/sveltekit/providers/mastodon";

export const { handle, signIn, signOut } = SvelteKitAuth({
	providers: [Google],
});
