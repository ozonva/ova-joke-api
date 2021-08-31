// +build test_unit

package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJokeString(t *testing.T) {
	type args struct {
		j Joke
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ZV joke",
			args: args{
				j: Joke{},
			},
			want: "",
		},
		{
			name: "joke with zv author",
			args: args{
				j: Joke{
					ID:     12,
					Author: &Author{},
					Text: "Since 1940, the year Chuck Norris was born, roundhouse kick" +
						" related deaths have increased 13,000 percent.",
				},
			},
			want: "\"Since 1940, the year Chuck Norris was born, roundhouse kick" +
				" related deaths have increased 13,000 percent.\"",
		},
		{
			name: "joke with normal author",
			args: args{
				j: Joke{
					ID: 12,
					Author: &Author{
						ID:   34,
						Name: "Sasha99",
					},
					Text: "Chuck Norris once roundhouse kicked a coal mine and turned it into a diamond mine.",
				},
			},
			want: "\"Chuck Norris once roundhouse kicked a coal mine and turned it into a diamond mine.\" ©Sasha99",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.args.j.String())
		})
	}
}

func TestJokeMarshal(t *testing.T) {
	type args struct {
		j Joke
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ZV joke",
			args: args{
				j: Joke{},
			},
			want: `{"id":0,"text":""}`,
		},
		{
			name: "joke with zv author",
			args: args{
				j: Joke{
					ID: 12,
					Text: "Since 1940, the year Chuck Norris was born, roundhouse kick related deaths have" +
						" increased 13,000 percent.",
				},
			},
			want: `{"id":12,"text":"Since 1940, the year Chuck Norris was born, roundhouse kick related deaths have` +
				` increased 13,000 percent."}`,
		},
		{
			name: "joke with normal author",
			args: args{
				j: Joke{
					ID: 12,
					Author: &Author{
						ID:   34,
						Name: "Sasha99",
					},
					Text: "Chuck Norris once roundhouse kicked a coal mine and turned it into a diamond mine.",
				},
			},
			want: `{"id":12,"text":"Chuck Norris once roundhouse kicked a coal mine and turned it into a` +
				` diamond mine.","author":{"id":34,"name":"Sasha99"}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.args.j)

			require.NoError(t, err)
			require.Equal(t, tt.want, string(got))
		})
	}
}

func TestNewJoke(t *testing.T) {
	a := NewAuthor(1, "author#1")

	type args struct {
		id   JokeID
		text string
		a    *Author
	}
	tests := []struct {
		name string
		args args
		want *Joke
	}{
		{
			name: "zv joke",
			args: args{},
			want: &Joke{
				ID:     0,
				Text:   "",
				Author: nil,
			},
		},
		{
			name: "simple joke",
			args: args{
				id:   1,
				text: "some joke",
				a:    a,
			},
			want: &Joke{
				ID:     1,
				Text:   "some joke",
				Author: a,
			},
		},
		{
			name: "without author",
			args: args{
				id:   1,
				text: "some joke",
			},
			want: &Joke{
				ID:     1,
				Text:   "some joke",
				Author: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, NewJoke(tt.args.id, tt.args.text, tt.args.a))
		})
	}
}
