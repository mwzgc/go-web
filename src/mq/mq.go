package mq

import (
	"go-web/src/config"

	stomp "github.com/go-stomp/stomp"
)

var address string

func init() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	value, _ := cfg.GetValue("mq", "address")
	address = value
}

// ActiveMQ ...
type ActiveMQ struct {
	Addr string
}

// NewActiveMQ with addr ...
func NewActiveMQ() *ActiveMQ {
	return &ActiveMQ{address}
}

// Check ...
func (mq *ActiveMQ) Check() error {
	conn, err := mq.Connect()
	if err == nil {
		defer conn.Disconnect()
		return nil
	}

	return err
}

// Connect ...
func (mq *ActiveMQ) Connect() (*stomp.Conn, error) {
	return stomp.Dial("tcp", mq.Addr)
}

// Send ...
func (mq *ActiveMQ) Send(destination string, msg string) error {
	conn, err := mq.Connect()
	if err != nil {
		panic(err)
	}
	defer conn.Disconnect()
	return conn.Send(
		destination,  // destination
		"text/plain", // content-type
		[]byte(msg))  // body
}

// Subscribe ...
func (mq *ActiveMQ) Subscribe(destination string, handler func(err error, msg string)) error {
	conn, err := mq.Connect()
	if err != nil {
		panic(err)
	}

	sub, err := conn.Subscribe(destination, stomp.AckAuto)
	if err != nil {
		return err
	}
	defer conn.Disconnect()
	defer sub.Unsubscribe()
	for {
		m := <-sub.C
		handler(m.Err, string(m.Body))
	}
	// return err
}
