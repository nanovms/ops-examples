Install HVM via:

https://github.com/HigherOrderCO/HVM#getting-started

```
ops run /home/eyberg/.cargo/bin/hvm -c config.json
```

config.json:

```
{
  "Args": ["run", "(@x(+ x 1) 41)"]
}
```
