package sliceutil

import "testing"

func TestEqualBool(t *testing.T) {
	type args struct {
		a []bool
		b []bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{name: "testcase1", args: args{a: nil, b: []bool{true}}, want: false, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{a: []bool{}, b: []bool{}}, want: false, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{a: []bool{true, false}, b: []bool{true}}, want: false, wantErr: ErrSliceLengthIsNotEqual},
		{name: "testcase4", args: args{a: []bool{true, false}, b: []bool{true, false}}, want: true, wantErr: nil},
		{name: "testcase5", args: args{a: []bool{true, false}, b: []bool{true, true}}, want: false, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EqualBool(tt.args.a, tt.args.b)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("EqualBool() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("EqualBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualInt(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{name: "testcase1", args: args{a: nil, b: []int{1}}, want: false, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{a: []int{}, b: []int{}}, want: false, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{a: []int{1, 2}, b: []int{1}}, want: false, wantErr: ErrSliceLengthIsNotEqual},
		{name: "testcase4", args: args{a: []int{1, 2}, b: []int{1, 2}}, want: true, wantErr: nil},
		{name: "testcase5", args: args{a: []int{1, 2}, b: []int{1, 1}}, want: false, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EqualInt(tt.args.a, tt.args.b)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("EqualBool() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("EqualBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualInt32(t *testing.T) {
	type args struct {
		a []int32
		b []int32
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{name: "testcase1", args: args{a: nil, b: []int32{1}}, want: false, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{a: []int32{}, b: []int32{}}, want: false, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{a: []int32{1, 2}, b: []int32{1}}, want: false, wantErr: ErrSliceLengthIsNotEqual},
		{name: "testcase4", args: args{a: []int32{1, 2}, b: []int32{1, 2}}, want: true, wantErr: nil},
		{name: "testcase5", args: args{a: []int32{1, 2}, b: []int32{1, 1}}, want: false, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EqualInt32(tt.args.a, tt.args.b)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("EqualBool() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("EqualBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualInt64(t *testing.T) {
	type args struct {
		a []int64
		b []int64
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{name: "testcase1", args: args{a: nil, b: []int64{1}}, want: false, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{a: []int64{}, b: []int64{}}, want: false, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{a: []int64{1, 2}, b: []int64{1}}, want: false, wantErr: ErrSliceLengthIsNotEqual},
		{name: "testcase4", args: args{a: []int64{1, 2}, b: []int64{1, 2}}, want: true, wantErr: nil},
		{name: "testcase5", args: args{a: []int64{1, 2}, b: []int64{1, 1}}, want: false, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EqualInt64(tt.args.a, tt.args.b)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("EqualBool() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("EqualBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualFloat32(t *testing.T) {
	type args struct {
		a []float32
		b []float32
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{name: "testcase1", args: args{a: nil, b: []float32{1.0}}, want: false, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{a: []float32{}, b: []float32{}}, want: false, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{a: []float32{1.1, 2.1}, b: []float32{1.1}}, want: false, wantErr: ErrSliceLengthIsNotEqual},
		{name: "testcase4", args: args{a: []float32{1.01, 2}, b: []float32{1.01, 2}}, want: true, wantErr: nil},
		{name: "testcase5", args: args{a: []float32{1, 2}, b: []float32{1, 1.1}}, want: false, wantErr: nil},
		{name: "testcase6", args: args{a: []float32{1.00000001, 2}, b: []float32{1.00000002, 1}}, want: false, wantErr: nil},
		{name: "testcase7", args: args{a: []float32{1.00000000000000001, 1}, b: []float32{1.00000000000000002, 1}}, want: true, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EqualFloat32(tt.args.a, tt.args.b)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("EqualBool() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("EqualBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualFloat64(t *testing.T) {
	type args struct {
		a []float64
		b []float64
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{name: "testcase1", args: args{a: nil, b: []float64{1.0}}, want: false, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{a: []float64{}, b: []float64{}}, want: false, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{a: []float64{1.1, 2.1}, b: []float64{1.1}}, want: false, wantErr: ErrSliceLengthIsNotEqual},
		{name: "testcase4", args: args{a: []float64{1.01, 2}, b: []float64{1.01, 2}}, want: true, wantErr: nil},
		{name: "testcase5", args: args{a: []float64{1, 2}, b: []float64{1, 1.1}}, want: false, wantErr: nil},
		{name: "testcase6", args: args{a: []float64{1.00000001, 2}, b: []float64{1.00000002, 1}}, want: false, wantErr: nil},
		{name: "testcase7", args: args{a: []float64{1.00000000000000001, 1}, b: []float64{1.00000000000000002, 1}}, want: true, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EqualFloat64(tt.args.a, tt.args.b)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("EqualBool() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("EqualBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualBytes(t *testing.T) {
	type args struct {
		a []byte
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{name: "testcase1", args: args{a: nil, b: []byte{1}}, want: false, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{a: []byte{}, b: []byte{}}, want: false, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{a: []byte{1, 2}, b: []byte{1}}, want: false, wantErr: ErrSliceLengthIsNotEqual},
		{name: "testcase4", args: args{a: []byte{1, 2}, b: []byte{1, 2}}, want: true, wantErr: nil},
		{name: "testcase5", args: args{a: []byte{1, 2}, b: []byte{1, 1}}, want: false, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EqualBytes(tt.args.a, tt.args.b)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("EqualBool() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("EqualBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type testarg struct {
	a int
	b int
}

func testargCmp(a, b interface{}) bool {
	if a.(testarg).a == b.(testarg).a {
		return true
	}

	return false
}

func TestEqualSlice(t *testing.T) {

	type args struct {
		a  []testarg
		b  []testarg
		fn comparator
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{name: "testcase1", args: args{a: []testarg{}, b: []testarg{}, fn: testargCmp}, want: false, wantErr: ErrSliceLengthZero},
		{name: "testcase2", args: args{a: []testarg{{a: 1, b: 2}, {a: 2, b: 4}}, b: []testarg{{a: 1, b: 3}, {a: 2, b: 5}}, fn: testargCmp}, want: true, wantErr: nil},
		{name: "testcase3", args: args{a: []testarg{{a: 1, b: 2}, {a: 2, b: 4}}, b: []testarg{{a: 3, b: 2}, {a: 2, b: 4}}, fn: testargCmp}, want: false, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EqualSlice(tt.args.a, tt.args.b, tt.args.fn)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("EqualSlice() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("EqualSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}
