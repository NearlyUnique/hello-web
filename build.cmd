:poor mans build script

@echo off
setlocal ENABLEDELAYEDEXPANSION


if "%1"=="" (
    echo USAGE: build target
    goto END
)
set DEST=%1
if exist %DEST% (
    echo about to delete %DEST%
    pause
    rd %DEST% /q/s
) else (
    echo what %DEST%.
)
md %DEST%
pushd %DEST%
git init
popd
set part=0

pushd %DEST%
:create home command
echo @git co . > home.cmd
echo @git co part0 >> home.cmd
:create end command
echo @git co . > end.cmd
echo @git co master >> end.cmd
popd

for /F "tokens=*" %%A in (index.txt) do (
    copy %%A %DEST%
    pushd %DEST%
    
    :create next command
    set /a next=!part!+1
    echo @git co . > next.cmd
    echo @git co part!next! >> next.cmd
    :create prev command
    set /a prev=!part!-1
    echo @git co . > prev.cmd
    echo @git co part!prev! >> prev.cmd

    git add .
    git commit -m "Part !part!: %%A"
    git branch part!part!

    set /a part+=1

    popd
)
:END