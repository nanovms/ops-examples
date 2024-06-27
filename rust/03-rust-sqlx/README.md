Rust SQLX Example
====================

This is a minimal example to showcase running the sqlx crate with Nanos.

Create the example skeleton and copy in the Rust source:

```sh
$ cargo new b
$ cargo build --release
```

Now we create the filesystem layout with the required libraries:

```sh
mkdir -p lib/x86_64-linux-gnu
mkdir -p etc
cp /lib/x86_64-linux-gnu/libnss_compat.so.2 lib/x86_64-linux-gnu/.
cp /lib/x86_64-linux-gnu/libnss_files.so.2 lib/x86_64-linux-gnu/.
```

You'll need to copy over an /etc/nsswitch.conf:

```sh
cp /etc/nsswitch.conf etc/.
```

```
# /etc/nsswitch.conf
#
# Example configuration of GNU Name Service Switch functionality.
# If you have the `glibc-doc-reference' and `info' packages installed,
try:
# `info libc "Name Service Switch"' for information about this file.

passwd:         compat systemd
group:          compat systemd
shadow:         compat
gshadow:        files

hosts:          files dns mymachines
networks:       files

protocols:      db files
services:       db files
ethers:         db files
rpc:            db files

netgroup:       nis
```

and if you are just playing around locally to test you can add a host to
your /etc/hosts (or use something that is resolveable).

Your layout should look like so:

```sh
eyberg@box:~/sql/b$ tree lib
lib
└── x86_64-linux-gnu
    ├── libnss_compat.so.2
    └── libnss_files.so.2

1 directory, 2 files
eyberg@box:~/sql/b$ tree etc/
etc/
├── hosts
└── nsswitch.conf

0 directories, 2 files
```

Now we can run it:

```sh
$ ops run -c config.json target/release/b
```
