package activitypub

import "testing"

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
