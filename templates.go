package main

import (
	"html/template"
	"io"
)

// indextmpl is the HTML template to generate for the index page of the static
// site (route "/").
var indextmpl = template.Must(
	template.New("index").Parse(`<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta http-equiv="content-type" content="text/html" charset="UTF-8">
        <title>{{.Hostname}}</title>
		<meta name="generator" content="staticgovanityurls (https://staticgovanityurls.nicolaspl.com)">
		<meta name="viewport" content="width=device-width,height=device-height,initial-scale=1.0,user-scalable=yes">
    </head>
    <body>
		<h1>{{.Hostname}}</h1>
		<ul>
			{{range .Paths}}<li><a href="https://pkg.go.dev/{{.}}">{{.}}</a></li>{{end}}
		</ul>
    </body>
</html>`),
)

// executeIndex generates the Index template using the given variables.
// paths is a list of import path (containing both hostname and prefix).
func executeIndex(o io.Writer, hostname string, paths []string) error {
	return indextmpl.Execute(o, struct {
		Hostname string
		Paths    []string
	}{
		Hostname: hostname,
		Paths:    paths,
	})
}

// pathtmpl is the HTML template to generate for the page of a module.
var pathtmpl = template.Must(
	template.New("path").Parse(`<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta http-equiv="content-type" content="text/html" charset="UTF-8">
			<meta name="generator" content="staticgovanityurls (https://staticgovanityurls.nicolaspl.com)">
			<meta name="viewport" content="width=device-width,height=device-height,initial-scale=1.0,user-scalable=yes">
			<meta name="go-import" content="{{.Prefix}} {{.Vcs}} {{.Repo}}">
			<meta name="go-source" content="{{.Prefix}} {{.Repo}} {{.Dir}} {{.File}}">
			<meta http-equiv="refresh" content="0; url=https://pkg.go.dev/{{.Prefix}}">
			<title>{{.Prefix}}</title>
		</head>
		<body>
			<h1>{{.Prefix}}</h1>
			<ul>
				<li><a href="https://pkg.go.dev/{{.Prefix}}">Documentation</a></li>
				<li><a href="{{.Repo}}">Source ({{.Vcs}})</a></li>
			</ul>
		</body>
	</html>`),
)

// executePath generates the path template using the given variables.
func executePath(o io.Writer, prefix string, vcs VCS, repo string, dir string, file string) error {
	return pathtmpl.Execute(o, struct {
		Prefix string
		Repo   string
		Dir    string
		File   string
		Vcs    VCS
	}{
		Prefix: prefix,
		Repo:   repo,
		Vcs:    vcs,
		Dir:    dir,
		File:   file,
	})
}