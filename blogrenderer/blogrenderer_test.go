package blogrenderer_test

import (
	"bytes"
	"testing"

	"github.com/maroda/blogrenderer"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "Hello World",
			Body:        "This is the body of the post.",
			Description: "Post Description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>Hello World</h1>`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

}
