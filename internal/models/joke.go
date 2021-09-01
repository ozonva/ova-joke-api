package models

import (
	"fmt"
	"strings"
)

type (
	JokeID   = uint64
	AuthorID = uint64
)

type Joke struct {
	ID       uint64 `json:"id"`
	Text     string `json:"text"`
	AuthorID uint64 `json:"authorId,omitempty"`
}

func (j Joke) String() string {
	sb := strings.Builder{}
	if j.ID != 0 {
		sb.WriteString(fmt.Sprintf("ID: %d", j.ID))
	}

	if j.Text != "" {
		if sb.Len() > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("Text: %q", j.Text))
	}

	if j.AuthorID != 0 {
		if sb.Len() > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("AuthorID: %d", j.AuthorID))
	}

	return sb.String()
}

func NewJoke(id JokeID, text string, authorID uint64) *Joke {
	return &Joke{
		ID:       id,
		Text:     text,
		AuthorID: authorID,
	}
}
