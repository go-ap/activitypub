package activitypub

import (
	"testing"
	"time"
)

func TestCreateActivityNew(t *testing.T) {
	var testValue = ObjectId("test")
	var now time.Time

	c := CreateActivityNew(testValue, nil)
	now = time.Now()
	if c.Activity.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c.Activity.Id, testValue)
	}
	if c.Activity.Type != CreateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c.Activity.Type, CreateType)
	}
	if now.Sub(c.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c.Published, now)
	}
}
