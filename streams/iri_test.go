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
	if id := i.GetID(); !id.IsValid() || id != ObjectID(i) {
		t.Errorf("ObjectID %q (%T) should equal %q (%T)", id, id, i, ObjectID(i))
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

func TestFlattenToIRI(t *testing.T) {
	t.Skipf("TODO")
}

func TestIRI_URL(t *testing.T) {
	t.Skipf("TODO")
}

func TestIRIs_Contains(t *testing.T) {
	t.Skipf("TODO")
}

func TestIRI_Equals(t *testing.T) {
	{
		i1 := IRI("http://example.com")
		i2 := IRI("http://example.com")
		// same host same scheme
		if !i1.Equals(i2, true) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example.com/ana/are/mere")
		i2 := IRI("http://example.com/ana/are/mere")
		// same host, same scheme and same path
		if !i1.Equals(i2, true) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("https://example.com")
		i2 := IRI("http://example.com")
		// same host different scheme
		if !i1.Equals(i2, false) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example.com/ana/are/mere")
		i2 := IRI("https://example.com/ana/are/mere")
		// same host, different scheme and same path
		if !i1.Equals(i2, false) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("https://example.com?ana=mere")
		i2 := IRI("http://example.com?ana=mere")
		// same host different scheme, same query
		if !i1.Equals(i2, false) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("https://example.com?ana=mere&foo=bar")
		i2 := IRI("http://example.com?foo=bar&ana=mere")
		// same host different scheme, same query - different order
		if !i1.Equals(i2, false) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example.com/ana/are/mere?foo=bar&ana=mere")
		i2 := IRI("https://example.com/ana/are/mere?ana=mere&foo=bar")
		// same host, different scheme and same path, same query different order
		if !i1.Equals(i2, false) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("https://example.com?ana=mere")
		i2 := IRI("http://example.com?ana=mere")
		// same host different scheme, same query
		if !i1.Equals(i2, false) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("https://example.com?ana=mere&foo=bar")
		i2 := IRI("http://example.com?foo=bar&ana=mere")
		// same host different scheme, same query - different order
		if !i1.Equals(i2, false) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example.com/ana/are/mere?foo=bar&ana=mere")
		i2 := IRI("https://example.com/ana/are/mere?ana=mere&foo=bar")
		// same host, different scheme and same path, same query different order
		if !i1.Equals(i2, false) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	///
	{
		i1 := IRI("http://example.com")
		i2 := IRI("https://example.com")
		// same host different scheme
		if i1.Equals(i2, true) {
			t.Errorf("%s should not equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example1.com")
		i2 := IRI("http://example.com")
		// different host same scheme
		if i1.Equals(i2, true) {
			t.Errorf("%s should not equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example.com/ana/1are/mere")
		i2 := IRI("http://example.com/ana/are/mere")
		// same host, same scheme and different path
		if i1.Equals(i2, true) {
			t.Errorf("%s should not equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example.com?ana1=mere")
		i2 := IRI("http://example.com?ana=mere")
		// same host same scheme, different query key
		if i1.Equals(i2, false) {
			t.Errorf("%s should not equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example.com?ana=mere")
		i2 := IRI("http://example.com?ana=mere1")
		// same host same scheme, different query value
		if i1.Equals(i2, false) {
			t.Errorf("%s should not equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("https://example.com?ana=mere&foo=bar")
		i2 := IRI("http://example.com?foo=bar1&ana=mere")
		// same host different scheme, different query value - different order
		if i1.Equals(i2, false) {
			t.Errorf("%s should not equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example.com/ana/are/mere?foo=bar&ana=mere")
		i2 := IRI("https://example.com/ana/are/mere?ana=mere&foo1=bar")
		// same host, different scheme and same path, differnt query key different order
		if i1.Equals(i2, false) {
			t.Errorf("%s should not equal %s", i1, i2)
		}
	}
}
