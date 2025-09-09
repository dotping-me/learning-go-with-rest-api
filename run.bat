@echo off
echo.

echo [1] Cleanup^!
go clean -cache
go clean -modcache
echo Cache cleaned
echo.

echo [2] Installing Dependencies^!
go mod tidy
echo.

echo [3] Compiling Templates^!
templ generate
echo.

echo [4] Building Binaries^!
go build -o main.exe ./backend
if errorlevel 1 (
    echo Build failed^!
    pause
    exit /b 1
)
echo Binaries built successfully^!
echo.

echo [5] Running Server^!
echo Starting main.exe...
start "" main.exe