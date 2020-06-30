package activitypub

import (
	"testing"
)

func TestLinkNew(t *testing.T) {
	var testValue = ID("test")
	var testType ActivityVocabularyType

	l := LinkNew(testValue, testType)

	if l.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", l.ID, testValue)
	}
	if l.Type != LinkType {
		t.Errorf("APObject Type '%v' different than expected '%v'", l.Type, LinkType)
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

func TestLink_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestLink_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestLink_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestLink_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestMentionNew(t *testing.T) {
	t.Skipf("TODO")
}

func TestLink_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}
