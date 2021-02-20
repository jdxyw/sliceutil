package sliceutil

import "reflect"

func reverseSlice(s interface{}) {
	l := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)

	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func ReverseInt(s []int) {
	reverseSlice(s)
}

func ReverseInt32(s []int32) {
	reverseSlice(s)
}

func ReverseInt64(s []int64) {
	reverseSlice(s)
}

func ReverseFloat32(s []float32) {
	reverseSlice(s)
}

func ReverseFloat64(s []float64) {
	reverseSlice(s)
}

func ReverseString(s []string) {
	reverseSlice(s)
}

func ReverseByte(s []byte) {
	reverseSlice(s)
}

func ReverseBool(s []bool) {
	reverseSlice(s)
}
