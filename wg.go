package safewrapper

import (
	"sync"
)

type wrapWgImpl struct {
	wg           sync.WaitGroup
	recoveryFunc func()
}

// NewWaitGroup create safe waitGroup
func NewWaitGroup(opts ...Option) *wrapWgImpl {
	w := &wrapWgImpl{}
	option := &Options{}
	for _, o := range opts {
		o(option)
	}

	w.recoveryFunc = option.RecoveryFunc
	if w.recoveryFunc == nil {
		w.recoveryFunc = defaultRecovery
	}

	return w
}

// Wrap fn func in goroutine to run
// Note: panic is not caught here, because we know for sure that panic will not be thrown,
// so we don't need to catch processing
func (w *wrapWgImpl) Wrap(fn func()) {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()

		fn()
	}()
}

// WrapWithRecover exec func with recover
func (w *wrapWgImpl) WrapWithRecover(fn func()) {
	w.wg.Add(1)
	go func() {
		defer w.recoveryFunc()
		defer w.wg.Done()

		fn()
	}()
}

// Wait this func wait for a set of goroutines to complete execution
func (w *wrapWgImpl) Wait() {
	w.wg.Wait()
}

// Done wg counter minus 1
func (w *wrapWgImpl) Done() {
	w.wg.Done()
}
