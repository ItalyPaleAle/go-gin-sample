const MiniCssExtractPlugin = require('mini-css-extract-plugin')
const {CleanWebpackPlugin} = require('clean-webpack-plugin')
const HtmlWebpackPlugin = require('html-webpack-plugin')
const path = require('path')

const mode = process.env.NODE_ENV || 'development'
const prod = mode === 'production'

module.exports = {
	entry: {
		bundle: ['./src/main.js']
	},
	resolve: {
		alias: {
			svelte: path.resolve('node_modules', 'svelte')
		},
		extensions: ['.mjs', '.js', '.svelte'],
		mainFields: ['svelte', 'browser', 'style', 'module', 'main']
	},
	output: {
		path: __dirname + '/public',
		publicPath: '/',
		filename: prod ? '[name].[contenthash:8].js' : '[name].js',
        chunkFilename: prod ? '[name].[contenthash:8].js' : '[name].js',
        crossOriginLoading: 'anonymous'
	},
	module: {
		rules: [
			{
                test: /\.(svelte)$/,
                exclude: [],
                use: {
                    loader: 'svelte-loader',
                    options: {
                        hotReload: true,
                        dev: !prod,
                    }
                }
            },
			{
				test: /\.css$/,
				use: [
					/**
					 * MiniCssExtractPlugin doesn't support HMR.
					 * For developing, use 'style-loader' instead.
					 * */
					prod ? MiniCssExtractPlugin.loader : 'style-loader',
					{loader: 'css-loader', options: {importLoaders: 1}},
                    'postcss-loader'
				]
			}
		]
	},
	mode,
	plugins: [
        new CleanWebpackPlugin({
            cleanOnceBeforeBuildPatterns: ['**/*', '!assets', '!assets/*']
        }),

        new MiniCssExtractPlugin({
            filename: '[name].[contenthash:8].css'
		}),
		
		new HtmlWebpackPlugin({
			template: 'src/index.html',
			minify: prod ? {
				collapseBooleanAttributes: true,
				collapseWhitespace: true,
				conservativeCollapse: true,
				decodeEntities: true,
				html5: true,
				keepClosingSlash: false,
				processConditionalComments: true,
				removeComments: true,
				removeEmptyAttributes: true,
			} : false,
		})
	],
	devServer: {
        contentBase: path.join(__dirname, 'public'),
        port: 3000
    },
    devtool: prod ? false : 'source-map'
}
