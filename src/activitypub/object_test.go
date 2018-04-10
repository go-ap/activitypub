package activitypub

import (
	"reflect"
	"testing"
)

func TestObjectNew(t *testing.T) {
	var testValue = ObjectID("test")
	var testType = ArticleType

	o := ObjectNew(testValue, testType)

	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != testType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, testType)
	}

	n := ObjectNew(testValue, "")
	if n.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", n.ID, testValue)
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
	mRes := "{\"de\":\"test\",\"en\":\"test\"}"
	if string(result) != mRes {
		t.Errorf("Different results '%v' vs. '%v'", string(result), mRes)
	}

	s := make(map[LangRef]string)
	s["en"] = "test"
	n1 := NaturalLanguageValue(s)
	result1, err1 := n1.MarshalJSON()
	if err1 != nil {
		t.Errorf("Failed marshaling '%v'", err1)
	}
	mRes1 := "\"test\""
	if string(result1) != mRes1 {
		t.Errorf("Different results '%v' vs. '%v'", string(result1), mRes1)
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

func TestObject_Link(t *testing.T) {
	o := ObjectNew("test", ObjectType)
	if !reflect.DeepEqual(Link{}, o.Link()) {
		t.Errorf("%#v should be an empty Link object", o.Link())
	}
}

func TestObjectsArr_Append(t *testing.T) {
	d := make(ObjectsArr, 0)

	val := apObject{ID: ObjectID("grrr")}

	d.Append(val)

	if len(d) != 1 {
		t.Errorf("Objects array should have exactly an element")
	}
	if !reflect.DeepEqual(d[0], val) {
		t.Errorf("First item in object array does not match %q", val.ID)
	}
}

func TestRecipientsDeduplication(t *testing.T) {
	bob := PersonNew("bob")
	alice := PersonNew("alice")
	foo := OrganizationNew("foo")
	bar := GroupNew("bar")

	first := make(ObjectsArr, 0)
	if len(first) != 0 {
		t.Errorf("Objects array should have exactly an element")
	}

	first.Append(bob)
	first.Append(alice)
	first.Append(foo)
	first.Append(bar)
	if len(first) != 4 {
		t.Errorf("Objects array should have exactly 4(four) elements, not %d", len(first))
	}

	first.Append(bar)
	first.Append(alice)
	first.Append(foo)
	first.Append(bob)
	if len(first) != 8 {
		t.Errorf("Objects array should have exactly 8(eight) elements, not %d", len(first))
	}

	RecipientsDeduplication(&first)
	if len(first) != 4 {
		t.Errorf("Objects array should have exactly 4(four) elements, not %d", len(first))
	}

	second := make(ObjectsArr, 0)
	second.Append(bar)
	second.Append(foo)

	RecipientsDeduplication(&first, &second)
	if len(first) != 4 {
		t.Errorf("First Objects array should have exactly 8(eight) elements, not %d", len(first))
	}
	if len(second) != 0 {
		t.Errorf("Second Objects array should have exactly 0(zero) elements, not %d", len(second))
	}
}
