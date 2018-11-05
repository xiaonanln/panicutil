package panicutil

import (
	"context"
	"log"
)

func RecoverPanic(f func()) (err interface{}) {
	defer func() {
		err = recover()
	}()

	f()
	return
}

func RunForever(ctx context.Context, f func()) {
	for {
		if ctx.Err() != nil {
			break
		}

		err := RecoverPanic(f)
		log.Printf("RunForever: panic: %v", err)
	}
}
