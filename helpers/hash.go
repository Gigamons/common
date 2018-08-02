package helpers

import (
	"crypto/md5"
	"io/ioutil"
)

// MD5File reads a File and generates a md5 hash out of it.
func MD5File(Path string) ([]byte, error) {
	b, err := ioutil.ReadFile(Path)
	if err != nil {
		return nil, err
	}
	return MD5(b)
}

// MD5String Converts a String to a MD5 Hash.
func MD5String(String string) ([]byte, error) {
	return MD5([]byte(String))
}

// MD5 Converts a []byte Array to a MD5 Hash.
func MD5(I []byte) (out []byte, err error) {
	hash := md5.New()
	_, err = hash.Write(I)
	out = hash.Sum(nil)
	return
}

// https://stackoverflow.com/questions/15311969/checking-the-equality-of-two-slices i was to lazy, so i've used this.

// Verify to verify if both byte arrays are the same.
func Verify(a, b []byte) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
