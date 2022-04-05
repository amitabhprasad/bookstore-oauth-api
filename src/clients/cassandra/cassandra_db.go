package cassandra

import (
	"os"

	"github.com/amitabhprasad/bookstore-util-go/logger"
	"github.com/gocql/gocql"
)

var (
	session        *gocql.Session
	cassandra_host = os.Getenv("cassandra_host")
)

func init() {
	// connect to cassandra cluster:
	if cassandra_host == "" {
		cassandra_host = "127.0.0.1"
	}
	log := logger.GetLogger()
	log.Print("Connecting to url " + cassandra_host)
	cluster := gocql.NewCluster(cassandra_host)
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
