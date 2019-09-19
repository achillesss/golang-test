package golangtest

import (
	"fmt"
	"testing"
	"time"
)

func receiveSignal(done chan struct{}, signal chan int) {
	for {
		select {
		case <-done:
			fmt.Printf("signal lenght: %d\n", len(signal))
			if len(signal) == 0 {
				return
			}

		case n := <-signal:
			fmt.Printf("receive: %d\n", n)
		}
	}
}

func TestChan(t *testing.T) {
	var done = make(chan struct{})
	var signal = make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			signal <- i
		}
	}()

	go receiveSignal(done, signal)
	go close(done)

	time.Sleep(time.Second * 10)
}
