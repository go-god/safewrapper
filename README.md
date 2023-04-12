# safewrapper
wrap go func and waitgroup exec goroutine In a safe way.
# Why encapsulate this package? 
- if panic occurs during the execution of go func,the entire program may exit if the current program does not recover more than once.
- Because goroutine cannot capture panic across coroutines, panic must be handled well in the current execution of func.
- This approach is `a defensive approach to programming`.

# How do you use it?
```go
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
	// compatible with the go library's waitGroup operations
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
```

