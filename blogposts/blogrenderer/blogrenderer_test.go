package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/thantko20/blogposts"
	"github.com/thantko20/blogposts/blogrenderer"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogposts.Post{
			Title:       "A Post",
			Description: "This is a description",
			Body:        "This is a **post**",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post to html", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := postRenderer.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())

	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogposts.Post{{Title: "Hello world"}, {Title: "Goodbye world"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogposts.Post{
			Title:       "A Post",
			Description: "This is a description",
			Body:        "This is a post",
			Tags:        []string{"go", "tdd"},
		}
	)
	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
