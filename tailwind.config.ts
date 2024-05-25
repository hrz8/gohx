import type { Config } from "tailwindcss";

export default {
  content: [
    "./**/*.templ",
    './node_modules/flowbite/**/*.js',
  ],
  plugins: [
    require('flowbite/plugin'),
  ],
} satisfies Config;
