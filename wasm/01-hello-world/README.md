Wasm Hello World
==================

First build your wasm:

```sh
emcc hi.c -s WASM=1 -o hi.html
```

Now run it:

```
$ ops load wasmer_0.1.4 -c config.json
```
