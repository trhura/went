## About

`went` is a small go program, that let you go to recently visited
directories by its basename without using its full/ path, when wrapped
around the shell bulitin cd.

+ zero-config
+

## Usage

TODO

## Installation

Assuming you have installed golang, and configured `$GOPATH`.

1. Get sources from github.
```shell
go get github.com/trhura/went
```
1. Compile.
```shell
cd $GOPATH/src/github.com/trhura/went && go build went.go
```
1. Wrap the shell (Append this in your `.bashrc` or `.zshrc`)
```bash
function went {
        builtin cd $($GOPATH/src/github.com/trhura/went/went $@)
}

alias cd=went
```
