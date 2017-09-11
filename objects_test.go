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

func TestActivityNew(t *testing.T) {
	var testValue = ObjectId("test")
	var testType string = "Accept"

	a := ActivityNew(testValue, testType)

	if a.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", a.Id, testValue)
	}
	if a.Type != testType {
		t.Errorf("Activity Type '%v' different than expected '%v'", a.Type, testType)
	}
}
