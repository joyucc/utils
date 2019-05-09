package gstring

import (
	"reflect"
	"testing"
)

func TestS_String(t *testing.T) {
	tests := []struct {
		name string
		s    S
		want string
	}{
		{
			"test",
			"test",
			"test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("S.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestS_Bytes(t *testing.T) {
	tests := []struct {
		name string
		s    S
		want []byte
	}{
		{
			"test",
			"test",
			[]byte("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("S.Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestS_Int64(t *testing.T) {
	tests := []struct {
		name string
		s    S
		want int64
	}{
		{
			"Int64",
			"123456",
			123456,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Int64(); got != tt.want {
				t.Errorf("S.Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestS_Int(t *testing.T) {
	tests := []struct {
		name string
		s    S
		want int
	}{
		{
			"Int64",
			"123456",
			123456,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Int(); got != tt.want {
				t.Errorf("S.Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestS_Uint(t *testing.T) {
	tests := []struct {
		name string
		s    S
		want uint
	}{
		{
			"Int64",
			"123456",
			123456,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Uint(); got != tt.want {
				t.Errorf("S.Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestS_Uint64(t *testing.T) {
	tests := []struct {
		name string
		s    S
		want uint64
	}{
		{
			"Int64",
			"123456",
			123456,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Uint64(); got != tt.want {
				t.Errorf("S.Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestS_Float64(t *testing.T) {
	tests := []struct {
		name string
		s    S
		want float64
	}{
		{
			"Int64",
			"123456",
			123456,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Float64(); got != tt.want {
				t.Errorf("S.Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestS_ToJSON(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		s       S
		args    args
		wantErr bool
	}{
		{
			"Int64",
			"json",
			struct{ v interface{} }{v: 100},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.ToJSON(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("S.ToJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
