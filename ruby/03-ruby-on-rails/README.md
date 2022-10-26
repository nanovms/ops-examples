Ruby 3.04 and Ruby on Rails 7.0.4 Hello World
==================

This tutorial is in companion to this video:
https://www.youtube.com/watch?v=nT0XpL-7Fnw

Install the Gems:

```sh
gem install rails
rails new blog
cd blog && bin/rails server
```

Copy over the following libs from /usr/lib:

```
➜  blog git:(main) ✗ tree lib
usr
└── lib
    ├── libcrypto.so.1.1
    ├── libdl.so.2
    ├── libm.so.6
    ├── libpthread.so.0
    ├── libssl.so
    ├── libssl.so.1.1
    └── libyaml-0.so.2
```

and this one from /x86_64-linux-gnu/:

```
1 directory, 7 files
➜  blog git:(main) ✗ tree lib
lib
├── assets
├── tasks
└── x86_64-linux-gnu
    └── libm.so.6
```

Stub out the call to getlogin_r with a fake /proc/self/loginuid:
(ensure there is no newline)

```
cat proc/self/loginuid
0%
```

Comment-out port and add bind lines in config/puma.rb:

```
# port ENV.fetch("PORT") { 3000 }
bind "tcp://0.0.0.0:#{ENV['PORT'] || 3000}"
```

Run it:

```sh
$  ops run /usr/bin/ruby -c config.json -p 3000
```
