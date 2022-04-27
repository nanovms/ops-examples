C++ Hello World
================

Depending on your libc version and your distro of choice you'll want to
include libresolv (you might not have to do this on newer versions):

```sh
cc main.cpp -o main
```

config.json:
```sh
{
        "Files": ["/usr/lib/x86_64-linux-gnu/libresolv.so.2"]
}
```

If you're unsure you can do a

```
ops run --trace main &>/tmp/out
```

and grep for library load failures. Most OPS packages auto-include this
for interpreted languages and other newer compiled languages deal with
this differently.

```
ops run -c config.json main
```
