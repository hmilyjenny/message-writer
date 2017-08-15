package cassandra

import "github.com/gocql/gocql"

var tables []string = []string{
	`CREATE TABLE IF NOT EXISTS messages_by_channel (
		channel timeuuid,
		id timeuuid,
		content_type text,
		publisher text,
		protocol text,
		payload blob,
		PRIMARY KEY ((channel), id)
	) WITH default_time_to_live = 86400`,
}

// Connect establishes connection to the Cassandra cluster.
func Connect(hosts []string, keyspace string) (*gocql.Session, error) {
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum

	return cluster.CreateSession()
}

// Initialize creates tables used by the service.
func Initialize(session *gocql.Session) error {
	for _, table := range tables {
		if err := session.Query(table).Exec(); err != nil {
			return err
		}
	}

	return nil
}
