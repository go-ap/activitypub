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
}

func TestCreateActivityNewWithApplication(t *testing.T) {
	testValue := ObjectID("my:note")
	n := ObjectNew("my:note", NoteType)
	a := ApplicationNew("some::app::")

	c1 := CreateActivityNew(testValue, *a, n)
	now := time.Now()
	if c1.Activity.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c1.Activity.ID, testValue)
	}
	if c1.Activity.Type != CreateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c1.Activity.Type, CreateType)
	}
	if now.Sub(c1.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c1.Published, now)
	}
	if c1.Activity.Actor.Object().ID != a.ID {
		t.Errorf("Actor ID %q different than expected %q", c1.Activity.Actor.Object().ID, a.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Actor.Object(), a.Object()) {
		t.Errorf("Actor %#v different than expected %#v", c1.Activity.Actor.Object(), a.Object())
	}
	if !reflect.DeepEqual(c1.Activity.Actor, Actor(*a)) {
		t.Errorf("Actor %#v\n\n different than expected\n\n %#v", c1.Activity.Actor, Actor(*a))
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

func TestCreateActivityNewWithGroup(t *testing.T) {
	testValue := ObjectID("my:note")
	n := ObjectNew("my:note", NoteType)
	g := GroupNew("users")

	c1 := CreateActivityNew(testValue, *g, n)
	now := time.Now()
	if c1.Activity.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c1.Activity.ID, testValue)
	}
	if c1.Activity.Type != CreateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c1.Activity.Type, CreateType)
	}
	if now.Sub(c1.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c1.Published, now)
	}
	if c1.Activity.Actor.Object().ID != g.ID {
		t.Errorf("Actor ID %q different than expected %q", c1.Activity.Actor.Object().ID, g.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Actor.Object(), g.Object()) {
		t.Errorf("Actor %#v different than expected %#v", c1.Activity.Actor.Object(), g.Object())
	}
	if !reflect.DeepEqual(c1.Activity.Actor, Actor(*g)) {
		t.Errorf("Actor %#v\n\n different than expected\n\n %#v", c1.Activity.Actor, Actor(*g))
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

func TestCreateActivityNewWithOrganization(t *testing.T) {
	testValue := ObjectID("my:note")
	n := ObjectNew("my:note", NoteType)
	o := OrganizationNew("users")

	c1 := CreateActivityNew(testValue, *o, n)
	now := time.Now()
	if c1.Activity.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c1.Activity.ID, testValue)
	}
	if c1.Activity.Type != CreateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c1.Activity.Type, CreateType)
	}
	if now.Sub(c1.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c1.Published, now)
	}
	if c1.Activity.Actor.Object().ID != o.ID {
		t.Errorf("Actor ID %q different than expected %q", c1.Activity.Actor.Object().ID, o.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Actor.Object(), o.Object()) {
		t.Errorf("Actor %#v different than expected %#v", c1.Activity.Actor.Object(), o.Object())
	}
	if !reflect.DeepEqual(c1.Activity.Actor, Actor(*o)) {
		t.Errorf("Actor %#v\n\n different than expected\n\n %#v", c1.Activity.Actor, Actor(*o))
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

func TestCreateActivityNewWithPerson(t *testing.T) {
	testValue := ObjectID("my:note")
	n := ObjectNew("my:note", NoteType)
	b := PersonNew("bob")

	c1 := CreateActivityNew(testValue, *b, n)
	now := time.Now()
	if c1.Activity.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c1.Activity.ID, testValue)
	}
	if c1.Activity.Type != CreateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c1.Activity.Type, CreateType)
	}
	if now.Sub(c1.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c1.Published, now)
	}
	if c1.Activity.Actor.Object().ID != b.ID {
		t.Errorf("Actor ID %q different than expected %q", c1.Activity.Actor.Object().ID, b.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Actor.Object(), b.Object()) {
		t.Errorf("Actor %#v different than expected %#v", c1.Activity.Actor.Object(), b.Object())
	}
	if !reflect.DeepEqual(c1.Activity.Actor, Actor(*b)) {
		t.Errorf("Actor %#v\n\n different than expected\n\n %#v", c1.Activity.Actor, Actor(*b))
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

func TestCreateActivityNewWithService(t *testing.T) {
	testValue := ObjectID("my:note")
	n := ObjectNew("my:note", NoteType)
	s := ServiceNew("::zz::")

	c1 := CreateActivityNew(testValue, *s, n)
	now := time.Now()
	if c1.Activity.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c1.Activity.ID, testValue)
	}
	if c1.Activity.Type != CreateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c1.Activity.Type, CreateType)
	}
	if now.Sub(c1.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c1.Published, now)
	}
	if c1.Activity.Actor.Object().ID != s.ID {
		t.Errorf("Actor ID %q different than expected %q", c1.Activity.Actor.Object().ID, s.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Actor.Object(), s.Object()) {
		t.Errorf("Actor %#v different than expected %#v", c1.Activity.Actor.Object(), s.Object())
	}
	if !reflect.DeepEqual(c1.Activity.Actor, Actor(*s)) {
		t.Errorf("Actor %#v\n\n different than expected\n\n %#v", c1.Activity.Actor, Actor(*s))
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

func TestCreateActivityNewWithActor(t *testing.T) {
	testValue := ObjectID("my:note")
	n := ObjectNew("my:note", NoteType)
	a := ActorNew("bob", ActorType)

	c1 := CreateActivityNew(testValue, *a, n)
	now := time.Now()
	if c1.Activity.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c1.Activity.ID, testValue)
	}
	if c1.Activity.Type != CreateType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c1.Activity.Type, CreateType)
	}
	if now.Sub(c1.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c1.Published, now)
	}
	if c1.Activity.Actor.Object().ID != a.ID {
		t.Errorf("Actor ID %q different than expected %q", c1.Activity.Actor.Object().ID, a.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Actor.Object(), a.Object()) {
		t.Errorf("Actor %#v different than expected %#v", c1.Activity.Actor.Object(), a.Object())
	}
	if !reflect.DeepEqual(c1.Activity.Actor, *a) {
		t.Errorf("Actor %#v\n\n different than expected\n\n %#v", c1.Activity.Actor, *a)
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
