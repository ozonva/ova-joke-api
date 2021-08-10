package hellower

import (
	"fmt"
	"io"
)

func makeHighlighted(s string) string {
	if len(s) == 0 {
		return ""
	}

	return fmt.Sprintf("\033[0;32m%s\033[0m", s)
}

func makeHelloPhrase(s string) string {
	if len(s) == 0 {
		return "Hello from unknown! 🚀\n"
	}

	return fmt.Sprintf("Hello from %s! 🚀\n", s)
}

// SayHelloFrom writes hello message from f into w.
func SayHelloFrom(w io.Writer, f string) error {
	data := []byte(makeHelloPhrase(makeHighlighted(f)))

	if n, err := w.Write(data); err != nil {
		return fmt.Errorf("partial write, %d bytes writed: %w", n, err)
	}

	return nil
}
