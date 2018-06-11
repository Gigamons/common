package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/url"
)

// DownloadBeatmapbyName Downloads a Beatmap from OSU!
func DownloadBeatmapbyName(BeatmapName string) (string, error) {
	uri := url.URL{Host: "osu.ppy.sh", Path: fmt.Sprintf("osu/%s", BeatmapName)}
	b, err := Download("http:" + uri.String())

	md5sum := md5.New()
	md5sum.Write(b)

	fhash := fmt.Sprintf("data/map/%s.osu", hex.EncodeToString(md5sum.Sum(nil)))
	err = ioutil.WriteFile(fhash, b, 0644)
	if err != nil {
		return "", err
	}
	return fhash, nil
}
