// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"appengine"
	"appengine/urlfetch"
)

var tmpl = template.Must(template.New("").Parse(`
<html>
 <head>
  <meta name="go-import" content="{{.PkgPath}} git {{.Repo}}">
 </head>
</html>
`))

const (
	githubOrg       = "azul3d"
	repoAliasHost   = "azul3d.org"
	repoAliasScheme = "http"
	fileHost        = "azul3d.github.io"
)

// isTip is short-hand for:
//  return version == "v0" || version == "dev"
func isTip(version string) bool {
	return version == "v0" || version == "dev"
}

// Pulls version tag from the URL. It would be at the last part of the URL
// like so:
//  foobar.org/something/something/maybe/here.v0
//  foobar.org/something/something/maybe/here.dev
//  foobar.org/something/something/maybe/here.v1.2
//
// NOT like:
//  foobar.org/something/something/maybe/here.v1.2/info/refs
//
// Always returns "dev" for any .dev or .v0 string.
func versionFromEnd(p string) string {
	if strings.HasSuffix(p, "dev") || strings.HasSuffix(p, "v0") {
		return "dev"
	}
	// he.re.v1.2
	split := strings.Split(path.Base(p), ".v")
	if len(split) > 1 {
		return "v" + split[len(split)-1]
	}
	return ""
}

// Takes a string like:
//  cmd/foo/bar.dev
//  cmd/foo/bar.v1
// and returns:
//  cmd-foo-bar
func gitRepoName(path string) string {
	// Change cmd/foo to cmd-foo
	path = strings.Replace(path, "/", "-", -1)
	// Strip version from end.
	return strings.Split(path, ".")[0]
}

func handleGoTool(ctx appengine.Context, w http.ResponseWriter, r *http.Request) bool {
	// Clean the URL.
	u := path.Clean(r.URL.Path)

	// Parse the query.
	query, _ := url.ParseQuery(r.URL.RawQuery)

	// If the client is the 'go get' tool, then we serve them a small page that
	// just contains the go-import meta tag -- that's all.
	if r.Method == "GET" && len(query.Get("go-get")) > 0 {
		repo := *r.URL
		repo.Host = repoAliasHost
		repo.Scheme = repoAliasScheme
		repo.Path = path.Join(u, "repo")
		repo.RawQuery = ""
		tmpl.Execute(w, map[string]interface{}{
			"Repo":    repo.String(),
			"PkgPath": path.Join(repoAliasHost, u),
		})
		return true
	}

	// If the client asks for /info/refs then we fetch them from the git repo
	// and serve them to the client.
	if r.Method == "GET" && strings.HasSuffix(u, "/info/refs") && query.Get("service") == "git-upload-pack" {
		// Strip /repo/info/refs from URL.
		fp := strings.Split(u, "/")
		fp = fp[1 : len(fp)-3]
		fps := path.Join(fp...)
		version := versionFromEnd(fps)
		repoName := gitRepoName(strings.TrimSuffix(fps, version))
		//ctx.Infof("u=%q version=%q repoName=%q\n", u, version, repoName)

		if len(version) == 0 {
			// The path doesn't have a version in it, we can't serve this
			// request.
			ctx.Infof("Request without version in URL.\n")
			w.WriteHeader(http.StatusNotFound)
			return true
		}

		// Create URL to target repo's /info/refs
		target := &url.URL{
			Scheme:   "http",
			Host:     "github.com",
			Path:     path.Join(githubOrg, repoName+".git", "/info/refs"),
			RawQuery: "service=git-upload-pack",
		}

		// Fetch info/refs from target repository.
		//ctx.Infof("fetchRefs from %s\n", target.String())
		refs, err := fetchRefs(urlfetch.Client(ctx), target.String())
		if err != nil {
			ctx.Infof("Failed to fetch remote refs: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return true
		}

		if !isTip(version) {
			//ctx.Infof("\n\nHack git refs:\n\n%s\n", string(refs.data))
			// Hack the git refs to the given version.
			err = refs.hack(version)
			if err != nil {
				ctx.Infof("%v\n", err)
				w.WriteHeader(http.StatusNotFound)
				return true
			}
			//ctx.Infof("\n\nAFTER HACK:\n\n%s\n", string(refs.data))
		}

		w.Header().Set("Content-Type", "application/x-git-upload-pack-advertisement")
		w.Write(refs.data)
		return true
	}

	// If the client wants to POST to /git-upload-pack we redirect their
	// request to the actual git repo.
	if r.Method == "POST" && strings.HasSuffix(u, "/git-upload-pack") {
		// Strip /repo/git-upload-pack from URL.
		fp := strings.Split(u, "/")
		fp = fp[1 : len(fp)-2]
		fps := path.Join(fp...)
		version := versionFromEnd(fps)
		repoName := gitRepoName(strings.TrimSuffix(fps, version))
		//ctx.Infof("u=%q version=%q repoName=%q\n", u, version, repoName)

		// Create URL to target repo's /git-upload-pack
		target := &url.URL{
			Scheme: "http",
			Host:   "github.com",
			Path:   path.Join(githubOrg, repoName+".git", "/git-upload-pack"),
		}

		w.Header().Set("Location", target.String())
		w.WriteHeader(http.StatusMovedPermanently)
		return true
	}

	// /info/refs for service=git-receive-pack is just forwarded to the repo
	// directly. This occurs when pushing changes via git.
	if r.Method == "GET" && strings.HasSuffix(u, "/info/refs") && query.Get("service") == "git-receive-pack" {
		// Strip /repo/info/refs from URL.
		fp := strings.Split(u, "/")
		fp = fp[1 : len(fp)-3]
		fps := path.Join(fp...)
		version := versionFromEnd(fps)
		repoName := gitRepoName(strings.TrimSuffix(fps, version))
		//ctx.Infof("u=%q version=%q repoName=%q\n", u, version, repoName)

		// Create URL to target repo's /info/refs
		target := &url.URL{
			Scheme:   "http",
			Host:     "github.com",
			Path:     path.Join(githubOrg, repoName+".git", "/info/refs"),
			RawQuery: "service=git-receive-pack",
		}

		http.Redirect(w, r, target.String(), http.StatusSeeOther)
		return true
	}

	return false
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	ctx.Infof("%v %v\n", r.Method, r.URL)

	// If it's the Go tool (or git HTTP, etc) then we let that function handle
	// it.
	if handleGoTool(ctx, w, r) {
		return
	}

	// Just proxy the request to the file host then (it's an actual user -- not
	// the Go tool).
	if r.URL.Scheme == "" {
		r.URL.Scheme = "http"
	}
	r.RequestURI = ""
	delete(r.Header, "Content-Length")
	r.URL.Host = fileHost

	resp, err := urlfetch.Client(ctx).Do(r)
	if err != nil {
		ctx.Infof("GET error: %v\n", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.Infof("GET read error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var modtime time.Time
	stringTime := resp.Header.Get("Last-Modified")
	if stringTime != "" {
		modtime, err = time.Parse(http.TimeFormat, stringTime)
		if err != nil {
			ctx.Infof("%v\n", err)
		}
	}
	http.ServeContent(w, r, path.Base(r.URL.Path), modtime, bytes.NewReader(respData))
	return
}

func init() {
	http.HandleFunc("/", handler)
}

type gitRefs struct {
	data []byte
}

var (
	ErrRepoNotFound    = errors.New("git repository not found")
	ErrVersionNotFound = errors.New("failed to find version in git refs")
)

func (r *gitRefs) hack(version string) error {
	var mrefi, mrefj, vrefi, vrefj int

	vhead := "refs/heads/" + version
	vtag := "refs/tags/" + version

	data := r.data
	sdata := string(r.data)
	for i, j := 0, 0; i < len(data); i = j {
		size, err := strconv.ParseInt(sdata[i:i+4], 16, 32)
		if err != nil {
			return fmt.Errorf("cannot parse refs line size: %s", string(data[i:i+4]))
		}
		if size == 0 {
			size = 4
		}
		j = i + int(size)
		if j > len(sdata) {
			return fmt.Errorf("incomplete refs data received from repo")
		}
		if sdata[0] == '#' {
			continue
		}

		hashi := i + 4
		hashj := strings.IndexByte(sdata[hashi:j], ' ')
		if hashj < 0 || hashj != 40 {
			continue
		}
		hashj += hashi

		namei := hashj + 1
		namej := strings.IndexAny(sdata[namei:j], "\n\x00")
		if namej < 0 {
			namej = j
		} else {
			namej += namei
		}

		name := sdata[namei:namej]

		if name == "refs/heads/master" {
			mrefi = hashi
			mrefj = hashj
		}

		if strings.HasPrefix(name, "refs/heads/v") || strings.HasPrefix(name, "refs/tags/v") {
			// Annotated tag is peeled off and overrides the same version just parsed.
			name = strings.TrimSuffix(name, "^{}")
			if name == vtag || name == vhead {
				vrefi = hashi
				vrefj = hashj
			}
		}

		//if mrefi > 0 && vrefi > 0 {
		//	break
		//}
	}

	if mrefi == 0 || vrefi == 0 {
		return ErrVersionNotFound
	}

	copy(data[mrefi:mrefj], data[vrefi:vrefj])
	return nil
}

func fetchRefs(client *http.Client, refsURL string) (*gitRefs, error) {
	resp, err := client.Get(refsURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusNotFound {
			return nil, ErrRepoNotFound
		} else {
			return nil, fmt.Errorf("error from repo: %v", resp.Status)
		}
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &gitRefs{
		data: data,
	}, nil
}
