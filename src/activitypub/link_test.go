package activitypub

import "testing"

func TestLinkNew(t *testing.T) {
	var testValue = ObjectId("test")
	var testType string

	l := LinkNew(testValue, testType)

	if l.Id != testValue {
		t.Errorf("Object Id '%v' different than expected '%v'", l.Id, testValue)
	}
	if l.Type != LinkType {
		t.Errorf("Object Type '%v' different than expected '%v'", l.Type, LinkType)
	}
}

func TestValidLinkType(t *testing.T) {
	var invalidType string = "RandomType"

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
