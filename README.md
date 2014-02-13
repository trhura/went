## About

`went` is a small `go` program, that let you go to recently visited
directories by its basename without using its full/ path, when wrapped
around the shell bulitin `cd`.

## Usage

TODO

## Installation

Assuming you have installed golang, and configured `$GOPATH`.

+ Get sources from github.
```sh
go get github.com/trhura/went
```

+ Compile.
```sh
cd $GOPATH/src/github.com/trhura/went && go build went.go
```

+ Wrap the shell (Append this in your `.bashrc` or `.zshrc`)
```bash
function went {
        builtin cd $($GOPATH/src/github.com/trhura/went/went $@)
}

alias cd=went
```
