Golang HTTP Hello World
==================

 1. Build our golang application.

on mac:
```sh
$ GOOS=linux GOARCH=amd64 go build
```

on linux:

```sh
$ go build
```

 2. Run our application within `ops`
```sh
$ ops run -p 8080 03-http-hello-world
```
