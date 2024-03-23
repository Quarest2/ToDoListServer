package pg_storage

import (
	log "github.com/sirupsen/logrus"
	"runtime"
)

func ErrorHandler(err error, message string) {
	if err == nil {
		return
	}

	_, file, no, ok := runtime.Caller(1)

	if ok {
		log.Printf("Error (%s) in %s:%d: %s\n", err.Error(), file, no, message)
	} else {
		log.Printf("Error (%s): %s\n", err.Error(), message)
	}
}
