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
	obj.Name.Set("en", pub.Content("test"))
	obj.Name.Set("fr", pub.Content("teste"))

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
	id := pub.ID("test_object")
	obj := pub.AcceptNew(id, nil)
	obj.Name.Set("en", pub.Content("Accept New"))

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
	if !strings.Contains(string(data), obj.Name.Get("en").String()) {
		t.Errorf("Could not find name %s in output %s", obj.Name.Get("en"), data)
	}
	if !strings.Contains(string(data), string(obj.Type)) {
		t.Errorf("Could not find activity type %#v in output %s", obj.Type, data)
	}
}
