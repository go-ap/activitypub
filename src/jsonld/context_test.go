package jsonld

import (
	"bytes"
	"strings"
	"testing"
)

func TestRef_MarshalText(t *testing.T) {
	test := "test"
	a := Ref(test)

	out, err := a.MarshalText()
	if err != nil {
		t.Errorf("Error %s", err)
	}
	if bytes.Compare(out, []byte(test)) != 0 {
		t.Errorf("Invalid result '%s', expected '%s'", out, test)
	}
}

func TestContext_Ref(t *testing.T) {
	url := "test"
	c := Context{url}

	if c.Ref() != Ref(url) {
		t.Errorf("Invalid result %#v, expected %#v", c.Ref(), Ref(url))
	}
}

func TestContext_MarshalJSON(t *testing.T) {
	url := "test"
	c := Context{url}

	out, err := c.MarshalJSON()
	if err != nil {
		t.Errorf("%s", err)
	}
	if !strings.Contains(string(out), url) {
		t.Errorf("Json doesn't contain %#v, %#v", url, string(out))
	}
}
