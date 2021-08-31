//go:build test_unit
// +build test_unit

package hellower

import (
	"bytes"
	"fmt"
	"testing"
)

type ErrorWriter struct{}

var alwaysWriteError = fmt.Errorf("alwayse write error")

func (e ErrorWriter) Write(_ []byte) (int, error) {
	return 0, alwaysWriteError
}

func Test_makeHighlighted(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple string",
			args: args{
				s: "simple string",
			},
			want: "\033[0;32msimple string\033[0m",
		},
		{
			name: "empty string",
			args: args{
				s: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeHighlighted(tt.args.s); got != tt.want {
				t.Errorf("makeHighlighted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSayHelloFrom(t *testing.T) {
	type args struct {
		from string
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name: "simple say hello",
			args: args{
				from: "Chuck Norris",
			},
			wantW:   "Hello from \033[0;32mChuck Norris\033[0m! 🚀\n",
			wantErr: false,
		},
		{
			name: "say hello from nobody",
			args: args{
				from: "",
			},
			wantW:   "Hello from unknown! 🚀\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := SayHelloFrom(w, tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("SayHelloFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("SayHelloFrom() gotW = %v, want %v", gotW, tt.wantW)
			}
		})
	}

	t.Run("Fail writer", func(t *testing.T) {
		w := ErrorWriter{}

		err := SayHelloFrom(w, "Chuck is not the best")
		if err == nil {
			t.Errorf("SayHelloFrom() error = %v, wantErr %v", err, alwaysWriteError)
			return
		}
	})
}

func Test_makeHelloPhrase(t *testing.T) {
	type args struct {
		from string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple string",
			args: args{
				from: "simple string",
			},
			want: "Hello from simple string! 🚀\n",
		},
		{
			name: "empty string",
			args: args{
				from: "",
			},
			want: "Hello from unknown! 🚀\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeHelloPhrase(tt.args.from); got != tt.want {
				t.Errorf("makeHelloPhrase() = %v, want %v", got, tt.want)
			}
		})
	}
}
