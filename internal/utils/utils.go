package utils

import (
	"errors"
	"fmt"
)

var ErrorFlipMapDuplicateKey = errors.New("duplicate key")

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// ChunkSlice split data into parts of sz length, last part's length may be less than sz.
func ChunkSlice(data []string, sz int) [][]string {
	if sz < 1 {
		sz = len(data)
	}

	chunksCnt := len(data) / sz
	if len(data)%sz > 0 {
		chunksCnt++
	}

	result := make([][]string, 0, chunksCnt)
	for i := 0; i*sz < len(data); i++ {
		chunk := make([]string, minInt(sz, len(data)-i*sz))
		copy(chunk, data[i*sz:minInt((i+1)*sz, len(data))])
		result = append(result, chunk)
	}

	return result
}

// FlipMap returns new map where key and values swapped. When given map m has several same values by
// different keys it's panic.
func FlipMap(m map[string]string) map[string]string {
	reverseMap := make(map[string]string)
	for k, v := range m {
		// when keys duplicates, trigger panic to prevent data overwriting
		if _, ok := reverseMap[v]; ok {
			panic(fmt.Errorf("duplicate key %s in reverse map: %w", v, ErrorFlipMapDuplicateKey))
		}

		reverseMap[v] = k
	}

	return reverseMap
}

// FilterByBlacklist returns all elements from data which not in blackList.
func FilterByBlacklist(data []string, blackList []string) []string {
	dict := make(map[string]struct{})
	for _, v := range blackList {
		dict[v] = struct{}{}
	}

	return filterBlackSet(data, dict)
}

func filterBlackSet(data []string, dict map[string]struct{}) []string {
	return filterByPredicate(data, func(s string) bool {
		_, ok := dict[s]
		return !ok
	})
}

// filterByPredicate return all elements from data which gives f(data[i]) true.
func filterByPredicate(data []string, f func(s string) bool) []string {
	filteredSlice := make([]string, 0)

	for i := range data {
		if f(data[i]) {
			filteredSlice = append(filteredSlice, data[i])
		}
	}

	return filteredSlice
}
