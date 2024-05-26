import type { Config } from "tailwindcss";

export default {
  content: [
    "./**/*.templ",
    './node_modules/flowbite/**/*.js',
  ],
  darkMode: 'class',
  plugins: [
    require('flowbite/plugin'),
  ],
} satisfies Config;
