#!/bin/bash

# loop through the examples dir and build them into wasm files
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/simple.wasm ./examples/simple/
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/aabb.wasm ./examples/aabb/
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/font.wasm ./examples/font/
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/tiles.wasm ./examples/tiles/
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/sweptaabb.wasm ./examples/sweptaabb/
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/sweptaabb2.wasm ./examples/sweptaabb2/
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/staticball.wasm ./examples/staticball/
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/dynamicball.wasm ./examples/dynamicball/
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/paddleball.wasm ./examples/paddleball/
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/sfx.wasm ./examples/sfx/
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/sequence.wasm ./examples/sequence/
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/tone.wasm ./examples/tone/
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/gameoflife.wasm ./examples/gameoflife/
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/texture.wasm ./examples/texture/
env GOOS=js GOARCH=wasm go build -o ./.dist/examples/spectrum.wasm ./examples/spectrum/

cp .dist/examples/* .dist/web/
cp .dist/examples/* .dist/web/examples/

cp $(go env GOROOT)/misc/wasm/wasm_exec.js .dist/web/
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .dist/web/examples/
