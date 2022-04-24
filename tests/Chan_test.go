package golangtest

import (
	"context"
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

	time.Sleep(time.Second)
}

type pauseTest struct {
	ctx         context.Context
	cancel      context.CancelFunc
	pauseChan   chan struct{}
	restoreChan chan struct{}
}

func (p *pauseTest) run() {
	for {
		select {
		case <-p.ctx.Done():
			fmt.Printf("exit\n")
			return

		case <-p.pauseChan:
			fmt.Printf("paused.\n")
			select {
			case <-p.ctx.Done():
			case <-p.restoreChan:
				fmt.Printf("restored.\n")
			}

		default:
			fmt.Printf("work.\n")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func (p *pauseTest) pause() {
	fmt.Printf("PAUSE\n")
	select {
	case <-p.pauseChan:
		fmt.Printf("already paused.\n")
	default:
		close(p.pauseChan)
		p.restoreChan = make(chan struct{})
	}
}

func (p *pauseTest) restore() {
	fmt.Printf("RESOTRE\n")
	select {
	case <-p.restoreChan:
		fmt.Printf("already stored.\n")
	default:
		close(p.restoreChan)
		p.pauseChan = make(chan struct{})
	}
}

func TestChanPause(t *testing.T) {
	var p pauseTest
	p.ctx, p.cancel = context.WithCancel(context.Background())
	p.restoreChan = make(chan struct{})
	p.pauseChan = make(chan struct{})
	go p.run()
	p.restore()
	time.Sleep(time.Second)
	p.pause()
	time.Sleep(time.Second)
	p.restore()
	time.Sleep(time.Second)
	p.pause()
	time.Sleep(time.Second)
	p.cancel()
	time.Sleep(time.Second)
}
