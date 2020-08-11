package golangtest

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type RandomString struct {
	allowedChars string
	length       int
	rand         *rand.Rand
}

const (
	RandomStringAllowedChars = `123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz`
)

func NewRandomString(length int, allowedChars string) *RandomString {
	var r RandomString
	r.length = length
	r.allowedChars = allowedChars
	r.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	return &r
}

func (r *RandomString) NextString() string {
	var randomChars []byte
	var allowedLength = len(r.allowedChars)
	for i := 0; i < r.length; i++ {
		var x = r.NextNum()
		randomChars = append(randomChars, r.allowedChars[x%allowedLength])
	}
	return string(randomChars)
}

func (r *RandomString) NextNum() int {
	return r.rand.Int()
}

func TestRand(t *testing.T) {
	var r = NewRandomString(32, RandomStringAllowedChars)
	var n = 1000
	var result []string
	for i := 0; i < n; i++ {
		var s = r.NextString()
		result = append(result, s)
	}
	fmt.Printf("result: %s\n", result)
}
