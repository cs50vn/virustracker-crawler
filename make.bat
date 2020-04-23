@echo off

rem Current dir
rem Host type
rem Build type: all, release, debug

set ARCH_TYPE=all
set BUILD_TYPE=%1%

python scripts/build.py %CD% windows %BUILD_TYPE%

