// +build test_unit

package utils

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/ozonva/ova-joke-api/internal/domain/joke"
)

func makeJokeCollection(sz int) joke.Collection {
	jokes := make(joke.Collection, 0, sz)
	for i := 0; i < 10; i++ {
		jokes = append(jokes, joke.New(joke.ID(i+1), "joke#"+strconv.Itoa(i), nil))
	}

	return jokes
}

func TestSplitToBulks(t *testing.T) {
	jokes := makeJokeCollection(10)

	type args struct {
		c  joke.Collection
		sz int
	}
	tests := []struct {
		name string
		args args
		want []joke.Collection
	}{
		{
			name: "simple case",
			args: args{
				c:  jokes[:4], // 0, 1, 2, 3
				sz: 3,
			},
			want: []joke.Collection{
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
			want: []joke.Collection{
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
			want: []joke.Collection{
				jokes[:4], // 0, 1, 2, 3
			},
		},
		{
			name: "empty",
			args: args{
				c:  joke.Collection{},
				sz: 10,
			},
			want: []joke.Collection{},
		},
		{
			name: "empty and zero sz",
			args: args{
				c:  joke.Collection{},
				sz: 0,
			},
			want: []joke.Collection{},
		},
		{
			name: "nil slice",
			args: args{
				c:  nil,
				sz: 10,
			},
			want: []joke.Collection{},
		},
		{
			name: "negative sz",
			args: args{
				c:  jokes[:4],
				sz: -2,
			},
			want: []joke.Collection{
				jokes[:4],
			},
		},
		{
			name: "zero sz",
			args: args{
				c:  jokes[:4],
				sz: 0,
			},
			want: []joke.Collection{
				jokes[:4],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitToBulks(tt.args.c, tt.args.sz); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitToBulks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildIndex(t *testing.T) {
	jokes := makeJokeCollection(10)

	type args struct {
		c joke.Collection
	}
	tests := []struct {
		name    string
		args    args
		want    map[joke.ID]*joke.Joke
		wantErr bool
	}{
		{
			name: "simple case",
			args: args{
				c: jokes[0:3],
			},
			want: map[joke.ID]*joke.Joke{
				1: jokes[0],
				2: jokes[1],
				3: jokes[2],
			},
		},
		{
			name: "empty map",
			args: args{
				c: joke.Collection{},
			},
			want: map[joke.ID]*joke.Joke{},
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
			if (err != nil) != tt.wantErr {
				t.Errorf("BuildIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildIndex() got = %v, want %v", got, tt.want)
			}
		})
	}
}