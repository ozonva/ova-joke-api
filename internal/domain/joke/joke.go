package joke

import (
	"fmt"

	"github.com/ozonva/ova-joke-api/internal/domain/author"
)

// compile time interface check
var _ fmt.Stringer = Joke{}

type (
	ID         = uint64
	Collection = []*Joke
)

type Joke struct {
	ID     ID             `json:"id"`
	Text   string         `json:"text"`
	Author *author.Author `json:"author,omitempty"`
}

func (j Joke) String() string {
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
