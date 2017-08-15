package main

import (
	"os"
	"strings"

	"github.com/mainflux/message-writer/cassandra"
)

const (
	sep         string = ","
	defCluster  string = "127.0.0.1"
	defKeyspace string = "message_writer"
	envCluster  string = "MESSAGE_WRITER_DB_CLUSTER"
	envKeyspace string = "MESSAGE_WRITER_DB_KEYSPACE"
)

type config struct {
	Cluster  string
	Keyspace string
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func main() {
	cfg := config{
		Cluster:  getenv(envCluster, defCluster),
		Keyspace: getenv(envKeyspace, defKeyspace),
	}

	hosts := strings.Split(cfg.Cluster, sep)
	session, err := cassandra.Connect(hosts, cfg.Keyspace)
	if err != nil {
		os.Exit(1)
	}
	defer session.Close()

	if err = cassandra.Initialize(session); err != nil {
		os.Exit(1)
	}
}
