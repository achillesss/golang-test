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

func testCloseChan(ch chan struct{}) {
	go func() {
		select {
		case signal := <-ch:
			fmt.Printf("select receive close signal: %v\n", signal)
		}
	}()

	go func() {
		for signal := range ch {
			fmt.Printf("for range receive close signal: %v\n", signal)
		}
	}()

	go func() {
		var signal = <-ch
		fmt.Printf("normal receive close signal: %v\n", signal)
	}()

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

	var closeTest = make(chan struct{})
	testCloseChan(closeTest)
	close(closeTest)

	time.Sleep(time.Second * 10)
}
