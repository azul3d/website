// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"azul3d.org/semver.v1"
)

const (
	githubOrg = "azul3d"
	host      = "azul3d.org"
	fileHost  = "azul3d.github.io"
	certFile  = "azul3d.org.pem"
	keyFile   = "azul3d.org.key"
)

var (
	lastIdlePurge = time.Now()
	pkgHandler    = &semver.Handler{
		Matcher:  semver.MatcherFunc(compatMatcher),
		Host:     host,
	}
	githubMatcher = semver.GitHub(githubOrg)
)

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

	// Now let the github matcher perform the matching.
	return githubMatcher.Match(u)
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

	// Default to HTTPS scheme.
	if r.URL.Scheme == "" {
		r.URL.Scheme = "https"
	}

	// Just proxy the request to the file host then (it's an actual user -- not
	// the Go tool).
	r.RequestURI = ""
	delete(r.Header, "Content-Length")

	// Change Host in URL so that the request goes to the file host.
	r.URL.Host = fileHost

	// Force the Host header to be the file host (github.io uses this header).
	r.Host = fileHost

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Printf("GET error: %v\n", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer resp.Body.Close()

	// Copy headers over.
	hdr := w.Header()
	for k, v := range resp.Header {
		hdr[k] = v
	}

	// Peek to detect the content type.
	br := bufio.NewReaderSize(resp.Body, 1024)
	if r.Method != "HEAD" && len(resp.Header.Get("If-Modified-Since")) != 0 {
		ident, err := br.Peek(512)
		if err != nil {
			log.Printf("Proxy peek error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		hdr["Content-Type"] = []string{http.DetectContentType(ident)}
	}

	// Write the header / status code.
	w.WriteHeader(resp.StatusCode)

	// Copy the response to the user.
	_, err = io.Copy(w, br)
	if err != nil {
		log.Printf("Proxy copy error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

var (
	addr    = flag.String("http", ":80", "HTTP address to serve on")
	tlsaddr = flag.String("https", ":443", "HTTPS address to serve on")
)

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)

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
