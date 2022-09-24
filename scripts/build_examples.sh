#!/bin/bash

# loop through the examples dir and build them into wasm files
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/aabb.wasm ./examples/aabb/
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/font.wasm ./examples/font/
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/tiles.wasm ./examples/tiles/
cp .dist/examples/* .dist/web/
cp .dist/examples/* .dist/web/examples/

cp $(go env GOROOT)/misc/wasm/wasm_exec.js .dist/web/
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .dist/web/examples/
