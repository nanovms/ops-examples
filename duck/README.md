Duck unikernel hello world

Duck is a modern, compiled and batteries included programming language
for full-stack web-development, that's built on top of the go ecosystem,
it leverages the concepts of duck-typing onto a fast and reliable
platform while introducing JSX like server-side templating and
client-side react components as first-class citizens.

You need to ensure that your go version matches whatever is being built
in nightly. For instance as of this writing go 1.25.5 is required.

Grab duck:

```
curl -fsSL https://duckup.sh | bash
duckup update
```

Generate a new project (or use this one):

```
dargo new
```

Compile && run:

```
dargo compile src/main.duck
ops run .dargo/duck_out
```
