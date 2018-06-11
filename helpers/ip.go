package helpers

import (
	"strconv"
	"strings"

	"github.com/Gigamons/common/consts"
	"github.com/pquerna/ffjson/ffjson"
)

var jsondecoder = ffjson.NewDecoder()

// We're going to use http://ip.zxq.co/, because why not ?

// GetIPInfo is gonna get the IP Information where the User is located at.
func GetIPInfo() *consts.GeoIP {
	GeoIP := consts.GeoIP{}
	f, err := Download("http://ip.zxq.co/")
	if err != nil {
		return nil
	}
	jsondecoder.DecodeFast(f, &GeoIP)

	Location := strings.Split(GeoIP.LocRaw, ",")
	if len(Location) > 1 {
		GeoIP.Location.Lat, err = strconv.ParseFloat(Location[0], 64) // We don't care it it throws an err
		GeoIP.Location.Lon, err = strconv.ParseFloat(Location[1], 64)
	}
	return &GeoIP
}
