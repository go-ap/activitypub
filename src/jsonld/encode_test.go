package jsonld

import (
	"encoding/json"
	"testing"
)

type mockBase struct {
	Id   string
	Name string
	Type string
}

type mockTypeA struct {
	*mockBase
	PropA string
	PropB float32
}

func TestMarshal(t *testing.T) {
	a := mockTypeA{&mockBase{"base_id", "MockObjA", "mock_obj"}, "prop_a", 0.001}
	b := mockTypeA{}

	c := Context{URL: "http://www.habarnam.ro"}
	var err error
	var out []byte

	out, err = Marshal(a, &c)
	if err != nil {
		t.Errorf("%s", err)
	}
	err = json.Unmarshal(out, &b)
	if err != nil {
		t.Errorf("%s", err)
	}
	if a.Id != b.Id {
		t.Errorf("Id isn't equal '%s' expected '%s'", a.Id, b.Id)
	}
	if a.Name != b.Name {
		t.Errorf("Name isn't equal '%s' expected '%s'", a.Name, b.Name)
	}
	if a.Type != b.Type {
		t.Errorf("Type isn't equal '%s' expected '%s'", a.Type, b.Type)
	}
	if a.PropA != b.PropA {
		t.Errorf("PropA isn't equal '%s' expected '%s'", a.PropA, b.PropA)
	}
	if a.PropB != b.PropB {
		t.Errorf("PropB isn't equal %f expected %f", a.PropB, b.PropB)
	}
}
