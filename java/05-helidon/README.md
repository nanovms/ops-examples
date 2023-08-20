Helidon is a cloud-native, open‑source set of Java libraries for writing
microservices that run on a fast web core powered by Netty.

Two great features that pair well with unikernels is the fast boot time
and the utilization of Java virtual threads.

You can run a helidon unikernel using the following commands:

```
mvn archetype:generate -DinteractiveMode=false \
    -DarchetypeGroupId=io.helidon.archetypes \
    -DarchetypeArtifactId=helidon-quickstart-se \
    -DarchetypeVersion=1.4.12 \
    -DgroupId=io.helidon.examples \
    -DartifactId=helidon-quickstart-se \
    -Dpackage=io.helidon.examples.quickstart.se
mvn package
ops pkg load eyberg/java:20.0.1 -c config.json -p 8080
```
