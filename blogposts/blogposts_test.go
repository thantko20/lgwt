package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/thantko20/blogposts"
)

type StudFailingFS struct{}

func (s *StudFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("something happened!")
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tag1, tag2
---
Hello
and world!`

		secondBody = `Title: Post 2
Description: Description 2
Tags: tag3, tag4
---
Peace
Sympathy
and Love`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFs(fs)
	if err != nil {
		t.Fatal(err)
	}
	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tag1", "tag2"},
		Body: `Hello
and world!`,
	})
	assertPost(t, posts[1], blogposts.Post{
		Title:       "Post 2",
		Description: "Description 2",
		Tags:        []string{"tag3", "tag4"},
		Body: `Peace
Sympathy
and Love`,
	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
