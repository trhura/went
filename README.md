## About

`went` is a small `python` script, that let you go to recently visited
directories by its basename without using its full path, when wrapped
around the shell bulitin `cd`.

Although there are similar and more feature-savy tools, like
[autojump] (https://github.com/joelthelion/autojump), `went` differs
from other tools in its simplicity and unobtrusiveness.

Being a simple wrapper around `cd`, there is no flag/option/setting to
remember, and `want` tries not to interfere with normal `cd` usages.
Actually, most of the time, you wouldn't even notice it is there. It
just silenty keep track of where you have been and let's you revisit
those places quickly.

## Usage

![Usage](doc/usage.png)

Very simple. Just use `cd dirname` to go to last visited directory
with that name. If there is more than one visited path with the same
name, use `cd .` to iterate through those paths.

A few other useful shortcuts, like using `cd ..[...]` to go up parent
directories, are also available.

## Installation

Simpy put went.py under any of your $PATH and make it executable.
```sh
curl https://raw.githubusercontent.com/trhura/went/master/went.py -o /usr/bin/went.py && chmod +x /usr/bin/went.py
```

+ Wrap the shell builtin shell `cd`. (Add this in your `.bashrc` or `.zshrc`)
```bash
function went {
        builtin cd "$(/usr/bin/went.py $@)"
}

alias cd=went
```

## Troubleshooting

+ If case you run into some bugs, removing history configuration file should resolve most issues. 

```bash
rm -f $HOME/.went.directories
```
