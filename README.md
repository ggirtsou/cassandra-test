```bash
docker build -t test-cassandra .
docker-compose up -d # wait for cassandra to settle and then observe logs in tests container
```

Troubleshooting

```
$ docker run --rm -it --network cassandra-test_app_net test-cassandra bash
$ nc -vz cassandra.test-cassandra_app_net 9042

# note the values in the config: https://github.com/ggirtsou/cassandra-test/blob/master/dev/cassandra.yaml#L612
# broadcast address: https://github.com/ggirtsou/cassandra-test/blob/master/dev/cassandra.yaml#L626

# or try with cassandra.test-cassandra_app_net, cassandra.docker.internal, cassandra as host
$ cqlsh 172.16.239.12 9042 cassandra -u cassandra -p cassandra

# see if Go can connect to Cassandra (has IP hardcoded in test file)
$ GO111MODULE=on CGO_ENABLED=0 go test -v ./...
```
From host:
```
# another thing you can try is to install cqlsh / DataGrip locally and try to connect
cqlsh 127.0.0.1 9042 cassandra -u cassandra -p cassandra --cqlversion="3.4.4" # run from host - this works!

# I also tried using an older versions of gocql as suggested in various comments here: https://github.com/gocql/gocql/issues/946
```

Docker engine version: `v19.03.13` 
OS: Mac
