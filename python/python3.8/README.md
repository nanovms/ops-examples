# Python 3.8 Flask Hello World with Requests

## Setup

```bash
python -m venv .venv --prompt nanovm
source .venv/bin/activate
pip install --upgrade pip
pip install -r requirements.txt
```

## Run

```bash
ops pkg load python_3.8.6 -c config.json -p 8080
```

## How to add more Python dependencies

1. `source .venv/bin/activate`
2. `pip install <DEPENDENCY>`

## How to add more C dependencies

You'll need `libgcc_s.so.1` and `libz.so.1` in [usr/lib64/](./usr/lib64/) for this example to work.

### If you are running Linux x86_64

Copy the dependencies over using [config.json](config.json).

### Another option

1. Download x86_64 binary RPM packages from [https://pkgs.org/](https://pkgs.org/), e.g., for a CentOS distro.
2. Get [rpm2cpio](https://man7.org/linux/man-pages/man8/rpm2cpio.8.html) and [cpio](https://en.wikipedia.org/wiki/Cpio).
3. `rpm2cpio ./<PACKAGE_NAME>.rpm | cpio -idmv`
4. Tidy up the extracted binaries and adapt [config.json](config.json).