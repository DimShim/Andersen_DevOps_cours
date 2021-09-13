#! /bin/bash


# define "help" message
help=$'
usage: ./fish.sh [PROCESS NAME or PID] [LINES LIMIT]

Shows the names of the organizations with which the PROCESS has
established a connection.

Examples:
./fish.sh firefox 5
'


# argument - help
if [[ $1 == '-h' || $1 == '--help' ]]; then
  # the only argument is 'help'
  echo "$help"
  exit 0
fi

# input isn't correct, description 
if [[ -z "$1" || -z "$2" ]]
then
  echo "Please enter process name or PID also lines limit!"
  echo "You can use the HELP! (-h or --help)."
  exit
fi

# frame
separator_line="+-------------------------------------------------------+"
echo $separator_line
printf "%-55s | \n" \
  "|     Organization name with established connection"
echo $separator_line


# script form task
ss -tunap | awk -v nameOrPID=$1 '$0~nameOrPID {print $6}' | cut -d: -f1 | 
sort | uniq -c | sort | tail -n$2 | grep -oP '(\d+\.){3}\d+' | 
while read IP ; do whois $IP | awk -F':' '/^Organization/ {print $2}' ; done

echo $separator_line
