OCaml Hello World
==================

 1. Build our ocaml application as a stand-alone (ocamlrun might
    work as well - haven't really tried).

```sh
$ ocamlc -custom -o hex hello.ml
```

 2. Run our application within `ops`
```sh
ops run -c config.json hex
```
