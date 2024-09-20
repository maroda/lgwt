package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

// fstest.MapFS implements fs.FS, so it can be used as a test here
func NewPostsFromFS(filesys fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(filesys, ".")
	if err != nil {
		return nil, err
	}

	// Create our slice of posts
	var posts []Post

	// Range through it, and use the name for getPost
	for _, f := range dir {
		post, err := getPost(filesys, f.Name())
		if err != nil {
			return nil, err //todo: needs clarification, should we totally fail if one file fails? or just ignore?
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// getPost could take a fs.DirEntry, but since we're only using it for getting Name,
// we can pass that as a string when we find it in NewPostsFromFS
// func getPost(filesys fs.FS, f fs.DirEntry) (Post, error) {
func getPost(filesys fs.FS, fileName string) (Post, error) {
	postFile, err := filesys.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

// Separator Prefixes used below
const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

// fs.File could be used as the func arg here
// But to decouple it better, fs.File implements io.Reader
// so it's better to use that.
// Uses Scanner to read in a multi-line file.
func newPost(postFile io.Reader) (Post, error) {
	// postData, err := io.ReadAll(postFile)
	// if err != nil {
	//	return Post{}, err
	//}
	scanner := bufio.NewScanner(postFile)

	// Anon function to run a scan on a line and return its text
	//readLine := func() string {
	//	scanner.Scan()
	//	return scanner.Text()
	//}

	// Still readLine, but use the *Separator consts as an input
	// Then use TrimPrefix instead of the slice operator
	readLine := func(prefixName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefixName)
	}

	// Run scanner on the first three labels we know
	title := readLine(titleSeparator)
	description := readLine(descriptionSeparator)
	tags := strings.Split(readLine(tagsSeparator), ", ")

	// The read marker will move with each line calling readLine()
	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        readBody(scanner),
	}, nil

	// return Post{Title: title, Description: description}, nil

	// This cuts out the "Title: " text by slicing the string
	// post := Post{Title: string(postData)[7:]}
	// return post, nil
}

// Separate out functionality to keep code readable
func readBody(s *bufio.Scanner) string {
	s.Scan() // Ignore the separator line (---)
	// Scan the rest of the lines of the file into an io.Writer
	buf := bytes.Buffer{}
	// scanner.Scan() returns a bool, so this will run
	// as long as there are lines to scan
	for s.Scan() {
		fmt.Fprintln(&buf, s.Text())
	}

	// Chop the final newline off the body
	return strings.TrimSuffix(buf.String(), "\n")
}
