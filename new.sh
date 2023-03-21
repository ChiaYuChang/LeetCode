#! /bin/bash

[[ $# -ne 1 ]] && echo "Usage: $0 question string" && exit 1

filename=$1

declare pkgname='leetcode'

OriOIFS=$IFS
IFS='.'
EXT=.go

# split file name
af=($filename) && IFS=$OriOIFS
declare -i qNmbr=${af[0]}
declare qName=${af[1]}

# format string
declare fmtQNmbr=$(printf '%04d' ${qNmbr})
declare fmtQName=${qName// /_}

# new file
touch ${fmtQNmbr}${fmtQName}${EXT}
touch ${fmtQNmbr}${fmtQName}_test${EXT}

echo "package ${pkgname}" >${fmtQNmbr}${fmtQName}${EXT}
echo "package ${pkgname}_test" >${fmtQNmbr}${fmtQName}_test${EXT}
