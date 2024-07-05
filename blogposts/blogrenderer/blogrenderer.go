package blogrenderer

import (
	"embed"
	"io"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/thantko20/blogposts"
)

var (
	//go:embed templates/*
	postTemplates embed.FS
)

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	parserExtenstions := parser.CommonExtensions | parser.OrderedListStart | parser.Autolink
	mdParser := parser.NewWithExtensions(parserExtenstions)
	return &PostRenderer{templ, mdParser}, err
}

func (p *PostRenderer) Render(w io.Writer, post blogposts.Post) error {
	err := p.templ.ExecuteTemplate(w, "blog.gohtml", newPostViewModel(post, p.mdParser))
	return err
}

func (p *PostRenderer) RenderIndex(w io.Writer, posts []blogposts.Post) error {
	err := p.templ.ExecuteTemplate(w, "index.gohtml", posts)
	return err
}

type postViewModel struct {
	blogposts.Post
	HTMLBody string
}

func newPostViewModel(post blogposts.Post, mdParser *parser.Parser) postViewModel {
	pvm := postViewModel{Post: post}
	pvm.HTMLBody = string(markdown.ToHTML([]byte(post.Body), mdParser, nil))
	return pvm
}
