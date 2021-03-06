package helpers

import (
	"fmt"
	"strconv"

	"github.com/Gigamons/common/consts"
)

// CalculateAccuracy calculates every playModes accuracy.
func CalculateAccuracy(count300 uint64, count100 uint64, count50 uint64, countMiss uint64, countGeki uint64, countKatu uint64, playMode byte) float64 {
	var thp uint64
	var th uint64
	if count300 == 0 && count100 == 0 && count50 == 0 && countMiss == 0 && countGeki == 0 && countKatu == 0 {
		return 1
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

// ToHumanAcc is converting the Accuracy to a human readable Number. (Way smaller then E.G 0.895023981 // should be 89.50 % acc)
func ToHumanAcc(acc float64) float64 {
	a := fmt.Sprintf("%.2f", acc*100)
	r, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return acc * 100
	}
	return r
}
