@echo off
setlocal
echo.

REM Parameter --clean passed
if "%~1"=="--clean" (
    echo [Step] Cleanup^!
    call go clean -cache
    call go clean -modcache
    echo Cache cleaned
    echo.

    echo [Step] Installing Dependencies^!
    call go mod tidy
    echo.
)

echo [Step] Compiling Templates^!
del /s /q frontend\templates\*_templ.go >nul 2>&1 REM Delete old templates
call templ generate ./frontend/templates -v
echo.

echo [Step] Building Binaries^!
call go build -o main.exe ./backend
if errorlevel 1 (
    echo Build failed^!
    pause
    exit /b 1
)
echo Binaries built successfully^!
echo.

echo [Step] Running Server^!
echo Starting main.exe...
start "" main.exe

endlocal