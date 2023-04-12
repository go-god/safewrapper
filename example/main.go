package main

import (
	"log"
	"runtime/debug"

	"github.com/go-god/safewrapper"
)

func main() {
	// safe go func exec
	done := make(chan struct{}, 1)
	safewrapper.Go(func() {
		defer close(done)

		log.Println("abc")
		panic("hello")
	}, func(r interface{}) {
		log.Println("exec recover:", r)
		log.Printf("full stack:%s\n", string(debug.Stack()))
	})

	<-done
	log.Println("ok")

	// safe waitGroup to wait some goroutine to finish
	var wg safewrapper.WaitGroup
	for i := 0; i < 10; i++ {
		// Note: i must be a duplicate, otherwise it will always be a value when wg executes func,
		// so it re-uses the index variable
		var index = i
		wg.WrapWithRecover(func() {
			if index%2 == 0 {
				panic("exec panic")
			}

			log.Println("current index: ", index)
		}, func(r interface{}) {
			log.Printf("current index:%d exec panic: %v", index, r)
		})
	}

	// Note: panic is not caught here, because we know for sure that panic will not be thrown,
	// so we don't need to catch processing
	wg.Wrap(func() {
		log.Println("hello")
	})

	wg.Wait()
}
