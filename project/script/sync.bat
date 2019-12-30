@echo off
set src=%1%
set target=%2%

echo %src%
echo %target%

if not exist %target% (
	md %target%
	git clone %src%    %target%
)

cd %target%

git fetch -p