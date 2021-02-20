package sliceutil

import (
	"bytes"
	"math"
	"reflect"
)

const (
	float64Epsilon = 1e-9
)

func verify(a, b interface{}) error {
	ai := reflect.ValueOf(a)
	bi := reflect.ValueOf(b)

	if ai.IsNil() || bi.IsNil() {
		return ErrSliceIsNil
	}

	if ai.Len() == 0 || bi.Len() == 0 {
		return ErrSliceLengthZero
	}

	if ai.Len() != bi.Len() {
		return ErrSliceLengthIsNotEqual
	}

	return nil
}

// EqualBool return true if two boolean slices are equal.
func EqualBool(a, b []bool) (bool, error) {

	if err := verify(a, b); err != nil {
		return false, err
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false, nil
		}
	}

	return true, nil
}

// EqualInt return true if two int slices are equal.
func EqualInt(a, b []int) (bool, error) {
	if err := verify(a, b); err != nil {
		return false, err
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false, nil
		}
	}

	return true, nil
}

// EqualInt32 return true if two int32 slices are equal.
func EqualInt32(a, b []int32) (bool, error) {
	if err := verify(a, b); err != nil {
		return false, err
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false, nil
		}
	}

	return true, nil
}

// EqualInt64 return true if two int64 slices are equal.
func EqualInt64(a, b []int64) (bool, error) {
	if err := verify(a, b); err != nil {
		return false, err
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false, nil
		}
	}

	return true, nil
}

// EqualFloat32 return true if two float32 slices are equal.
func EqualFloat32(a, b []float32) (bool, error) {
	if err := verify(a, b); err != nil {
		return false, err
	}

	for i, _ := range a {
		if math.Abs(float64(a[i])-float64(b[i])) > float64Epsilon {
			return false, nil
		}
	}

	return true, nil
}

// EqualFloat64 return true if two float63 slices are equal.
func EqualFloat64(a, b []float64) (bool, error) {
	if err := verify(a, b); err != nil {
		return false, err
	}

	for i, _ := range a {
		if math.Abs(a[i]-b[i]) > float64Epsilon {
			return false, nil
		}
	}

	return true, nil
}

// EqualBytes return true if two byte slices are equal.
func EqualBytes(a, b []byte) (bool, error) {
	if err := verify(a, b); err != nil {
		return false, err
	}

	if bytes.Equal(a, b) {
		return true, nil
	}

	return false, nil
}

type comparator func(a, b interface{}) bool

// EqualSlice can compare two slices of any type and use customized comparator.
func EqualSlice(a, b interface{}, fn comparator) (bool, error) {
	if err := verify(a, b); err != nil {
		return false, err
	}

	ai := reflect.ValueOf(a)
	bi := reflect.ValueOf(b)

	for i := 0; i < ai.Len(); i++ {
		if fn(ai.Index(i).Interface(), bi.Index(i).Interface()) != true {
			return false, nil
		}
	}

	return true, nil
}
