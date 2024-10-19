package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

type NatsPublisher struct {
	conn *nats.Conn
}

// NewNatsPublisher creates new NatsPublisher
func NewNatsPublisher(natsURL string) (*NatsPublisher, error) {
	conn, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}
	return &NatsPublisher{conn: conn}, nil
}

// Publish sends a message to NATS
func (n *NatsPublisher) Publish(subject string, data []byte) error {
	err := n.conn.Publish(subject, data)
	if err != nil {
		log.Printf("Error publishing message to NATS: %v", err)
		return err
	}
	log.Printf("Message published to NATS successfully on subject: %s", subject)
	return nil
}

// Close closes the NATS connection
func (n *NatsPublisher) Close() {
	if n.conn != nil {
		n.conn.Close()
		log.Println("NATS connection closed")
	}
}
