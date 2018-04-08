// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonld

import (
	"testing"
)

var validTests = []struct {
	data string
	ok   bool
}{
	{`foo`, false},
	{`}{`, false},
	{`{]`, false},
	{`{}`, true},
	{`{"foo":"bar"}`, true},
	{`{"foo":"bar","bar":{"baz":["qux"]}}`, true},
}

func TestValid(t *testing.T) {
	for _, tt := range validTests {
		if ok := Valid([]byte(tt.data)); ok != tt.ok {
			t.Errorf("Valid(%#q) = %v, want %v", tt.data, ok, tt.ok)
		}
	}
}
