#! /bin/bash


# define "help" message
help=$'Shows the names of the organizations with which the PROCESS has
established a connection.'


# there is only one argument
  if [[ $1 == '-h' || $1 == '--help' ]]; then
    # the only argument is 'help'
    echo "$help"
    exit 0
  fi


# final output
separator_line="+-------------------------------------------------------+"
echo $separator_line
printf "%-55s | \n" \
  "|     Organization name with established connection"
echo $separator_line


# script form task
ss -tunap | awk '/firefox/ {print $6}' | cut -d: -f1 | 
sort | uniq -c | sort | tail -n5 | grep -oP '(\d+\.){3}\d+' | 
while read IP ; do whois $IP | awk -F':' '/^Organization/ {print $2}' ; done

echo $separator_line
