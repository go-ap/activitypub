package jsonld

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"bytes"
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

	url := "http://www.habarnam.ro"
	c := Context{URL: url}
	var err error
	var out []byte

	out, err = Marshal(a, &c)
	if err != nil {
		t.Errorf("%s", err)
	}
	if !strings.Contains(string(out), url) {
		t.Errorf("Context url not found '%s' in %s", url, out)
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

func TestMarshalNullContext(t *testing.T) {
	var a = struct {
		PropA string
		PropB float64
	} {"test", 0.0004}

	outL, errL := Marshal(a, nil )
	if errL != nil {
		t.Errorf("%s", errL)
	}
	outJ, errJ := Marshal(a, nil )
	if errJ != nil {
		t.Errorf("%s", errJ)
	}
	if !bytes.Equal(outL, outJ) {
		t.Errorf("Json output should be euqlal '%s', received '%s'", outL, outJ)
	}
}

func TestIsEmpty(t *testing.T) {
	var a int = 0
	if !IsEmpty(reflect.ValueOf(a)) {
		t.Errorf("Invalid empty valid %s", a)
	}
	if !IsEmpty(reflect.ValueOf(uint(a))) {
		t.Errorf("Invalid empty valid %s", uint(a))
	}
	var b float64 = 0
	if !IsEmpty(reflect.ValueOf(b)) {
		t.Errorf("Invalid empty valid %s", b)
	}
	var c string = ""
	if !IsEmpty(reflect.ValueOf(c)) {
		t.Errorf("Invalid empty valid %s", c)
	}
	var d []byte = nil
	if !IsEmpty(reflect.ValueOf(d)) {
		t.Errorf("Invalid empty valid %v", d)
	}
	var e *interface{} = nil
	if !IsEmpty(reflect.ValueOf(e)) {
		t.Errorf("Invalid empty valid %v", e)
	}
	f := struct {
		a string
		b int
	}{}
	if !IsEmpty(reflect.ValueOf(f)) {
		t.Errorf("Invalid empty valid %v", f)
	}
}

func TestPayloadWithContext_MarshalJSON(t *testing.T) {
	t.SkipNow()
}
func TestPayloadWithContext_UnmarshalJSON(t *testing.T) {
	t.SkipNow()
}
