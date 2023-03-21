#! /bin/bash

filename="words.txt"

function count_words_freq() {
	[[ ! -f $1 ]] && echo "file $1 not found" >&2 && return 1

	exec {f}<$1
	local -A count

	while read -r <&$f || [[ -n $REPLY ]]; do
		for word in $REPLY; do
			if [[ -n ${count[$word]} ]]; then
				count[$word]=$((${count[$word]} + 1))
			else
				count[$word]=1
			fi
		done
	done
	exec {f}<&-

	for key in ${!count[@]}; do
		printf "%s %d\n" $key ${count[$key]}
	done | sort -r -nk 2
}

count_words_freq $filename
