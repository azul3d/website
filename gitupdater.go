// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"os/exec"
)

// GitUpdater can update a git repository and tell you if anything was actually
// changed.
//
// Pulling changes into a git repository is not enough without knowing if
// something was pulled, because we will restart the process if changes have
// occured.
type GitUpdater struct {
	Dir string
}

// localSHA returns the SHA of the local repository:
//
//  git rev-parse HEAD
//
func (g *GitUpdater) localSHA() (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = g.Dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

// remoteSHA queries the origin of the repository for the HEAD SHA:
//
//  git ls-remote origin HEAD
//
// Only the first 40 bytes (SHA length) are returned.
func (g *GitUpdater) remoteSHA() (string, error) {
	cmd := exec.Command("git", "ls-remote", "origin", "HEAD")
	cmd.Dir = g.Dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return string(out.Bytes()[:40]), nil
}

// Update checks if the local and remote SHA's of HEAD match, if they don't
// then it fast-forwards the repository and returns true. False is always
// returned otherwise.
func (g *GitUpdater) Update() (bool, error) {
	// If the local and remote are identical, then no changes have occured in
	// HEAD.
	local, err := g.localSHA()
	if err != nil {
		return false, err
	}
	remote, err := g.remoteSHA()
	if err != nil {
		return false, err
	}
	if local == remote {
		// No changes in HEAD.
		return false, nil
	}

	// Pull and fast-forward now.
	cmd := exec.Command("git", "pull", "--ff-only")
	cmd.Dir = g.Dir
	err = cmd.Run()
	if err != nil {
		return false, err
	}
	return true, nil
}
