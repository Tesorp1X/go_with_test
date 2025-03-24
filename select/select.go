package pselect

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	statrA := time.Now()
	http.Get(a)
	aDuration := time.Since(statrA)

	statB := time.Now()
	http.Get(b)
	bDuration := time.Since(statB)

	if aDuration < bDuration {
		winner = a
	} else {
		winner = b
	}

	return
}
