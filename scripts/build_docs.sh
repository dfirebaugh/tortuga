#!/bin/bash

source ./scripts/download_mdbook.sh
mkdir -p .dist/web

bin/mdbook build docs/tortuga --dest-dir ../../.dist/web

source ./scripts/build_examples.sh
source ./scripts/adapt_js_module.sh
