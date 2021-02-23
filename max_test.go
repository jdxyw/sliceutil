package sliceutil

import "testing"

func TestMaxInt(t *testing.T) {
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
		{name: "testcase3", args: args{s: []int{1, 2, 3, 4, 5}}, want: 5, wantErr: nil},
		{name: "testcase4", args: args{s: []int{-5, -9, -3, -1, -4, -6}}, want: -1, wantErr: nil},
		{name: "testcase5", args: args{s: []int{-1, -2, 3, 4, -9, -8}}, want: 4, wantErr: ErrSliceLengthZero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxInt(tt.args.s)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("MaxInt() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("MaxInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt32(t *testing.T) {
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
		{name: "testcase3", args: args{s: []int32{1, 2, 3, 4, 5}}, want: 5, wantErr: nil},
		{name: "testcase4", args: args{s: []int32{-5, -9, -3, -1, -4, -6}}, want: -1, wantErr: nil},
		{name: "testcase5", args: args{s: []int32{-1, -2, 3, 4, -9, -8}}, want: 4, wantErr: ErrSliceLengthZero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxInt32(tt.args.s)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("MaxInt32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("MaxInt32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt64(t *testing.T) {
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
		{name: "testcase3", args: args{s: []int64{1, 2, 3, 4, 5}}, want: 5, wantErr: nil},
		{name: "testcase4", args: args{s: []int64{-5, -9, -3, -1, -4, -6}}, want: -1, wantErr: nil},
		{name: "testcase5", args: args{s: []int64{-1, -2, 3, 4, -9, -8}}, want: 4, wantErr: ErrSliceLengthZero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxInt64(tt.args.s)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("MaxInt64() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("MaxInt64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxFloat32(t *testing.T) {
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
		{name: "testcase3", args: args{s: []float32{1, 2, 3, 4, 5}}, want: 5, wantErr: nil},
		{name: "testcase4", args: args{s: []float32{-5, -9, -3, -1, -4, -6}}, want: -1, wantErr: nil},
		{name: "testcase5", args: args{s: []float32{-1, -2, 3, 4, -9, -8}}, want: 4, wantErr: ErrSliceLengthZero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxFloat32(tt.args.s)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("MaxFloat32() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("MaxFloat32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxFloat64(t *testing.T) {
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
		{name: "testcase3", args: args{s: []float64{1, 2, 3, 4, 5}}, want: 5, wantErr: nil},
		{name: "testcase4", args: args{s: []float64{-5, -9, -3, -1, -4, -6}}, want: -1, wantErr: nil},
		{name: "testcase5", args: args{s: []float64{-1, -2, 3, 4, -9, -8}}, want: 4, wantErr: ErrSliceLengthZero},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxFloat64(tt.args.s)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("MaxFloat64() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("MaxFloat64() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxByte(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name    string
		args    args
		want    byte
		wantErr error
	}{
		{name: "testcase1", args: args{s: nil}, want: 0, wantErr: ErrSliceIsNil},
		{name: "testcase2", args: args{s: []byte{}}, want: 0, wantErr: ErrSliceLengthZero},
		{name: "testcase3", args: args{s: []byte{1, 2, 3, 4, 5}}, want: 5, wantErr: nil},
		{name: "testcase4", args: args{s: []byte{5, 9, 3, 1, 4, 6}}, want: 9, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MaxByte(tt.args.s)
			if err != nil {
				if err != tt.wantErr {
					t.Errorf("MaxFloat64() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}
			if got != tt.want {
				t.Errorf("MaxFloat64() got = %v, want %v", got, tt.want)
			}
		})
	}
}