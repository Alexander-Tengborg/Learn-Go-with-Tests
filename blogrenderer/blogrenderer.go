package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"
)

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func (p Post) SanitizedTitle() string {
	return strings.ReplaceAll(strings.ToLower(p.Title), " ", "-")
}

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, post Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", post)

}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}
