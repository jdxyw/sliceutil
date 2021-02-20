package sliceutil

import (
	"reflect"
	"testing"
)

func TestReverseInt(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "testcase1", args: args{s: []int{1, 2, 3, 4, 5}}, want: []int{5, 4, 3, 2, 1}},
		{name: "testcase2", args: args{s: []int{}}, want: []int{}},
		{name: "testcase3", args: args{s: []int{1, 2, 3, 4, 5, 6}}, want: []int{6, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseInt(tt.args.s)
			if reflect.DeepEqual(tt.args.s, tt.want) != true {
				t.Errorf("ReverseInt() got = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}

func TestReverseInt32(t *testing.T) {
	type args struct {
		s []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{name: "testcase1", args: args{s: []int32{1, 2, 3, 4, 5}}, want: []int32{5, 4, 3, 2, 1}},
		{name: "testcase2", args: args{s: []int32{}}, want: []int32{}},
		{name: "testcase3", args: args{s: []int32{1, 2, 3, 4, 5, 6}}, want: []int32{6, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseInt32(tt.args.s)
			if reflect.DeepEqual(tt.args.s, tt.want) != true {
				t.Errorf("ReverseInt32() got = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}

func TestReverseInt64(t *testing.T) {
	type args struct {
		s []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{name: "testcase1", args: args{s: []int64{1, 2, 3, 4, 5}}, want: []int64{5, 4, 3, 2, 1}},
		{name: "testcase2", args: args{s: []int64{}}, want: []int64{}},
		{name: "testcase3", args: args{s: []int64{1, 2, 3, 4, 5, 6}}, want: []int64{6, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseInt64(tt.args.s)
			if reflect.DeepEqual(tt.args.s, tt.want) != true {
				t.Errorf("ReverseInt64() got = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}

func TestReverseFloat32(t *testing.T) {
	type args struct {
		s []float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{name: "testcase1", args: args{s: []float32{1, 2, 3, 4, 5}}, want: []float32{5, 4, 3, 2, 1}},
		{name: "testcase2", args: args{s: []float32{}}, want: []float32{}},
		{name: "testcase3", args: args{s: []float32{1, 2, 3, 4, 5, 6}}, want: []float32{6, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseFloat32(tt.args.s)
			if reflect.DeepEqual(tt.args.s, tt.want) != true {
				t.Errorf("ReverseFloat32() got = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}

func TestReverseFloat64(t *testing.T) {
	type args struct {
		s []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{name: "testcase1", args: args{s: []float64{1, 2, 3, 4, 5}}, want: []float64{5, 4, 3, 2, 1}},
		{name: "testcase2", args: args{s: []float64{}}, want: []float64{}},
		{name: "testcase3", args: args{s: []float64{1, 2, 3, 4, 5, 6}}, want: []float64{6, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseFloat64(tt.args.s)
			if reflect.DeepEqual(tt.args.s, tt.want) != true {
				t.Errorf("ReverseFloat64() got = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "testcase1", args: args{s: []string{"java", "php", "python"}}, want: []string{"python", "php", "java"}},
		{name: "testcase2", args: args{s: []string{}}, want: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseString(tt.args.s)
			if reflect.DeepEqual(tt.args.s, tt.want) != true {
				t.Errorf("ReverseString() got = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}

func TestReverseByte(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "testcase1", args: args{s: []byte{1, 2, 3, 4, 5}}, want: []byte{5, 4, 3, 2, 1}},
		{name: "testcase2", args: args{s: []byte{}}, want: []byte{}},
		{name: "testcase3", args: args{s: []byte{1, 2, 3, 4, 5, 6}}, want: []byte{6, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseByte(tt.args.s)
			if reflect.DeepEqual(tt.args.s, tt.want) != true {
				t.Errorf("ReverseByte() got = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}

func TestReverseBool(t *testing.T) {
	type args struct {
		s []bool
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{name: "testcase1", args: args{s: []bool{true, true, false, false, false}}, want: []bool{false, false, false, true, true}},
		{name: "testcase2", args: args{s: []bool{}}, want: []bool{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseBool(tt.args.s)
			if reflect.DeepEqual(tt.args.s, tt.want) != true {
				t.Errorf("ReverseBool() got = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
