const colors = require("tailwindcss/colors")

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.{html,js,ts}"],
  plugins: [require("@tailwindcss/typography")],
}
