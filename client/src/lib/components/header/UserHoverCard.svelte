<script lang="ts">
	import * as HoverCard from "$lib/components/ui/hover-card";
	import { Button } from "$lib/components/ui/button";
	import PersonIcon from "~icons/radix-icons/person";
	import GoogleIcon from "~icons/logos/google-icon";

	export let open: boolean;
	export let onOpenChange: (value: boolean) => void;

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
	<HoverCard.Content class="flex flex-col gap-4">
		<div>Not logged in</div>
		<div>
			<Button>
				<GoogleIcon font-size="24" /> &nbsp; Login with Google
			</Button>
		</div>
	</HoverCard.Content>
</HoverCard.Root>
