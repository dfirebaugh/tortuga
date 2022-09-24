#!/bin/bash

# loop through the examples dir and build them into wasm files
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/aabb.wasm ./examples/aabb/
cp .dist/examples/* .dist/web/
