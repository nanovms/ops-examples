A modern, jvm-based, full-stack framework for building modular, easily
testable microservice and serverless applications.

Install sdkman:

```
curl -s "https://get.sdkman.io" | bash
```

Install micronaut cli tool:
```
sdk install micronaut
```

create a hello world app framework:

```
mn create-app hello-world
```

```
./gradlew run
```

You should be able to hit the endpoint.

```
```

Now let's build it:

```
./gradlew assemble
```

We can run it as a jar:

```
java -jar build/libs/hello-world-0.1-all.jar
```

Now we can try running it via ops:

```
ops pkg load eyberg/java:20.0.1 -c config.json -p 8080
```

```
curl -XGET http://127.0.0.1:8080/
```
