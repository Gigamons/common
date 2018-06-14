package helpers

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"time"
)

func Pseudorandombytes(n int64) []byte {
	out := bytes.NewBuffer(nil)
	for i := int64(0); i < n; i++ {
		logic := rand.Int63n(time.Now().Unix() + time.Now().UnixNano())
		le := logic&^time.Now().Unix() | logic%(i+10^logic)
		be := logic | time.Now().UnixNano() | logic%le
		binary.Write(out, binary.LittleEndian, le)
		binary.Write(out, binary.BigEndian, be)
	}
	return out.Bytes()[:n]
}
