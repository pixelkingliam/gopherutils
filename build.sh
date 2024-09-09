#!/bin/bash

# Directory to store hashes
HASH_DIR="/tmp/go_build_hashes"
BUILD_DIR="/tmp/go_build"
mkdir -p "$HASH_DIR"
mkdir -p "$BUILD_DIR"
VERSION_NUMBER="1.0"
GIT_HASH=$(git rev-parse --short HEAD)
# Flag for force rebuild
FORCE_REBUILD=false
RELEASE=false
FINAL_VERSION="??"
YEAR=$(date +%Y)
# Parse command-line arguments
while [[ "$#" -gt 0 ]]; do
    case $1 in
        -f|--force) FORCE_REBUILD=true ;;
        --release) RELEASE=true ;;
        *) echo "Unknown parameter passed: $1"; exit 1 ;;
    esac
    shift
done
if [ $RELEASE == true ]; then
    FINAL_VERSION="$VERSION_NUMBER"
else
    FINAL_VERSION="$VERSION_NUMBER ($GIT_HASH)"
fi
echo -e $FINAL_VERSION

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
    if [ -f "$hash_file" ] && [ "$FORCE_REBUILD" = false ]; then
        # Read stored hash
        stored_hash=$(cat "$hash_file")

        # Compare hashes
        if [ "$current_hash" == "$stored_hash" ]; then
            echo -e "[BUILD.SH]=> Skipping ${file}, no changes"
            continue
        fi
    fi

    # Compile the file

    cp "$file" "$BUILD_DIR/$filename"
    sed -e 's|//VArg|Version         bool `short:"v" long:"version" description:"Shows program version."\`|' \
        -e "s|//VCode|if (options.Version) {\n    println(\"$filename_no_ext (Gopherutils) $FINAL_VERSION\\\\nCopyright © $YEAR Pixel\\\\nLicense MIT: MIT License <https://opensource.org/license/mit>\\\\n\\\\nWritten by Pixel\")\n return}|" \
        "$file" > "$BUILD_DIR/$filename"
    echo -e "[BUILD.SH]=> Compiling ${file}"
    go build -o "bin/$filename_no_ext" "$BUILD_DIR/$filename"

    # Store the new hash
    echo "$current_hash" > "$hash_file"
done

echo -e "[BUILD.SH]=> Finished Compiling."
