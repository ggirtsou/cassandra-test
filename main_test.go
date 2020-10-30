package main

import (
	"fmt"
	"testing"

	"github.com/gocql/gocql"
)

func TestCassandraConnect(t *testing.T) {
	cluster := gocql.NewCluster("172.16.239.12")
	cluster.Port = 9042
	cluster.CQLVersion = "3.4.4"
	cluster.ProtoVersion = 4
	cluster.Keyspace = "system" // that keyspace is not modifiable by user
	cluster.Consistency = gocql.Quorum
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "cassandra",
		Password: "cassandra",
	}

	_, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra init done")
}
