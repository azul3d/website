// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mdattr parses Markdown file attributes.
//
// In an effort to provide attributes associated with Markdown files this
// package parses them at the start of the file.
//
// The overal format is ([s] is whitespace):
//
//  +[s]name[s]=[s]value[s]
//
// That is, prefixed and suffixed whitespace is stripped. For example, given a
// Markdown document in the form of:
//
//  +Name   = Graham Joira
//  +  VarTwo = 2
//  +Var3   = The quick brown fox jumped over the lazy dog..
//
//
//  # Header 1
//  Paragraph 1!
//
//  ## Header 2
//  Paragraph 2!
//
// Parse would return the attribute map (note lack of whitespace on end of "Joira   "):
//
//  map[string]string{
//      "Name": "Graham Joira",
//      "VarTwo": "2",
//      "Var3": "The quick brown fox jumped over the lazy dog..",
//  }
//
// And Parse would return the data (note preceding newline removal):
//
//  # Header 1
//  Paragraph 1!
//
//  ## Header 2
//  Paragraph 2!
//
package mdattr // import "azul3d.org/website/mdattr"
