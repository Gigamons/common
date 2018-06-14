package helpers

import (
	"fmt"
	"strconv"

	"github.com/Gigamons/common/consts"
)

// CalculateAccuracy calculates every playModes accuracy.
func CalculateAccuracy(count300 int64, count100 int64, count50 int64, countMiss int64, countGeki int64, countKatu int64, playMode int8) float64 {
	var thp int64
	var th int64
	if count300 == 0 {
		count300 = 1
	}
	switch playMode {
	case consts.STD:
		thp = count50*50 + count100*100 + count300*300
		th = countMiss + count50 + count300 + count100
		return float64(thp) / float64((th * 300))
	case consts.Taiko:
		thp = count50*50 + count300*100
		th = countMiss + count100 + count300
		return float64(thp) / float64(th*100)
	case consts.CTB:
		thp = count300 + count100 + count50
		th = thp + count300 + countKatu
		return float64(thp) / float64(th)
	case consts.Mania:
		thp = count50*50 + count100*100 + countKatu*200 + count300*300 + countGeki*300
		th = countMiss + count50 + count100 + count300 + countGeki + countKatu
		return float64(thp) / (float64(th) * 300)
	default:
		return 0
	}
}

// ToHumanAcc is converting the Accuracy to a human readable Number, (Way smaller then E.G 0.895023981 // should be 89.50 % acc)
func ToHumanAcc(acc float64) float64 {
	a := fmt.Sprintf("%.2f", acc*100)
	r, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return acc * 100
	}
	return r
}
