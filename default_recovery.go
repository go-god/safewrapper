package safewrapper

import "log"

// defaultRecoveryFunc default recover func.
func defaultRecoveryFunc(r interface{}) {
	log.Printf("safe wrapper exec recover:%v\n", r)
}
