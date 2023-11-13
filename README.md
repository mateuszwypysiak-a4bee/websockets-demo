# websockets-demo

## Server

1. Install GO: https://go.dev/doc/install
2. run `go run pkg/main/main.go`

## React demo

React demo is located in `/static/react`. It's normal Vite project. Can be served using Vite commands, or by GO server as static files.

To serve via GO servet after Vite build remove forward `/` from script and style links `index.html`.

## Available clients

- [native client](http://localhost:8080/static/native.html)
- [react client](http://localhost:8080/static/react/dist/)
