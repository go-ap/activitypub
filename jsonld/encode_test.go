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
	Ctx = &Context{URL: Ref(url)}
	var err error
	var out []byte

	out, err = Marshal(a)
	if err != nil {
		t.Errorf("%s", err)
	}
	if !strings.Contains(string(out), url) {
		t.Errorf("Context url not found '%s' in %s", url, out)
	}
	err = Unmarshal(out, &b)
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
		t.Errorf("Json output should be euqlal '%s', received '%s'", outL, outJ)
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

func TestPayloadWithContext_MarshalJSON(t *testing.T) {
	empty := payloadWithContext{}
	eData, eErr := empty.MarshalJSON()

	if eErr != nil {
		t.Errorf("Error: %s", eErr)
	}
	Ctx = nil
	n, _ := Marshal(nil)
	if bytes.Compare(eData, n) != 0 {
		t.Errorf("Empty payload should resolve to null json value '%s', received '%s'", n, eData)
	}

	var a interface{}
	a = 1
	p := payloadWithContext{Obj: &a}
	pData, pErr := p.MarshalJSON()

	if pErr != nil {
		t.Errorf("Error: %s", pErr)
	}
	av, _ := Marshal(a)
	if bytes.Compare(pData, av) != 0 {
		t.Errorf("Empty payload should resolve to value '%#v', received '%s'", av, pData)
	}
}
