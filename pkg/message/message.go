package message

import (
	"crypto/sha256"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type MessageBus interface {
	Subscribe() chan Request
	Publish(msg Request)
}

type Request interface {
	GetPayload() interface{}
	GetRequestor() string
	GetTopic() string
	Hash() string
	String() string
}

type Topic interface {
	String() string
}

type Message struct {
	Topic     Topic
	Requestor string
	Payload   interface{}
}

func (m Message) GetTopic() string {
	return m.Topic.String()
}
func (m Message) GetRequestor() string {
	return m.Requestor
}
func (m Message) GetPayload() interface{} {
	if m.Payload == nil {
		log.Error("message was malformed")
		return errors.New("message was malformed")
	}
	return m.Payload
}

func (m Message) Hash() string {
	h := sha256.New()
	h.Write([]byte(m.String()))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (m Message) String() string {
	return fmt.Sprintf("%#v", m)
}
