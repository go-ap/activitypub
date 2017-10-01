package tests

import (
	"testing"
	"activitypub"
	"jsonld"
)

func TestAcceptSerialization(t *testing.T) {
	o := activitypub.AcceptNew("https://localhost/myactivity")
	o.Name = make(activitypub.NaturalLanguageValue, 1)
	o.Name["en"] = "test"

	ctx := jsonld.Context{URL:"https://www.w3.org/ns/activitystreams"}

	bytes, err := jsonld.Marshal(o, &ctx)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	t.Logf("%s", bytes)
}
