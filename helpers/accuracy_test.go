package helpers

import (
	"testing"

	"github.com/Gigamons/common/consts"
)

const (
	Count300  = uint64(500)
	Count100  = uint64(421)
	Count50   = uint64(120)
	CountMiss = uint64(3)
	CountGeki = uint64(250)
	CountKatu = uint64(220)
)

var (
	Accuracy = 0.0
)

func TestCalculateAccuracy(t *testing.T) {
	t.Log("Testing Accuracy for STD")
	Accuracy = CalculateAccuracy(Count300, Count100, Count50, CountMiss, CountGeki, CountKatu, consts.STD)
	if Accuracy == 0 {
		t.Error("Accuracy is 0")
	}
	t.Logf("Result is %v", Accuracy)
	t.Log("Testing Accuracy for Taiko")
	Accuracy = CalculateAccuracy(Count300, Count100, Count50, CountMiss, CountGeki, CountKatu, consts.Taiko)
	if Accuracy == 0 {
		t.Error("Accuracy is 0")
	}
	t.Logf("Result is %v", Accuracy)
	t.Log("Testing Accuracy for CTB")
	Accuracy = CalculateAccuracy(Count300, Count100, Count50, CountMiss, CountGeki, CountKatu, consts.CTB)
	if Accuracy == 0 {
		t.Error("Accuracy is 0")
	}
	t.Logf("Result is %v", Accuracy)
	t.Log("Testing Accuracy for Mania")
	Accuracy = CalculateAccuracy(Count300, Count100, Count50, CountMiss, CountGeki, CountKatu, consts.Mania)
	if Accuracy == 0 {
		t.Error("Accuracy is 0")
	}
	t.Logf("Result is %v", Accuracy)
}

func TestToHumanAcc(t *testing.T) {
	t.Log("Calculate Accuracy for Mania")
	Accuracy = CalculateAccuracy(Count300, Count100, Count50, CountMiss, CountGeki, CountKatu, consts.Mania)
	if Accuracy == 0 {
		t.Error("Accuracy is 0")
	}
	t.Logf("Result is %v", Accuracy)
	Result := ToHumanAcc(Accuracy)
	if Result == 0 {
		t.Error("Accuracy is 0")
	}
	t.Logf("Result is %v", Result)
}
