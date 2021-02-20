package sliceutil

import "reflect"

func verifyRm(s interface{}, i int) error {
	si := reflect.ValueOf(s)

	if si.IsNil() {
		return ErrSliceIsNil
	}

	if si.Len() == 0 {
		return ErrSliceLengthZero
	}

	if i < 0 || i > si.Len()-1 {
		return ErrSliceIndexOutOfBound
	}

	return nil
}

// RemoveIndexInt removes an element at index i
func RemoveIndexInt(s []int, i int) ([]int, error) {
	if err := verifyRm(s, i); err != nil {
		return s, err
	}

	return append(s[:i], s[i+1:]...), nil
}

// RemoveIndexInt32 removes an element at index i
func RemoveIndexInt32(s []int32, i int) ([]int32, error) {
	if err := verifyRm(s, i); err != nil {
		return s, err
	}

	return append(s[:i], s[i+1:]...), nil
}

// RemoveIndexInt64 removes an element at index i
func RemoveIndexInt64(s []int64, i int) ([]int64, error) {
	if err := verifyRm(s, i); err != nil {
		return s, err
	}

	return append(s[:i], s[i+1:]...), nil
}

// RemoveIndexFloat32 removes an element at index i
func RemoveIndexFloat32(s []float32, i int) ([]float32, error) {
	if err := verifyRm(s, i); err != nil {
		return s, err
	}

	return append(s[:i], s[i+1:]...), nil
}

// RemoveIndexFloat64 removes an element at index i
func RemoveIndexFloat64(s []float64, i int) ([]float64, error) {
	if err := verifyRm(s, i); err != nil {
		return s, err
	}

	return append(s[:i], s[i+1:]...), nil
}

// RemoveIndexByte removes an element at index i
func RemoveIndexByte(s []byte, i int) ([]byte, error) {
	if err := verifyRm(s, i); err != nil {
		return s, err
	}

	return append(s[:i], s[i+1:]...), nil
}

// RemoveIndexBool removes an element at index i
func RemoveIndexBool(s []bool, i int) ([]bool, error) {
	if err := verifyRm(s, i); err != nil {
		return s, err
	}

	return append(s[:i], s[i+1:]...), nil
}
