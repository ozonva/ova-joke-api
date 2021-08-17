package joke

import (
	"fmt"

	"github.com/ozonva/ova-joke-api/internal/domain/author"
)

type (
	ID         = uint64
	Collection = []*Joke
)

type Joke struct {
	ID     ID             `json:"id"`
	Text   string         `json:"text"`
	Author *author.Author `json:"author,omitempty"`
}

func String(j Joke) string {
	if j.Text == "" {
		return ""
	}

	if copyright := j.Author.Copyright(); copyright != "" {
		return fmt.Sprintf("%q %s", j.Text, copyright)
	}

	return fmt.Sprintf("%q", j.Text)
}

func New(id ID, text string, a *author.Author) *Joke {
	return &Joke{
		ID:     id,
		Text:   text,
		Author: a,
	}
}
