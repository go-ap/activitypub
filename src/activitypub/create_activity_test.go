package activitypub

import (
	"reflect"
	"testing"
	"time"
)

func TestCreateActivityNew(t *testing.T) {
	var testValue = ObjectID("test")
	var now time.Time

	c := CreateActivityNew(testValue, nil, nil)
	now = time.Now()
	if c.Activity.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c.Activity.ID, testValue)
	}
	if c.Activity.Type != CreateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c.Activity.Type, CreateType)
	}
	if now.Sub(c.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c.Published, now)
	}

	testValue = ObjectID("my:note")
	n := ObjectNew("my:note", NoteType)
	b := PersonNew("bob")

	c1 := CreateActivityNew(testValue, *b, n)
	now = time.Now()
	if c1.Activity.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c1.Activity.ID, testValue)
	}
	if c1.Activity.Type != CreateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c1.Activity.Type, CreateType)
	}
	if now.Sub(c.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c1.Published, now)
	}
	if c1.Activity.Actor.Object().ID != b.ID {
		t.Errorf("Actor ID %q different than expected %q", c1.Activity.Actor.Object().ID, b.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Actor.Object(), b.Object()) {
		t.Errorf("Actor %#v different than expected %#v", c1.Activity.Actor.Object(), b.Object())
	}
	if !reflect.DeepEqual(c1.Activity.Actor, *b) {
		t.Errorf("Actor %#v\n\n different than expected\n\n %#v", c1.Activity.Actor, *b)
	}
	if c1.Activity.Object.Object().ID != n.ID {
		t.Errorf("Object %q different than expected %q", c1.Activity.Object.Object().ID, n.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Object.Object(), n.Object()) {
		t.Errorf("Object %#v different than expected %#v", c1.Activity.Object.Object(), n.Object())
	}
	if !reflect.DeepEqual(c1.Activity.Object, n) {
		t.Errorf("Object %#v different than expected %#v", c1.Activity.Object, n)
	}
}
