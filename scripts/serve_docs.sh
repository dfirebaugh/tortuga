#!/bin/bash

source ./scripts/build_docs.sh

function cleanBuild() {
    source ./scripts/build_examples.sh
    source ./scripts/adapt_js_module.sh
}

# this is a little wonky,
#   we need to do stuff to our build after it's built, but 
#   we want the server to be the last thing called
#   so that we don't have stray processes running after we ctrl+c
cleanBuild & bin/mdbook serve -p 8888 docs/tortuga --dest-dir ../../.dist/web --open
