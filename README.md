Usage
-----

Save a `docker run` command:

```
dr --name saved -i -t -v $(pwd):/app -p 50015:50015 clusterflunk/clusterflunk pserve development
```

Run a saved `docker run` command:

```
dr saved
```