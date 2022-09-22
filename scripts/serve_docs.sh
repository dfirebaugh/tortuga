#!/bin/bash

source ./scripts/download_mdbook.sh

bin/mdbook serve -p 8888 docs/tortuga --open
