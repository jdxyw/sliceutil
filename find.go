package sliceutil

import "reflect"

func findFist(s interface{}, a interface{}) (int, bool) {
	si := reflect.ValueOf(s)

	if si.IsNil() || si.Len() == 0 {
		return -1, false
	}

	for i := 0; i < si.Len(); i++ {
		if reflect.DeepEqual(si.Index(i).Interface(), a) {
			return i, true
		}
	}

	return -1, false
}

// FindFistInt returns the first specified element index. If can't find anything, return -1 and false
func FindFistInt(s []int, a int) (int, bool) {
	return findFist(s, a)
}

// FindFistInt32 returns the first index of specified element of type int32. If can't find anything, return -1 and false
func FindFistInt32(s []int32, a int32) (int, bool) {
	return findFist(s, a)
}

// FindFistInt64 returns the first index of specified element of type int64. If can't find anything, return -1 and false
func FindFistInt64(s []int64, a int64) (int, bool) {
	return findFist(s, a)
}

// FindFistBool returns the first index of specified element of type bool. If can't find anything, return -1 and false
func FindFistBool(s []bool, a bool) (int, bool) {
	return findFist(s, a)
}

// FindFistByte returns the first index of specified element of type byte. If can't find anything, return -1 and false
func FindFistByte(s []byte, a byte) (int, bool) {
	return findFist(s, a)
}

// FindFistString returns the first index of specified element of type string. If can't find anything, return -1 and false
func FindFistString(s []string, a string) (int, bool) {
	return findFist(s, a)
}
