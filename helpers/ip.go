package helpers

import (
	"strconv"
	"strings"

	"github.com/Gigamons/common/logger"

	"github.com/Gigamons/common/consts"
	"github.com/pquerna/ffjson/ffjson"
)

// We're going to use http://ip.zxq.co/, because why not ?
// GetIPInfo is gonna get the IP Information where the User is located at.
func GetIPInfo(ip string) *consts.GeoIP {
	GeoIP := consts.GeoIP{}
	f, err := Download("http://ip.zxq.co/" + ip)
	if err != nil {
		logger.Errorln(err)
		return nil
	}
	ffjson.Unmarshal(f, &GeoIP)

	Location := strings.Split(GeoIP.LocRaw, ",")
	if len(Location) > 1 {
		GeoIP.Location.Lat, err = strconv.ParseFloat(Location[0], 64) // We don't care it it throws an err
		GeoIP.Location.Lon, err = strconv.ParseFloat(Location[1], 64)
	}
	return &GeoIP
}
