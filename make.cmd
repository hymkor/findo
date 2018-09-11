setlocal
call :"%1"
endlocal
exit /b

:""
:"all"
    go fmt
    for %%I in (386 amd64) do (
        set GOARCH=%%I
        mkdir cmd\%%I
        go build -o cmd\%%I\findo.exe -ldflags "-s -w"
    )
    exit /b

:"package"
    for %%I in (386 amd64) do zip -j findo-%%I-%DATE:/=%.zip cmd\%%I\findo.exe
    exit /b
