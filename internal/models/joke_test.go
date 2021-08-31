//go:build test_unit
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
			name: "joke with zv ID",
			args: args{
				j: Joke{
					Text: "Since 1940, the year Chuck Norris was born, roundhouse kick" +
						" related deaths have increased 13,000 percent.",
					AuthorID: 24,
				},
			},
			want: "Text: \"Since 1940, the year Chuck Norris was born, roundhouse kick" +
				" related deaths have increased 13,000 percent.\", AuthorID: 24",
		},
		{
			name: "joke with zv author",
			args: args{
				j: Joke{
					ID: 12,
					Text: "Since 1940, the year Chuck Norris was born, roundhouse kick" +
						" related deaths have increased 13,000 percent.",
				},
			},
			want: "ID: 12, Text: \"Since 1940, the year Chuck Norris was born, roundhouse kick" +
				" related deaths have increased 13,000 percent.\"",
		},
		{
			name: "joke with zv text",
			args: args{
				j: Joke{
					ID:       12,
					AuthorID: 24,
				},
			},
			want: "ID: 12, AuthorID: 24",
		},
		{
			name: "joke full filled",
			args: args{
				j: Joke{
					ID:       12,
					Text:     "Chuck Norris once roundhouse kicked a coal mine and turned it into a diamond mine.",
					AuthorID: 24,
				},
			},
			want: "ID: 12, Text: \"Chuck Norris once roundhouse kicked a coal mine and turned it into a diamond mine.\", AuthorID: 24",
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
					ID:       12,
					Text:     "Chuck Norris once roundhouse kicked a coal mine and turned it into a diamond mine.",
					AuthorID: AuthorID(24),
				},
			},
			want: `{"id":12,"text":"Chuck Norris once roundhouse kicked a coal mine and turned it into a` +
				` diamond mine.","authorId":24}`,
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
	type args struct {
		id       JokeID
		text     string
		authorID AuthorID
	}
	tests := []struct {
		name string
		args args
		want *Joke
	}{
		{
			name: "zv joke",
			args: args{},
			want: &Joke{},
		},
		{
			name: "simple joke",
			args: args{
				id:       1,
				text:     "some joke",
				authorID: 24,
			},
			want: &Joke{
				ID:       1,
				Text:     "some joke",
				AuthorID: 24,
			},
		},
		{
			name: "without id",
			args: args{
				text:     "some joke",
				authorID: 24,
			},
			want: &Joke{
				AuthorID: 24,
				Text:     "some joke",
			},
		},
		{
			name: "without text",
			args: args{
				id:       1,
				authorID: 24,
			},
			want: &Joke{
				ID:       1,
				AuthorID: 24,
			},
		},
		{
			name: "without author",
			args: args{
				id:   1,
				text: "some joke",
			},
			want: &Joke{
				ID:   1,
				Text: "some joke",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, NewJoke(tt.args.id, tt.args.text, tt.args.authorID))
		})
	}
}
