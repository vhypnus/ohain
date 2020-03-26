@echo off

set dir=%1%
set srcCommitId=%2%
set targetCommitId=%3%

echo %dir%
echo %srcCommitId%
echo %targetCommitId%

if exist %dir% (
	cd %dir%
	git fetch -p
	git diff %srcCommitId%  %targetCommitId%
)