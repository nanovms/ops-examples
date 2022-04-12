A note about using golang...
============================

When building your executable, we want to always make sure we are building for
the OS of `ops` rather than your machine's OS. So always make sure you are
building for linux.

```sh
$ GOOS=linux GOARCH=amd64 go build
```
