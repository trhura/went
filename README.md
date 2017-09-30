## Name

went - a smarter cd for terminal addicts

## About

Ever got tired of typing full directory paths to navigate in terminal?
No more! `went` is a small `python` script, which takes you to your
recently visited directories by merely typing directory name.

The beauty of `went` over similar tools is its simplicity and ease of
use. It is a simple wrapper around built-in `cd` command. 

So, there is no additional command nor flags to remember, and actually,
most of the time, you wouldn't even notice it is there. It just silenty
keep track of where you have been and let's you revisit those places
quickly.

## Usage

Very simple. Just use `cd dirname` to go to last visited directory
with that name. If there is more than one visited path with the same
name, use `cd .` to iterate through those paths.

![Usage](doc/usage.png)

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
