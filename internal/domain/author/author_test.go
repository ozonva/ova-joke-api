// +build test_unit

package author

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	type args struct {
		a Author
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ZV Author",
			args: args{
				a: Author{},
			},
			want: "",
		},
		{
			name: "simple",
			args: args{
				a: Author{
					ID:   42,
					Name: "lev1828",
				},
			},
			want: "lev1828",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.a); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyright(t *testing.T) {
	type args struct {
		a Author
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ZV author",
			args: args{
				a: Author{},
			},
			want: "",
		},
		{
			name: "normal author",
			args: args{
				a: Author{
					ID:   42,
					Name: "lev1828",
				},
			},
			want: "©lev1828",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Copyright(tt.args.a); got != tt.want {
				t.Errorf("Copyright() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		id   ID
		name string
	}
	tests := []struct {
		name string
		args args
		want *Author
	}{
		{
			name: "zv author",
			args: args{},
			want: &Author{
				ID:   0,
				Name: "",
			},
		},
		{
			name: "simple author",
			args: args{
				id:   1,
				name: "author name #1",
			},
			want: &Author{
				ID:   1,
				Name: "author name #1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.id, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
