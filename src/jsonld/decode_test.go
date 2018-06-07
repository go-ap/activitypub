package jsonld

import (
	"strconv"
	"testing"

	ap "activitypub"
)

func TestUnmarshalWithEmptyJsonObject(t *testing.T) {
	obj := mockTypeA{}
	err := Unmarshal([]byte("{}"), &obj)
	if err != nil {
		t.Error(err)
	}
	if obj.Id != "" {
		t.Errorf("Id should have been an empty string, found %s", obj.Id)
	}
	if obj.Name != "" {
		t.Errorf("Name should have been an empty string, found %s", obj.Name)
	}
	if obj.Type != "" {
		t.Errorf("Type should have been an empty string, found %s", obj.Type)
	}
	if obj.PropA != "" {
		t.Errorf("PropA should have been an empty string, found %s", obj.PropA)
	}
	if obj.PropB != 0 {
		t.Errorf("PropB should have been 0.0, found %f", obj.PropB)
	}
}

type mockWithContext struct {
	mockTypeA
	Context Context `jsonld:"@context"`
}

func TestUnmarshalWithEmptyJsonObjectWithStringContext(t *testing.T) {
	obj := mockWithContext{}
	url := "http://www.habarnam.ro"
	data := []byte(`{"@context": "` + url + `" }`)
	err := Unmarshal(data, &obj)
	if err != nil {
		t.Error(err)
	}
	if obj.Context.Ref() != Ref(url) {
		t.Errorf("@context should have been %q, found %q", url, obj.Context.Collapse())
	}
	if obj.Id != "" {
		t.Errorf("Id should have been an empty string, found %s", obj.Id)
	}
	if obj.Name != "" {
		t.Errorf("Name should have been an empty string, found %s", obj.Name)
	}
	if obj.Type != "" {
		t.Errorf("Type should have been an empty string, found %s", obj.Type)
	}
	if obj.PropA != "" {
		t.Errorf("PropA should have been an empty string, found %s", obj.PropA)
	}
	if obj.PropB != 0 {
		t.Errorf("PropB should have been 0.0, found %f", obj.PropB)
	}
}

func TestUnmarshalWithEmptyJsonObjectWithObjectContext(t *testing.T) {
	obj := mockWithContext{}
	url := "http://www.habarnam.ro"
	data := []byte(`{"@context": { "@url": "` + url + `"} }`)
	err := Unmarshal(data, &obj)
	if err != nil {
		t.Error(err)
	}
	if obj.Context.Ref() != Ref(url) {
		t.Errorf("@context should have been %q, found %q", url, obj.Context.Collapse())
	}
	if obj.Id != "" {
		t.Errorf("Id should have been an empty string, found %s", obj.Id)
	}
	if obj.Name != "" {
		t.Errorf("Name should have been an empty string, found %s", obj.Name)
	}
	if obj.Type != "" {
		t.Errorf("Type should have been an empty string, found %s", obj.Type)
	}
	if obj.PropA != "" {
		t.Errorf("PropA should have been an empty string, found %s", obj.PropA)
	}
	if obj.PropB != 0 {
		t.Errorf("PropB should have been 0.0, found %f", obj.PropB)
	}
}

func TestUnmarshalWithEmptyJsonObjectWithOneLanguageContext(t *testing.T) {
	obj := mockWithContext{}
	url := "http://www.habarnam.ro"
	langEn := "en-US"

	data := []byte(`{"@context": { "@url": "` + url + `", "@language": "` + langEn + `"} }`)
	err := Unmarshal(data, &obj)
	if err != nil {
		t.Error(err)
	}
	if obj.Context.Ref() != Ref(url) {
		t.Errorf("@context should have been %q, found %q", url, obj.Context.Collapse())
	}
	if obj.Context.Language != ap.LangRef(langEn) {
		t.Errorf("@context should have been %q, found %q", url, obj.Context.Collapse())
	}
	if obj.Id != "" {
		t.Errorf("Id should have been an empty string, found %s", obj.Id)
	}
	if obj.Name != "" {
		t.Errorf("Name should have been an empty string, found %s", obj.Name)
	}
	if obj.Type != "" {
		t.Errorf("Type should have been an empty string, found %s", obj.Type)
	}
	if obj.PropA != "" {
		t.Errorf("PropA should have been an empty string, found %s", obj.PropA)
	}
	if obj.PropB != 0 {
		t.Errorf("PropB should have been 0.0, found %f", obj.PropB)
	}
}
func TestUnmarshalWithEmptyJsonObjectWithFullObject(t *testing.T) {
	obj := mockWithContext{}
	url := "http://www.habarnam.ro"
	langEn := "en-US"
	propA := "ana"
	var propB float32 = 6.66
	typ := "test"
	name := "test object #1"
	id := "777sdad"

	data := []byte(`{
		"@context": { "@url": "` + url + `", "@language": "` + langEn + `"}, 
		"PropA": "` + propA + `", 
		"PropB": ` + strconv.FormatFloat(float64(propB), 'f', 2, 32) + `, 
		"Id" : "` + id + `",
		"Name" : "` + name + `",
		"Type" : "` + typ + `"
	}`)
	err := Unmarshal(data, &obj)
	if err != nil {
		t.Error(err)
	}
	if obj.Context.Ref() != Ref(url) {
		t.Errorf("@context should have been %q, found %q", url, obj.Context.Collapse())
	}
	if obj.Context.Language != ap.LangRef(langEn) {
		t.Errorf("@context should have been %q, found %q", url, obj.Context.Collapse())
	}
	if obj.Id != id {
		t.Errorf("Id should have been %q, found %q", id, obj.Id)
	}
	if obj.Name != name {
		t.Errorf("Name should have been %q, found %q", name, obj.Name)
	}
	if obj.Type != typ {
		t.Errorf("Type should have been %q, found %q", typ, obj.Type)
	}
	if obj.PropA != propA {
		t.Errorf("PropA should have been %q, found %q", propA, obj.PropA)
	}
	if obj.PropB != propB {
		t.Errorf("PropB should have been %f, found %f", propB, obj.PropB)
	}
}
