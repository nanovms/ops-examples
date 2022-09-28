React Hello World
==================

First let's create our application and run it locally:

```sh
npx create-react-app helloworld && cd helloworld
npm start
npm run build
```

Then we can use serve to get a prod build:

```sh
npm install -g serve

serve -s build
```

Finally we can wrap it up with a node package and config:

```sh
$ ops pkg load eyberg/node:v16.5.0 -c config.json -p 3000
```
