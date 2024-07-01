// #v-ifdef DEMO
// @ts-ignore
import service from "./demo";
// #v-else
// @ts-ignore
// biome-ignore lint/suspicious/noRedeclare: conditional compile
import service from "./prod";
// #v-endif

export default service;
