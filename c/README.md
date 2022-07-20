OpenCL C Example
====================

This is a short example showing OpenCL usage taken from:

https://github.com/KhronosGroup/OpenCL-Guide/blob/main/chapters/getting_started_linux.md

make:

```
gcc -Wall -Wextra -D CL_TARGET_OPENCL_VERSION=100 main.c -o HelloOpenCL -lOpenCL
```

```
eyberg@box:~/cltest$ ops run -c config.json HelloOpenCL
booting /home/eyberg/.ops/images/HelloOpenCL.img ...
en1: assigned 10.0.2.15
1 platform(s) found
```

I copied over these libs from their respective location:

```
eyberg@box:~/cltest$ tree usr
usr
└── lib
    └── x86_64-linux-gnu
        ├── libbsd.so.0
        ├── libclang-cpp.so.12
        ├── libdrm.so.2
        ├── libedit.so.2
        ├── libelf.so.1
        ├── libexpat.so.1
        ├── libffi.so.8
        ├── libgcc_s.so.1
        ├── libicudata.so.67
        ├── libicuuc.so.67
        ├── libLLVM-12.so.1
        ├── liblzma.so.5
        ├── libmd.so.0
        ├── libMesaOpenCL.so.1
        ├── libm.so.6
        ├── libstdc++.so.6
        ├── libtinfo.so.6
        ├── libxml2.so.2
        ├── libz.so.1
        └── libzstd.so.1

2 directories, 20 files
```
