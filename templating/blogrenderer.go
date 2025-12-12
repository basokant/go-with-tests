package blogrenderer

import (
	"embed"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
	"io"
	"strings"

	blogposts "github.com/basokant/go-with-tests/files"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type postViewModel struct {
	blogposts.Post
	SanitisedTitle string
	HTMLBody       template.HTML
}

func sanitisedTitle(title string) string {
	return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
}

func NewPostViewModel(p blogposts.Post) postViewModel {
	vm := postViewModel{Post: p}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	doc := parser.NewWithExtensions(extensions)

	vm.SanitisedTitle = sanitisedTitle(p.Title)
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), doc, nil))

	return vm
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
	return r.templ.ExecuteTemplate(w, "blog.gohtml", NewPostViewModel(p))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []blogposts.Post) error {
	postViewModels := make([]postViewModel, 0, len(posts))
	for _, post := range posts {
		postViewModels = append(postViewModels, NewPostViewModel(post))
	}

	return r.templ.ExecuteTemplate(w, "index.gohtml", postViewModels)
}
