########################################
## Auto completeion for went
########################################

_cd()
{
    local cur
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"

    DIRNAMES=$(cut -d, -f1 $HOME/.went.recentf | tr "\n" " ")
    COMPREPLY=( $(compgen -W "${DIRNAMES}" -- ${cur}) )
    return 0
}

complete -o default -F _cd cd
