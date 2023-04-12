package safewrapper

// Go for go func safely
// exec go func with recovery func
func Go(fn func(), recoveryFunc ...func(r interface{})) {
	var recoveryHandler = defaultRecoveryFunc
	if len(recoveryFunc) > 0 && recoveryFunc[0] != nil {
		recoveryHandler = recoveryFunc[0]
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				recoveryHandler(r)
			}
		}()

		fn()
	}()
}
