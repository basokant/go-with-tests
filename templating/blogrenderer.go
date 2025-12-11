package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	blogposts "github.com/basokant/go-with-tests/files"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostViewModel struct {
	Title, SanitisedTitle, Description, Body string
	Tags                                     []string
}

func sanitisedTitle(title string) string {
	return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
}

func NewPostViewModel(p blogposts.Post) PostViewModel {
	return PostViewModel{
		Title:          p.Title,
		SanitisedTitle: sanitisedTitle(p.Title),
		Description:    p.Description,
		Body:           p.Body,
		Tags:           p.Tags,
	}
}

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p blogposts.Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", p)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []blogposts.Post) error {
	postViewModels := make([]PostViewModel, 0, len(posts))
	for _, post := range posts {
		postViewModels = append(postViewModels, NewPostViewModel(post))
	}

	return r.templ.ExecuteTemplate(w, "index.gohtml", postViewModels)
}
