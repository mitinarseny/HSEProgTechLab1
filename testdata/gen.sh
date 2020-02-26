#!/bin/sh

i=1
while [ "${i}" -le 2 ]; do
  for n in 100 1000 10000 100000; do
    echo "[G] ${n}_${i}.json"
    jg student.yaml \
    -a "${n}" \
    -f names=names.txt,faculties=faculties.txt,specialities=specialities.txt \
    -o "${n}_${i}.json"
  done
  i=$((i + 1))
done