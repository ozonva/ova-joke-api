// +build test_unit

package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAuthorString(t *testing.T) {
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
			name: "only AuthorID",
			args: args{
				a: Author{ID: 42},
			},
			want: "AuthorID: 42",
		},
		{
			name: "only Name",
			args: args{
				a: Author{Name: "lev1828"},
			},
			want: "Name: \"lev1828\"",
		},
		{
			name: "simple",
			args: args{
				a: Author{
					ID:   42,
					Name: "lev1828",
				},
			},
			want: "AuthorID: 42, Name: \"lev1828\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.args.a.String())
		})
	}
}

func TestAuthorCopyright(t *testing.T) {
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
			require.Equal(t, tt.want, tt.args.a.Copyright())
		})
	}
}

func TestNewAuthor(t *testing.T) {
	type args struct {
		id   AuthorID
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
			require.Equal(t, tt.want, NewAuthor(tt.args.id, tt.args.name))
		})
	}
}
