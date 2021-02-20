package sliceutil

import "reflect"

func loopSliceAndCmp(s interface{}, a interface{}) bool {
	si := reflect.ValueOf(s)
	for i := 0; i < si.Len(); i++ {
		if reflect.DeepEqual(si.Index(i).Interface(), a) {
			return true
		}
	}

	return false
}

func ContainsInt(s []int, a int) bool {
	return loopSliceAndCmp(s, a)
}

func ContainsInt32(s []int32, a int32) bool {
	return loopSliceAndCmp(s, a)
}

func ContainsInt64(s []int64, a int64) bool {
	return loopSliceAndCmp(s, a)
}

func ContainsBool(s []bool, a bool) bool {
	return loopSliceAndCmp(s, a)
}

func ContainsString(s []string, a string) bool {
	return loopSliceAndCmp(s, a)
}

func ContainsByte(s []byte, a byte) bool {
	return loopSliceAndCmp(s, a)
}
