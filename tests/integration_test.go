package tests

import (
	"activitypub"
	"jsonld"
	"testing"
)

func TestAcceptSerialization(t *testing.T) {
	o := activitypub.AcceptNew("https://localhost/myactivity")
	o.Name = make(activitypub.NaturalLanguageValue, 1)
	o.Name["en"] = "test"

	ctx := jsonld.Context{URL: "https://www.w3.org/ns/activitystreams"}

	bytes, err := jsonld.Marshal(o, &ctx)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	t.Logf("%s", bytes)
}

func TestCreateActivityHTTPSerialization(t *testing.T) {
	id := activitypub.ObjectId("test_object")
	o := activitypub.AcceptNew(id)
	o.Name["en"] = "Accept New"

	baseUri := string(activitypub.ActivityBaseURI)
	c := jsonld.Context{
		URL: baseUri + o.Type,
	}

	out, err := jsonld.Marshal(o, &c)
	//out, err := jsonld.Marshal(o)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", out)
}
