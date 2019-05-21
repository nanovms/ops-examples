Nginx Hello World Config
==================

You should look at the first example first before trying this one.

In earlier examples we showed how to add html to your filesystem. Now
let's edit nginx's configuration and make it listen on a different port.

In this example we tell OPS to load the entire usr directory specified
(which has a new nginx.conf) that listens on port 8084.

```
ops load nginx_1.15.6 -c config.json -p 8084
```

Now if we load it we see that we can listen on a different port.

```
curl -XGET http://127.0.0.1:8084/
```
