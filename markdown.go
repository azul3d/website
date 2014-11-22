// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	bf "github.com/russross/blackfriday"
)

// mdRender renders the given markdown file as unsanitized HTML.
func mdRender(input []byte) []byte {
	// set up the HTML renderer
	htmlFlags := 0
	htmlFlags |= bf.HTML_USE_XHTML
	htmlFlags |= bf.HTML_USE_SMARTYPANTS
	htmlFlags |= bf.HTML_SMARTYPANTS_FRACTIONS
	htmlFlags |= bf.HTML_SMARTYPANTS_LATEX_DASHES
	renderer := bf.HtmlRenderer(htmlFlags, "", "")

	// set up the parser
	extensions := 0
	extensions |= bf.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= bf.EXTENSION_TABLES
	extensions |= bf.EXTENSION_FENCED_CODE
	extensions |= bf.EXTENSION_AUTOLINK
	extensions |= bf.EXTENSION_STRIKETHROUGH
	extensions |= bf.EXTENSION_SPACE_HEADERS
	extensions |= bf.EXTENSION_HEADER_IDS

	return bf.Markdown(input, renderer, extensions)
}
