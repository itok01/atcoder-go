$dir = Get-Location

for ($i = 0; $i -lt $args.Length; $i++) {
    if (!(Test-Path ($dir.Path + '/' + $args[$i]))) {
        New-Item -ItemType Directory ($dir.Path + '/' + $args[$i])
    }
    $dir = Resolve-Path ($dir.Path + '/' + $args[$i])
}

$mainGo = 'package main

import "fmt"

func main() {
}'

Write-Output $mainGo | Out-File ($dir.Path + "/main.go")
