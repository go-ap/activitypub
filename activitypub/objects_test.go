package activitypub

import (
	"testing"
)

func TestObjectNew(t *testing.T) {
	var testValue = ObjectId("test")
	var testType = ArticleType

	o := ObjectNew(testValue, testType)

	if o.Id != testValue {
		t.Errorf("Object Id '%v' different than expected '%v'", o.Id, testValue)
	}
	if o.Type != testType {
		t.Errorf("Object Type '%v' different than expected '%v'", o.Type, testType)
	}

	n := ObjectNew(testValue, "")
	if n.Id != testValue {
		t.Errorf("Object Id '%v' different than expected '%v'", n.Id, testValue)
	}
	if n.Type != ObjectType {
		t.Errorf("Object Type '%v' different than expected '%v'", n.Type, ObjectType)
	}

}

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

func TestValidObjectType(t *testing.T) {
	var invalidType string = "RandomType"

	if ValidObjectType(ObjectType) {
		t.Errorf("Generic Object Type '%v' should not be valid", ObjectType)
	}
	if ValidObjectType(invalidType) {
		t.Errorf("Object Type '%v' should not be valid", invalidType)
	}
	for _, validType := range validObjectTypes {
		if !ValidObjectType(validType) {
			t.Errorf("Object Type '%v' should be valid", validType)
		}
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

func TestBaseObject_Serialize(t *testing.T) {
	//var testValue = ObjectId("test")
	//var testType = ArticleType
	//
	//o := ObjectNew(testValue, testType)
	//if o.Serialize() != "" {
	//	t.Errorf("Invalid serialize result")
	//}
}