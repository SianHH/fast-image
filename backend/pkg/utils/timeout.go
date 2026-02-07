package utils

import (
	"context"
	"errors"
	"time"
)

func Timeout(timeout time.Duration, do func() error) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- do()
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return errors.New("operation timeout")
	}
}
