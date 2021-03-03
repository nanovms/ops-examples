Nginx Hello World
==================

 1. Run nginx with the built-in configuration:

```sh
$ ops pkg load nginx_1.15.6 -p 8083
```

You'll notice the webserver is running but all it returns is error
pages. Now let's do it with a basic configuration:

```sh
ops pkg load nginx_1.15.6 -c config.json -p 8083
```

This tells OPS to load index.html into the root of the nginx filesystem.
We can now hit it with:

```
curl -XGET http://127.0.0.1:8083/index.html
```
