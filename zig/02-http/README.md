Zig HTTP Unikernel
=======

This zig webserver uses the zap framework. We set the workers to 0 to
avoid forking.

Create a new app:

```
git clone https://github.com/zigzap/zap.git
zig build hello
```

Run:

```
ops run -c config.json -p 3000 ./zig-out/bin/hello
```
