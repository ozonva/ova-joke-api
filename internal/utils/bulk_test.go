// +build test_unit

package utils

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ozonva/ova-joke-api/internal/models"
)

func makeJokeCollection(sz int) []models.Joke {
	jokes := make([]models.Joke, 0, sz)
	for i := 1; i < sz+1; i++ {
		jokes = append(jokes, *models.NewJoke(models.JokeID(i), "joke#"+strconv.Itoa(i), nil))
	}

	return jokes
}

func TestSplitToBulks(t *testing.T) {
	jokes := makeJokeCollection(10)

	type args struct {
		c  []models.Joke
		sz int
	}
	tests := []struct {
		name string
		args args
		want [][]models.Joke
	}{
		{
			name: "simple case",
			args: args{
				c:  jokes[:4], // 0, 1, 2, 3
				sz: 3,
			},
			want: [][]models.Joke{
				jokes[0:3], // 0, 1, 2
				jokes[3:4], // 3
			},
		},
		{
			name: "full filled",
			args: args{
				c:  jokes[:4], // 0, 1, 2, 3
				sz: 2,
			},
			want: [][]models.Joke{
				jokes[:2],  // 0, 1
				jokes[2:4], // 2, 3
			},
		},
		{
			name: "less then chunk size",
			args: args{
				c:  jokes[:4], // 0, 1, 2, 3
				sz: 10,
			},
			want: [][]models.Joke{
				jokes[:4], // 0, 1, 2, 3
			},
		},
		{
			name: "empty",
			args: args{
				c:  []models.Joke{},
				sz: 10,
			},
			want: [][]models.Joke{},
		},
		{
			name: "empty and zero sz",
			args: args{
				c:  []models.Joke{},
				sz: 0,
			},
			want: [][]models.Joke{},
		},
		{
			name: "nil slice",
			args: args{
				c:  nil,
				sz: 10,
			},
			want: [][]models.Joke{},
		},
		{
			name: "negative sz",
			args: args{
				c:  jokes[:4],
				sz: -2,
			},
			want: [][]models.Joke{
				jokes[:4],
			},
		},
		{
			name: "zero sz",
			args: args{
				c:  jokes[:4],
				sz: 0,
			},
			want: [][]models.Joke{
				jokes[:4],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, SplitToBulks(tt.args.c, tt.args.sz))
		})
	}
}

func TestBuildIndex(t *testing.T) {
	jokes := makeJokeCollection(10)

	type args struct {
		c []models.Joke
	}
	tests := []struct {
		name    string
		args    args
		want    map[models.JokeID]models.Joke
		wantErr bool
	}{
		{
			name: "simple case",
			args: args{
				c: jokes[0:3],
			},
			want: map[models.JokeID]models.Joke{
				1: jokes[0],
				2: jokes[1],
				3: jokes[2],
			},
		},
		{
			name: "empty map",
			args: args{
				c: []models.Joke{},
			},
			want: map[models.JokeID]models.Joke{},
		},
		{
			name: "with duplicate key in collection",
			args: args{
				c: append(jokes, jokes[0]),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BuildIndex(tt.args.c)

			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, got)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}
