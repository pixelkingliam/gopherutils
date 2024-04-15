#!/bin/bash

for file in source/*.go; do
    echo -e "[BUILD.SH]=> Compiling ${file}"

    filename=$(basename -- "$file")
    filename_no_ext="${filename%.*}"
    go build -o "bin/$filename_no_ext" "$file"
done
echo -e "[BUILD.SH]=> Finished Compiling."