package blogrenderer

import (
	"fmt"
	"io"

	blogposts "github.com/basokant/go-with-tests/files"
)

func Render(w io.Writer, p blogposts.Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>\n<p>%s</p>\n", p.Title, p.Description)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(w, "Tags: <ul>")
	if err != nil {
		return err
	}

	for _, tag := range p.Tags {
		_, err = fmt.Fprintf(w, "<li>%s</li>", tag)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprint(w, "</ul>")
	if err != nil {
		return err
	}

	return err
}
