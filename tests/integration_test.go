package tests

import (
	"testing"

	a "github.com/go-ap/activitystreams"
	j "github.com/go-ap/jsonld"

	"strings"
)

func TestAcceptSerialization(t *testing.T) {
	obj := a.AcceptNew("https://localhost/myactivity", nil)
	obj.Name = make(a.NaturalLanguageValues, 1)
	obj.Name.Set("en", "test")
	obj.Name.Set("fr", "teste")

	uri := "https://www.w3.org/ns/activitystreams"
	p := j.WithContext(j.IRI(uri))

	data, err := p.Marshal(obj)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if !strings.Contains(string(data), uri) {
		t.Errorf("Could not find context url %#v in output %s", p.Context, data)
	}
	if !strings.Contains(string(data), string(obj.ID)) {
		t.Errorf("Could not find id %#v in output %s", string(obj.ID), data)
	}
	if !strings.Contains(string(data), string(obj.Name.Get("en"))) {
		t.Errorf("Could not find name %#v in output %s", string(obj.Name.Get("en")), data)
	}
	if !strings.Contains(string(data), string(obj.Name.Get("fr"))) {
		t.Errorf("Could not find name %#v in output %s", string(obj.Name.Get("fr")), data)
	}
	if !strings.Contains(string(data), string(obj.Type)) {
		t.Errorf("Could not find activity type %#v in output %s", obj.Type, data)
	}
}

func TestCreateActivityHTTPSerialization(t *testing.T) {
	id := a.ObjectID("test_object")
	obj := a.AcceptNew(id, nil)
	obj.Name.Set("en", "Accept New")

	uri := string(a.ActivityBaseURI)

	data, err := j.WithContext(j.IRI(uri)).Marshal(obj)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if !strings.Contains(string(data), uri) {
		t.Errorf("Could not find context url %#v in output %s", j.GetContext(), data)
	}
	if !strings.Contains(string(data), string(obj.ID)) {
		t.Errorf("Could not find id %#v in output %s", string(obj.ID), data)
	}
	if !strings.Contains(string(data), obj.Name.Get("en")) {
		t.Errorf("Could not find name %s in output %s", obj.Name.Get("en"), data)
	}
	if !strings.Contains(string(data), string(obj.Type)) {
		t.Errorf("Could not find activity type %#v in output %s", obj.Type, data)
	}
}
