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
	})

	<-done
	log.Println("ok")
}

/*
=== RUN   TestGo
2023/04/12 19:05:15 abc
2023/04/12 19:05:15 wrapper exec recover:hello
2023/04/12 19:05:15 ok
--- PASS: TestGo (0.00s)
PASS
*/
