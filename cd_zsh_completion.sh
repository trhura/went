#compdef cd

typeset -A opt_args

_cd() {
    local -a cmd recents
    recents=( $(cut -d, -f1 $HOME/.went.recentf | tr "\n" " ") )
    compadd "$@" -a recents
    _files
}

_cd "$@"
