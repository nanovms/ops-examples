This example shows how to run an Astro javascript unikernel.

```
npm create astro@latest
npm run build

cp /usr/lib/x86_64-linux-gnu/librt.so.1 usr/lib/x86_64-linux-gnu/.
cp /etc/hosts etc/.
```

```
 ops pkg load eyberg/node:20.5.0 -c config.json -p 4321
```

(This particular node package left out librt which you should copy into
usr.)

```
eyberg@venus:~/astr/tender-tower$ tree etc/
etc/
└── hosts

1 directory, 1 file
eyberg@venus:~/astr/tender-tower$ tree usr
usr
└── lib
    └── x86_64-linux-gnu
        └── librt.so.1

3 directories, 1 file
```
