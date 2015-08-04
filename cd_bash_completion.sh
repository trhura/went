########################################
## Auto completeion for went
########################################

_cd()
{
    local cur
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
	
    WENT_COMPLS=$(cut -d, -f1 $HOME/.went.recentf | tr "\n" " ")
	DIR_COMPLS=$(ls -dF "$cur"* 2> /dev/null)
	COMPLS=$(echo "$DIR_COMPLS" "$WENT_COMPLS")
    COMPREPLY=($(compgen -W "${COMPLS}" -- ${cur}))
    return 0
}

complete -o dirnames -o nospace -F _cd cd
