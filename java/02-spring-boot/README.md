# Spring Boot Hello World

This example shows how to run a simple spring boot example with ops.

To package and deploy run the following command.

Verify you can build the jar and run it:

```
mvn dependency:tree
mvn spring-boot:run
```

```
$ ops pkg load java_1.8.0_191 -p 8080 -c config.json
```
