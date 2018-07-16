package activitypub

import "testing"

func TestURI_GetLink(t *testing.T) {
	val := "http://example.com"
	u := URI(val)
	if u.GetLink() != URI(val) {
		t.Errorf("URI %q should equal %q", u, val)
	}
}

func TestURI_String(t *testing.T) {
	val := "http://example.com"
	u := URI(val)
	if u.String() != val {
		t.Errorf("URI %q should equal %q", u, val)
	}
}

func TestIRI_GetLink(t *testing.T) {
	val := "http://example.com"
	u := IRI(val)
	if u.GetLink() != URI(val) {
		t.Errorf("URI %q should equal %q", u, val)
	}
}

func TestIRI_String(t *testing.T) {
	val := "http://example.com"
	u := IRI(val)
	if u.String() != val {
		t.Errorf("URI %q should equal %q", u, val)
	}
}
