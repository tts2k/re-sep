// @ts-nocheck

// #v-ifdef PLATFORM_SELF
import service from "./prod";
// #v-endif

// #v-ifdef PLATFORM_VERCEL
// biome-ignore lint/suspicious/noRedeclare: conditional compile
import service from "./vercel";
// #v-endif

// #v-ifdef PLATFORM_DEMO
// biome-ignore lint/suspicious/noRedeclare: conditional compile
import service from "./demo";
// #v-endif

export default service;
