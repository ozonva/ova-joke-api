package models

import (
	"fmt"
)

// compile time interface check.
var _ fmt.Stringer = Joke{}

type (
	JokeID     = uint64
	Collection = []*Joke
)

type Joke struct {
	ID     JokeID  `json:"id"`
	Text   string  `json:"text"`
	Author *Author `json:"author,omitempty"`
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

func NewJoke(id JokeID, text string, a *Author) *Joke {
	return &Joke{
		ID:     id,
		Text:   text,
		Author: a,
	}
}
