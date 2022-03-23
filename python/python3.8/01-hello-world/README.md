```
ops pkg load eyberg/python:3.8.6 -p 8083 -c config.json
```

On Google Cloud:
```
ops image create -t gcp -g my-project -c config.json --package eyberg/python:3.8.6
```
