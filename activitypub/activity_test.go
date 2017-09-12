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

	g := ActivityNew(testValue, "")

	if g.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", g.Id, testValue)
	}
	if g.Type != ActivityType {
		t.Errorf("Activity Type '%v' different than expected '%v'", g.Type, ActivityType)
	}
}

func TestValidActivityType(t *testing.T) {
	var invalidType string = "RandomType"

	if ValidActivityType(ActivityType) {
		t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
	}
	if ValidActivityType(invalidType) {
		t.Errorf("Activity Type '%v' should not be valid", invalidType)
	}
	for _, validType := range validActivityTypes {
		if !ValidActivityType(validType) {
			t.Errorf("Activity Type '%v' should be valid", validType)
		}
	}
}