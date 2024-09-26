package blogrenderer_test

import (
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/maroda/blogrenderer"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := blogrenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		// This is a way to create a "golden copy" of a successful test to refer to,
		// instead of doing all the got/want and having to get the string right.
		// This can match templates to templates.
		approvals.VerifyString(t, buf.String())

		/*
			err := blogrenderer.Render(&buf, aPost)

			if err != nil {
				t.Fatal(err)
			}

			got := buf.String()
			want := `<h1>hello world</h1><p>This is a description</p>Tags: <ul><li>go</li><li>tdd</li></ul>`
			if got != want {
				t.Errorf("got '%s' want '%s'", got, want)
			}
		*/
	})
}
