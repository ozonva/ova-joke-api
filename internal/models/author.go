package models

import "fmt"

// compile time interface check.
var _ fmt.Stringer = Author{}

type AuthorID = uint64

type Author struct {
	ID   AuthorID `json:"id"`
	Name string   `json:"name"`
}

func (a Author) String() string {
	if a.ID == 0 {
		if a.Name == "" {
			return ""
		}

		return fmt.Sprintf("Name: %q", a.Name)
	}

	if a.Name == "" {
		return fmt.Sprintf("AuthorID: %d", a.ID)
	}

	return fmt.Sprintf("AuthorID: %d, Name: %q", a.ID, a.Name)
}

func (a Author) Copyright() string {
	if a.Name != "" {
		return fmt.Sprintf("©%s", a.Name)
	}

	return ""
}

func NewAuthor(id AuthorID, name string) *Author {
	return &Author{
		ID:   id,
		Name: name,
	}
}
