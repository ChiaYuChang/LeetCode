#! /bin/bash

datafile=$1
if [[ -z $1 ]]; then
    datafile="file.txt"
fi

declare -i n
n=$2
if [[ -z $2 ]]; then
    n=10
fi

echo '> Method 1 Use for loop'

function get_line_n() {
    # test input
    [[ ! -f $1 ]] && echo "file $1 not found" >&2 && return 1
    [[ $2 -lt 1 ]] && echo "n should be greater than 1" >&2 && return 1

    exec {f}<$1
    local -i i
    local IFSbk

    # reaf file
    IFSbk=$IFS
    while
        IFS=""
        read <&$f || [[ -n $REPLY ]]
    do
        let i++
        if [[ i -eq $2 ]]; then
            echo $REPLY
        fi
    done

    # cleanup
    IFS=$IFSbk
    exec {f}<&-
}
get_line_n $datafile $n

echo '> Method 2 Use awk'
awk 'NR == '${n}' {print $0}' $datafile

echo '> Method 3 Use sed'
sed -n $(printf "%dp" ${n}) $datafile
