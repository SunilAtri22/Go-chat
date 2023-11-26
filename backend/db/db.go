package db

import (
	"github.com/gocql/gocql"
)

type Database struct {
	session *gocql.Session
}

func NewDatabase() (*Database, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "chat_app"
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return &Database{session: session}, nil
}

func (d *Database) Close() {
	d.session.Close()
}

func (d *Database) GetSession() *gocql.Session {
	return d.session
}
