package main

import (
	"fmt"
	"html/template"
	"strings"
)

const (
	happyBuilder   = `data:image/png;base64,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`
	sadBuilder     = `data:image/png;base64,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`
	workingBuilder = `data:image/png;base64,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`

	imgHappyBuilder   = `https://assets-cdn.github.com/images/icons/emoji/unicode/1f49a.png`
	imgSadBuilder     = `https://assets-cdn.github.com/images/icons/emoji/unicode/1f534.png`
	imgWorkingBuilder = `https://assets-cdn.github.com/images/icons/emoji/unicode/2754.png`
)

var templates = template.New("pages")

func init() {
	templates.Funcs(template.FuncMap(map[string]interface{}{
		"builderImg": func(b *Build) template.HTML {
			return template.HTML(builderImgForBuild(b))
		},
		"revLink": func(b *Build) template.HTML {
			return template.HTML(githubRevisionLink(b))
		},
	}))
	template.Must(templates.New("index.html").Parse(`
      <html>
      <head>
      <title>All builds: jekyll-build-server</title>
      <style type="text/css">
        body { font-family: monospace; }
      </style>
      </head>
      <body>
        all builds: {{range .}}<br>
        {{builderImg .}} <a href="/{{.Id}}">{{.Id}}</a> created at {{.CreatedAt}}{{end}}
      </body>
      </html>
    `))
	template.Must(templates.New("build.show.html").Parse(`
      <html>
      <head>
      <title>Build {{.Id}}: jekyll-build-server</title>
      <style type="text/css">
        body { font-family: monospace; }
      </style>
      </head>
      <body>
      <p>&larr; <a href="/">all builds </a></p>
      <h3>{{.Id}}</h3>
      <p>
          revision: {{revLink .}}<br>
          created: {{.CreatedAt}}<br>
          completed: {{.CompletedAt}}<br>
          success: {{.Success}} {{builderImg .}}<br>
          output:
      </p>
      <pre>{{.Output}}</pre>
      </body>
      </html>
    `))
}

func builderImgForBuild(b *Build) string {
	return fmt.Sprintf(
		"<img width='15' height='15' style='vertical-align: middle' data-success='%v' src='%s'>",
		b.Success,
		builderIconForBuild(b))
}

func builderIconForBuild(b *Build) string {
	if b.Success {
		return happyBuilder
	} else {
		if b.CompletedAt == "" {
			return workingBuilder
		} else {
			return sadBuilder
		}
	}
}

func githubRevisionLink(b *Build) string {
	pieces := strings.Split(b.Id, "-")
	revision := pieces[len(pieces)-1]
	repo := strings.Join(pieces[0:len(pieces)-1], "-")
	return linkTo(fmt.Sprintf("https://github.com/%s/commit/%s", repo, revision), revision)
}

func linkTo(url, text string) string {
	if text == "" {
		text = url
	}
	return fmt.Sprintf("<a href='%s'>%s</a>", url, text)
}
