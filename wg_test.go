package safewrapper

import (
	"log"
	"testing"
)

func TestNewWaitGroup(t *testing.T) {
	wg := NewWaitGroup()
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

	wg.Wrap(func() {
		log.Println("hello")
	})

	wg.Wait()
}

/*
=== RUN   TestNewWaitGroup
2023/04/12 19:05:33 hello
2023/04/12 19:05:33 wrapper exec recover:exec panic
2023/04/12 19:05:33 wrapper exec recover:exec panic
2023/04/12 19:05:33 wrapper exec recover:exec panic
2023/04/12 19:05:33 wrapper exec recover:exec panic
2023/04/12 19:05:33 current index:  3
2023/04/12 19:05:33 wrapper exec recover:exec panic
2023/04/12 19:05:33 current index:  9
2023/04/12 19:05:33 current index:  7
2023/04/12 19:05:33 current index:  5
2023/04/12 19:05:33 current index:  1
--- PASS: TestNewWaitGroup (0.00s)
PASS
*/
