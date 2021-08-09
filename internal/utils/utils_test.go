// +build test_unit

package utils

import (
	"errors"
	"reflect"
	"testing"
)

func Test_chunkSlice(t *testing.T) {
	type args struct {
		data []string
		sz   int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "simple case",
			args: args{
				data: []string{"one", "two", "three", "four"},
				sz:   3,
			},
			want: [][]string{
				{"one", "two", "three"},
				{"four"},
			},
		},
		{
			name: "full filled",
			args: args{
				data: []string{"one", "two", "three", "four"},
				sz:   2,
			},
			want: [][]string{
				{"one", "two"},
				{"three", "four"},
			},
		},
		{
			name: "less then chunk size",
			args: args{
				data: []string{"one", "two", "three", "four"},
				sz:   10,
			},
			want: [][]string{
				{"one", "two", "three", "four"},
			},
		},
		{
			name: "empty",
			args: args{
				data: []string{},
				sz:   10,
			},
			want: [][]string{},
		},
		{
			name: "nil slice",
			args: args{
				data: nil,
				sz:   10,
			},
			want: [][]string{},
		},
		{
			name: "negative sz",
			args: args{
				data: []string{"one", "two", "three", "four"},
				sz:   -2,
			},
			want: [][]string{
				{"one", "two", "three", "four"},
			},
		},
		{
			name: "zero sz",
			args: args{
				data: []string{"one", "two", "three", "four"},
				sz:   0,
			},
			want: [][]string{
				{"one", "two", "three", "four"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChunkSlice(tt.args.data, tt.args.sz); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chunkSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlipMap(t *testing.T) {
	type args struct {
		m map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "simple case",
			args: args{
				m: map[string]string{
					"Pushkin": "Aleksandr",
					"Tolstoy": "Lev",
					"Brodsky": "Iosif",
				},
			},
			want: map[string]string{
				"Aleksandr": "Pushkin",
				"Lev":       "Tolstoy",
				"Iosif":     "Brodsky",
			},
		},
		{
			name: "empty map",
			args: args{
				m: map[string]string{},
			},
			want: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlipMap(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlipMap() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("panic on values duplicate", func(t *testing.T) {
		var actualErr error
		func() {
			defer func() {
				if rv := recover(); rv != nil {
					if err, ok := rv.(error); ok {
						actualErr = err
					} else {
						t.Errorf("FlipMap() recover value is not error compartable")
					}
				} else {
					t.Errorf("FlipMap() don't panic, but expected")
				}
			}()
			given := map[string]string{
				"Pushkin": "Aleksandr",
				"Blok":    "Aleksandr",
			}

			FlipMap(given)
		}()

		if !errors.Is(actualErr, ErrorFlipMapDuplicateKey) {
			t.Errorf(
				"FlipMap() panic returns unexpected error type: want=%v, got=%v",
				ErrorFlipMapDuplicateKey,
				actualErr,
			)
		}
	})
}

func Test_filterByPredicate(t *testing.T) {
	lenFilter := func(s string) bool {
		return len(s) < 4
	}

	type args struct {
		data []string
		f    func(s string) bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "simple case",
			args: args{
				data: []string{"one", "two", "three", "four", "five"},
				f:    lenFilter,
			},
			want: []string{"one", "two"},
		},
		{
			name: "empty data",
			args: args{
				data: []string{},
				f:    lenFilter,
			},
			want: []string{},
		},
		{
			name: "with repeats in data",
			args: args{
				data: []string{"one", "two", "three", "one", "six", "two", "seven"},
				f:    lenFilter,
			},
			want: []string{"one", "two", "one", "six", "two"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterByPredicate(tt.args.data, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterByPredicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterBlackSet(t *testing.T) {
	type args struct {
		data []string
		dict map[string]struct{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "simple case",
			args: args{
				data: []string{"one", "two", "three", "four", "five"},
				dict: map[string]struct{}{
					"two":  {},
					"four": {},
				},
			},
			want: []string{"one", "three", "five"},
		},
		{
			name: "with duplicate values",
			args: args{
				data: []string{"four", "one", "two", "three", "two", "four", "one"},
				dict: map[string]struct{}{
					"two":  {},
					"four": {},
				},
			},
			want: []string{"one", "three", "one"},
		},
		{
			name: "with empty dict",
			args: args{
				data: []string{"one", "two", "three", "four", "five"},
				dict: map[string]struct{}{},
			},
			want: []string{"one", "two", "three", "four", "five"},
		},
		{
			name: "with no intersections",
			args: args{
				data: []string{"one", "two", "three", "four", "five"},
				dict: map[string]struct{}{
					"six":   {},
					"seven": {},
				},
			},
			want: []string{"one", "two", "three", "four", "five"},
		},
		{
			name: "with fun intersection",
			args: args{
				data: []string{"one", "two", "one", "one", "two"},
				dict: map[string]struct{}{
					"one": {},
					"two": {},
				},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filterBlackSet(tt.args.data, tt.args.dict); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterBlackSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterByBlacklist(t *testing.T) {
	type args struct {
		data   []string
		values []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "simple case",
			args: args{
				data:   []string{"one", "two", "three", "four", "five"},
				values: []string{"three", "four", "five"},
			},
			want: []string{"one", "two"},
		},
		{
			name: "empty data",
			args: args{
				data:   []string{},
				values: []string{"one", "two"},
			},
			want: []string{},
		},
		{
			name: "nil data",
			args: args{
				data:   nil,
				values: []string{"one", "two"},
			},
			want: []string{},
		},
		{
			name: "nil black list",
			args: args{
				data:   []string{"one", "two"},
				values: nil,
			},
			want: []string{"one", "two"},
		},
		{
			name: "empty black list",
			args: args{
				data:   []string{"one", "two", "three", "four", "five"},
				values: []string{},
			},
			want: []string{"one", "two", "three", "four", "five"},
		},
		{
			name: "with duplicates in values",
			args: args{
				data:   []string{"one", "two", "three", "four", "five"},
				values: []string{"one", "one", "one"},
			},
			want: []string{"two", "three", "four", "five"},
		},
		{
			name: "with duplicates in data",
			args: args{
				data:   []string{"one", "two", "one", "four", "one", "five"},
				values: []string{"one"},
			},
			want: []string{"two", "four", "five"},
		},
		{
			name: "with no intersections",
			args: args{
				data:   []string{"one", "two", "three", "four", "five"},
				values: []string{"six", "seven"},
			},
			want: []string{"one", "two", "three", "four", "five"},
		},
		{
			name: "with full intersection",
			args: args{
				data:   []string{"one", "two", "one", "two", "one"},
				values: []string{"one", "two"},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterByBlacklist(tt.args.data, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterByBlacklist() = %v, want %v", got, tt.want)
			}
		})
	}
}
