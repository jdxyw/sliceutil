package sliceutil

import (
	"reflect"
	"testing"
)

func TestRemoveIndexInt(t *testing.T) {
	type args struct {
		s []int
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr error
	}{
		{name: "testcase1", args: args{s: []int{}, i: 0}, want: []int{}, wantErr: ErrSliceLengthZero},
		{name: "testcase2", args: args{s: nil, i: 0}, want: nil, wantErr: ErrSliceIsNil},
		{name: "testcase3", args: args{s: []int{1, 2, 3}, i: 4}, want: []int{1, 2, 3}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase4", args: args{s: []int{1, 2, 3, 4, 5}, i: 0}, want: []int{2, 3, 4, 5}, wantErr: nil},
		{name: "testcase5", args: args{s: []int{1, 2, 3, 4, 5}, i: 4}, want: []int{1, 2, 3, 4}, wantErr: nil},
		{name: "testcase6", args: args{s: []int{1, 2, 3, 4, 5}, i: 2}, want: []int{1, 2, 4, 5}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveIndexInt(tt.args.s, tt.args.i)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("RemoveIndexInt() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIndexInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveIndexInt32(t *testing.T) {
	type args struct {
		s []int32
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []int32
		wantErr error
	}{
		{name: "testcase1", args: args{s: []int32{}, i: 0}, want: []int32{}, wantErr: ErrSliceLengthZero},
		{name: "testcase2", args: args{s: nil, i: 0}, want: nil, wantErr: ErrSliceIsNil},
		{name: "testcase3", args: args{s: []int32{1, 2, 3}, i: 4}, want: []int32{1, 2, 3}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase4", args: args{s: []int32{1, 2, 3, 4, 5}, i: 0}, want: []int32{2, 3, 4, 5}, wantErr: nil},
		{name: "testcase5", args: args{s: []int32{1, 2, 3, 4, 5}, i: 4}, want: []int32{1, 2, 3, 4}, wantErr: nil},
		{name: "testcase6", args: args{s: []int32{1, 2, 3, 4, 5}, i: 2}, want: []int32{1, 2, 4, 5}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveIndexInt32(tt.args.s, tt.args.i)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("RemoveIndexInt() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIndexInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveIndexInt64(t *testing.T) {
	type args struct {
		s []int64
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []int64
		wantErr error
	}{
		{name: "testcase1", args: args{s: []int64{}, i: 0}, want: []int64{}, wantErr: ErrSliceLengthZero},
		{name: "testcase2", args: args{s: nil, i: 0}, want: nil, wantErr: ErrSliceIsNil},
		{name: "testcase3", args: args{s: []int64{1, 2, 3}, i: 4}, want: []int64{1, 2, 3}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase4", args: args{s: []int64{1, 2, 3, 4, 5}, i: 0}, want: []int64{2, 3, 4, 5}, wantErr: nil},
		{name: "testcase5", args: args{s: []int64{1, 2, 3, 4, 5}, i: 4}, want: []int64{1, 2, 3, 4}, wantErr: nil},
		{name: "testcase6", args: args{s: []int64{1, 2, 3, 4, 5}, i: 2}, want: []int64{1, 2, 4, 5}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveIndexInt64(tt.args.s, tt.args.i)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("RemoveIndexInt() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIndexInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveIndexFloat32(t *testing.T) {
	type args struct {
		s []float32
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []float32
		wantErr error
	}{
		{name: "testcase1", args: args{s: []float32{}, i: 0}, want: []float32{}, wantErr: ErrSliceLengthZero},
		{name: "testcase2", args: args{s: nil, i: 0}, want: nil, wantErr: ErrSliceIsNil},
		{name: "testcase3", args: args{s: []float32{1, 2, 3}, i: 4}, want: []float32{1, 2, 3}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase4", args: args{s: []float32{1, 2, 3, 4, 5}, i: 0}, want: []float32{2, 3, 4, 5}, wantErr: nil},
		{name: "testcase5", args: args{s: []float32{1, 2, 3, 4, 5}, i: 4}, want: []float32{1, 2, 3, 4}, wantErr: nil},
		{name: "testcase6", args: args{s: []float32{1, 2, 3, 4, 5}, i: 2}, want: []float32{1, 2, 4, 5}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveIndexFloat32(tt.args.s, tt.args.i)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("RemoveIndexInt() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIndexInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveIndexFloat64(t *testing.T) {
	type args struct {
		s []float64
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		wantErr error
	}{
		{name: "testcase1", args: args{s: []float64{}, i: 0}, want: []float64{}, wantErr: ErrSliceLengthZero},
		{name: "testcase2", args: args{s: nil, i: 0}, want: nil, wantErr: ErrSliceIsNil},
		{name: "testcase3", args: args{s: []float64{1, 2, 3}, i: 4}, want: []float64{1, 2, 3}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase4", args: args{s: []float64{1, 2, 3, 4, 5}, i: 0}, want: []float64{2, 3, 4, 5}, wantErr: nil},
		{name: "testcase5", args: args{s: []float64{1, 2, 3, 4, 5}, i: 4}, want: []float64{1, 2, 3, 4}, wantErr: nil},
		{name: "testcase6", args: args{s: []float64{1, 2, 3, 4, 5}, i: 2}, want: []float64{1, 2, 4, 5}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveIndexFloat64(tt.args.s, tt.args.i)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("RemoveIndexInt() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIndexInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveIndexByte(t *testing.T) {
	type args struct {
		s []byte
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr error
	}{
		{name: "testcase1", args: args{s: []byte{}, i: 0}, want: []byte{}, wantErr: ErrSliceLengthZero},
		{name: "testcase2", args: args{s: nil, i: 0}, want: nil, wantErr: ErrSliceIsNil},
		{name: "testcase3", args: args{s: []byte{1, 2, 3}, i: 4}, want: []byte{1, 2, 3}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase4", args: args{s: []byte{1, 2, 3, 4, 5}, i: 0}, want: []byte{2, 3, 4, 5}, wantErr: nil},
		{name: "testcase5", args: args{s: []byte{1, 2, 3, 4, 5}, i: 4}, want: []byte{1, 2, 3, 4}, wantErr: nil},
		{name: "testcase6", args: args{s: []byte{1, 2, 3, 4, 5}, i: 2}, want: []byte{1, 2, 4, 5}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveIndexByte(tt.args.s, tt.args.i)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("RemoveIndexInt() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIndexInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveIndexBool(t *testing.T) {
	type args struct {
		s []bool
		i int
	}
	tests := []struct {
		name    string
		args    args
		want    []bool
		wantErr error
	}{
		{name: "testcase1", args: args{s: []bool{}, i: 0}, want: []bool{}, wantErr: ErrSliceLengthZero},
		{name: "testcase2", args: args{s: nil, i: 0}, want: nil, wantErr: ErrSliceIsNil},
		{name: "testcase3", args: args{s: []bool{true, false, true}, i: 4}, want: []bool{true, false, true}, wantErr: ErrSliceIndexOutOfBound},
		{name: "testcase4", args: args{s: []bool{true, false, true, false, false}, i: 0}, want: []bool{false, true, false, false}, wantErr: nil},
		{name: "testcase5", args: args{s: []bool{true, false, true, false, false}, i: 4}, want: []bool{true, false, true, false}, wantErr: nil},
		{name: "testcase6", args: args{s: []bool{true, false, true, false, false}, i: 2}, want: []bool{true, false, false, false}, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RemoveIndexBool(tt.args.s, tt.args.i)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("RemoveIndexInt() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIndexInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
