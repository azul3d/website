// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mdattr

import(
	"testing"
	"bytes"
)

func testParse(file, want string, wantAttr map[string]string, t *testing.T) {
	// Parse file.
	attr, data, err := Parse(bytes.NewBufferString(file))
	if err != nil {
		t.Fatal(err)
	}

	// Validate data.
	if !bytes.Equal(data, []byte(want)) {
		t.Logf("Got %q\n", string(data))
		t.Fatalf("Want %q\n", want)
	}

	// Check for nil/non-nil status.
	if len(attr) != len(wantAttr) {
		t.Log("Got", len(attr), "attributes")
		t.Fatal("Want", len(wantAttr), "attributes")
	}
	for k, v := range attr {
		if v != wantAttr[k] {
			t.Logf("Got  attr %q=%q", k, v)
			t.Fatalf("Want attr %q=%q", k, wantAttr[k])
		}
	}
}

func TestParseAttr(t *testing.T) {
	md := "# Header1\nWe live forever!\n"
	file := "+  Hello  =  the fox is quick\n+Foo   = bar  bar \n\n\n" + md
	wantAttr := map[string]string{
		"Hello": "the fox is quick",
		"Foo": "bar  bar",
	}
	testParse(file, md, wantAttr, t)
}
func TestParseNoAttr(t *testing.T) {
	s := "# Header1\nWe live forever!\n"
	testParse(s, s, nil, t)
}
