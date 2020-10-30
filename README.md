## Problem

Could not connect to Cassandra from another container. `nc` reported port open and Cassandra was up and running in its container. Both containers were running in the same user defined network.

## Solution

In [gocql](https://github.com/gocql/gocql/), do not define aggressive timeouts. Removed

```go
cluster.Timeout = 300 // in ms
cluster.ConnectTimeout = 300 // in ms
```

You can connect to Cassandra from host using 127.0.0.1 9042 with `cqlsh` or [DataGrip](https://www.jetbrains.com/datagrip/).

## Setup

```bash
docker build -t test-cassandra .
docker-compose up -d # wait for cassandra to settle and then observe logs in tests container
```

## Troubleshooting

```bash
$ docker run --rm -it --network cassandra-test_app_net test-cassandra bash
$ nc -vz cassandra 9042

# note the values in the config: https://github.com/ggirtsou/cassandra-test/blob/master/dev/cassandra.yaml#L612
# broadcast address: https://github.com/ggirtsou/cassandra-test/blob/master/dev/cassandra.yaml#L626

# or try with cassandra.test-cassandra_app_net, cassandra.docker.internal, cassandra as host
$ cqlsh 172.16.239.12 9042 cassandra -u cassandra -p cassandra

# see if Go can connect to Cassandra (has IP hardcoded in test file)
$ GO111MODULE=on CGO_ENABLED=0 go test -v ./...
```
From host:
```bash
# another thing you can try is to install cqlsh / DataGrip locally and try to connect
$ cqlsh 127.0.0.1 9042 cassandra -u cassandra -p cassandra --cqlversion="3.4.4" # run from host - this works!
```

`9042` listens on IPv6, [disabling TCP6](https://cassandra.apache.org/doc/latest/configuration/cassandra_config_file.html#listen-interface-prefer-ipv6) in Cassandra doesn't have any effect on this.
```
root@f96fd59f5db5:/# netstat -ltnp | grep -w ':9042'
tcp6       0      0 :::9042                 :::*                    LISTEN      -
```

* Docker engine version: `v19.03.13`
* OS: Mac

## References

* I also tried using older versions of gocql as suggested in various comments here: https://github.com/gocql/gocql/issues/946
* https://github.com/gocql/gocql/issues/997
* https://stackoverflow.com/questions/32072680/cassandra-is-forcing-to-open-cql-native-port-on-ipv6-interface
* https://github.com/gocql/gocql/issues/997
* https://stackoverflow.com/questions/41373122/getting-cassandra-to-listen-on-9042-properly
* https://groups.google.com/a/lists.datastax.com/g/java-driver-user/c/TERfynufkfg
* https://support.datastax.com/hc/en-us/articles/207433496-Unable-to-connect-to-DSE-via-native-protocol-over-IPv6
* https://cassandra.apache.org/doc/latest/configuration/cassandra_config_file.html#listen-interface-prefer-ipv6
* https://github.com/gocql/gocql/issues/30
