.NET Hello World
==================

Create a new console app:

```sh
dotnet new console -o myApp
dotnet build
```

Publish as a [self-contained](https://docs.microsoft.com/en-us/dotnet/core/deploying/runtime-patch-selection) app:

```sh
dotnet publish -c release --self-contained --runtime linux-x64 -o ./publish
```

or as a [single file](https://docs.microsoft.com/en-us/dotnet/core/deploying/single-file) (and self-contained) app:

```sh
dotnet publish -c release --self-contained --runtime linux-x64 -o ./publish -p:PublishSingleFile=true
```

Run with [this `config.json`](./config.json):

```sh
ops run -c config.json publish/myApp
```
