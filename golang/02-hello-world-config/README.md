Golang Hello World
==================

First, we need to compile our executable, but we need to make sure we compile
for the `ops` environment and not the one our machine is. For example, if
you're using a mac, we want to compile for linux rather than mac. If you're
using linux, the compilation should be the same.

 1. Build our golang application.

```sh
$ GOOS=linux go build hello.go
```

 2. Run our application within `ops`
```sh
$ ops run -c config.json hello
```
