#! /usr/bin/env python
# Author: Thura Hlaing <trhura@gmail.com>
# Time-stamp: <2017-11-17 10:41:04 (thurahlaing)>

__author__ = "Thura Hlaing <trhura@gmail.com>"
__copyright__ = "Copyright 2017, Planet Earth"

import os
import sys
import pickle
import os.path

class RecentlyVisitedDirectories:
    pathsbydir = {}
    CONFIG_FILE = os.path.expandvars("$HOME/.went.directories")
    
    def get_recently_visited (self, dirname):
        if not dirname in self.pathsbydir: return ""
        
        path = self.pathsbydir[dirname][0]
        if os.path.exists(path):
            return path
        
        del self.pathsbydir[dirname][0]
        return self.get_recently_visited(dirname)

    def get_next_recently_visited (self, dirname):
        if not dirname in self.pathsbydir: return ""

        # shift the list by putting index 0 to the end
        self.pathsbydir[dirname] = self.pathsbydir[dirname][1:] + [self.pathsbydir[dirname][0]]
        return self.get_recently_visited(dirname)
    
    def add_recently_visited (self, path):
        if not(os.path.exists(path) and os.path.isdir(path)): return

        dirname = os.path.basename(path)
        if not dirname in self.pathsbydir:
            self.pathsbydir[dirname] = [path]
        else:
            self.pathsbydir[dirname] = [path] + [p for p in self.pathsbydir[dirname] if p != path]
    
    def __init__(self):
        try:
            with open(self.CONFIG_FILE, 'rb') as configFile:
                self.pathsbydir = pickle.load (configFile)
        except: pass

    def __save__(self):
        for k, v in self.pathsbydir.items():
            if not v: del self.pathsbydir[k] # if empty paths, remove
            self.pathsbydir[k] = v[:3] # only save the firs three paths
            
        with open(self.CONFIG_FILE, 'wb') as configFile:
            pickle.dump (self.pathsbydir, configFile)

WORKING_DIRECTORY = os.getcwd()
RECENT_DIRECTORIES = RecentlyVisitedDirectories() 

def go_home_if_nowhere(destination):
    """ Go to home directory if no argument was passed. """
    return os.path.expandvars("$HOME") if not destination else ""

def go_next_place_if_dot(destination):
    """ Go to recently visited directory with same name if a single dot"""
    if destination == ".":
        return RECENT_DIRECTORIES.get_next_recently_visited(os.path.basename(WORKING_DIRECTORY))

def go_parents_if_dots(destination):
    """ Go to parent directories for .. and ..."""
    if not all(c == '.' for c in destination): return ""

    path = WORKING_DIRECTORY
    for c in range(1, len(destination)): path = os.path.dirname(path)
    return path

def go_there_if_path (destination):
    """ Go to specified path for absolute and relative paths (aka contains slash). """
    if os.path.isabs(destination): return destination
    
    abspath = os.path.abspath(destination)
    return abspath if os.path.exists(abspath) else ""

def go_to_last_visited_place(destination):
    return RECENT_DIRECTORIES.get_recently_visited(destination)

DIRECTIONS = [
    go_home_if_nowhere,
    go_next_place_if_dot,
    go_parents_if_dots,
    go_there_if_path,
    go_to_last_visited_place,
]

def go(path):
    RECENT_DIRECTORIES.add_recently_visited(path)
    print(path)
    
def determine_wheretogo(destination):
    for direction in DIRECTIONS:
        path = direction(destination)
        if path: go(path); break
    else:
        # Ok, we don't know where to go, but `cd` may know
        go(destination) 
        
def main():
    destination = sys.argv[1] if len(sys.argv) > 1 else None
    determine_wheretogo(destination)
    RECENT_DIRECTORIES.__save__()

if __name__ == "__main__": main() 
