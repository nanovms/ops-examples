GlassFish Unikernel Hello World

```
wget https://mirrors.jevincanders.net/eclipse/ee4j/glassfish/glassfish-7.0.25.zip
unzip glassfish-7.0.25.zip
mv glassfish7/* .
rmdir glassfish7
```

```
ops pkg load eyberg/java:21.0.4 -c config.json  -p 8080
```
