package signal

import (
	"os"
	"os/signal"
	"syscall"
)

func Free() <-chan os.Signal {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	return signalChan
}
