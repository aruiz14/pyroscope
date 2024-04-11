# pproffix

## TL;DR

Download the source code and compile into the current directory (a working Go toolchain is required to be installed):

```
tmp="$(mktemp -d)"
curl -LSsLf https://github.com/aruiz14/pyroscope/archive/main.tar.gz | tar zx -C "${tmp}" --strip-components=1
GOBIN="${PWD}" go install -C "${tmp}" ./tools/pproffix
rm -rf "${tmp}"
```

## Usage

```
./pproffix [-normalize] [-concurrency=N] [-suffix=] <list of pprof files>
```