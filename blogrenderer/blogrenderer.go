package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

const (
	// This template should match 'want'
	postTemplate = `<h2>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
)

var (
	//go:embed templates/*
	postsT embed.FS
)

func Render(w io.Writer, p Post) error {
	// templ, err := template.New("blog").Parse(postTemplate)
	templ, err := template.ParseFS(postsT, "templates/*.gohtml")
	if err != nil {
		return err
	}

	// if err := templ.Execute(w, p); err != nil {
	if err := templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil

	// The above html/template rendering does exactly what this does:
	/*
		_, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(w, "<p>%s</p>", p.Description)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(w, "Tags: <ul>")
		if err != nil {
			return err
		}

		for _, tag := range p.Tags {
			_, err = fmt.Fprintf(w, "<li>%s</li>", tag)
			if err != nil {
				return err
			}
		}

		_, err = fmt.Fprint(w, "</ul>")
		if err != nil {
			return err
		}

		return err
	*/
}
