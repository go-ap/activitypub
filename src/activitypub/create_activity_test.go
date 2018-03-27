package activitypub

import (
	"reflect"
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

	testValue = ObjectId("my:note")
	n := ObjectNew("my:note", NoteType)
	c1 := CreateActivityNew(testValue, n)
	now = time.Now()
	if c1.Activity.Id != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c1.Activity.Id, testValue)
	}
	if c1.Activity.Type != CreateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c1.Activity.Type, CreateType)
	}
	if now.Sub(c.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c1.Published, now)
	}
	if !reflect.DeepEqual(c1.Activity.Object.(*apObject), n) {
		t.Errorf("Object \n'%#v' different than expected \n'%#v'", c1.Activity.Object, n)
	}
}
