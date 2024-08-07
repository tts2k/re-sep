<script lang="ts" context="module">
	/* Margin */
	const PreviewMarginPresets = {
		left: ["pl-0", "pl-5p", "pl-10p", "pl-20p", "pl-30p", "pl-40p"],
		right: ["pr-0", "pr-5p", "pr-10p", "pr-20p", "pr-30p", "pr-40p"],
	} as const;
</script>

<script lang="ts">
	import {
		FontPreset,
		getFontSizeArray,
		FontSizeTag as Tag,
	} from "$lib/stylePresets";
	import { previewConfig } from "../store/previewConfig";

	export let scale: number;
	export let showBorder: boolean;

	$: fontSizePreset = getFontSizeArray($previewConfig.fontSize - 1);
	$: fontFamily = FontPreset[$previewConfig.font];
	$: justified = $previewConfig.justify ? "text-justify" : "";
	$: marginLeft =
		PreviewMarginPresets.left[($previewConfig.margin?.left || 1) - 1];
	$: marginRight =
		PreviewMarginPresets.right[($previewConfig.margin?.right || 1) - 1];

	$: scaleStyle = `transform: scale(${scale / 100});`;

	/*
	 * Border is mandatory for viewing margin on small scale so by default it's
	 * gonna be on a accent color.
	 * For preview, margin will be displayed as padding instead and use
	 * border as a way to preview margin size
	 */
	$: borderColor = showBorder ? "outline-foreground" : "outline-accent";
</script>

<article
	class="origin-center transition-transform ease-in-out duration-300
	{fontFamily} {justified} outline outline-4 {borderColor} {marginLeft}
	{marginRight}"
	style={scaleStyle}
>
	<h1 class={fontSizePreset[Tag.H1]}>Unix Philosophy</h1>

	<p class={fontSizePreset[Tag.TEXT]}>
		The Unix philosophy, originated by Ken Thompson, is a set of cultural
		norms and philosophical approaches to minimalist, modular software
		development. It is based on the experience of leading developers of the
		Unix operating system. Early Unix developers were important in bringing
		the concepts of modularity and reusability into software engineering
		practice, spawning a "software tools" movement. Over time, the leading
		developers of Unix (and programs that ran on it) established a set of
		cultural norms for developing software; these norms became as important
		and influential as the technology of Unix itself, and have been termed
		the "Unix philosophy."
	</p>

	<p class={fontSizePreset[Tag.TEXT]}>
		The Unix philosophy emphasizes building simple, compact, clear, modular,
		and extensible code that can be easily maintained and repurposed by
		developers other than its creators. The Unix philosophy favors
		composability as opposed to monolithic design.
	</p>

	<h2 class={fontSizePreset[Tag.H2]}>Origin</h2>

	<h3 class={fontSizePreset[Tag.H3]}>
		The Unix philosophy is documented by Doug McIlroy in the Bell System
		Technical Journal from 1978
	</h3>

	<h4 class={fontSizePreset[Tag.H4]}>
		1. Make each program do one thing well. To do a new job, build afresh
		rather than complicate old programs by adding new "features".
	</h4>

	<h4 class={fontSizePreset[Tag.H4]}>
		2. Expect the output of every program to become the input to another, as
		yet unknown, program. Don't clutter output with extraneous information.
		Avoid stringently columnar or binary input formats. Don't insist on
		interactive input
	</h4>

	<h4 class={fontSizePreset[Tag.H4]}>
		3. Design and build software, even operating systems, to be tried early,
		ideally within weeks. Don't hesitate to throw away the clumsy parts and
		rebuild them.
	</h4>

	<h4 class={fontSizePreset[Tag.H4]}>
		4. Use tools in preference to unskilled help to lighten a programming
		task, even if you have to detour to build the tools and expect to throw
		some of them out after you've finished using them.
	</h4>

	<p class={fontSizePreset[Tag.TEXT]}>
		It was later summarized by Peter H. Salus in
		<em>A Quarter-Century of Unix (1994)</em>
	</p>

	<ul class={fontSizePreset[Tag.TEXT]}>
		<li>Write programs that do one thing and do it well.</li>
		<li>Write programs to work together.</li>
		<li>
			Write programs to handle text streams, because that is a universal
			interface.
		</li>
	</ul>

	<p class={fontSizePreset[Tag.TEXT]}>
		In their Unix paper of 1974, Ritchie and Thompson quote the following
		design considerations:
	</p>

	<ul class={fontSizePreset[Tag.TEXT]}>
		<li>Make it easy to write, test, and run programs.</li>
		<li>Interactive use instead of batch processing.</li>
		<li>
			Economy and elegance of design due to size constraints ("salvation
			through suffering").
		</li>
		<li>
			Self-supporting system: all Unix software is maintained under Unix.
		</li>
	</ul>

	<h2 class={fontSizePreset[Tag.H2]}>The UNIX Programming Environment</h2>

	<p class={fontSizePreset[Tag.TEXT]}>
		In their preface to the 1984 book,
		<em> The UNIX Programming Environment </em>, Brian Kernighan and Rob
		Pike, both from Bell Labs, give a brief description of the Unix design
		and the Unix philosophy:
	</p>

	<blockquote class={fontSizePreset[Tag.TEXT]}>
		Even though the UNIX system introduces a number of innovative programs
		and techniques, no single program or idea makes it work well. Instead,
		what makes it effective is the approach to programming, a philosophy of
		using the computer. Although that philosophy can't be written down in a
		single sentence, at its heart is the idea that the power of a system
		comes more from the relationships among programs than from the programs
		themselves. Many UNIX programs do quite trivial things in isolation,
		but, combined with other programs, become general and useful tools.
	</blockquote>

	<p class={fontSizePreset[Tag.TEXT]}>
		The authors further write that their goal for this book is "to
		communicate the UNIX programming philosophy."
	</p>

	<footer class="text-right mt-5 font-sans italic">
		Source:
		<a
			class="hover:underline"
			href="https://en.wikipedia.org/wiki/Unix_philosophy"
			target="_blank"
		>
			Wikipedia page of UNIX Philosophy
		</a>
	</footer>
</article>

<!-- shadcn typography -->
<style lang="postcss">
	h1 {
		@apply scroll-m-20 font-extrabold tracking-tight text-center mb-32;
	}

	h2 {
		@apply scroll-m-20 border-b pb-2 font-semibold tracking-tight
		transition-colors mt-10;
	}

	h3 {
		@apply scroll-m-20 font-semibold tracking-tight mt-10;
	}

	h4 {
		@apply scroll-m-20 font-semibold tracking-tight mt-10;
	}

	p {
		@apply [&:not(:first-child)]:mt-6;
	}

	ul {
		@apply my-6 ml-6 list-disc;
	}

	li {
		@apply mt-2;
	}

	blockquote {
		@apply mt-6 border-l-2 pl-6 italic;
	}
</style>
