package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/ozonva/ova-joke-api/internal/models"
)

var (
	jokesFile string
	namesFile string
	outFile   string
)

func init() {
	flag.StringVar(&jokesFile, "jokes", "tools/jokegen/jokes.txt", "file with list of jokes")
	flag.StringVar(&namesFile, "names", "tools/jokegen/names.txt", "file with list of author's names")
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

// makeUserCollection generates collection of author.Author objects with names from names.
func makeUserCollection(names []string) []*models.Author {
	users := make([]*models.Author, 0, len(names))
	for i, name := range names {
		users = append(users, models.NewAuthor(
			models.AuthorID(i+1),
			name,
		))
	}

	return users
}

// makeJokeCollection generates joke.Joke objects with texts from jokes slice and authors from ac.
func makeJokeCollection(jokes []string, ac []*models.Author) []models.Joke {
	rand.Seed(time.Now().UnixNano())

	collection := make([]models.Joke, 0, len(jokes))
	for i, text := range jokes {
		collection = append(collection, *models.NewJoke(models.JokeID(i+1), text, ac[rand.Intn(len(ac))])) // nolint:gosec
	}

	return collection
}

// writeJokesAsJSON serialize []joke.Joke into JSON and write to file.
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

//go:generate go run jokegen.go -jokes=jokes.txt -names=names.txt -out=generated/jokes.json
func main() {
	flag.Parse()

	const (
		JOKES = "jokes"
		NAMES = "names"
	)

	type fileData struct {
		path string
		data []string
	}

	files := map[string]*fileData{
		JOKES: {path: jokesFile},
		NAMES: {path: namesFile},
	}

	for t, file := range files {
		data, err := readFile(file.path)
		if err != nil {
			panic(err)
		}

		files[t].data = data
	}

	userCollection := makeUserCollection(files[NAMES].data)
	jokeCollection := makeJokeCollection(files[JOKES].data, userCollection)

	err := writeJokesAsJSON(outFile, jokeCollection)
	if err != nil {
		panic(err)
	}
}
