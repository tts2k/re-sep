<script lang="ts">
	import * as HoverCard from "$lib/components/ui/hover-card";
	import { Button } from "$lib/components/ui/button";
	import PersonIcon from "~icons/radix-icons/person";
	import GoogleIcon from "~icons/logos/google-icon";
	//import MastodonIcon from "~icons/logos/mastodon-icon";
	import { signIn, signOut } from "@auth/sveltekit/client";
	import { page } from "$app/stores";
	import { toast } from "svelte-sonner";

	export let open: boolean;
	export let onOpenChange: (value: boolean) => void;
	export let loading = false;

	const signOutWrapper = () => {
		toast.promise(signOut());
	};

	$: triggerOpacity = open ? "opacity-100" : "opacity-0";
</script>

<HoverCard.Root bind:open bind:onOpenChange openDelay={300}>
	<HoverCard.Trigger>
		<div
			class="{triggerOpacity} hover:opacity-100 transition-opacity
			ease-in-out duration-300"
		>
			<Button variant="ghost" size="icon">
				<PersonIcon font-size="24" />
			</Button>
		</div>
	</HoverCard.Trigger>
	{#if $page.data.session}
		<HoverCard.Content class="flex flex-col gap-4">
			<div>Logged in as {$page.data.session.user?.name}</div>
			<div>
				<Button on:click={signOutWrapper} {loading}>Log out</Button>
			</div>
		</HoverCard.Content>
	{:else}
		<HoverCard.Content class="flex flex-col gap-4">
			<div>Not logged in</div>
			<Button on:click={() => signIn("google")} {loading}>
				<GoogleIcon font-size="24" /> &nbsp; Login with Google
			</Button>
		</HoverCard.Content>
	{/if}
</HoverCard.Root>
