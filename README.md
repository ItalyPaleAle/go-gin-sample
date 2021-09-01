# Go/Gin sample app

This is a sample web app written in Go and using the [Gin](https://github.com/gin-gonic/gin) framework. After launching it, you'll see a random Taylor Swift quote every time, as obtained from [taylor.rest](https://taylor.rest).

This app contains two parts:

1. A backend API server written in Go and using Gin.
2. A frontend static web app, built using Svelte 3 and Tailwind CSS, and bundled with Webpack.

The app demonstrates creating a Go app that contains both a RESTful API and a frontend application. It supports TLS.

To run this sample app, you'll need:

- Go 1.15 or higher
- Node.js (recommended version 14 or higher)

## Building the frontend

Before you can run the app, you need to build the frontend.

In the **frontend** directory, install the dependencies with

```sh
npm ci
```

You can then build the app with:

```sh
npm run build
```

You can also start a development server with:

```sh
npm run dev
```

## Running and building the app

After having built the frontend, you can run the Go application with:

```sh
go run .
```

You can build a self-contained binary by using the `build.sh` script. In addition to running `go build`, this also fetches and runs [pkger](https://github.com/markbates/pkger) to embed all static assets (ie. the compiled frontend app) into the binary.

```sh
./build.sh
```

## Config and environmental variables

The app expects a TLS certificate in the `certs` folder (relative to the current working directory). In particular:

- `certs/cert.pem` should contain the TLS certificate (PEM-encoded)
- `certs/key.pem` should contain the TLS key (PEM-encoded)

> This repository contains a self-signed certificate in the `certs` folder

You can tweak the app's behavior with these environmental variables:

- **`DATA_ENDPOINT`**: endpoint of the `taylor.rest` service (default: `https://api.taylor.rest/`)
- **`NO_TLS`**: set this to `1` (ie. `NO_TLS=1`) to disable TLS
- **`TLS_CERT`**: path to a TLS certificate (default: `certs/cert.pem`)
- **`TLS_KEY`**: path to a TLS key (default: `certs/key.pem`)
- **`BIND`**: address to bind to; the default value is `127.0.0.1` but you might need to change this to `0.0.0.0` to listen on all interfaces
- **`HTTP_PORT`**: port for the HTTP server, without TLS (default: `8080`)
- **`HTTPS_PORT`**: port for the HTTPS server, with TLS (default: `8443`)

HELLO WORLD