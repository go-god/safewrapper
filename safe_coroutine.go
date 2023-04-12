package safewrapper

// Go for go func safely
func Go(fn func(), opts ...Option) {
	option := &Options{}
	for _, o := range opts {
		o(option)
	}

	if option.RecoveryFunc == nil {
		option.RecoveryFunc = defaultRecovery
	}

	go func() {
		defer option.RecoveryFunc()

		fn()
	}()
}
