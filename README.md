# `lah`
command line application that can be used alternative to `ls -lah`

### Purpose

We already have `ls` which is builtin command for almost all *nix I used. `ls` is used in many automation. The purpose of `lah` project is not to replace `ls` tool. This tool exist to complement `ls -lah` use case. I expect this command line tools will help linux noobs to understand and grasp different columns and their values used in `ls -lah`.

### Usage

```
Usage: lah [FILE]

List information about the FILEs (the current directory by default).

[FILE] Directory you want to list files for
```

### Preview

This is how the response looks like.

```bash
$ lah
FIL 0644  0 269 B hasanlock:staff .editorconfig
DIR 0755 12 448 B hasanlock:staff .git
FIL 0644  0 269 B hasanlock:staff .gitignore
FIL 0644  0 1.1KB hasanlock:staff LICENSE
FIL 0644  0 113 B hasanlock:staff README.md
FIL 0644  0 111 B hasanlock:staff Taskfile.yml
SOC 0755  0   0 B hasanlock:staff aSocket.sock
FIL 0644  0 2.5KB hasanlock:staff dispfmt.go
FIL 0644  0  36 B hasanlock:staff go.mod
FIL 0755  0 2.2MB hasanlock:staff lah
FIL 0644  0 786 B hasanlock:staff lah.go
DIR 0644  0   0 B hasanlock:staff sample_pipe
DIR 0755  1  96 B hasanlock:staff some
```

### Upcoming [WIP]
```
@todo provide install guide
@todo if user pass a file, not dir
@todo if user want to list multiple dir
```
