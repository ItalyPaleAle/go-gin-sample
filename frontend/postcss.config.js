const prod = process.env.NODE_ENV == 'production'

module.exports = {
    plugins: [
        require('postcss-import')(),
        require('postcss-url')(),
        require('tailwindcss'),
        ...(prod ? [require('@fullhuman/postcss-purgecss')({

            // Specify the paths to all of the template files in your project
            content: [
                './src/**/*.svelte',
                './src/**/*.html',
                './src/**/*.js',
            ],

            // Include any special characters you're using in this regular expression
            defaultExtractor: content => content.match(/[\w-/:]+(?<!:)/g) || []
        })] : []),
        ...(prod ? [require('cssnano')] : []),
    ],
}