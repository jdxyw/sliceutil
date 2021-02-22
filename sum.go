package sliceutil

import "reflect"

func verifySum(s interface{}) error {
	si := reflect.ValueOf(s)

	if si.IsNil() {
		return ErrSliceIsNil
	}

	if si.Len() == 0 {
		return ErrSliceLengthZero
	}

	return nil
}

// SumInt returns sum of a int slice.
func SumInt(s []int) (int, error) {
	if err := verifySum(s); err != nil {
		return 0, err
	}

	var total int = 0
	for _, a := range s {
		total += a
	}

	return total, nil
}

// SumInt32 returns sum of a int32 slice.
func SumInt32(s []int32) (int32, error) {
	if err := verifySum(s); err != nil {
		return 0, err
	}

	var total int32 = 0
	for _, a := range s {
		total += a
	}

	return total, nil
}

// SumInt64 returns sum of a int64 slice.
func SumInt64(s []int64) (int64, error) {
	if err := verifySum(s); err != nil {
		return 0, err
	}

	var total int64 = 0
	for _, a := range s {
		total += a
	}

	return total, nil
}

// SumFloat32 returns sum of a float32 slice.
func SumFloat32(s []float32) (float32, error) {
	if err := verifySum(s); err != nil {
		return 0, err
	}

	var total float32 = 0
	for _, a := range s {
		total += a
	}

	return total, nil
}

// SumFloat64 returns sum of a float64 slice.
func SumFloat64(s []float64) (float64, error) {
	if err := verifySum(s); err != nil {
		return 0, err
	}

	var total float64 = 0
	for _, a := range s {
		total += a
	}

	return total, nil
}
