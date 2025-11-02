Fil-C unikernel example.

OPS will try and load all the libraries that it finds onto your disk,
however, the loader itself (ld-yolo-x86_64) is different in Fil-C. It is a symlink to libyoloc.so .

For example these libraries will be auto-found and loaded:

```
eyberg@venus:~/tr$ ldd hi
        linux-vdso.so.1 (0x00007fff53f03000)
        libc.so => /home/eyberg/fi/filc-0.674-linux-x86_64/build/bin/../../pizfix/lib/libc.so (0x00007879c0c00000)
        libpizlo.so => /home/eyberg/fi/filc-0.674-linux-x86_64/build/bin/../../pizfix/lib/libpizlo.so (0x00007879c0800000)
        libyoloc.so => /home/eyberg/fi/filc-0.674-linux-x86_64/build/bin/../../pizfix/lib/libyoloc.so (0x00007879c11e1000)
```

Howver, this will not be:

```
eyberg@venus:~/tr$ file hi
hi: ELF 64-bit LSB pie executable, x86-64, version 1 (SYSV), dynamically linked, interpreter /home/eyberg/fi/filc-0.674-linux-x86_64/build/bin/../../pizfix/lib/ld-yolo-x86_64.so, with debug_info, not stripped
```

As OPS expects that not to really change. In this scenario you simply
need to include it manually.

```
mkdir -p home/eyberg/fi/filc-0.674-linux-x86_64/pizfix/lib
mkdir -p home/eyberg/fi/filc-0.674-linux-x86_64/build/bin
```

Then copy it over into your location:

```
cp /home/eyberg/fi/filc-0.674-linux-x86_64/build/bin/../../pizfix/lib/ld-yolo-x86_64.so home/eyberg/fi/filc-0.674-linux-x86_64/pizfix/lib/.
```

Let's compile:

Just adjust the path to your filc installation.

```
~/fi/filc-0.674-linux-x86_64/build/bin/clang -O2 -g -o main main.c
```
