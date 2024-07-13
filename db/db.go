package db

import (
	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

var Cluster *gocql.ClusterConfig

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	scyllaHosts := os.Getenv("SCYLLA_HOSTS")
	scyllaKeyspace := os.Getenv("SCYLLA_KEYSPACE")

	Cluster = gocql.NewCluster(strings.Split(scyllaHosts, ",")...)
	Cluster.Keyspace = scyllaKeyspace
	Cluster.Consistency = gocql.Quorum
}
