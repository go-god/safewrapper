package safewrapper

import (
	"log"
	"testing"
)

func TestGo(t *testing.T) {
	done := make(chan struct{}, 1)
	Go(func() {
		defer close(done)

		log.Println("abc")
		panic("hello")
	}, func(r interface{}) {
		log.Println("exec recover:", r)
	})

	<-done
	log.Println("ok")
}

/*
=== RUN   TestGo
2023/04/12 20:39:26 abc
2023/04/12 20:39:26 exec recover: hello
2023/04/12 20:39:26 ok
--- PASS: TestGo (0.00s)
PASS
*/
