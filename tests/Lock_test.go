package golangtest

import (
	"sync"
	"testing"
)

func TestLock(t *testing.T) {
	var l sync.RWMutex
	l.Lock()
	go l.Lock()
	go l.Unlock()
	l.Unlock()
}
