package helpers

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"time"
)

// Pseudorandombytes Generates a pseudo byte array. good for crypto.
func Pseudorandombytes(n int) []byte {
	out := bytes.NewBuffer(nil)
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		logic := rand.Int63n(time.Now().Unix() + time.Now().UnixNano())
		le := logic&^time.Now().Unix() | logic
		be := logic | time.Now().UnixNano() | logic%le
		binary.Write(out, binary.LittleEndian, le)
		binary.Write(out, binary.BigEndian, be)
	}
	return out.Bytes()[:n]
}

// RandomString is just generate a random string. aA-zZ 0-9
func RandomString(n int) (ret string) {
	runeArr := []rune("aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ123456789")
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		ret += string(runeArr[rand.Intn(len(runeArr))])
	}
	return
}
