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
	i := IRI("http://example.com")
	if id := i.GetID(); id == nil || *id != ObjectID(i) {
		t.Errorf("ObjectID %q (%T) should equal %q (%T)", *id, id, i, ObjectID(i))
	}
}

func TestIRI_GetType(t *testing.T) {
	i := IRI("http://example.com")
	if i.GetType() != LinkType {
		t.Errorf("Invalid type for %T object %s, expected %s", i, i.GetType(), LinkType)
	}
}

func TestIRI_IsLink(t *testing.T) {
	i := IRI("http://example.com")
	if i.IsLink() != true {
		t.Errorf("%T.IsLink() returned %t, expected %t", i, i.IsLink(), true)
	}
}

func TestIRI_IsObject(t *testing.T) {
	i := IRI("http://example.com")
	if i.IsObject() != false {
		t.Errorf("%T.IsObject() returned %t, expected %t", i, i.IsObject(), false)
	}
}

func TestIRI_UnmarshalJSON(t *testing.T) {
	val := "http://example.com"
	i := IRI("")

	err := i.UnmarshalJSON([]byte(val))
	if err != nil {
		t.Error(err)
	}
	if val != i.String() {
		t.Errorf("%T invalid value after Unmarshal %q, expected %q", i, i, val)
	}
}
