package sliceutil

import (
	"math"
	"testing"
)

func TestSumInt(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{name: "testcase1", args: args{s: nil}, want: 0, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{s: []int{}}, want: 0, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{s: []int{1, 1, 1, 1}}, want: 4, wantErr: nil},
		{name: "testcase4", args: args{s: []int{1, 1, 1, 1, 5, 5, 5}}, want: 19, wantErr: nil},
		{name: "testcase5", args: args{s: []int{1, 1, 1, 1, -5, -5, -5}}, want: -11, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumInt(tt.args.s)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("SumInt() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("SumInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumInt32(t *testing.T) {
	type args struct {
		s []int32
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr error
	}{
		{name: "testcase1", args: args{s: nil}, want: 0, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{s: []int32{}}, want: 0, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{s: []int32{1, 1, 1, 1}}, want: 4, wantErr: nil},
		{name: "testcase4", args: args{s: []int32{1, 1, 1, 1, 5, 5, 5}}, want: 19, wantErr: nil},
		{name: "testcase5", args: args{s: []int32{1, 1, 1, 1, -5, -5, -5}}, want: -11, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumInt32(tt.args.s)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("SumInt32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("SumInt32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumInt64(t *testing.T) {
	type args struct {
		s []int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr error
	}{
		{name: "testcase1", args: args{s: nil}, want: 0, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{s: []int64{}}, want: 0, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{s: []int64{1, 1, 1, 1}}, want: 4, wantErr: nil},
		{name: "testcase4", args: args{s: []int64{1, 1, 1, 1, 5, 5, 5}}, want: 19, wantErr: nil},
		{name: "testcase5", args: args{s: []int64{1, 1, 1, 1, -5, -5, -5}}, want: -11, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumInt64(tt.args.s)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("SumInt64() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("SumInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumFloat32(t *testing.T) {
	type args struct {
		s []float32
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr error
	}{
		{name: "testcase1", args: args{s: nil}, want: 0, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{s: []float32{}}, want: 0, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{s: []float32{1.1, 1.2, 1.3, 1.1}}, want: 4.7, wantErr: nil},
		{name: "testcase4", args: args{s: []float32{1.1, 1.1, 1.2, 1.2, 5, 5, 5.1}}, want: 19.7, wantErr: nil},
		{name: "testcase5", args: args{s: []float32{1, 1, 1.1, 1.1, -5.1, -5.1, -5.1}}, want: -11.1, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumFloat32(tt.args.s)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("SumFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if math.Abs(float64(got)-float64(tt.want)) > 0.00001 {
				t.Errorf("SumFloat32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumFloat64(t *testing.T) {
	type args struct {
		s []float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr error
	}{
		{name: "testcase1", args: args{s: nil}, want: 0, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{s: []float64{}}, want: 0, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{s: []float64{1.1, 1.2, 1.3, 1.1}}, want: 4.7, wantErr: nil},
		{name: "testcase4", args: args{s: []float64{1.1, 1.1, 1.2, 1.2, 5, 5, 5.1}}, want: 19.7, wantErr: nil},
		{name: "testcase5", args: args{s: []float64{1, 1, 1.1, 1.1, -5.1, -5.1, -5.1}}, want: -11.1, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumFloat64(tt.args.s)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("SumFloat64() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if math.Abs(got-tt.want) > 0.00001 {
				t.Errorf("SumFloat64() got = %v, want %v", got, tt.want)
			}
		})
	}
}
