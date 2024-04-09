Get-ChildItem source\*.go | ForEach-Object {
    $filename = $_.Name
    $filename_no_ext = [System.IO.Path]::GetFileNameWithoutExtension($filename)
    go build -o "bin\$filename_no_ext" $_.FullName
}