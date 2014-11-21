// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mdattr

import(
	"bytes"
	"bufio"
	"io"
)

// Parse parses the reader for attributes and returns a map of attributes and a
// data slice which simply strips the attributes and preceding whitespace of
// the given reader.
//
// Any IO errors encountered during scanning are returned, along with a nil map
// and data slice.
func Parse(r io.Reader) (attrs map[string]string, data []byte, err error) {
	attrs = make(map[string]string)
	scanner := bufio.NewScanner(r)

	// Scan all attributes into the map.
	for scanner.Scan() {
		b := scanner.Bytes()
		ns := bytes.TrimSpace(b)
		if len(ns) == 0 {
			continue
		}
		if ns[0] != '+' {
			// With whitespace removed, the line has another non-plus character
			// so stop scanning for attributes.
			data = append(data, b...)
			data = append(data, '\n')
			break
		}

		ns = ns[1:] // Trim prefixed +
		split := bytes.SplitN(ns, []byte{'='}, 2)
		if len(split) != 2 {
			continue // Ignore invalid lines.
		}
		varName := string(bytes.TrimSpace(split[0]))
		varValue := string(bytes.TrimSpace(split[1]))
		attrs[varName] = varValue
	}

	// Attribute scanning stops at the first blank line, scan the rest.
	for scanner.Scan() {
		b := scanner.Bytes()
		data = append(data, b...)
		data = append(data, '\n')
	}

	// Pass any scanner errors to the caller.
	err = scanner.Err()
	if err != nil {
		return nil, nil, err
	}
	return
}
