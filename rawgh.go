// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"sync"
	"time"
	"net/http"
	"net/url"
	"io/ioutil"
	"path"
)

// ghFile represents a single github file cache entry.
type ghFile struct {
	t time.Time // The time at which the file was last requested.
	data []byte // Literal file data.
}

// RawGH can fetch and cache requests for small files from:
//
//  https://raw.githubusercontent.com
//
// Small files are a constraint, because the file is read into memory
// completely.
type RawGH struct {
	sync.Mutex

	// GitHub user, repository, and branch. Can modify these in-between each
	// Fetch.
	//
	// If branch is an empty string: "master" is used.
	User, Repo, Branch string

	// Duration after which files are considered stale and requested again upon
	// the next Fetch.
	//
	// If zero, 5 * time.Minute is used.
	PurgeTime time.Duration

	// HTTP client used for fetching files, http.DefaultClient is used if this
	// is nil.
	Client *http.Client

	// Cached files by URL.
	cache map[string]ghFile
}

// Fetch returns the file at:
//
//  https://raw.githubusercontent.com/[User]/[Repo]/[Branch]/[fp]
//
// The same data slice is returned to multiple callers, so it should be
// considered read-only.
func (r *RawGH) Fetch(fp string) ([]byte, error) {
	r.Lock()
	defer r.Unlock()

	u := &url.URL{
		Scheme: "https",
		Host: "raw.githubusercontent.com",
		Path: path.Join(r.User, r.Repo, r.Branch, fp),
	}
	us := u.String()

	// Assume 5 minute purge time by default.
	purgeTime := r.PurgeTime
	if purgeTime == 0 {
		purgeTime = 5 * time.Minute
	}

	// If the file is cached and not expired, use that version.
	cached, ok := r.cache[us]
	if ok && time.Since(cached.t) < purgeTime {
		return cached.data, nil
	}

	// Fetch the file.
	client := r.Client
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Get(us)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Cache the file for later.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r.cache[us] = ghFile{
		t: time.Now(),
		data: body,
	}
	return body, nil
}
