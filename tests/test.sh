#!/bin/zsh
# todo support any SHELL
# todo use tmp file for that
nasm -o0 -f bin $1 && hexdump `basename $1 .asm`