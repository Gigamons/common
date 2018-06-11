package helpers

import "os"

// Exists is gonna check if a file exists, (IF return true else false)
func Exists(Path string) bool {
	if _, err := os.Stat(Path); os.IsExist(err) {
		return true
	}
	return false
}

// NotExists is gonna check if a file doesn't exists, (IF Does not return true else false)
func NotExists(Path string) bool {
	if _, err := os.Stat(Path); os.IsNotExist(err) {
		return true
	}
	return false
}
