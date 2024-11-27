Fresh Deno Unikernel

```
deno run -A -r https://fresh.deno.dev fresh-site 
deno task build
deno compile --include static/ --include _fresh/ --include deno.json --include fresh.gen.ts -A main.ts
ops run -p 8000 fresh-site
```
