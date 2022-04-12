Go Read-Write Mounted Volume
==================

1. build our Go application
```sh
$ GOOS=linux GOARCH=amd64 go build -o volume-example
```

2. create volume using `ops volume create <volume_name> -d <path_to_data_source>`
```sh
ops volume create not-empty -d files/
```

3. get list of created volumes to find the UUID of relevant volume to attach
```sh
ops volume list

NAME		UUID		PATH
not-empty	some-uuid	path-to-volume
```

4. run our application with volume attached using `ops run <path_to_binary> --mounts <volume_UUID:mount_path>`
```sh
ops run volume-example --mounts some-uuid:/mnt
```
this sample application runs binary created from `volume.go` twice. on first run, you should expect to see this in the output
```
reading mnt/hello.txt:
hello
```
and on the next run, you should expect to see this in the output
```
reading mnt/hello.txt:
hello
hi
```
