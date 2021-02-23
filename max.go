package sliceutil

import "reflect"

func verifymax(s interface{}) error {
	si := reflect.ValueOf(s)

	if si.IsNil() {
		return ErrSliceIsNil
	}

	if si.Len() == 0 {
		return ErrSliceLengthZero
	}

	return nil
}

// MaxInt returns the max element in a int slice.
func MaxInt(s []int) (int, error) {
	if err := verifymax(s); err != nil {
		return 0, err
	}

	ret := s[0]

	for _, a := range s {
		if a > ret {
			ret = a
		}
	}

	return ret, nil
}

// MaxInt32 returns the max element in a int32 slice.
func MaxInt32(s []int32) (int32, error) {
	if err := verifymax(s); err != nil {
		return 0, err
	}

	ret := s[0]

	for _, a := range s {
		if a > ret {
			ret = a
		}
	}

	return ret, nil
}

// MaxInt64 returns the max element in a int64 slice.
func MaxInt64(s []int64) (int64, error) {
	if err := verifymax(s); err != nil {
		return 0, err
	}

	ret := s[0]

	for _, a := range s {
		if a > ret {
			ret = a
		}
	}

	return ret, nil
}

// MaxFloat32 returns the max element in a float32 slice.
func MaxFloat32(s []float32) (float32, error) {
	if err := verifymax(s); err != nil {
		return 0, err
	}

	ret := s[0]

	for _, a := range s {
		if a > ret {
			ret = a
		}
	}

	return ret, nil
}

// MaxFloat64 returns the max element in a float32 slice.
func MaxFloat64(s []float64) (float64, error) {
	if err := verifymax(s); err != nil {
		return 0, err
	}

	ret := s[0]

	for _, a := range s {
		if a > ret {
			ret = a
		}
	}

	return ret, nil
}

// MaxByte returns the max element in a float32 slice.
func MaxByte(s []byte) (byte, error) {
	if err := verifymax(s); err != nil {
		return 0, err
	}

	ret := s[0]

	for _, a := range s {
		if a > ret {
			ret = a
		}
	}

	return ret, nil
}
