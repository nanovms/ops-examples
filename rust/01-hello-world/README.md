Rust Hello World
==================

If on Linux:

```sh
$ ops run main
```

If on Mac:

If you are on a mac you'll need to cross-compile your rust program to
Linux first as OPS/Nanos uses ELFs. There are 2 ways of doing this:

Static or dynamic. Using MUSL will be the easiest to do but dynamic will
get you full ASLR and will probably be much faster.

### Static with MUSL

Install the toolchains first:
```
brew tap SergioBenitez/osxct
brew install FiloSottile/musl-cross/musl-cross
```

Add the target:
```
rustup target add x86_64-unknown-linux-musl
```

Build:

```
TARGET_CC=x86_64-linux-musl-gcc \
rustc --target=x86_64-unknown-linux-musl \
-C linker=x86_64-linux-musl-gcc main.rs
```

To run:

```
➜ ops run main
 100% |████████████████████████████████████████|  [0s:0s]
booting /Users/eyberg/.ops/images/main.img ...
assigned: 10.0.2.15
yo!
exit status 1
```

### Dynamic

The process here is a bit different as you'll need to copy over any
libraries you are linked to and put them in the correct spot.

Install the toolchain from the same brew tap:
```
brew tap SergioBenitez/osxct
brew install x86_64-unknown-linux-gnu
```

Add the target:
```
rustup target add x86_64-unknown-linux-gnu
```

Compile:
```
TARGET_CC=x86_64-unknown-linux-gnu-gcc \
rustc --target=x86_64-unknown-linux-gnu \
-C linker=x86_64-unknown-linux-gnu-gcc main.rs
```

Sample Config:
```
{
  "Dirs": ["lib64"]
}
```

Next you need to point at or copy the libraries over:

```
➜  tree
.
├── config.json
├── lib64
│   ├── ld-linux-x86-64.so.2
│   ├── libc.so.6
│   ├── libdl.so.2
│   ├── libgcc_s.so.1
│   ├── libpthread.so.0
│   └── librt.so.1
├── main
└── main.rs
```

From the toolchain I installed it was installed here:

```
cp /usr/local/Cellar/x86_64-unknown-linux-gnu/7.2.0/toolchain/x86_64-unknown-linux-gnu/sysroot/lib64/ld-linux-x86-64.so.2 lib64/
cp /usr/local/Cellar/x86_64-unknown-linux-gnu/7.2.0/toolchain/x86_64-unknown-linux-gnu/sysroot/lib64/libc.so.6 lib64/
cp /usr/local/Cellar/x86_64-unknown-linux-gnu/7.2.0/toolchain/x86_64-unknown-linux-gnu/sysroot/lib64/libdl.so.2 lib64/
cp /usr/local/Cellar/x86_64-unknown-linux-gnu/7.2.0/toolchain/x86_64-unknown-linux-gnu/sysroot/lib64/libgcc_s.so.1 lib64/
cp /usr/local/Cellar/x86_64-unknown-linux-gnu/7.2.0/toolchain/x86_64-unknown-linux-gnu/sysroot/lib64/libpthread.so.0 lib64/
cp /usr/local/Cellar/x86_64-unknown-linux-gnu/7.2.0/toolchain/x86_64-unknown-linux-gnu/sysroot/lib64/librt.so.1 lib64/
```

If you need to figure out what libraries are there you can use:

```
gobjdump -p main
```

You can also scp the binary to linux and run ldd on it:

```
eyberg@s1:~$ ldd main
        linux-vdso.so.1 (0x00007fff9d1e5000)
        libdl.so.2 => /lib/x86_64-linux-gnu/libdl.so.2 (0x00007f0b7b050000)
        librt.so.1 => /lib/x86_64-linux-gnu/librt.so.1 (0x00007f0b7b045000)
        libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007f0b7b022000)
        libgcc_s.so.1 => /lib/x86_64-linux-gnu/libgcc_s.so.1 (0x00007f0b7b007000)
        libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f0b7ae15000)
        /lib64/ld-linux-x86-64.so.2 (0x00007f0b7b2a1000)
```

Finally remember you can use the '-d' flag in ```ops run``` to get the
paths that it is looking in as well.

To run:

```
➜ LD_LIBRARY_PATH=lib64 ops run -c config.json main
 100% |████████████████████████████████████████|  [0s:0s]
booting /Users/eyberg/.ops/images/main.img ...
assigned: 10.0.2.15
yo!
exit status 1
```
