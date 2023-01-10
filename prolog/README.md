Create a stand-alone application:
```
swipl --stand_alone=true -o myexe -g hello -c hello.pl
```

Cp over the swi-prolog folder:

```
mkdir -p usr/lib/swi-prolog
cp -R /usr/lib/swi-prolog/* usr/lib/swi-prolog/.
```

```
ops run -c config.json myexe
```
