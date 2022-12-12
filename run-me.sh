#!/usr/bin/env bash
set -euo pipefail
clear
sleep 5s
figlet A Niko Original | lolcat
figlet TAS-Like Speedrun
figlet Cookie Clicker Any % | lolcat
sleep 3s
clear
figlet 3 | lolcat
sleep 0.75s
clear
figlet 2 | lolcat
sleep 0.75s
clear
figlet 1 | lolcat
sleep 0.75s
clear
if [[ `uname -a | grep -i linux | grep -i microsoft`  != "" ]]; then 
    echo "hacking for wsl (on every other system, including vanilla windows, this \"just works\")"
    set -x
    export PATH="${PATH}:/mnt/c/Program\ Files/Google/Chrome/Application/"
    export GOOS=windows
fi
go run .