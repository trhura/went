#!/bin/bash

source "./assert.sh"

# Build binary
go build went.go

# use tmp directory for recentf file
HOME="/tmp"
rm -f "$HOME/.went.recentf"

assert "./went $PWD" "$PWD"
assert "./went $(basename $PWD)" "$PWD"



assert_end
