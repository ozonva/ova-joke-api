// Package generates fixtures for jokes.
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ozonva/ova-joke-api/internal/models"
)

var (
	jokesFile string
	outFile   string
)

func init() {
	flag.StringVar(&jokesFile, "jokes", "tools/jokegen/jokes.txt", "file with list of jokes")
	flag.StringVar(&outFile, "out", "tools/jokegen/generated/jokes.json", "file with list of jokes objects")
}

// readFile given by path into []string with line by line split.
func readFile(path string) (result []string, rerr error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %s, %w", path, err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			// not overwrite open/read error with close one
			if rerr == nil {
				rerr = err
			}
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		result = append(result, s.Text())
	}

	if s.Err() != nil {
		return nil, fmt.Errorf("unable to read file with scanner: %w", s.Err())
	}

	return result, nil
}

// makeJokeCollection generates models.Joke objects with texts from jokes slice and authors from ac.
func makeJokeCollection(jokes []string, maxAuthorID int) []models.Joke {
	collection := make([]models.Joke, 0, len(jokes))
	for i, text := range jokes {
		collection = append(collection, *models.NewJoke(models.JokeID(i+1), text, models.AuthorID((i)%maxAuthorID+1)))
	}

	return collection
}

// writeJokesAsJSON serialize []models.Joke into JSON and write to file.
func writeJokesAsJSON(path string, data []models.Joke) error {
	content, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("marshal failed: %w", err)
	}

	if err := ioutil.WriteFile(path, content, 0o600); err != nil {
		return fmt.Errorf("write into file %q failed: %w", path, err)
	}

	return nil
}

//go:generate go run jokegen.go -jokes=jokes.txt -out=generated/jokes.json
func main() {
	flag.Parse()

	const (
		JOKES = "jokes"
	)

	type fileData struct {
		path string
		data []string
	}

	files := map[string]*fileData{
		JOKES: {path: jokesFile},
	}

	for t, file := range files {
		data, err := readFile(file.path)
		if err != nil {
			panic(err)
		}

		files[t].data = data
	}

	jokeCollection := makeJokeCollection(files[JOKES].data, 10)

	err := writeJokesAsJSON(outFile, jokeCollection)
	if err != nil {
		panic(err)
	}
}
