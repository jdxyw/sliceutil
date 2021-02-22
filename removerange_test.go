package sliceutil

import (
	"reflect"
	"testing"
)

func TestRemoveRangeInt(t *testing.T) {
	type args struct {
		s    []int
		from int
		to   int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr error
	}{
		{name: "testcase1", args: args{s: nil, from: 0, to: 0}, want: nil, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{s: []int{}, from: 0, to: 0}, want: []int{}, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{s: []int{1, 2, 3, 4}, from: -1, to: 0}, want: []int{1, 2, 3, 4}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase4", args: args{s: []int{1, 2, 3, 4}, from: 4, to: 6}, want: []int{1, 2, 3, 4}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase5", args: args{s: []int{1, 2, 3, 4}, from: 2, to: 1}, want: []int{1, 2, 3, 4}, wantErr: ErrSliceRange},
		{name: "testcase6", args: args{s: []int{1, 2, 3, 4, 5, 6}, from: 2, to: 4}, want: []int{1, 2, 5, 6}, wantErr: nil},
		{name: "testcase7", args: args{s: []int{1, 2, 3, 4, 5, 6}, from: 0, to: 6}, want: []int{}, wantErr: nil},
		{name: "testcase8", args: args{s: []int{1, 2, 3, 4, 5, 6}, from: 0, to: 1}, want: []int{2, 3, 4, 5, 6}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveRangeInt(tt.args.s, tt.args.from, tt.args.to)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("RemoveRangeInt() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveRangeInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveRangeInt32(t *testing.T) {
	type args struct {
		s    []int32
		from int
		to   int
	}
	tests := []struct {
		name    string
		args    args
		want    []int32
		wantErr error
	}{
		{name: "testcase1", args: args{s: nil, from: 0, to: 0}, want: nil, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{s: []int32{}, from: 0, to: 0}, want: []int32{}, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{s: []int32{1, 2, 3, 4}, from: -1, to: 0}, want: []int32{1, 2, 3, 4}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase4", args: args{s: []int32{1, 2, 3, 4}, from: 4, to: 6}, want: []int32{1, 2, 3, 4}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase5", args: args{s: []int32{1, 2, 3, 4}, from: 2, to: 1}, want: []int32{1, 2, 3, 4}, wantErr: ErrSliceRange},
		{name: "testcase6", args: args{s: []int32{1, 2, 3, 4, 5, 6}, from: 2, to: 4}, want: []int32{1, 2, 5, 6}, wantErr: nil},
		{name: "testcase7", args: args{s: []int32{1, 2, 3, 4, 5, 6}, from: 0, to: 6}, want: []int32{}, wantErr: nil},
		{name: "testcase8", args: args{s: []int32{1, 2, 3, 4, 5, 6}, from: 0, to: 1}, want: []int32{2, 3, 4, 5, 6}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveRangeInt32(tt.args.s, tt.args.from, tt.args.to)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("RemoveRangeInt32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveRangeInt32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveRangeInt64(t *testing.T) {
	type args struct {
		s    []int64
		from int
		to   int
	}
	tests := []struct {
		name    string
		args    args
		want    []int64
		wantErr error
	}{
		{name: "testcase1", args: args{s: nil, from: 0, to: 0}, want: nil, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{s: []int64{}, from: 0, to: 0}, want: []int64{}, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{s: []int64{1, 2, 3, 4}, from: -1, to: 0}, want: []int64{1, 2, 3, 4}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase4", args: args{s: []int64{1, 2, 3, 4}, from: 4, to: 6}, want: []int64{1, 2, 3, 4}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase5", args: args{s: []int64{1, 2, 3, 4}, from: 2, to: 1}, want: []int64{1, 2, 3, 4}, wantErr: ErrSliceRange},
		{name: "testcase6", args: args{s: []int64{1, 2, 3, 4, 5, 6}, from: 2, to: 4}, want: []int64{1, 2, 5, 6}, wantErr: nil},
		{name: "testcase7", args: args{s: []int64{1, 2, 3, 4, 5, 6}, from: 0, to: 6}, want: []int64{}, wantErr: nil},
		{name: "testcase8", args: args{s: []int64{1, 2, 3, 4, 5, 6}, from: 0, to: 1}, want: []int64{2, 3, 4, 5, 6}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveRangeInt64(tt.args.s, tt.args.from, tt.args.to)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("RemoveRangeInt64() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveRangeInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveRangeByte(t *testing.T) {
	type args struct {
		s    []byte
		from int
		to   int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr error
	}{
		{name: "testcase1", args: args{s: nil, from: 0, to: 0}, want: nil, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{s: []byte{}, from: 0, to: 0}, want: []byte{}, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{s: []byte{1, 2, 3, 4}, from: -1, to: 0}, want: []byte{1, 2, 3, 4}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase4", args: args{s: []byte{1, 2, 3, 4}, from: 4, to: 6}, want: []byte{1, 2, 3, 4}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase5", args: args{s: []byte{1, 2, 3, 4}, from: 2, to: 1}, want: []byte{1, 2, 3, 4}, wantErr: ErrSliceRange},
		{name: "testcase6", args: args{s: []byte{1, 2, 3, 4, 5, 6}, from: 2, to: 4}, want: []byte{1, 2, 5, 6}, wantErr: nil},
		{name: "testcase7", args: args{s: []byte{1, 2, 3, 4, 5, 6}, from: 0, to: 6}, want: []byte{}, wantErr: nil},
		{name: "testcase8", args: args{s: []byte{1, 2, 3, 4, 5, 6}, from: 0, to: 1}, want: []byte{2, 3, 4, 5, 6}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveRangeByte(tt.args.s, tt.args.from, tt.args.to)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("RemoveRangeByte() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveRangeByte() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveRangeBool(t *testing.T) {
	type args struct {
		s    []bool
		from int
		to   int
	}
	tests := []struct {
		name    string
		args    args
		want    []bool
		wantErr error
	}{
		{name: "testcase1", args: args{s: nil, from: 0, to: 0}, want: nil, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{s: []bool{}, from: 0, to: 0}, want: []bool{}, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{s: []bool{true, false, true, false}, from: -1, to: 0}, want: []bool{true, false, true, false}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase4", args: args{s: []bool{true, false, true, false}, from: 4, to: 6}, want: []bool{true, false, true, false}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase5", args: args{s: []bool{true, false, true, false}, from: 2, to: 1}, want: []bool{true, false, true, false}, wantErr: ErrSliceRange},
		{name: "testcase6", args: args{s: []bool{true, false, true, false, false, false}, from: 2, to: 4}, want: []bool{true, false, false, false}, wantErr: nil},
		{name: "testcase7", args: args{s: []bool{true, false, true, false, false, false}, from: 0, to: 6}, want: []bool{}, wantErr: nil},
		{name: "testcase8", args: args{s: []bool{true, false, true, false, false, false}, from: 0, to: 1}, want: []bool{false, true, false, false, false}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveRangeBool(tt.args.s, tt.args.from, tt.args.to)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("RemoveRangeBool() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveRangeBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveRangeFloat32(t *testing.T) {
	type args struct {
		s    []float32
		from int
		to   int
	}
	tests := []struct {
		name    string
		args    args
		want    []float32
		wantErr error
	}{
		{name: "testcase1", args: args{s: nil, from: 0, to: 0}, want: nil, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{s: []float32{}, from: 0, to: 0}, want: []float32{}, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{s: []float32{1.1, 2, 3, 4}, from: -1, to: 0}, want: []float32{1.1, 2, 3, 4}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase4", args: args{s: []float32{1, 2.0, 3.3, 4}, from: 4, to: 6}, want: []float32{1, 2.0, 3.3, 4}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase5", args: args{s: []float32{1, 2.1, 3.2, 4.3}, from: 2, to: 1}, want: []float32{1, 2.1, 3.2, 4.3}, wantErr: ErrSliceRange},
		{name: "testcase6", args: args{s: []float32{1, 2.1, 3.2, 4.3, 5, 6}, from: 2, to: 4}, want: []float32{1, 2.1, 5, 6}, wantErr: nil},
		{name: "testcase7", args: args{s: []float32{1, 2.1, 3.2, 4.3, 5, 6}, from: 0, to: 6}, want: []float32{}, wantErr: nil},
		{name: "testcase8", args: args{s: []float32{1, 2.1, 3.2, 4.3, 5, 6}, from: 0, to: 1}, want: []float32{2.1, 3.2, 4.3, 5, 6}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveRangeFloat32(tt.args.s, tt.args.from, tt.args.to)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("RemoveRangeFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveRangeFloat32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveRangeFloat64(t *testing.T) {
	type args struct {
		s    []float64
		from int
		to   int
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr error
	}{
		{name: "testcase1", args: args{s: nil, from: 0, to: 0}, want: nil, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{s: []float64{}, from: 0, to: 0}, want: []float64{}, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{s: []float64{1.1, 2, 3, 4}, from: -1, to: 0}, want: []float64{1.1, 2, 3, 4}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase4", args: args{s: []float64{1, 2.0, 3.3, 4}, from: 4, to: 6}, want: []float64{1, 2.0, 3.3, 4}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase5", args: args{s: []float64{1, 2.1, 3.2, 4.3}, from: 2, to: 1}, want: []float64{1, 2.1, 3.2, 4.3}, wantErr: ErrSliceRange},
		{name: "testcase6", args: args{s: []float64{1, 2.1, 3.2, 4.3, 5, 6}, from: 2, to: 4}, want: []float64{1, 2.1, 5, 6}, wantErr: nil},
		{name: "testcase7", args: args{s: []float64{1, 2.1, 3.2, 4.3, 5, 6}, from: 0, to: 6}, want: []float64{}, wantErr: nil},
		{name: "testcase8", args: args{s: []float64{1, 2.1, 3.2, 4.3, 5, 6}, from: 0, to: 1}, want: []float64{2.1, 3.2, 4.3, 5, 6}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveRangeFloat64(tt.args.s, tt.args.from, tt.args.to)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("RemoveRangeFloat64() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveRangeFloat64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveRangeString(t *testing.T) {
	type args struct {
		s    []string
		from int
		to   int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr error
	}{
		{name: "testcase1", args: args{s: nil, from: 0, to: 0}, want: nil, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{s: []string{}, from: 0, to: 0}, want: []string{}, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{s: []string{"11", "221", "33", "444"}, from: -1, to: 0}, want: []string{"11", "221", "33", "444"}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase4", args: args{s: []string{"11", "221", "33", "444"}, from: 4, to: 6}, want: []string{"11", "221", "33", "444"}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase5", args: args{s: []string{"11", "221", "33", "444"}, from: 2, to: 1}, want: []string{"11", "221", "33", "444"}, wantErr: ErrSliceRange},
		{name: "testcase6", args: args{s: []string{"11", "221", "33", "444", "555", "666"}, from: 2, to: 4}, want: []string{"11", "221", "555", "666"}, wantErr: nil},
		{name: "testcase7", args: args{s: []string{"11", "221", "33", "444", "555", "666"}, from: 0, to: 6}, want: []string{}, wantErr: nil},
		{name: "testcase8", args: args{s: []string{"11", "221", "33", "444", "555", "666"}, from: 0, to: 1}, want: []string{"221", "33", "444", "555", "666"}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveRangeString(tt.args.s, tt.args.from, tt.args.to)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("RemoveRangeString() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveRangeString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
