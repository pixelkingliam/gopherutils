# Directory to store hashes
$hashDir = "$env:TEMP\go_build_hashes"
if (-not (Test-Path $hashDir)) {
    New-Item -Path $hashDir -ItemType Directory | Out-Null
}

# Flag for force rebuild
$forceRebuild = $false

# Parse command-line arguments
param (
[switch]$Force
)
$forceRebuild = $Force

# Function to calculate MD5 hash of a file
function Calculate-MD5 {
    param (
        [string]$filePath
    )

    $hashAlgorithm = [System.Security.Cryptography.MD5]::Create()
    $fileStream = [System.IO.File]::OpenRead($filePath)
    $hashBytes = $hashAlgorithm.ComputeHash($fileStream)
    $fileStream.Close()

    return [BitConverter]::ToString($hashBytes) -replace '-'
}

$sourceFiles = Get-ChildItem -Path "source" -Filter "*.go"

foreach ($file in $sourceFiles) {
    $filename = $file.Name
    $filenameNoExt = [System.IO.Path]::GetFileNameWithoutExtension($filename)
    $hashFile = Join-Path -Path $hashDir -ChildPath "$filename.md5"

    # Calculate current hash
    $currentHash = Calculate-MD5 -filePath $file.FullName

    # Check if hash file exists
    if (-not $forceRebuild -and (Test-Path $hashFile)) {
        # Read stored hash
        $storedHash = Get-Content -Path $hashFile

        # Compare hashes
        if ($currentHash -eq $storedHash) {
            Write-Host "[BUILD.PS1]=> Skipping $filename, no changes"
            continue
        }
    }

    # Compile the file
    Write-Host "[BUILD.PS1]=> Compiling $filename"
    & "go.exe" build -o "bin\$filenameNoExt.exe" $file.FullName

    # Store the new hash
    $currentHash | Set-Content -Path $hashFile
}

Write-Host "[BUILD.PS1]=> Finished Compiling."
