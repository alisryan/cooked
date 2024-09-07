import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: "#D9D9D9",
        primaryHover: "#C9C9C9",
        secondary: "#393939",
        secondaryHover: "#292929",
        tertiary: "#000000"
      }
    },
  },
  plugins: [],
};
export default config;
