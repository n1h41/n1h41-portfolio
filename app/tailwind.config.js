/** @type {import('tailwindcss').Config} */
const withMT = require("@material-tailwind/html/utils/withMT");

module.exports = withMT({
  content: ["./views/**/*.{html,js,templ}"],
  theme: {
    extend: {},
  },
  plugins: [],
});
