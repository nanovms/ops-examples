Next.JS Hello World
==================

For this example we'll follow along with the example shown here:

https://nextjs.org/learn/basics/create-nextjs-app/setup

First let's generate the application:

```
npx create-next-app@latest nextjs-blog --use-npm --example "https://github.com/vercel/next-learn/tree/main/basics/learn-starter"
```

The instructions say to do a 'npm run dev' but let's look a litle closer
at what is going on here.

`npm run dev` is cli sugar for:

```
node /home/eyberg/ntut/nextjs-blog/node_modules/.bin/next dev
```

however, '.bin/next' is a symlink. For instance on my box:

```
eyberg@box:~/ntut/nextjs-blog$ ls -lh node_modules/.bin/next
lrwxrwxrwx 1 eyberg eyberg 21 Sep 28 21:05 node_modules/.bin/next -> ../next/dist/bin/next
```

So in our config we'll fix that.

Next let's build the app so we have a .next directory which we need to
include in our unikernel, along with node_modules, pages, public, and
styles.

```
node node_modules/next/dist/bin/next build
```

Then we can use serve to get a prod build:

```sh
ops pkg load eyberg/node:20.5.0 -c config.json -p 3000
```
