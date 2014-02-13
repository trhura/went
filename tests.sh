#!/bin/bash

source "./assert.sh"

##############################
### Build go binaries and setup
##############################
go build went.go
WENT=$(pwd)/went

HOME="/tmp"
rm -f "$HOME/.went.recentf"
TESTROOT=$(pwd)/wenttmp ##### WILL BE DELETED #####

##############################
### Test going up using `..[...]` notation
##############################
LEAFDIR=$TESTROOT/4/3/2/1
mkdir -p $LEAFDIR

builtin cd $LEAFDIR
assert "basename $($WENT ..)" "2"
builtin cd $LEAFDIR
assert "basename $($WENT ...)" "3"
builtin cd $LEAFDIR
assert "basename $($WENT ....)" "4"

########################################
### Test keeping track and revisiting recent dirs
########################################
cd $TESTROOT

SAMEDIR="same"
SAMEDIR_1=$TESTROOT/1/$SAMEDIR
SAMEDIR_2=$SAMEDIR_1/2/$SAMEDIR
SAMEDIR_3=$SAMEDIR_2/3/$SAMEDIR
mkdir -p $SAMEDIR_3

assert "$WENT $SAMEDIR_1" "$SAMEDIR_1" # absolute path
cd $HOME && assert "$WENT $SAMEDIR" "$SAMEDIR_1" # last dir

assert "$WENT $SAMEDIR_2" "$SAMEDIR_2" # absolute path
 # if `same` exists in cwd, go to that instead of last visited
cd $TESTROOT/1 && assert "$WENT $SAMEDIR" "$SAMEDIR_1"

assert "$WENT $SAMEDIR_2" "$SAMEDIR_2" # absolute path
cd $TESTROOT && assert "$WENT $SAMEDIR" "$SAMEDIR_2"

assert "$WENT $SAMEDIR_3" "$SAMEDIR_3" # absolute path
 # if `same` exists in cwd, go to that instead of last visited
cd $SAMEDIR_1/2 && assert "$WENT $SAMEDIR" "$SAMEDIR_2"

assert "$WENT $SAMEDIR_3" "$SAMEDIR_3" # absolute path
 # if `same` exists in cwd, go to that instead of last visited
cd $SAMEDIR_1 && assert "$WENT $SAMEDIR" "$SAMEDIR_3"

########################################
### Circling
########################################
assert "$WENT $SAMEDIR_3" "$SAMEDIR_3"
assert "$WENT $SAMEDIR_2" "$SAMEDIR_2"
assert "$WENT $SAMEDIR_1" "$SAMEDIR_1"

cd $HOME && assert "$WENT $SAMEDIR" "$SAMEDIR_1"
cd $SAMEDIR_1 && assert "$WENT ." "$SAMEDIR_2"
cd $SAMEDIR_2 && assert "$WENT ." "$SAMEDIR_3"
cd $SAMEDIR_3 && assert "$WENT ." "$SAMEDIR_1"

########################################
### Invalid/Deleted Files
########################################
assert "$WENT $SAMEDIR_1" "$SAMEDIR_1"
assert "$WENT $SAMEDIR_2" "$SAMEDIR_2"
assert "$WENT $SAMEDIR_3" "$SAMEDIR_3"
rm -r $SAMEDIR_3

cd $TESTROOT && assert "$WENT $SAMEDIR" "$SAMEDIR_2"
cd $SAMEDIR_2 && assert "$WENT ." "$SAMEDIR_1"
cd $SAMEDIR_1 && assert "$WENT ." "$SAMEDIR_2"
rm -r $SAMEDIR_2

cd $TESTROOT && assert "$WENT $SAMEDIR" "$SAMEDIR_1"
cd $SAMEDIR_1 && assert "$WENT ." "$SAMEDIR_1"

# Cleanup
rm -rf $TESTROOT

# THE END
assert_end
