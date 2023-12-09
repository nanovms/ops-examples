This is a simple pytorch unikernel example.

```
python3 -m venv .venv --prompt nanovm
source .venv/bin/activate
pip install torch
pip install numpy
ops pkg load eyberg/python:3.10.6 -c config.json -a hi.py
```

tree:

```
cp /usr/lib/x86_64-linux-gnu/libgcc_s.so.1 usr/lib/.
cp /usr/lib/x86_64-linux-gnu/libstdc++.so.6 usr/lib/.
cp /usr/lib/x86_64-linux-gnu/libutil.so.1 usr/lib/.
```

```
eyberg@venus:~/pytorch$ tree usr
usr
└── lib
    ├── libgcc_s.so.1
    ├── libstdc++.so.6
    └── libutil.so.1
```
