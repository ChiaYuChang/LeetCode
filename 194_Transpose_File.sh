#! /bin/bash

declare -i ncol
ncol=$(head -1 file.txt | awk '{print NF}')

for i in $(seq 1 ${ncol}); do
    l=$(awk -v col="${i}" '{ print $col }' file.txt)
    echo ${l[*]}
done

unset ncol
unset l
