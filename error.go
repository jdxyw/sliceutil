package sliceutil

import "errors"

var (
	ErrSliceLengthZero       = errors.New("slice length is zero")
	ErrSliceIsNil            = errors.New("slice is nil")
	ErrSliceLengthIsNotEqual = errors.New("slice's length is not equal")
	ErrSliceIndexOutOfBound  = errors.New("slice index is out of bound")
)
