package test

import (
	"fmt"
	"go-web/src/godis"
	"go-web/src/mq"
	"testing"
)

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		godis.GetValue("user")
		// var u User
		// err := json.Unmarshal([]byte(uStr), &u)
		// if err != nil {
		// 	panic(err)
		// }

		// fmt.Printf("user = %v", u)
	}
}

func TestSendMq(b *testing.T) {
	mq := mq.NewActiveMQ()
	mq.Send("hello", "hello world")
}

func TestGetMq(b *testing.T) {
	mq := mq.NewActiveMQ()
	mq.Subscribe("hello", func(err error, msg string) {
		fmt.Println(msg)
	})
}
