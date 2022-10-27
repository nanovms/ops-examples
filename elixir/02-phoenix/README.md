This is a quickstart for running the Phoenix web framework on Elixir.

```ops run -c erl.json _build/dev/rel/demo/erts-11.0/bin/erlexec```

This was pieced together from looking at:
(ther might be a better way of dealing with this)

```
PHX_SERVER=true ./_build/dev/rel/demo/bin/demo start
```

and cutting a release:

```
mix phx.gen.release
mix release
```

erl.json:
(note: I'm not confident all of this is necessary but I'll let an
elixir/phoenix pro correct this.)
```
{
  "Env": {
    "HOME": "/",
    "ROOTDIR": "/usr/local/lib/erlang",
    "BINDIR": "/usr/local/lib/erlang/erts-11.0/bin",
    "EMU": "beam",
    "PROGNAME": "hi",
    "MIX_DEBUG": "true",
    "MIX_ENV": "dev",
    "MIX_HOME": "/home/eyberg/.mix",
    "HEX_OFFLINE": "true",
    "PHX_SERVER": "true",
    "RELEASE_SYS_CONFIG": "/_build/dev/rel/demo/releases/0.1.0/sys",
    "RELEASE_ROOT": "/_build/dev/rel/demo",
    "LC_ALL": "C.UTF-8",
    "LANG": "en_US.UTF-8"
  },
  "BaseVolumeSz": "200m",
  "Files": ["mix.exs", "mix.lock", "mix"],
  "Dirs": ["assets", "_build", "config", "etc", "lib", "priv", "usr", "test"],
  "MapDirs": {
    "/home/eyberg/.mix/*": "/home/eyberg/.mix",
    "/home/eyberg/elix/*": "/home/eyberg/elix",
    "/home/eyberg/.hex/*": "/home/eyberg/.hex",
    "/home/eyberg/.hex/*": "/.hex",
    "/usr/local/lib/erlang/*": "/usr/local/lib/erlang",
    "/usr/lib/locale/C.utf8/*": "/usr/lib/locale/C.utf8",
    "/usr/lib/x86_64-linux-gnu/gconv/*":"/usr/lib/x86_64-linux-gnu/gconv"
 },
  "Args": ["/_build/dev/rel/demo/erts-11.0/bin/erlexec", "-elixir", "ansi_enabled", "true",
  "-noshell", "-s", "elixir", "start_cli", "-mode", "embedded", "-setcookie",
  "I5OQ66G32JL7IHZHTNOATBVLYFMKG2HDPB2Z3KD3EZ5JM52Y52EA====", "-sname", "demo", "-config",
  "/_build/dev/rel/demo/releases/0.1.0/sys", "-boot",
  "/_build/dev/rel/demo/releases/0.1.0/start", "-boot_var", "RELEASE_LIB",
  "/_build/dev/rel/demo/lib", "-args_file",
  "/_build/dev/rel/demo/releases/0.1.0/vm.args", "-extra", "--no-halt"]
}
```

You'll want to copy the following from their respective locations:

```
eyberg@s1:~/m/demo$ tree usr etc
usr
├── lib
│   ├── locale
│   │   └── locale-archive
│   └── x86_64-linux-gnu
│       └── libcrypto.so.1.1
└── share
    └── locale
        └── locale.alias
etc
└── locale.alias
```

Finally adjust your webserver to listen on any-interface in this config
file: _build/dev/rel/demo/releases/0.1.0/runtime.exs

```
config :demo, DemoWeb.Endpoint, server: true, http: [ip: {0, 0, 0, 0}, port: 4000]
```
