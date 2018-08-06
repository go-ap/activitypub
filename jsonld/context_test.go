package jsonld

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestRef_MarshalText(t *testing.T) {
	test := "test"
	a := IRI(test)

	out, err := a.MarshalText()
	if err != nil {
		t.Errorf("Error %s", err)
	}
	if bytes.Compare(out, []byte(test)) != 0 {
		t.Errorf("Invalid result '%s', expected '%s'", out, test)
	}
}

func TestContext_MarshalJSON(t *testing.T) {
	url := "test"
	c := Context{NilTerm: IRI(url)}

	out, err := c.MarshalJSON()
	if err != nil {
		t.Errorf("%s", err)
	}
	if !strings.Contains(string(out), url) {
		t.Errorf("Json doesn't contain %s, %s", url, string(out))
	}
	jUrl, _ := json.Marshal(url)
	if !bytes.Equal(jUrl, out) {
		t.Errorf("Strings should be equal %s, %s", jUrl, out)
	}

	asTerm := "testingTerm##"
	asUrl := "https://activitipubrocks.com"
	c2 := Context{NilTerm: IRI(url), Term(asTerm): IRI(asUrl)}
	out, err = c2.MarshalJSON()
	if err != nil {
		t.Errorf("%s", err)
	}
	if !strings.Contains(string(out), url) {
		t.Errorf("Json doesn't contain URL %s, %s", url, string(out))
	}
	if !strings.Contains(string(out), asUrl) {
		t.Errorf("Json doesn't contain URL %s, %s", asUrl, string(out))
	}
	if !strings.Contains(string(out), asTerm) {
		t.Errorf("Json doesn't contain Term %s, %s", asTerm, string(out))
	}

	testTerm := "test_term"
	c3 := Context{Term(testTerm): IRI(url), Term(asTerm): IRI(asUrl)}
	out, err = c3.MarshalJSON()
	if err != nil {
		t.Errorf("%s", err)
	}
	if !strings.Contains(string(out), url) {
		t.Errorf("Json doesn't contain URL %s, %s", url, string(out))
	}
	if !strings.Contains(string(out), asUrl) {
		t.Errorf("Json doesn't contain URL %s, %s", asUrl, string(out))
	}
	if !strings.Contains(string(out), asTerm) {
		t.Errorf("Json doesn't contain Term %s, %s", asTerm, string(out))
	}
	if !strings.Contains(string(out), testTerm) {
		t.Errorf("Json doesn't contain Term %s, %s", testTerm, string(out))
	}
}
