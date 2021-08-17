package author

import "fmt"

type ID = uint64

type Author struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
}

func (a Author) String() string {
	if a.ID == 0 {
		if a.Name == "" {
			return ""
		}

		return fmt.Sprintf("Name: %q", a.Name)
	}

	if a.Name == "" {
		return fmt.Sprintf("ID: %d", a.ID)
	}

	return fmt.Sprintf("ID: %d, Name: %q", a.ID, a.Name)
}

func (a Author) Copyright() string {
	if a.Name != "" {
		return fmt.Sprintf("©%s", a.Name)
	}

	return ""
}

func New(id ID, name string) *Author {
	return &Author{
		ID:   id,
		Name: name,
	}
}
