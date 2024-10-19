package nats

import (
	"github.com/nats-io/nats.go"
	"log"
)

type NatsClient struct {
	conn *nats.Conn
}

// NewNatsClient создает нового клиента NATS.
func NewNatsClient(natsURL string) (*NatsClient, error) {
	conn, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}
	return &NatsClient{conn: conn}, nil
}

// Publish публикует сообщение в заданный NATS subject.
func (n *NatsClient) Publish(subject string, data []byte) error {
	err := n.conn.Publish(subject, data)
	if err != nil {
		log.Printf("Ошибка при отправке сообщения в NATS: %v", err)
		return err
	}
	log.Printf("Сообщение успешно отправлено в NATS: %s", subject)
	return nil
}
