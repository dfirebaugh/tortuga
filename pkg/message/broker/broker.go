package broker

import "github.com/dfirebaugh/tortuga/pkg/message"

type Broker struct {
	stopCh    chan struct{}
	publishCh chan message.Request
	subCh     chan chan message.Request
	unsubCh   chan chan message.Request
	lastMsg   string
}

func NewBroker() *Broker {
	return &Broker{
		stopCh:    make(chan struct{}),
		publishCh: make(chan message.Request, 1),
		subCh:     make(chan chan message.Request, 1),
		unsubCh:   make(chan chan message.Request, 1),
	}
}

func (b *Broker) Start() {
	subs := map[chan message.Request]struct{}{}
	for {
		select {
		case <-b.stopCh:
			return
		case msgCh := <-b.subCh:
			subs[msgCh] = struct{}{}
		case msgCh := <-b.unsubCh:
			delete(subs, msgCh)
		case msg := <-b.publishCh:
			for msgCh := range subs {
				// msgCh is buffered, use non-blocking send to protect the broker:
				select {
				case msgCh <- msg:
				default:
				}
			}
		}
	}
}

func (b *Broker) Stop() {
	close(b.stopCh)
}

func (b *Broker) Subscribe() chan message.Request {
	msgCh := make(chan message.Request, 5)
	b.subCh <- msgCh
	return msgCh
}

func (b *Broker) Unsubscribe(msgCh chan message.Request) {
	b.unsubCh <- msgCh
}

func (b *Broker) Publish(msg message.Request) {
	if b.lastMsg == msg.Hash() {
		return
	}

	b.lastMsg = msg.Hash()
	b.publishCh <- msg
}
