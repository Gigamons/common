package helpers

import (
	"reflect"
	"testing"
)

func TestPseudorandombytes(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name     string
		args     args
		dontwant []byte
	}{
		{
			name:     "Pseudorandombytes",
			args:     args{n: 2 ^ 20},
			dontwant: nil,
		},
		{
			name:     "Pseudorandombytes",
			args:     args{n: 2 ^ 10},
			dontwant: nil,
		},
		{
			name:     "Pseudorandombytes",
			args:     args{n: 2 ^ 5},
			dontwant: nil,
		},
		{
			name:     "Pseudorandombytes",
			args:     args{n: 2 ^ 1},
			dontwant: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pseudorandombytes(tt.args.n); reflect.DeepEqual(got, tt.dontwant) {
				t.Errorf("Pseudorandombytes() = %v, doesn't want %v", got, tt.dontwant)
			} else {
				t.Logf("Pseudorandombytes() = %x, length = %v", got, len(got))
			}
		})
	}
}
