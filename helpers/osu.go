package helpers

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func DownloadBeatmapbyName(BeatmapName string) (string, error) {
	uri := url.URL{Host: "osu.ppy.sh", Path: fmt.Sprintf("osu/%s", BeatmapName)}
	h, err := http.Get("http:" + uri.String())
	if err != nil {
		return "", err
	}
	defer h.Body.Close()

	b, err := ioutil.ReadAll(h.Body)
	if err != nil {
		return "", err
	}

	md5sum := md5.New()
	md5sum.Write(b)

	fhash := fmt.Sprintf("data/map/%x.osu", md5sum.Sum(nil))
	err = ioutil.WriteFile(fhash, b, 0644)
	if err != nil {
		return "", err
	}
	return fhash, nil
}
