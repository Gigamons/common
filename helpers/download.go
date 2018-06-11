package helpers

import (
	"io/ioutil"
	"net/http"
)

// Download is downloading the content and returns the body. mostly used for downloading a file.
func Download(URL string) (file []byte, err error) {
	response, err := http.Get(URL)
	if err != nil {
		return
	}
	defer response.Body.Close()
	file, err = ioutil.ReadAll(response.Body)
	return
}
