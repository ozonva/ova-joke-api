package author

import "fmt"

type ID = uint64

type Author struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
}

func String(a Author) string {
	return a.Name
}

func Copyright(a Author) string {
	if author := String(a); author != "" {
		return fmt.Sprintf("©%s", String(a))
	}

	return ""
}

func New(id ID, name string) *Author {
	return &Author{
		ID:   id,
		Name: name,
	}
}
