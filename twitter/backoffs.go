package twitter

import (
	"time"

	"github.com/cenkalti/backoff"
)

func newExponentialBackOff() *backoff.ExponentialBackOff {
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = 5 * time.Second
	b.Multiplier = 2.0
	b.MaxInterval = 320 * time.Second
	b.Reset()
	return b
}

func newAggressiveExponentialBackOff() *backoff.ExponentialBackOff {
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = 1 * time.Minute
	b.Multiplier = 2.0
	b.MaxInterval = 16 * time.Minute
	b.Reset()
	return b
}

func newBackoffWithMaxRetries(b backoff.BackOff, maxRetries uint64) backoff.BackOff {
	return backoff.WithMaxRetries(b, maxRetries)
}
