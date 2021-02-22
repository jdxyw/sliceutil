package sliceutil

import "testing"

func TestFindFistInt(t *testing.T) {
	type args struct {
		s []int
		a int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{name: "testcase1", args: args{s: []int{}, a: 1}, want: -1, want1: false},
		{name: "testcase2", args: args{s: nil, a: 1}, want: -1, want1: false},
		{name: "testcase3", args: args{s: []int{1, 2, 3, 4, 5}, a: 1}, want: 0, want1: true},
		{name: "testcase4", args: args{s: []int{1, 2, 3, 4, 5}, a: 5}, want: 4, want1: true},
		{name: "testcase5", args: args{s: []int{1, 2, 2, 3, 4, 2, 2, 4, 5}, a: 2}, want: 1, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindFistInt(tt.args.s, tt.args.a)
			if got != tt.want {
				t.Errorf("FindFistInt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindFistInt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFindFistInt32(t *testing.T) {
	type args struct {
		s []int32
		a int32
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{name: "testcase1", args: args{s: []int32{}, a: 1}, want: -1, want1: false},
		{name: "testcase2", args: args{s: nil, a: 1}, want: -1, want1: false},
		{name: "testcase3", args: args{s: []int32{1, 2, 3, 4, 5}, a: 1}, want: 0, want1: true},
		{name: "testcase4", args: args{s: []int32{1, 2, 3, 4, 5}, a: 5}, want: 4, want1: true},
		{name: "testcase5", args: args{s: []int32{1, 2, 2, 3, 4, 2, 2, 4, 5}, a: 2}, want: 1, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindFistInt32(tt.args.s, tt.args.a)
			if got != tt.want {
				t.Errorf("FindFistInt32() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindFistInt32() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFindFistInt64(t *testing.T) {
	type args struct {
		s []int64
		a int64
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{name: "testcase1", args: args{s: []int64{}, a: 1}, want: -1, want1: false},
		{name: "testcase2", args: args{s: nil, a: 1}, want: -1, want1: false},
		{name: "testcase3", args: args{s: []int64{1, 2, 3, 4, 5}, a: 1}, want: 0, want1: true},
		{name: "testcase4", args: args{s: []int64{1, 2, 3, 4, 5}, a: 5}, want: 4, want1: true},
		{name: "testcase5", args: args{s: []int64{1, 2, 2, 3, 4, 2, 2, 4, 5}, a: 2}, want: 1, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindFistInt64(tt.args.s, tt.args.a)
			if got != tt.want {
				t.Errorf("FindFistInt64() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindFistInt64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFindFistBool(t *testing.T) {
	type args struct {
		s []bool
		a bool
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{name: "testcase1", args: args{s: []bool{}, a: false}, want: -1, want1: false},
		{name: "testcase2", args: args{s: nil, a: true}, want: -1, want1: false},
		{name: "testcase3", args: args{s: []bool{true, false, false, false, false}, a: true}, want: 0, want1: true},
		{name: "testcase4", args: args{s: []bool{false, false, false, false, true}, a: true}, want: 4, want1: true},
		{name: "testcase5", args: args{s: []bool{false, true, true, false, false, false, false, false, false}, a: true}, want: 1, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindFistBool(tt.args.s, tt.args.a)
			if got != tt.want {
				t.Errorf("FindFistBool() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindFistBool() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFindFistByte(t *testing.T) {
	type args struct {
		s []byte
		a byte
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{name: "testcase1", args: args{s: []byte{}, a: 1}, want: -1, want1: false},
		{name: "testcase2", args: args{s: nil, a: 1}, want: -1, want1: false},
		{name: "testcase3", args: args{s: []byte{1, 2, 3, 4, 5}, a: 1}, want: 0, want1: true},
		{name: "testcase4", args: args{s: []byte{1, 2, 3, 4, 5}, a: 5}, want: 4, want1: true},
		{name: "testcase5", args: args{s: []byte{1, 2, 2, 3, 4, 2, 2, 4, 5}, a: 2}, want: 1, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindFistByte(tt.args.s, tt.args.a)
			if got != tt.want {
				t.Errorf("FindFistByte() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindFistByte() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFindFistString(t *testing.T) {
	type args struct {
		s []string
		a string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{name: "testcase1", args: args{s: []string{}, a: "1"}, want: -1, want1: false},
		{name: "testcase2", args: args{s: nil, a: "1"}, want: -1, want1: false},
		{name: "testcase3", args: args{s: []string{"1", "2", "3", "4", "5"}, a: "1"}, want: 0, want1: true},
		{name: "testcase4", args: args{s: []string{"1", "2", "3", "4", "5"}, a: "5"}, want: 4, want1: true},
		{name: "testcase5", args: args{s: []string{"1", "2", "2", "3", "4", "2", "2", "4", "5"}, a: "2"}, want: 1, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindFistString(tt.args.s, tt.args.a)
			if got != tt.want {
				t.Errorf("FindFistString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindFistString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
