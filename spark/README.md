# Apache Spark
[Apache Spark](https://spark.apache.org/) is a unified analytics engine for large-scale data processing.

## Simple Example

### Obtain the Spark Package
Spark can be found in the `ops` package manager:
```
$ ops pkg list

+--------------------+-----------+------------+------------+--------------------------------+
|    PACKAGENAME     |  VERSION  |  LANGUAGE  |  RUNTIME   |          DESCRIPTION           |
+--------------------+-----------+------------+------------+--------------------------------+
| spark_3.0.0        | 3.0.0     | java       | jvm        | spark - unified analytics      |
|                    |           |            |            | engine for large-scale data    |
|                    |           |            |            | processing                     |
+--------------------+-----------+------------+------------+--------------------------------+
```


Start by downloading the `spark_3.0.0` package by running:
```
$ ops pkg get spark_3.0.0
```

It can now be found in the local cache under `~/.ops/packages`. Further details about the package can be obtained by running:
```
$ ops pkg describe spark_3.0.0

Information for spark_3.0.0 package:
# Spark package for the Nanos unikernel

## Configuration

## Running

This has only been tested with 2G ram and up.

$ ops pkg load spark_3.0.0 -p 8080 -p 7077 -c config.json
```

### Launch the Spark Master
This package can be launched by running:
```
$ ops pkg load spark_3.0.0 -p 8080 -p 7077
```

This will start a Spark Master with Web UI at http://0.0.0.0:8080/ and workers can attach to the Spark Master via [spark://0.0.0.0:7077](park://0.0.0.0:7077).

### Attaching a Worker
In order to attach a worker node we first have to download [Apache Spark 3.0.0](https://archive.apache.org/dist/spark/spark-3.0.0/spark-3.0.0-bin-hadoop2.7.tgz).

1. Decompress the downloaded package:
```
$ tar xvzf ~/Downloads/spark-3.0.0-bin-hadoop2.7.tgz -C ~/.
```

2. Start the worker node:
```
$ ~/spark-3.0.0-bin-hadoop2.7/sbin/start-slave.sh spark://0.0.0.0:7077
```

Go to http://0.0.0.0:8080/ to check the status of your Apache Spark cluster.

### Launch Example
There does exist an example application within the downloaded Apache Spark directory from the previous section. This application can be deployed onto the Spark cluster by running:
```
./spark-submit --class org.apache.spark.examples.SparkPi --master spark://0.0.0.0:7077 --executor-memory 1G --total-executor-cores 1 ../examples/jars/spark-examples_2.12-3.0.0.jar 10000
```
