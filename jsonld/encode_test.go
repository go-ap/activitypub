package jsonld

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

type mockBase struct {
	Id   string
	Name string
	Type string
}

type mockTypeA struct {
	mockBase
	PropA string
	PropB float32
}

func TestMarshal(t *testing.T) {
	a := mockTypeA{mockBase{"base_id", "MockObjA", "mock_obj"}, "prop_a", 0.001}
	b := mockTypeA{}

	url := "http://www.habarnam.ro"
	p := WithContext(IRI(url))

	var err error
	var out []byte

	out, err = p.Marshal(a)
	if err != nil {
		t.Errorf("%s", err)
	}
	if !strings.Contains(string(out), string(ContextKw)) {
		t.Errorf("Context name not found %q in %s", ContextKw, out)
	}
	if !strings.Contains(string(out), url) {
		t.Errorf("Context url not found %q in %s", url, out)
	}
	err = Unmarshal(out, &b)
	if err != nil {
		t.Errorf("%s", err)
	}
	if a.Id != b.Id {
		t.Errorf("Id isn't equal %q expected %q in %s", a.Id, b.Id, out)
	}
	if a.Name != b.Name {
		t.Errorf("Name isn't equal %q expected %q", a.Name, b.Name)
	}
	if a.Type != b.Type {
		t.Errorf("Type isn't equal %q expected %q", a.Type, b.Type)
	}
	if a.PropA != b.PropA {
		t.Errorf("PropA isn't equal %q expected %q", a.PropA, b.PropA)
	}
	if a.PropB != b.PropB {
		t.Errorf("PropB isn't equal %f expected %f", a.PropB, b.PropB)
	}
}

func TestMarshalNullContext(t *testing.T) {
	var a = struct {
		PropA string
		PropB float64
	}{"test", 0.0004}

	outL, errL := Marshal(a)
	if errL != nil {
		t.Errorf("%s", errL)
	}
	outJ, errJ := Marshal(a)
	if errJ != nil {
		t.Errorf("%s", errJ)
	}
	if !bytes.Equal(outL, outJ) {
		t.Errorf("Json output should be euqlal %q, received %q", outL, outJ)
	}
}

func TestIsEmpty(t *testing.T) {
	var a int
	if !isEmptyValue(reflect.ValueOf(a)) {
		t.Errorf("Invalid empty value %v", a)
	}
	if !isEmptyValue(reflect.ValueOf(uint(a))) {
		t.Errorf("Invalid empty value %v", uint(a))
	}
	var b float64
	if !isEmptyValue(reflect.ValueOf(b)) {
		t.Errorf("Invalid empty value %v", b)
	}
	var c string
	if !isEmptyValue(reflect.ValueOf(c)) {
		t.Errorf("Invalid empty value %s", c)
	}
	var d []byte
	if !isEmptyValue(reflect.ValueOf(d)) {
		t.Errorf("Invalid empty value %v", d)
	}
	var e *interface{}
	if !isEmptyValue(reflect.ValueOf(e)) {
		t.Errorf("Invalid empty value %v", e)
	}
	g := false
	if !isEmptyValue(reflect.ValueOf(g)) {
		t.Errorf("Invalid empty value %v", g)
	}
	h := true
	if isEmptyValue(reflect.ValueOf(h)) {
		t.Errorf("Invalid empty value %v", h)
	}
}

func TestWithContext_MarshalJSON(t *testing.T) {
	tv := "value_test"
	v := struct{ Test string }{Test: tv}

	data, err := WithContext(IRI("http://example.com")).Marshal(v)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Contains(data, []byte(ContextKw)) {
		t.Errorf("%q not found in %s", ContextKw, data)
	}
	m := reflect.TypeOf(v)
	mv := reflect.ValueOf(v)
	for i := 0; i < m.NumField(); i++ {
		f := m.Field(i)
		v := mv.Field(i)
		if !bytes.Contains(data, []byte(f.Name)) {
			t.Errorf("%q not found in %s", f.Name, data)
		}
		if !bytes.Contains(data, []byte(v.String())) {
			t.Errorf("%q not found in %s", v.String(), data)
		}
	}

}
