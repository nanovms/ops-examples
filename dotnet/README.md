.Net Hello World
==================

Generate a New .Net Application:

```
dotnet new console -o myApp
dotnet build
dotnet publish -c release --self-contained --runtime linux-x64
```

```sh
$ ops run -c config.json publish/myApp
```
