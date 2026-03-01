Next.JS Unikernel Hello World
==================

For this example we'll follow along with the example shown here:

```
npx create-next-app@latest      # "my-app" is default's project name
cd my-app
npm run build
```

```sh
 # to be run inside nextJS project directory (e.g "my-app")
 cp ../config.json .
 cp ../proc .
 cp ../index.js .
 ops pkg load eyberg/node:20.5.0 -p 3000 -c config.json
```
