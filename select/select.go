package pselect

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	aDuration := measureDuration(a)
	bDuration := measureDuration(b)

	if aDuration < bDuration {
		winner = a
	} else {
		winner = b
	}

	return
}

func measureDuration(url string) time.Duration {
	statr := time.Now()
	http.Get(url)
	return time.Since(statr)
}
