package sliceutil

import "testing"

func TestContainsInt(t *testing.T) {
	type args struct {
		s []int
		a int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "testcase1", args: args{s: []int{1, 2, 3, 4, 5, 6}, a: 3}, want: true},
		{name: "testcase2", args: args{s: []int{1, 2, 3, 4, 5, 6}, a: 6}, want: true},
		{name: "testcase3", args: args{s: []int{1, 2, 3, 4, 5, 6}, a: 1}, want: true},
		{name: "testcase4", args: args{s: []int{1, 2, 3, 4, 5, 6}, a: 8}, want: false},
		{name: "testcase5", args: args{s: []int{}, a: 3}, want: false},
		{name: "testcase6", args: args{s: []int{1, 2, 3, 4, 5, 6}, a: -3}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt(tt.args.s, tt.args.a); got != tt.want {
				t.Errorf("ContainsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt32(t *testing.T) {
	type args struct {
		s []int32
		a int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "testcase1", args: args{s: []int32{1, 2, 3, 4, 5, 6}, a: 3}, want: true},
		{name: "testcase2", args: args{s: []int32{1, 2, 3, 4, 5, 6}, a: 6}, want: true},
		{name: "testcase3", args: args{s: []int32{1, 2, 3, 4, 5, 6}, a: 1}, want: true},
		{name: "testcase4", args: args{s: []int32{1, 2, 3, 4, 5, 6}, a: 8}, want: false},
		{name: "testcase5", args: args{s: []int32{}, a: 3}, want: false},
		{name: "testcase6", args: args{s: []int32{1, 2, 3, 4, 5, 6}, a: -3}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt32(tt.args.s, tt.args.a); got != tt.want {
				t.Errorf("ContainsInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt64(t *testing.T) {
	type args struct {
		s []int64
		a int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "testcase1", args: args{s: []int64{1, 2, 3, 4, 5, 6}, a: 3}, want: true},
		{name: "testcase2", args: args{s: []int64{1, 2, 3, 4, 5, 6}, a: 6}, want: true},
		{name: "testcase3", args: args{s: []int64{1, 2, 3, 4, 5, 6}, a: 1}, want: true},
		{name: "testcase4", args: args{s: []int64{1, 2, 3, 4, 5, 6}, a: 8}, want: false},
		{name: "testcase5", args: args{s: []int64{}, a: 3}, want: false},
		{name: "testcase6", args: args{s: []int64{1, 2, 3, 4, 5, 6}, a: -3}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt64(tt.args.s, tt.args.a); got != tt.want {
				t.Errorf("ContainsInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsBool(t *testing.T) {
	type args struct {
		s []bool
		a bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "testcase1", args: args{s: []bool{true, true, false, false, true}, a: false}, want: true},
		{name: "testcase2", args: args{s: []bool{true, true, false, false, true}, a: true}, want: true},
		{name: "testcase3", args: args{s: []bool{false, false, false}, a: true}, want: false},
		{name: "testcase4", args: args{s: []bool{}, a: true}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsBool(tt.args.s, tt.args.a); got != tt.want {
				t.Errorf("ContainsBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsString(t *testing.T) {
	type args struct {
		s []string
		a string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "testcase1", args: args{s: []string{"abc", "bcd", "cdf", "", "efg"}, a: "abc"}, want: true},
		{name: "testcase2", args: args{s: []string{"abc", "bcd", "cdf", "", "efg"}, a: "efg"}, want: true},
		{name: "testcase3", args: args{s: []string{"abc", "bcd", "cdf", "", "efg"}, a: ""}, want: true},
		{name: "testcase4", args: args{s: []string{"abc", "bcd", "cdf", " ", "efg"}, a: " "}, want: true},
		{name: "testcase5", args: args{s: []string{"abc", "bcd", "cdf", " ", "efg"}, a: "vcx"}, want: false},
		{name: "testcase6", args: args{s: []string{""}, a: "vcx"}, want: false},
		{name: "testcase7", args: args{s: []string{""}, a: ""}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsString(tt.args.s, tt.args.a); got != tt.want {
				t.Errorf("ContainsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsByte(t *testing.T) {
	type args struct {
		s []byte
		a byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "testcase1", args: args{s: []byte{1, 2, 3, 4, 5, 6}, a: 3}, want: true},
		{name: "testcase2", args: args{s: []byte{1, 2, 3, 4, 5, 6}, a: 6}, want: true},
		{name: "testcase3", args: args{s: []byte{1, 2, 3, 4, 5, 6}, a: 1}, want: true},
		{name: "testcase4", args: args{s: []byte{1, 2, 3, 4, 5, 6}, a: 8}, want: false},
		{name: "testcase5", args: args{s: []byte{}, a: 3}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsByte(tt.args.s, tt.args.a); got != tt.want {
				t.Errorf("ContainsByte() = %v, want %v", got, tt.want)
			}
		})
	}
}
