package utils

import (
	"fmt"

	"github.com/ozonva/ova-joke-api/internal/domain/joke"
)

// SplitToBulks chunk given collection to slices up to sz length.
func SplitToBulks(c []joke.Joke, sz int) [][]joke.Joke {
	if sz < 1 {
		if len(c) == 0 {
			return [][]joke.Joke{}
		}

		sz = len(c)
	}

	// both sz and len(c) != 0
	chunksCnt := (len(c) + sz - 1) / sz

	result := make([][]joke.Joke, 0, chunksCnt)
	for i := 0; i < len(c); i += sz {
		result = append(result, c[i:minInt(i+sz, len(c))])
	}

	return result
}

// BuildIndex convert given collection to map with id as key and collection value as value.
func BuildIndex(c []joke.Joke) (map[joke.ID]joke.Joke, error) {
	result := make(map[joke.ID]joke.Joke)
	for i := range c {
		if _, ok := result[c[i].ID]; ok {
			return nil, fmt.Errorf("%w, value with joke.ID = %d already exists", ErrorDuplicateKey, c[i].ID)
		}
		result[c[i].ID] = c[i]
	}

	return result, nil
}
