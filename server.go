package main

import (
	"bytes"
	goTemplate "html/template"
	"log"
	"net/http"
	"strings"
)

func serve(config appConfig) {
	http.HandleFunc("/", makeRedirector(config))
	log.Fatal(http.ListenAndServe(config.listenAddress, nil))
}

func urlExists(url string) bool {
	/* #nosec */
	if resp, err := http.Head(url); err != nil {
		log.Println(err)
	} else {
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			return true
		}
	}
	return false
}

func makeRedirector(config appConfig) http.HandlerFunc {
	var redirectHTML = `<!DOCTYPE html>
<html lang="en">
<meta charset="utf-8">
<meta name="go-import" content="{{.ImportPrefix}} git {{.VcsURL}}">
<meta name="go-source" content="{{.ImportPrefix}} {{.VcsURL}} {{.VcsURL}}/tree/master{/dir} {{.VcsURL}}/blob/master{/dir}/{file}#L{line}">
<meta http-equiv="refresh" content="0; url={{.VcsURL}}">

<div style="width: 40em; margin: auto">
	<h1><code>{{.ImportPrefix}}</code></h1>

	<p>
	The source is at <a href="{{.VcsURL}}">{{.VcsURL}}</a>; you should be redirected shortly.

	<p style="font-size: 70%">
	Powered by <a href="https://github.com/docwhat/go-importd">docwhat.org/go-importd</a>.
	</br>
	See the <a href="https://golang.org/cmd/go/">go command</a> documentation for info on <code>go get</code> <code>&lt;meta&gt;</code> redirects.
</div>
`

	template := goTemplate.Must(goTemplate.New("main").Parse(redirectHTML))

	type templateData struct {
		ImportPrefix string
		VcsURL       string
	}

	return func(resp http.ResponseWriter, req *http.Request) {
		repoName := strings.SplitN(strings.Trim(req.URL.Path, "/"), "/", 2)[0]

		if repoName == "" {
			http.Error(resp, "Not Found", 404)
			return
		}

		data := &templateData{ImportPrefix: (config.importDomain + repoName), VcsURL: (config.githubUserURL + repoName)}

		if !urlExists(data.VcsURL) {
			http.Error(resp, "Not Found", 404)
			return
		}

		var buf bytes.Buffer
		err := template.Execute(&buf, data)
		if err != nil {
			http.Error(resp, err.Error(), 500)
			return
		}

		if _, err := resp.Write(buf.Bytes()); err != nil {
			http.Error(resp, err.Error(), 500)
		}
	}
}
