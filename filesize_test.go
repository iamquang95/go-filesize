package filesize

import (
	"errors"
	"math"
	"testing"
)

const (
	eps = 1e-8
)

func TestConvert(t *testing.T) {
	invalidUnitErr := errors.New("Invalid unit")
	testcases := []struct {
		input          Byte
		unit           Unit
		expectedOutput float64
		expectedErr    error
	}{
		{Byte(0), B, float64(0), nil},
		{Byte(KB), KB, float64(1), nil},
		{Byte(MB), MB, float64(1), nil},
		{Byte(GB), GB, float64(1), nil},
		{Byte(TB), TB, float64(1), nil},
		{Byte(PB), PB, float64(1), nil},
		{Byte(EB), EB, float64(1), nil},
		{Byte(2048), KB, float64(2), nil},
		{Byte(1048576), KB, float64(1024), nil},
		{Byte(1048576), MB, float64(1), nil},
		{Byte(1581252608), MB, float64(1508), nil},
		{Byte(4608), KB, float64(4.5), nil},
		{Byte(21440476741632), TB, float64(19.5), nil},
		{Byte(1024), Unit(100), float64(0), invalidUnitErr},
	}
	for _, tc := range testcases {
		output, err := Byte(tc.input).Convert(tc.unit)
		if err != nil {
			if err.Error() != tc.expectedErr.Error() {
				t.Errorf("Covert(%f) throw unexpected error %s", float64(tc.input), err.Error())
			}
		} else {
			if !floatEqual(output, tc.expectedOutput) {
				t.Errorf("Covert(%f) => %f, expected %f", float64(tc.input), output, tc.expectedOutput)
			}
		}
	}
}

func TestToString(t *testing.T) {
	invalidUnitErr := errors.New("Invalid unit")
	testcases := []struct {
		input          Byte
		unit           Unit
		expectedOutput string
		expectedErr    error
	}{
		{Byte(0), B, "0B", nil},
		{Byte(KB), KB, "1.0KB", nil},
		{Byte(MB), MB, "1.0MB", nil},
		{Byte(GB), GB, "1.0GB", nil},
		{Byte(TB), TB, "1.0TB", nil},
		{Byte(PB), PB, "1.0PB", nil},
		{Byte(EB), EB, "1.0EB", nil},
		{Byte(2048), KB, "2.0KB", nil},
		{Byte(1048576), KB, "1024.0KB", nil},
		{Byte(1048576), MB, "1.0MB", nil},
		{Byte(1581252608), MB, "1508.0MB", nil},
		{Byte(4608), KB, "4.5KB", nil},
		{Byte(21440476741632), TB, "19.5TB", nil},
		{Byte(1024), Unit(100), "", invalidUnitErr},
	}
	for _, tc := range testcases {
		output, err := Byte(tc.input).ToString(tc.unit)
		if err != nil {
			if err.Error() != tc.expectedErr.Error() {
				t.Errorf("CovertToString(%f) throw unexpected error %s", float64(tc.input), err.Error())
			}
		} else {
			if output != tc.expectedOutput {
				t.Errorf("CovertToString(%f) => %s, expected %s", float64(tc.input), output, tc.expectedOutput)
			}
		}
	}
}

func floatEqual(lhs float64, rhs float64) bool {
	if math.Abs(lhs-rhs) < eps {
		return true
	}
	return false
}
