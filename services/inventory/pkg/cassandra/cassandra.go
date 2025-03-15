package cassandra

import (
	"github.com/gocql/gocql"
	"strconv"
)

func NewClient(user string, password string, host string, port string, keyspace string) (*gocql.Session, error) {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	cluster := gocql.NewCluster(host)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: user,
		Password: password,
	}
	cluster.Port = portInt
	cluster.Keyspace = keyspace

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}
