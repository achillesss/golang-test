package golangtest

import "testing"

type f func()

func TestFunction(t *testing.T) {
	var x = f(func() {
		print("x\n")
	})

	x()
}
