package tests

import (
	"strings"
	"testing"

	j "github.com/go-ap/jsonld"

	pub "github.com/go-ap/activitypub"
)

func TestAcceptSerialization(t *testing.T) {
	obj := pub.AcceptNew("https://localhost/myactivity", nil)
	obj.Name = make(pub.NaturalLanguageValues, 1)
	obj.Name[pub.English] = pub.Content("test")
	obj.Name[pub.French] = pub.Content("test")

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
	if !strings.Contains(string(data), string(obj.Name.Get(pub.English))) {
		t.Errorf("Could not find name %#v in output %s", string(obj.Name.Get(pub.English)), data)
	}
	if !strings.Contains(string(data), string(obj.Name.Get(pub.French))) {
		t.Errorf("Could not find name %#v in output %s", string(obj.Name.Get(pub.French)), data)
	}
	if !strings.Contains(string(data), string(obj.Type)) {
		t.Errorf("Could not find activity type %#v in output %s", obj.Type, data)
	}
}

func TestCreateActivityHTTPSerialization(t *testing.T) {
	id := pub.ID("test_object")
	obj := pub.AcceptNew(id, nil)
	obj.Name[pub.English] = pub.Content("Accept New")

	uri := string(pub.ActivityBaseURI)

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
	if !strings.Contains(string(data), obj.Name.Get(pub.English).String()) {
		t.Errorf("Could not find name %s in output %s", obj.Name.Get(pub.English), data)
	}
	if !strings.Contains(string(data), string(obj.Type)) {
		t.Errorf("Could not find activity type %#v in output %s", obj.Type, data)
	}
}
