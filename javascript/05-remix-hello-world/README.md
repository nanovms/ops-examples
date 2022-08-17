Remix Hello World
==================

This is largely based off of https://remix.run/docs/en/v1/tutorials/blog.

This is slightly different because of it being in a unikernel environment we don't support shelling out to run various commands (via npx) but it's fairly easy to translate those into what they actually run.

Note: This example assumes a linux env so you'll want to build it there or if you are on a mac use the crossbuild functionality as shown here: https://docs.ops.city/ops/crossbuilding .

Let's start by creating the application:

```
npx create-remix --template remix-run/indie-stack blog-tutorial
```

Generate your session secret:
```
openssl rand -hex 32
```

Next we'll going to create a tmp folder:

```
mkdir -p tmp
```

Then we'll going to grab the libraries that Prisma was installed with. These aren't found or copied into the disk image automatically because they are loaded at runtime.

How can you tell that you have libraries that load dynamically? You can try this command.

```
➜  blog-tutorial find . -name "*.so.node"
./node_modules/.prisma/client/libquery_engine-debian-openssl-1.1.x.so.node
./node_modules/@prisma/engines/libquery_engine-debian-openssl-1.1.x.so.node
./node_modules/prisma/libquery_engine-debian-openssl-1.1.x.so.node
```

We'll have this folder loaded but these libraries might have other dependencies.

We can find that out like so:

```
➜  blog-tutorial ldd node_modules/.prisma/client/libquery_engine-debian-openssl-1.1.x.so.node                                                                                    
        linux-vdso.so.1 (0x00007ffd9b1ab000)                                                                                                                                     
        libz.so.1 => /usr/lib/libz.so.1 (0x00007f178ac05000)                                                                                                                     
        libssl.so.1.1 => /usr/lib/libssl.so.1.1 (0x00007f178ab70000)                    
        libcrypto.so.1.1 => /usr/lib/libcrypto.so.1.1 (0x00007f1788a00000)                                                                                                       
        libgcc_s.so.1 => /usr/lib/libgcc_s.so.1 (0x00007f178ab50000)                    
        librt.so.1 => /usr/lib/librt.so.1 (0x00007f178ab4b000)                                                                                                                   
        libpthread.so.0 => /usr/lib/libpthread.so.0 (0x00007f178ab46000)                                                                                                         
        libm.so.6 => /usr/lib/libm.so.6 (0x00007f1788d19000)                                                                                                                     
        libdl.so.2 => /usr/lib/libdl.so.2 (0x00007f178ab3f000)                                                                                                                   
        libc.so.6 => /usr/lib/libc.so.6 (0x00007f1788600000)                                                                                                                     
        /usr/lib64/ld-linux-x86-64.so.2 (0x00007f178ac3f000) 
```

Now most of these are already going to be present on our image whether we used a pre-made package or we'll just running /usr/bin/node like in this example, however, you can compare and see that libssl and libcrypto are missing which happen to the ones that we need.

So we'll create a new dir and copy them in:

```
mkdir -p usr/lib
```

```
➜  blog-tutorial tree usr 
usr
├── lib
│   ├── libcrypto.so.1.1
│   ├── libc.so.6
│   ├── libdl.so.2
│   ├── libgcc_s.so.1
│   ├── libm.so.6
│   ├── libpthread.so.0
│   ├── librt.so.1
│   ├── libssl.so.1.1
│   └── libz.so.1
└── lib64
    └── ld-linux-x86-64.so.2

2 directories, 10 files
```

Then our config looks like this. You can see I just mapped over the entire node_modules npm folder which I found like so:

```
➜  ~ which npm
/usr/bin/npm
➜  ~ ls -lh /usr/bin/npm 
lrwxrwxrwx 1 root root 38 Jul 24 20:53 /usr/bin/npm -> ../lib/node_modules/npm/bin/npm-cli.js
```

Thus the entry command is actually reflected in the ```Args``` array and we set the required env vars. For instance if you don't set the session secret env var you might wind up with a very obtuse error message.

```
{
    "BaseVolumeSz": "250m",
    "Args": ["node_modules/@remix-run/serve/dist/cli.js", "build"],
    "Dirs": ["node_modules", "build", "public", "prisma", "usr", "tmp"],
    "Files": ["package.json", ".env"],
    "MapDirs": {"/usr/lib/node_modules/npm/*": "/usr/lib/node_modules/npm" },
    "Env": {
      "NODE_ENV": "development",
      "PORT": "8080",
      "SESSION_SECRET": "140bc6f511dc672b8fe726ccd0394891e07de7a5531d5b8f31925eabb35aa56f"
    }
}
```

and wah-lah - you're now running Remix as a unikernel.

```
➜  blog-tutorial ops run /usr/bin/node -n -c config.json -p 8080
booting /home/eyberg/.ops/images/node.img ...
en1: assigned 10.0.2.15
Remix App Server started at http://localhost:8080 (http://10.0.2.15:8080)
en1: assigned FE80::ACFD:F9FF:FE25:F03A
```

Next steps?

* Try creating a package on https://repo.ops.city of a remix base.
* Try deploying your remix application to AWS or GCE.
