import type { Action } from "svelte/action";
import type { Writable } from "svelte/store";

type ActionProps = {
	store: Writable<string>;
};

export const toc: Action<HTMLElement, ActionProps> = (
	container: HTMLElement,
	props: ActionProps,
) => {
	if (typeof IntersectionObserver === "undefined") {
		return;
	}

	const options = {
		rootMargin: "-50% 0px",
		threshold: 0,
	};

	const onObserver = (entries: IntersectionObserverEntry[]) => {
		for (const entry of entries) {
			if (entry.isIntersecting) {
				props.store.set(entry.target.id);
			}
		}
	};

	const observer = new IntersectionObserver(onObserver, options);

	// This should be pre-proccessed by the golang backend with goquery after scraping
	const headingLinks = container.querySelectorAll(
		"h1 > a, h2 > a, h3 > a, h4 > a, h5 > a, h6 > a",
	);

	for (const link of headingLinks) {
		observer.observe(link);
	}

	return {
		destroy() {
			observer.disconnect();
			props.store.set("");
		},
	};
};
