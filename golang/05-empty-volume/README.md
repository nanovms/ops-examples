Go Read-Write Mounted Volume
==================

1. build our Go application
```sh
$ GOOS=linux go build -o volume-example
```

2. create empty volume using `ops volume create <volume_name>`
```sh
ops volume create empty
```

3. get list of created volumes to find the UUID of relevant volume to attach
```sh
ops volume list 

NAME		UUID		PATH
empty		some-uuid	path-to-volume
```

4. run our application with volume attached using `ops run <path_to_binary> --mounts <volume_UUID:mount_path>`
```sh
ops run volume-example --mounts some-uuid:/mnt
```

this sample application runs binary created from `volume.go` twice. on first run, you should expect to see this in the output
```
open mnt/hello.txt: no such file or directory
```
as the attached volume is initially empty, but on the next run, file `hello.txt` has been successfully created
