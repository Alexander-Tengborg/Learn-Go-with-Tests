package blogrenderer

import (
	"html/template"
	"io"
)

const (
	postTemplate = "<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func Renderer(w io.Writer, post Post) error {
	templ, err := template.New("blog").Parse(postTemplate)
	if err != nil {
		return err
	}

	err = templ.Execute(w, post)
	if err != nil {
		return err
	}

	return nil
}
