package sliceutil

import "reflect"

func verifyRmr(s interface{}, from, to int) error {
	si := reflect.ValueOf(s)

	if si.Kind() != reflect.Slice {
		return ErrTypeNotSupport
	}

	if si.IsNil() {
		return ErrSliceIsNil
	}

	if si.Len() == 0 {
		return ErrSliceLengthZero
	}

	if from < 0 || from > si.Len()-1 || to < 0 || to > si.Len() {
		return ErrSliceIndexOutOfBound
	}

	if to < from {
		return ErrSliceRange
	}

	return nil
}

// RemoveRangeInt removes the elements between from and to index(exclusive, [from, to)) from a int slice.
func RemoveRangeInt(s []int, from, to int) ([]int, error) {
	if err := verifyRmr(s, from, to); err != nil {
		return s, err
	}

	return append(s[:from], s[to:]...), nil
}

// RemoveRangeInt32 removes the elements between from and to index(exclusive, [from, to)) from a int32 slice.
func RemoveRangeInt32(s []int32, from, to int) ([]int32, error) {
	if err := verifyRmr(s, from, to); err != nil {
		return s, err
	}

	return append(s[:from], s[to:]...), nil
}

// RemoveRangeInt64 removes the elements between from and to index(exclusive, [from, to)) from a int64 slice.
func RemoveRangeInt64(s []int64, from, to int) ([]int64, error) {
	if err := verifyRmr(s, from, to); err != nil {
		return s, err
	}

	return append(s[:from], s[to:]...), nil
}

// RemoveRangeByte removes the elements between from and to index(exclusive, [from, to)) from a byte slice.
func RemoveRangeByte(s []byte, from, to int) ([]byte, error) {
	if err := verifyRmr(s, from, to); err != nil {
		return s, err
	}

	return append(s[:from], s[to:]...), nil
}

// RemoveRangeBool removes the elements between from and to index(exclusive, [from, to)) from a bool slice.
func RemoveRangeBool(s []bool, from, to int) ([]bool, error) {
	if err := verifyRmr(s, from, to); err != nil {
		return s, err
	}

	return append(s[:from], s[to:]...), nil
}

// RemoveRangeString removes the elements between from and to index(exclusive, [from, to)) from a string slice.
func RemoveRangeString(s []string, from, to int) ([]string, error) {
	if err := verifyRmr(s, from, to); err != nil {
		return s, err
	}

	return append(s[:from], s[to:]...), nil
}

// RemoveRangeFloat32 removes the elements between from and to index(exclusive, [from, to)) from a float32 slice.
func RemoveRangeFloat32(s []float32, from, to int) ([]float32, error) {
	if err := verifyRmr(s, from, to); err != nil {
		return s, err
	}

	return append(s[:from], s[to:]...), nil
}

// RemoveRangeFloat64 removes the elements between from and to index(exclusive, [from, to)) from a float64 slice.
func RemoveRangeFloat64(s []float64, from, to int) ([]float64, error) {
	if err := verifyRmr(s, from, to); err != nil {
		return s, err
	}

	return append(s[:from], s[to:]...), nil
}

// RemoveRangeSlice removes the elements between from and to index(exclusive, [from, to)) from a slice.
func RemoveRangeSlice(s interface{}, from, to int) (interface{}, error) {
	if err := verifyRmr(s, from, to); err != nil {
		return s, err
	}

	si := reflect.ValueOf(s)
	ret := make([]interface{}, 0)

	for i := 0; i < si.Len(); i++ {
		if i < from || i >= to {
			ret = append(ret, si.Index(i))
		}
	}

	return ret, nil
}
