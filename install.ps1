param(
    [string]$Prefix = "$env:USERPROFILE\AppData\Local\Programs\cli-calculator\bin",
    [string]$AppName = "calc",
    [string]$BuildDir = "bin"
)

$ErrorActionPreference = "Stop"

$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$goModPath = Join-Path $scriptDir "go.mod"

if (-not (Test-Path $goModPath)) {
    throw "go.mod topilmadi: $goModPath"
}

$binaryName = $AppName
if (-not $binaryName.EndsWith(".exe")) {
    $binaryName = "$binaryName.exe"
}

$buildPath = Join-Path $scriptDir $BuildDir
$binaryPath = Join-Path $buildPath $binaryName

New-Item -ItemType Directory -Force -Path $buildPath | Out-Null
New-Item -ItemType Directory -Force -Path $Prefix | Out-Null

Push-Location $scriptDir
try {
    go build -o $binaryPath .
} finally {
    Pop-Location
}

Copy-Item -Force $binaryPath (Join-Path $Prefix $binaryName)

Write-Host "installed: $(Join-Path $Prefix $binaryName)"

$pathEntries = ($env:PATH -split ';') | Where-Object { $_ -ne "" }
if ($pathEntries -notcontains $Prefix) {
    Write-Host "PATH ga qoshing: `$env:PATH = `"$Prefix;`$env:PATH`""
}
