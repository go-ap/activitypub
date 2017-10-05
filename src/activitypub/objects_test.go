package activitypub

import (
	"testing"
)

func TestObjectNew(t *testing.T) {
	var testValue = ObjectId("test")
	var testType = ArticleType

	o := ObjectNew(testValue, testType)

	if o.Id != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.Id, testValue)
	}
	if o.Type != testType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, testType)
	}

	n := ObjectNew(testValue, "")
	if n.Id != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", n.Id, testValue)
	}
	if n.Type != ObjectType {
		t.Errorf("APObject Type '%v' different than expected '%v'", n.Type, ObjectType)
	}

}

func TestValidGenericType(t *testing.T) {
	for _, validType := range validGenericObjectTypes {
		if !ValidObjectType(validType) {
			t.Errorf("Generic Type '%v' should be valid", validType)
		}
	}
}

func TestValidObjectType(t *testing.T) {
	var invalidType ActivityVocabularyType = "RandomType"

	if ValidObjectType(invalidType) {
		t.Errorf("APObject Type '%v' should not be valid", invalidType)
	}
	for _, validType := range validObjectTypes {
		if !ValidObjectType(validType) {
			t.Errorf("APObject Type '%v' should be valid", validType)
		}
	}
}

func TestMarshalJSON(t *testing.T) {
	m := make(map[LangRef]string)
	m["en"] = "test"
	m["de"] = "test"

	n := NaturalLanguageValue(m)
	result, err := n.MarshalJSON()
	if err != nil {
		t.Errorf("Failed marshaling '%v'", err)
	}
	m_res := "{\"de\":\"test\",\"en\":\"test\"}"
	if string(result) != m_res {
		t.Errorf("Different results '%v' vs. '%v'", string(result), m_res)
	}

	s := make(map[LangRef]string)
	s["en"] = "test"
	n1 := NaturalLanguageValue(s)
	result1, err1 := n1.MarshalJSON()
	if err1 != nil {
		t.Errorf("Failed marshaling '%v'", err1)
	}
	m_res1 := "\"test\""
	if string(result1) != m_res1 {
		t.Errorf("Different results '%v' vs. '%v'", string(result1), m_res1)
	}
}

func TestNaturalLanguageValue_MarshalJSON(t *testing.T) {
	p := make(NaturalLanguageValue, 2)
	p["en"] = "the test"
	p["fr"] = "le test"

	js := "{\"en\":\"the test\",\"fr\":\"le test\"}"
	out, err := p.MarshalJSON()

	if err != nil {
		t.Errorf("Error: '%s'", err)
	}
	if js != string(out) {
		t.Errorf("Different marshal result '%s', instead of '%s'", out, js)
	}
	p1 := make(NaturalLanguageValue, 1)
	p1["en"] = "the test"

	out1, err1 := p1.MarshalJSON()

	if err1 != nil {
		t.Errorf("Error: '%s'", err1)
	}
	txt := "\"the test\""
	if txt != string(out1) {
		t.Errorf("Different marshal result '%s', instead of '%s'", out1, txt)
	}
}

func TestObject_IsLink(t *testing.T) {
	o := ObjectNew("test", ObjectType)
	if o.IsLink() {
		t.Errorf("%#v should not be a valid link", o.Type)
	}
	m := ObjectNew("test", AcceptType)
	if m.IsLink() {
		t.Errorf("%#v should not be a valid link", m.Type)
	}
}

func TestObject_IsObject(t *testing.T) {
	o := ObjectNew("test", ObjectType)
	if !o.IsObject() {
		t.Errorf("%#v should be a valid object", o.Type)
	}
	m := ObjectNew("test", AcceptType)
	if !m.IsObject() {
		t.Errorf("%#v should be a valid object", m.Type)
	}
}
