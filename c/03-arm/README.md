This is a simple hello world example to cross-compile for arm on x86.

```
aarch64-linux-gnu-gcc main.c -static -o arm
```

```
ops run arm --arch=arm64
```
