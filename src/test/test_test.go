package test

import (
	"go-web/src/godis"
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
