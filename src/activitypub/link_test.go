package activitypub

import (
	"reflect"
	"testing"
)

func TestLinkNew(t *testing.T) {
	var testValue = ObjectID("test")
	var testType ActivityVocabularyType

	l := LinkNew(testValue, testType)

	if l.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", l.ID, testValue)
	}
	if l.Type != LinkType {
		t.Errorf("APObject Type '%v' different than expected '%v'", l.Type, LinkType)
	}
}

func TestValidLinkType(t *testing.T) {
	var invalidType ActivityVocabularyType = "RandomType"

	if ValidLinkType(LinkType) {
		t.Errorf("Generic GetLink Type '%v' should not be valid", LinkType)
	}
	if ValidLinkType(invalidType) {
		t.Errorf("GetLink Type '%v' should not be valid", invalidType)
	}
	for _, validType := range validLinkTypes {
		if !ValidLinkType(validType) {
			t.Errorf("GetLink Type '%v' should be valid", validType)
		}
	}
}

func TestLink_IsLink(t *testing.T) {
	l := LinkNew("test", LinkType)
	if !l.IsLink() {
		t.Errorf("%#v should be a valid link", l.Type)
	}
	m := LinkNew("test", MentionType)
	if !m.IsLink() {
		t.Errorf("%#v should be a valid link", m.Type)
	}
}

func TestLink_IsObject(t *testing.T) {
	l := LinkNew("test", LinkType)
	if l.IsObject() {
		t.Errorf("%#v should not be a valid object", l.Type)
	}
	m := LinkNew("test", MentionType)
	if m.IsObject() {
		t.Errorf("%#v should not be a valid object", m.Type)
	}
}

func TestMention_IsLink(t *testing.T) {
	m := MentionNew("test")
	if !m.IsLink() {
		t.Errorf("%#v should be a valid Mention", m.Type)
	}
}

func TestMention_IsObject(t *testing.T) {
	m := MentionNew("test")
	if m.IsObject() {
		t.Errorf("%#v should not be a valid object", m.Type)
	}
}

func TestMention_Object(t *testing.T) {
	m := MentionNew("test")
	if !reflect.DeepEqual(BaseObject{}, m.GetObject()) {
		t.Errorf("%#v should be an empty object", m.GetObject())
	}
}

func TestMention_Link(t *testing.T) {
	m := MentionNew("test")
	if !reflect.DeepEqual(Link(*m), m.GetLink()) {
		t.Errorf("%#v should be a GetLink object", m.GetLink())
	}
}
