<script lang="ts">
	import * as HoverCard from "$lib/components/ui/hover-card";
	import { Button } from "$lib/components/ui/button";
	import PersonIcon from "~icons/radix-icons/person";
	import GoogleIcon from "~icons/logos/google-icon";
	import { toast } from "svelte-sonner";
	import { user, login } from "@/stores/user";
	export let open: boolean;
	export let onOpenChange: (value: boolean) => void;
	export let loading = false;

	$: triggerOpacity = open ? "opacity-100" : "opacity-0";

	const onLogin = async (provider: string) => {
		try {
			loading = true;
			await login(provider);
		} catch (err) {
			console.error(err);
			toast.error("Error: Server is not running");
			loading = false;
		}
	};

	const onLogout = async () => {};
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
	{#if $user.loggedIn}
		<HoverCard.Content class="flex flex-col gap-4">
			<div>Logged in as {$user.name}</div>
			<div>
				<Button on:click={onLogout} {loading}>Log out</Button>
			</div>
		</HoverCard.Content>
	{:else}
		<HoverCard.Content class="flex flex-col gap-4">
			<div>Not logged in</div>
			<div>
				<Button on:click={() => onLogin("google")} {loading}>
					<GoogleIcon font-size="24" /> &nbsp; Login with Google
				</Button>
			</div>
		</HoverCard.Content>
	{/if}
</HoverCard.Root>
