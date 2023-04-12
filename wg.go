package safewrapper

import (
	"sync"
)

// WaitGroup wrap sync.WaitGroup
// Compatible with the go library's waitGroup operations
type WaitGroup struct {
	wg sync.WaitGroup
}

// Wrap fn func in goroutine to run,It's relatively unsafe.
// Note: panic is not caught here, because we know for sure that panic will not be thrown,
// so we don't need to catch processing
func (w *WaitGroup) Wrap(fn func()) {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()

		fn()
	}()
}

// WrapWithRecover exec func with recover safely
func (w *WaitGroup) WrapWithRecover(fn func(), recoveryFunc ...func(r interface{})) {
	var recoveryHandler = defaultRecoveryFunc
	if len(recoveryFunc) > 0 && recoveryFunc[0] != nil {
		recoveryHandler = recoveryFunc[0]
	}

	w.wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				recoveryHandler(r)
			}
		}()
		defer w.wg.Done()

		fn()
	}()
}

// Wait this func wait for a set of goroutines to complete execution
func (w *WaitGroup) Wait() {
	w.wg.Wait()
}

// Done wg delta counter minus 1
func (w *WaitGroup) Done() {
	w.wg.Done()
}

// Add delta+1
func (w *WaitGroup) Add(delta int) {
	w.wg.Add(delta)
}
