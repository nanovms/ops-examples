This example shows how to run a Vite javascript unikernel.

```
npm create vite@latest
cd vite-project
npm install
```

```
 ops pkg load eyberg/node:20.5.0 -c config.json -p 5173
```

(This particular node package left out librt which you should copy into
usr:)

```
eyberg@venus:~/v/vite-project$ tree usr
usr
└── lib
    └── x86_64-linux-gnu
        └── librt.so.1
```
