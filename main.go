// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
	"path/filepath"
	"go/build"
	"os"
	"html/template"

	"azul3d.org/semver.v1"
	"azul3d.org/website/mdattr"
)

const (
	githubOrg = "azul3d"
	fileHost  = "azul3d.github.io"
	certFile  = "azul3d.org.pem"
	keyFile   = "azul3d.org.key"
)

var (
	lastIdlePurge = time.Now()
	pkgHandler    = &semver.Handler{
		Matcher: semver.MatcherFunc(compatMatcher),
		Host:    "azul3d.org",
	}
	githubMatcher = semver.GitHub(githubOrg)
	pages         = http.Dir("pages")
	src = &GitUpdater{
		Dir: gpPath("azul3d.org/website"),
	}
	contentDir = gpPath("azul3d.org/website/content")
	tmpls = template.Must(template.ParseGlob(path.Join(gpPath("azul3d.org/website/templates"), "*")))
)

// gpPath finds and returns the absolute path to the first directory found in
// the $GOPATH list.
//
//  $GOPATH=/home/joe/godev;/home/k/godev
//  gpPath("foobar") -> "/home/joe/godev/src/foobar"
//  gpPath("a/b")    -> "/home/k/godev/src/a/b"
//  gpPath("oops")   -> ""
//
func gpPath(relPath string) string {
	for _, p := range filepath.SplitList(build.Default.GOPATH) {
		p = filepath.Join(p, "src", relPath)
		if _, err := os.Stat(p); err == nil {
			return p
		}
	}
	return ""
}

func compatMatcher(u *url.URL) (r *semver.Repo, err error) {
	// Special case just for glfw.v3.1 -- we made a bad mistake here and this
	// is purely for backwards compatability so you can still write the (bad)
	// import:
	//
	//  "azul3d.org/native/glfw.v3.1"
	//
	// All imports like this should be updated to just ".v3" instead.
	if u.Path == "/native/glfw.v3.1" {
		u.Path = "/native/glfw.v3"
	}

	// Special case just for .dev paths -- previously we suggested importing
	// in-development packages under the .dev extension (which violates semver
	// specification):
	//
	//  "azul3d.org/tmx.dev"
	//
	// Now we redirect those to .v0 -- which has the *same effect*, but only a
	// different meaning under semver specification. In-development packages
	// today are imported under either .v0 (no stable release yet) or .vN-dev:
	//
	//  "azul3d.org/tmx.v2-dev"
	//
	if strings.HasSuffix(u.Path, ".dev") {
		u.Path = strings.TrimSuffix(u.Path, ".dev") + ".v0"
	}

	// Special case for /website, we want:
	//
	//  go get azul3d.org/website
	//
	// to simply download the latest (v0).
	if strings.HasPrefix(u.Path, "/website") {
		u.Path = u.Path[:len("/website")]
		u.Path = "/website.v0" + u.Path
	}

	// Now let the github matcher perform the matching.
	return githubMatcher.Match(u)
}

func mdHandler(w http.ResponseWriter, p string) bool {
	// foo/bar -> foo/bar/index.html
	_, file := path.Split(p)
	if len(file) == 0 {
		p = path.Join(p, "index.html")
	}

	// foo.html -> foo.md
	if strings.HasSuffix(p, ".html") {
		p = strings.TrimSuffix(p, ".html")
		p = p + ".md"
	}

open:
	// Open Markdown file.
	f, err := pages.Open(p)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return true
		}
		return false
	}
	defer f.Close()

	fi, err := f.(*os.File).Stat()
	if fi.IsDir() {
		// Directories have their index.md served.
		p = path.Join(p, "index.md")
		goto open
	}

	// Split attributes from the markdown file.
	attr, mdData, err := mdattr.Parse(f)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return true
	}

	tmplData := map[string]interface{}{
		"HTML": template.HTML(mdRender(mdData)),
		"Attr": attr,
	}

	tmp := tmpls.Lookup("markdown.tmpl")
	err = tmp.Execute(w, tmplData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return true
	}
	return true
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v\n", r.Method, r.URL)

	// Purge idle connections.
	if time.Since(lastIdlePurge) > 2*time.Hour {
		lastIdlePurge = time.Now()
		http.DefaultTransport.(*http.Transport).CloseIdleConnections()
	}

	// Give our semver handler the ability to handle the request.
	status, err := pkgHandler.Handle(w, r)
	if err != nil {
		log.Println(err)
	}
	if status == semver.Handled {
		return
	}
	if status == semver.PkgPage {
		// Package page, redirect them to godoc.org documentation.
		tmp := *r.URL
		tmp.Scheme = "https"
		tmp.Host = "godoc.org"
		tmp.Path = path.Join(pkgHandler.Host, tmp.Path)
		http.Redirect(w, r, tmp.String(), http.StatusSeeOther)
		return
	}

	// Let the markdown handler serve the request if it can.
	if mdHandler(w, r.URL.Path) {
		return
	}

	// Don't know what they want, 404.
	w.WriteHeader(http.StatusNotFound)
	if mdHandler(w, "404.html") {
		return
	}
}

var (
	addr    = flag.String("http", ":80", "HTTP address to serve on")
	tlsaddr = flag.String("https", ":443", "HTTPS address to serve on")
	update  = flag.Bool("update", true, "update via Git and shutdown server after pull")
)

func main() {
	flag.Parse()
	http.Handle("/content/", http.StripPrefix("/content/", http.FileServer(http.Dir(contentDir))))
	http.HandleFunc("/", handler)

	// Source code updater.
	if *update {
		go func() {
			for{
				time.Sleep(5 * time.Minute)
				updated, err := src.Update()
				if err != nil {
					log.Println("Update error:", err)
					continue
				}
				if updated {
					log.Println("Updated source code. Exiting server..")
					os.Exit(0)
				}
				log.Println("No updates.")
			}
		}()
	}

	// Start HTTPS server:
	go func() {
		if len(*tlsaddr) > 0 {
			log.Println("Serving on", *tlsaddr)
			err := http.ListenAndServeTLS(*tlsaddr, certFile, keyFile, nil)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	// Start HTTP server:
	log.Println("Serving on", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
