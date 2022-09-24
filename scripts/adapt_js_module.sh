#!/bin/bash

# mdbook builds all `additional-js` to `type="text/javascript"`
#   we want to add `wasm-view.js` as a module
#   this might be worth submitting a PR to mdbook repo to add support for js modules
function wasmJSModuleAdapter() {
    cd .dist/web/
    find . -type f \( -iname "*.html" -o -iname "*.html" \) -exec sed -i 's#<script type="text/javascript" src="wasm-view.js"></script>#<script type="module" src="wasm-view.js"></script>#' '{}' +
    find . -type f \( -iname "*.html" -o -iname "*.html" \) -exec sed -i 's#<script type="text/javascript" src="../wasm-view.js"></script>#<script type="module" src="../wasm-view.js"></script>#' '{}' +
    cd ../..
}

wasmJSModuleAdapter
