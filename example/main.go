package main

import (
	"log"

	"github.com/go-god/safewrapper"
)

func main() {
	// safe go func exec
	done := make(chan struct{}, 1)
	safewrapper.Go(func() {
		defer close(done)

		log.Println("abc")
		panic("hello")
	})

	<-done
	log.Println("ok")

	// safe waitGroup to wait some goroutine to finish
	wg := safewrapper.NewWaitGroup()
	for i := 0; i < 10; i++ {
		// Note: i must be a duplicate, otherwise it will always be a value when wg executes func,
		// so it re-uses the index variable
		var index = i
		wg.WrapWithRecover(func() {
			if index%2 == 0 {
				panic("exec panic")
			}

			log.Println("current index: ", index)
		})
	}

	// Note: panic is not caught here, because we know for sure that panic will not be thrown,
	// so we don't need to catch processing
	wg.Wrap(func() {
		log.Println("hello")
	})

	wg.Wait()
}
