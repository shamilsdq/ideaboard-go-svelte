import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";
import * as path from "path";

export default {
  // Consult https://svelte.dev/docs#compile-time-svelte-preprocess
  // for more information about preprocessors
  preprocess: vitePreprocess(),
  alias: {
    "@": path.resolve("./src"),
  },
};
