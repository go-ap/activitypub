package jsonld

import (
	"encoding/json"
	"testing"
)

type mock_base struct {
	Id   string
	Name string
	Type string
}
type mock_type_a struct {
	*mock_base
	PropA string
	PropB float32
}

func TestMarshal(t *testing.T) {
	a := mock_type_a{&mock_base{"base_id", "MockObjA", "mock_obj"}, "prop_a", 0.001}

	c := Context{URL: "http://www.habarnam.ro"}
	out, err := Marshal(a, &c)
	if err != nil {
		t.Errorf("%s", err)
	}
	outj, errj := json.Marshal(a)
	if errj != nil {
		t.Errorf("%s", errj)
	}
	if string(outj) != string(out) {
		t.Errorf("Invalid json serialization %s != %s", outj, out)
	}
	t.Logf("%s", out)
}
