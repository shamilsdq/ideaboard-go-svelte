import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";
import * as path from "path";
import dotenv from "dotenv";

dotenv.config({ path: "../.env" });

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  resolve: {
    alias: { "@": path.resolve(__dirname, "./src") },
  },
  define: { "process.env": process.env },
});
