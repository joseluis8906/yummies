/** @type {import('tailwindcss').Config} */
const { fontFamily } = require('tailwindcss/defaultTheme');
module.exports = {
    darkMode: 'media',
    content: ["./src/**/*.{html, js}"],
    theme: {
        extend: {
            fontFamily: {
                sans: ['Roboto', ...fontFamily.sans],
            },
        },
    },
    plugins: [
        require("@catppuccin/tailwindcss")({
            // prefix to use, e.g. `text-pink` becomes `text-ctp-pink`.
            // default is `false`, which means no prefix
            prefix: "ctp",
            // which flavour of colours to use by default, in the `:root`
            defaultFlavour: "mocha",
        }),
        require("@tailwindcss/typography"),
    ],
}

