package safewrapper

import (
	"log"
	"testing"
)

func TestNewWaitGroup(t *testing.T) {
	wg := WaitGroup{}
	for i := 0; i < 10; i++ {
		// Note: i must be a duplicate, otherwise it will always be a value when wg executes func,
		// so it re-uses the index variable
		var index = i
		wg.WrapWithRecover(func() {
			if index%2 == 0 {
				panic("mock panic")
			}

			log.Println("current index: ", index)
		}, func(r interface{}) {
			log.Printf("current index:%d exec panic: %v", index, r)
		})
	}

	wg.Wrap(func() {
		log.Println("hello")
	})

	wg.Wait()
}

/*
=== RUN   TestNewWaitGroup
2023/04/12 20:39:00 hello
2023/04/12 20:39:00 current index:  3
2023/04/12 20:39:00 current index:2 exec panic: mock panic
2023/04/12 20:39:00 current index:  7
2023/04/12 20:39:00 current index:  5
2023/04/12 20:39:00 current index:0 exec panic: mock panic
2023/04/12 20:39:00 current index:4 exec panic: mock panic
2023/04/12 20:39:00 current index:6 exec panic: mock panic
2023/04/12 20:39:00 current index:  9
2023/04/12 20:39:00 current index:8 exec panic: mock panic
2023/04/12 20:39:00 current index:  1
--- PASS: TestNewWaitGroup (0.00s)
PASS
*/
