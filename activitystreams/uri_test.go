package activitystreams

import "testing"

func TestIRI_GetLink(t *testing.T) {
	val := "http://example.com"
	u := IRI(val)
	if u.GetLink() != IRI(val) {
		t.Errorf("IRI %q should equal %q", u, val)
	}
}

func TestIRI_String(t *testing.T) {
	val := "http://example.com"
	u := IRI(val)
	if u.String() != val {
		t.Errorf("IRI %q should equal %q", u, val)
	}
}

func TestIRI_GetID(t *testing.T) {

}

func TestIRI_GetType(t *testing.T) {

}

func TestIRI_IsLink(t *testing.T) {

}

func TestIRI_IsObject(t *testing.T) {

}

func TestIRI_UnmarshalJSON(t *testing.T) {

}
