package sliceutil

import "errors"

var (
	ErrSliceLengthZero       = errors.New("slice length is zero")
	ErrSliceIsNil            = errors.New("slice is nil")
	ErrSliceLengthIsNotEqual = errors.New("slice's length is not equal")
	ErrSliceIndexOutOfBound  = errors.New("slice index is out of bound")
	ErrTypeNotSupport        = errors.New("typs is not supported, it should be slice")
	ErrSliceRange            = errors.New("range from shoule be less than to")
)
