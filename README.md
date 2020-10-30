```bash
docker build -t test-cassandra .
docker-compose up -d # wait for cassandra to settle and then observe logs in tests container
```

Troubleshooting

```
$ docker run --rm -it --network test-cassandra_app_net test-cassandra bash
$ nc -vz cassandra.test-cassandra_app_net 9042

# or try with 172.16.239.12, cassandra.docker.internal, cassandra as host
$ cqlsh cassandra.test-cassandra_app_net 9042 cassandra -u cassandra -p cassandra

# see if Go can connect to Cassandra (has IP hardcoded in test file)
$ go test -v ./...

# another thing you can try is to install cqlsh / DataGrip locally and try to connect
cqlsh 127.0.0.1 9042 cassandra -u cassandra -p cassandra --cqlversion="3.4.4" # run from host - this works!

# I also tried using an older versions of gocql as suggested in various comments here: https://github.com/gocql/gocql/issues/946
```

Docker engine version: `v19.03.13` 
OS: Mac
