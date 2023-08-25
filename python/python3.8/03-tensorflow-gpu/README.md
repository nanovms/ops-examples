# Python 3.8 TensorFlow Hello World with the GPU Klib on Google Cloud

In this example we'll run tensorflow inside Nanos using the GPU klib on
Google Cloud.

You may wish to follow along from this tutorial here:
https://nanovms.com/dev/tutorials/gpu-accelerated-computing-nanos-unikernels

You'll first need to build the nvidia klib:
(It is out of tree because it is so large.)

https://github.com/nanovms/gpu-nvidia

```
make NANOS_DIR=/path/to/nanos/kernel/source
```

Download the driver:

https://www.nvidia.com/Download/driverResults.aspx/191320/en-us/

```
~/Downloads/NVIDIA-Linux-x86_64-515.65.01.run -x
```

Copy over the firmware:

```
cp ~/Downloads/NVIDIA-Linux-x86_64-515.65.01/firmware/gsp.bin nvidia/515.65.01/
```

In your example unikernel repo you should have a layout that looks like
this:

## Setup

```bash
$ tree
dev
nvidia
└── 515.65.01
    └── gsp.bin
proc
└── self
    └── cmdline
usr
hi.py
keras_sequential.py
lib
└── x86_64-linux-gnu
    └── libnvidia-ptxjitcompiler.so.1
```

## Run

```bash
ops pkg load python_3.8.6 -c config.json
```
