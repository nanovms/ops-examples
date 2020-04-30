Rust Hello World
==================

If on Linux:

```sh
$ ops run main
```

If on Mac:

If you are on a mac you'll need to cross-compile your rust program to
Linux first as OPS/Nanos uses ELFs.

Install the toolchains first:
```
brew tap SergioBenitez/osxct
brew install FiloSottile/musl-cross/musl-cross
```

Add the target:
```
rustup target add x86_64-unknown-linux-musl
```

Put this into your .cargo/config:
```
[target.x86_64-unknown-linux-musl]
linker = "x86_64-linux-musl-gcc"
```

Build:

```
TARGET_CC=x86_64-linux-musl-gcc \
RUSTFLAGS="-C linker=x86_64-linux-musl-gcc" \
cargo build --target=x86_64-unknown-linux-musl
```
