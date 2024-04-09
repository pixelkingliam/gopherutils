#!/bin/bash

for file in source/*.go; do
    filename=$(basename -- "$file")
    filename_no_ext="${filename%.*}"
    go build -o "bin/$filename_no_ext" "$file"
done