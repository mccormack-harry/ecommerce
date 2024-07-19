const defaultTheme = require('tailwindcss/defaultTheme')

/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        'internal/templ/**/*.templ',
    ],
    corePlugins: {
        preflight: true,
    },
    theme: {
        extend: {
            fontFamily: {
                'sans': ['Inter', ...defaultTheme.fontFamily.sans],
            }
        }
    }
}