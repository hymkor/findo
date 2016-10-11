@echo off
setlocal
if not "%1" == "" goto %1
    set GOARCH=386
    go fmt
    go build
    goto end
:test
    findvbproj.exe $(OutputPath)$(AssemblyName).exe
    goto end
:end
endlocal
