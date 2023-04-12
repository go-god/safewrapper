package safewrapper

import "log"

// Option wrapper option
type Options struct {
	RecoveryFunc func()
}

// Options optional function
type Option func(o *Options)

// WithRecover set recover func
func WithRecover(recoveryFunc func()) Option {
	return func(o *Options) {
		o.RecoveryFunc = recoveryFunc
	}
}

// defaultRecovery default recover func.
func defaultRecovery() {
	if e := recover(); e != nil {
		log.Printf("wrapper exec recover:%v\n", e)
	}
}
