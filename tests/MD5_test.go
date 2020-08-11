package golangtest

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	var h1 = md5.New()
	var s = h1.Sum([]byte{'1'})
	s = h1.Sum([]byte{1, 2, 3, 4, 5})
	fmt.Printf("Sum: %s\n", hex.EncodeToString(s))
	s = h1.Sum(nil)
	fmt.Printf("Sum: %s\n", hex.EncodeToString(s))

	var h2 = md5.New()
	h2.Write([]byte{'1'})

	s = h2.Sum([]byte{'2'})
	fmt.Printf("Sum: %s\n", hex.EncodeToString(s))
	s = h2.Sum(nil)
	fmt.Printf("Sum: %s\n", hex.EncodeToString(s))

	var h3 = md5.New()
	h3.Write([]byte{'1', '2'})
	s = h3.Sum(nil)
	fmt.Printf("Sum: %s\n", hex.EncodeToString(s))
}
