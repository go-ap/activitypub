package activitypub

import "testing"

func TestLinkNew(t *testing.T) {
	var testValue = ObjectId("test")
	var testType ActivityVocabularyType

	l := LinkNew(testValue, testType)

	if l.Id != testValue {
		t.Errorf("Object Id '%v' different than expected '%v'", l.Id, testValue)
	}
	if l.Type != LinkType {
		t.Errorf("Object Type '%v' different than expected '%v'", l.Type, LinkType)
	}
}

func TestValidLinkType(t *testing.T) {
	var invalidType ActivityVocabularyType = "RandomType"

	if ValidLinkType(LinkType) {
		t.Errorf("Generic Link Type '%v' should not be valid", LinkType)
	}
	if ValidLinkType(invalidType) {
		t.Errorf("Link Type '%v' should not be valid", invalidType)
	}
	for _, validType := range validLinkTypes {
		if !ValidLinkType(validType) {
			t.Errorf("Link Type '%v' should be valid", validType)
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

func TestMention_IsMention(t *testing.T) {
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
