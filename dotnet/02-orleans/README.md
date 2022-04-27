.NET Orleans Hello World
==================

This example follows the HelloWorld found here:

https://github.com/dotnet/orleans/tree/main/samples/HelloWorld

This was tested with .Net 6 so switch all the '5.0.0' values in

https://github.com/dotnet/orleans/blob/main/samples/HelloWorld/HelloWorld.csproj

to 6.0.0 if that's what you have installed.

Publish as a [self-contained](https://docs.microsoft.com/en-us/dotnet/core/deploying/runtime-patch-selection) app:

```sh
dotnet publish -c release --self-contained --runtime linux-x64 -o ./publish
```

Run with [this `config.json`](./config.json):

```sh
ops run -c config.json publish/HelloWorld
```
