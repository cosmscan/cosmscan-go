package signal

import (
	"cosmscan-go/pkg/log"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

type SignalReceiver func(signal os.Signal)

// Handler handles os.Signal
// SIGINT & SIGTERM signals stop the process.
// SIGQUIT signal dumps stack to log and stop the process
type Handler struct {
	log       log.Interface
	receivers []SignalReceiver
	quit      chan struct{}
}

func NewHandler(log log.Interface, receivers ...SignalReceiver) *Handler {
	return &Handler{
		log:       log,
		receivers: receivers,
		quit:      make(chan struct{}),
	}
}

func (h *Handler) Stop() {
	close(h.quit)
}

// Loop creates a loop to receive os.Signal
func (h *Handler) Loop() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer signal.Stop(signals)
	buf := make([]byte, 1<<20)

	for {
		select {
		case <-h.quit:
			h.log.Log("msg", "signal handler stopped")
			return
		case sig := <-signals:
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM:
				h.log.Log("msg", "signal received, stopping", "signal", sig)
				for _, r := range h.receivers {
					r(sig)
				}
				return
			case syscall.SIGQUIT:
				n := runtime.Stack(buf, true)
				h.log.Log("msg", "received SIGQUIT, dumping stack")
				fmt.Printf("=== dump ===\n\n%s", buf[:n])
				return
			}
		}
	}
}
