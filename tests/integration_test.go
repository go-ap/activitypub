package tests

import (
	"testing"

	a "github.com/mariusor/activitypub.go/activitypub"
	j "github.com/mariusor/activitypub.go/jsonld"

	"strings"
)

func TestAcceptSerialization(t *testing.T) {
	obj := a.AcceptNew("https://localhost/myactivity", nil)
	obj.Name = make(a.NaturalLanguageValue, 1)
	obj.Name["en"] = "test"
	obj.Name["fr"] = "teste"

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
	if !strings.Contains(string(data), string(obj.Name["en"])) {
		t.Errorf("Could not find name %#v in output %s", string(obj.Name["en"]), data)
	}
	if !strings.Contains(string(data), string(obj.Name["fr"])) {
		t.Errorf("Could not find name %#v in output %s", string(obj.Name["fr"]), data)
	}
	if !strings.Contains(string(data), string(obj.Type)) {
		t.Errorf("Could not find activity type %#v in output %s", obj.Type, data)
	}
}

func TestCreateActivityHTTPSerialization(t *testing.T) {
	id := a.ObjectID("test_object")
	obj := a.AcceptNew(id, nil)
	obj.Name["en"] = "Accept New"

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
	if !strings.Contains(string(data), string(obj.Name["en"])) {
		t.Errorf("Could not find name %#v in output %s", string(obj.Name["en"]), data)
	}
	if !strings.Contains(string(data), string(obj.Type)) {
		t.Errorf("Could not find activity type %#v in output %s", obj.Type, data)
	}
}
