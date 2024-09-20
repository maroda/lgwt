package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/maroda/lgwt/blogposts"
)

/*
These stubs can be used this way:

_, err := blogposts.NewPostsFromFS(StubFailingFS{})

but... if we're not doing anything interesting with the error
it's not worth a test. but logging the error might be interesting.
*/
type StubFailingFS struct {
}

// This method fakes a failing filesystem
// The file is nil, the error is !nil
func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("i will always fail")
}

// Test extracting labeled blog parts from a source file
// for publishing into html
func TestNewBlogPosts(t *testing.T) {
	const (
		aBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		bBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello.md": {Data: []byte(aBody)},
		"world.md": {Data: []byte(bBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	// Does it accurately return the correct number of posts created?
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	// Does it match what we expect in the content?
	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
