[![Go Report Card](https://goreportcard.com/badge/github.com/zetamatta/findo)](https://goreportcard.com/report/github.com/zetamatta/findo)

findo - a command like `find`
==============================

The `findo` search files from the tree below the current directory.

```
$ findo -h
Usage of findo.exe:
  -1    Show nameonly without size and timestamp
  -d string
        Set start Directory (default ".")
  -f    Select fileonly not including directories
  -ignoredots
        Ignore files and directory starting with dot
  -in duration
        Files modified in the duration such as 300ms, -1.5h or 2h45m
  -notin duration
        Files modified not in the duration such as 300ms, -1.5h or 2h45m
  -l    Show size and timestamp
  -q    Enclose filename with double-quotations
  -v    verbose (use with -x)
  -x string
        Execute a command replacing {} to FILENAME
```

Example-1: no arguments and no redirect
-------------------------------------

`findo` prints the tree under the current directory.

```
$ findo
.git
       4,096 2018-06-22 18:08:28.6822833 +0900 JST
.git\COMMIT_EDITMSG
         347 2017-12-29 08:54:25.1692404 +0900 JST
.git\FETCH_HEAD
          96 2018-06-22 18:08:28.6197845 +0900 JST
.git\HEAD
          23 2017-11-28 21:04:59.3605057 +0900 JST
```

Example-2: when the standard output is redirected
-------------------------------------------------

Timestamp and size are omitted.

```
$ findo | more
.git
.git\COMMIT_EDITMSG
.git\FETCH_HEAD
.git\HEAD
    :
```

Example-3: A filename pattern(wildcard) is given.
---------------------------------------

```
$ findo H*
.git\HEAD
          23 2017-11-28 21:04:59.3605057 +0900 JST
.git\hooks
       4,096 2017-11-28 21:04:59.1064544 +0900 JST
.git\logs\HEAD
         716 2018-06-22 18:08:28.6822833 +0900 JST
    :
```

Example-4: Executing commands
-----------------------------

```
$ findo -q -x "echo {} & wc.exe -l {}" *.md *.go
"findo.go"
     95
"hoge and hoge.md"
      0
"readme.md"
     90
```

History
=======

(no binary)
-----------
- Support Linux not only Windows

20180912
--------
- Add options: `-x`,`-q`,`-d`,`-l`
- Support Multi patterns
- If stdout is terminal, output simple. and add `-l` option

20171128
--------
- First release
