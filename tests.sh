#!/bin/bash

source "./assert.sh"

# Build go binary
go build went.go
WENT=$(pwd)/went

# use tmp directory to store recently visited db for testing
HOME="/tmp"
rm -f "$HOME/.went.recentf"

# It should
assert "$WENT $PWD" "$PWD"
assert "$WENT $(basename $PWD)" "$PWD"

# Test going parent directories using `...`
mkdir -p /tmp/testwent/4/3/2/1
LEAFCHILD=/tmp/testwent/4/3/2/1
builtin cd $LEAFCHILD
assert "basename $($WENT ..)" "2"
builtin cd $LEAFCHILD
assert "basename $($WENT ...)" "3"
builtin cd $LEAFCHILD
assert "basename $($WENT ....)" "4"
rm -rf /tmp/testwent

assert_end
