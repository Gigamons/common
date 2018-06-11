package helpers

import "testing"

func TestDownload(t *testing.T) {
	file, err := Download("https://google.com")
	if err != nil {
		t.Error(err)
	}
	if len(file) < 64 {
		t.Error("Could not verify valid file.")
	}
	t.Log(string(file[:15]))
}
