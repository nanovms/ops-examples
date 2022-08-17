NPM Hello World
==================

Install a dependency:

```sh
npm i print-hello-world-test
```

Make a simple config containing the node_modules && your
package-lock.json:

```sh
➜  tt cat config.json
{
  "Dirs": ["node_modules"],
  "Files": ["package-lock.json"]
}
```

Run it:

```sh
➜  tt ops pkg load node_v14.2.0 -c config.json -a hi.js
booting /Users/eyberg/.ops/images/node ...
en1: assigned 10.0.2.15
hello world 1 2 3
exit status 1
```
