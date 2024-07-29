#!/bin/bash

# Directory to store hashes
HASH_DIR="/tmp/go_build_hashes"
mkdir -p "$HASH_DIR"

# Function to calculate MD5 hash of a file
calculate_md5() {
    ./bin/md5sum "$1" | awk '{ print $1 }'
}

for file in source/*.go; do
    filename=$(basename -- "$file")
    filename_no_ext="${filename%.*}"
    hash_file="$HASH_DIR/$filename.md5"

    # Calculate current hash
    current_hash=$(calculate_md5 "$file")

    # Check if hash file exists
    if [ -f "$hash_file" ]; then
        # Read stored hash
        stored_hash=$(cat "$hash_file")

        # Compare hashes
        if [ "$current_hash" == "$stored_hash" ]; then
            echo -e "[BUILD.SH]=> Skipping ${file}, no changes"
            continue
        fi
    fi

    # Compile the file
    echo -e "[BUILD.SH]=> Compiling ${file}"
    go build -o "bin/$filename_no_ext" "$file"

    # Store the new hash
    echo "$current_hash" > "$hash_file"
done

echo -e "[BUILD.SH]=> Finished Compiling."
