Rust RoAPI Example
====================

This is a minimal example to showcase running the Roapi example with Nanos.

The easiest example is to use the statically linked MUSL example found
here:

https://github.com/roapi/roapi/releases/download/roapi-http-v0.1.3/roapi-http-x86_64-unknown-linux-musl.tar.gz

```sh
{
  "Args": ["-a", "0.0.0.0:8080", "-t", "spacex:s3://rand-img/spacex-launches.json"],
  "Dirs": ["root"]
}
```

Place your ~/.aws/credentials and ~/.aws/config inside a local 'root'
folder like so:

```
eyberg@box:~/tt$ tree root/.aws/
root/.aws/
├── config
└── credentials
```

If you are compiling from source then you must manually copy over all
the libs that aren't available in ldd output:

```
eyberg@s1:~/r/roapi/target/debug$ cat config.json
{
  "Args": ["-a", "0.0.0.0:8080", "-t", "spacex:s3://rand-img/spacex-launches.json"],
  "Dirs": ["root", "usr", "etc", "lib"]
}
eyberg@s1:~/r/roapi/target/debug$ tree root/ usr/ etc/ lib
root/
usr/
└── lib
    ├── ssl
    │   └── openssl.cnf
    └── x86_64-linux-gnu
        ├── libcrypto.so.1.1
        └── libssl.so.1.1
etc/
└── nsswitch.conf
lib
└── x86_64-linux-gnu
    ├── libnss_compat.so.2
    ├── libnss_files.so.2
    └── libnss_systemd.so.2

4 directories, 7 files
```

If you are running into problems you can turn on '--trace' to see which
libraries you might be missing. Grep through the output and then copy them in:

```
ops run --trace -c config.json -p 8080 roapi-http  &>/tmp/o
```

Now we create the filesystem layout with the required libraries:

```sh
mkdir -p lib/x86_64-linux-gnu
mkdir -p etc
mkdir -p usr/lib/ssl/
mkdir -p usr/lib/x86_64-linux-gnu
cp /lib/x86_64-linux-gnu/libnss_compat.so.2 lib/x86_64-linux-gnu/.
cp /lib/x86_64-linux-gnu/libnss_files.so.2 lib/x86_64-linux-gnu/.
....
```
