package mq

import (
	"github.com/Unknwon/goconfig"
	stomp "github.com/go-stomp/stomp"
)

var address string

func init() {
	cfg, err := goconfig.LoadConfigFile("../../conf.ini")
	if err != nil {
		panic(err)
	}

	value, _ := cfg.GetValue("mq", "address")
	address = value
}

type ActiveMQ struct {
	Addr string
}

//New activeMQ with addr.
func NewActiveMQ() *ActiveMQ {
	return &ActiveMQ{address}
}

// Used for health check
func (this *ActiveMQ) Check() error {
	conn, err := this.Connect()
	if err == nil {
		defer conn.Disconnect()
		return nil
	} else {
		return err
	}
}

// Connect to activeMQ
func (this *ActiveMQ) Connect() (*stomp.Conn, error) {
	return stomp.Dial("tcp", this.Addr)
}

// Send msg to destination
func (this *ActiveMQ) Send(destination string, msg string) error {
	conn, err := this.Connect()
	if err != nil {
		panic(err)
	}
	defer conn.Disconnect()
	return conn.Send(
		destination,  // destination
		"text/plain", // content-type
		[]byte(msg))  // body
}

// Subscribe Message from destination
// func handler handle msg reveived from destination
func (this *ActiveMQ) Subscribe(destination string, handler func(err error, msg string)) error {

	conn, err := this.Connect()
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
	return err
}
