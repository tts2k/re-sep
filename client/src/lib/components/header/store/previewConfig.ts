import type { UserConfig } from "@/proto/user_config";
import { writable } from "svelte/store";
import { defaultConfig } from "@/defaultConfig";

export const previewConfig = writable<UserConfig>(defaultConfig);
