package blogrenderer

import (
	"embed"
	"io"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/thantko20/blogposts"
)

var (
	//go:embed templates/*
	postTemplates embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ}, err
}

func (p *PostRenderer) Render(w io.Writer, post blogposts.Post) error {
	body := markdown.ToHTML([]byte(post.Body), nil, nil)
	post.Body = string(body)
	err := p.templ.ExecuteTemplate(w, "blog.gohtml", post)
	return err
}
