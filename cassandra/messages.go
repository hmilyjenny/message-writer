package cassandra

import (
	"github.com/gocql/gocql"
	writer "github.com/mainflux/message-writer"
)

var _ writer.MessageRepository = (*msgRepository)(nil)

type msgRepository struct {
	session *gocql.Session
}

// NewMessageRepository instantiates Cassandra message repository.
func NewMessageRepository(session *gocql.Session) writer.MessageRepository {
	return &msgRepository{session}
}

func (repo *msgRepository) Save(msg writer.Message) error {
	cql := `INSERT INTO messages_by_channel (channel, id, content_type, publisher, protocol, payload)
			VALUES (?, now(), ?, ?, ?, ?)`

	return repo.session.Query(cql, msg.Channel, msg.ContentType, msg.Publisher, msg.Protocol, msg.Payload).Exec()
}
