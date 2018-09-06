package activitypub

import (
	"reflect"
	"testing"
	"time"
)

func TestLikeActivityNew(t *testing.T) {
	var testValue = ObjectID("test")
	var now time.Time

	c := LikeActivityNew(testValue, nil, nil)
	now = time.Now()
	if c.Activity.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c.Activity.ID, testValue)
	}
	if c.Activity.Type != LikeType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c.Activity.Type, LikeType)
	}
	if now.Sub(c.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c.Published, now)
	}
}

func TestLikeActivityNewWithApplication(t *testing.T) {
	testValue := ObjectID("my:note")
	n := ObjectNew("my:note", NoteType)
	a := ApplicationNew("some::app::")

	c1 := LikeActivityNew(testValue, *a, n)
	now := time.Now()
	if c1.Activity.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c1.Activity.ID, testValue)
	}
	if c1.Activity.Type != LikeType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c1.Activity.Type, LikeType)
	}
	if now.Sub(c1.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c1.Published, now)
	}
	if *c1.Activity.Actor.GetID() != a.ID {
		t.Errorf("Actor ID %q different than expected %q", *c1.Activity.Actor.GetID(), a.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Actor.GetID(), a.GetID()) {
		t.Errorf("Actor %#v different than expected %#v", c1.Activity.Actor.GetID(), a.GetID())
	}
	if !reflect.DeepEqual(c1.Activity.Actor, *a) {
		t.Errorf("Actor %#v\n\n different than expected\n\n %#v", c1.Activity.Actor, Actor(*a))
	}
	if *c1.Activity.Object.GetID() != n.ID {
		t.Errorf("GetID %q different than expected %q", *c1.Activity.Object.GetID(), n.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Object.GetID(), n.GetID()) {
		t.Errorf("GetID %#v different than expected %#v", c1.Activity.Object.GetID(), n.GetID())
	}
	if !reflect.DeepEqual(c1.Activity.Object, n) {
		t.Errorf("GetID %#v different than expected %#v", c1.Activity.Object, n)
	}
	in := c1.Activity.Actor.(Application).Liked.(*Liked)
	if in.TotalItems != 1 {
		t.Errorf("Liked collection of %q should have exactly one element, not %d", *c1.Activity.Actor.GetID(), in.TotalItems)
	}
	if len(in.OrderedItems) != 1 {
		t.Errorf("Liked collection length of %q should have exactly one element, not %d", *c1.Activity.Actor.GetID(), len(in.OrderedItems))
	}
	if in.TotalItems != uint(len(in.OrderedItems)) {
		t.Errorf("Liked collection length of %q should have same size as TotalItems, %d vs %d", *c1.Activity.Actor.GetID(), in.TotalItems, len(in.OrderedItems))
	}
	if !reflect.DeepEqual(in.OrderedItems[0].GetID(), n.GetID()) {
		t.Errorf("First item in Liked is does not match %q", *n.GetID())
	}
}

func TestLikeActivityNewWithGroup(t *testing.T) {
	testValue := ObjectID("my:note")
	n := ObjectNew("my:note", NoteType)
	g := GroupNew("users")

	c1 := LikeActivityNew(testValue, *g, n)
	now := time.Now()
	if c1.Activity.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c1.Activity.ID, testValue)
	}
	if c1.Activity.Type != LikeType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c1.Activity.Type, LikeType)
	}
	if now.Sub(c1.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c1.Published, now)
	}
	if *c1.Activity.Actor.GetID() != g.ID {
		t.Errorf("Actor ID %q different than expected %q", *c1.Activity.Actor.GetID(), g.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Actor.GetID(), g.GetID()) {
		t.Errorf("Actor %#v different than expected %#v", c1.Activity.Actor.GetID(), g.GetID())
	}
	if !reflect.DeepEqual(c1.Activity.Actor, *g) {
		t.Errorf("Actor %#v\n\n different than expected\n\n %#v", c1.Activity.Actor, Actor(*g))
	}
	if *c1.Activity.Object.GetID() != n.ID {
		t.Errorf("GetID %q different than expected %q", *c1.Activity.Object.GetID(), n.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Object.GetID(), n.GetID()) {
		t.Errorf("GetID %#v different than expected %#v", c1.Activity.Object.GetID(), n.GetID())
	}
	if !reflect.DeepEqual(c1.Activity.Object, n) {
		t.Errorf("GetID %#v different than expected %#v", c1.Activity.Object, n)
	}
}

func TestLikeActivityNewWithOrganization(t *testing.T) {
	testValue := ObjectID("my:note")
	n := ObjectNew("my:note", NoteType)
	o := OrganizationNew("users")

	c1 := LikeActivityNew(testValue, *o, n)
	now := time.Now()
	if c1.Activity.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c1.Activity.ID, testValue)
	}
	if c1.Activity.Type != LikeType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c1.Activity.Type, LikeType)
	}
	if now.Sub(c1.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c1.Published, now)
	}
	if *c1.Activity.Actor.GetID() != o.ID {
		t.Errorf("Actor ID %q different than expected %q", *c1.Activity.Actor.GetID(), o.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Actor.GetID(), o.GetID()) {
		t.Errorf("Actor %#v different than expected %#v", c1.Activity.Actor.GetID(), o.GetID())
	}
	if !reflect.DeepEqual(c1.Activity.Actor, *o) {
		t.Errorf("Actor %#v\n\n different than expected\n\n %#v", c1.Activity.Actor, Actor(*o))
	}
	if *c1.Activity.Object.GetID() != n.ID {
		t.Errorf("GetID %q different than expected %q", *c1.Activity.Object.GetID(), n.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Object.GetID(), n.GetID()) {
		t.Errorf("GetID %#v different than expected %#v", c1.Activity.Object.GetID(), n.GetID())
	}
	if !reflect.DeepEqual(c1.Activity.Object, n) {
		t.Errorf("GetID %#v different than expected %#v", c1.Activity.Object, n)
	}
}

func TestLikeActivityNewWithPerson(t *testing.T) {
	testValue := ObjectID("my:note")
	n := ObjectNew("my:note", NoteType)
	b := PersonNew("bob")

	c1 := LikeActivityNew(testValue, *b, n)
	now := time.Now()
	if c1.Activity.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c1.Activity.ID, testValue)
	}
	if c1.Activity.Type != LikeType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c1.Activity.Type, LikeType)
	}
	if now.Sub(c1.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c1.Published, now)
	}
	if *c1.Activity.Actor.GetID() != b.ID {
		t.Errorf("Actor ID %q different than expected %q", *c1.Activity.Actor.GetID(), b.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Actor.GetID(), b.GetID()) {
		t.Errorf("Actor %#v different than expected %#v", c1.Activity.Actor.GetID(), b.GetID())
	}
	if !reflect.DeepEqual(c1.Activity.Actor, *b) {
		t.Errorf("Actor %#v\n\n different than expected\n\n %#v", c1.Activity.Actor, Actor(*b))
	}
	if *c1.Activity.Object.GetID() != n.ID {
		t.Errorf("GetID %q different than expected %q", *c1.Activity.Object.GetID(), n.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Object.GetID(), n.GetID()) {
		t.Errorf("GetID %#v different than expected %#v", c1.Activity.Object.GetID(), n.GetID())
	}
	if !reflect.DeepEqual(c1.Activity.Object, n) {
		t.Errorf("GetID %#v different than expected %#v", c1.Activity.Object, n)
	}
}

func TestLikeActivityNewWithService(t *testing.T) {
	testValue := ObjectID("my:note")
	n := ObjectNew("my:note", NoteType)
	s := ServiceNew("::zz::")

	c1 := LikeActivityNew(testValue, *s, n)
	now := time.Now()
	if c1.Activity.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c1.Activity.ID, testValue)
	}
	if c1.Activity.Type != LikeType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c1.Activity.Type, LikeType)
	}
	if now.Sub(c1.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c1.Published, now)
	}
	if *c1.Activity.Actor.GetID() != s.ID {
		t.Errorf("Actor ID %q different than expected %q", *c1.Activity.Actor.GetID(), s.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Actor.GetID(), s.GetID()) {
		t.Errorf("Actor %#v different than expected %#v", c1.Activity.Actor.GetID(), s.GetID())
	}
	if !reflect.DeepEqual(c1.Activity.Actor, *s) {
		t.Errorf("Actor %#v\n\n different than expected\n\n %#v", c1.Activity.Actor, Actor(*s))
	}
	if *c1.Activity.Object.GetID() != n.ID {
		t.Errorf("GetID %q different than expected %q", *c1.Activity.Object.GetID(), n.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Object.GetID(), n.GetID()) {
		t.Errorf("GetID %#v different than expected %#v", c1.Activity.Object.GetID(), n.GetID())
	}
	if !reflect.DeepEqual(c1.Activity.Object, n) {
		t.Errorf("GetID %#v different than expected %#v", c1.Activity.Object, n)
	}
}

func TestLikeActivityNewWithActor(t *testing.T) {
	testValue := ObjectID("my:note")
	n := ObjectNew("my:note", NoteType)
	a := ActorNew("bob", ActorType)

	c1 := LikeActivityNew(testValue, *a, n)
	now := time.Now()
	if c1.Activity.ID != testValue {
		t.Errorf("Activity Id '%v' different than expected '%v'", c1.Activity.ID, testValue)
	}
	if c1.Activity.Type != LikeType {
		t.Errorf("Activity Type '%v' different than expected '%v'", c1.Activity.Type, LikeType)
	}
	if now.Sub(c1.Published).Round(time.Millisecond) != 0 {
		t.Errorf("Published time '%v' different than expected '%v'", c1.Published, now)
	}
	if *c1.Activity.Actor.GetID() != a.ID {
		t.Errorf("Actor ID %q different than expected %q", *c1.Activity.Actor.GetID(), a.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Actor.GetID(), a.GetID()) {
		t.Errorf("Actor %#v different than expected %#v", c1.Activity.Actor.GetID(), a.GetID())
	}
	if !reflect.DeepEqual(c1.Activity.Actor, *a) {
		t.Errorf("Actor %#v\n\n different than expected\n\n %#v", c1.Activity.Actor, *a)
	}
	if *c1.Activity.Object.GetID() != n.ID {
		t.Errorf("GetID %q different than expected %q", *c1.Activity.Object.GetID(), n.ID)
	}
	if !reflect.DeepEqual(c1.Activity.Object.GetID(), n.GetID()) {
		t.Errorf("GetID %#v different than expected %#v", c1.Activity.Object.GetID(), n.GetID())
	}
	if !reflect.DeepEqual(c1.Activity.Object, n) {
		t.Errorf("GetID %#v different than expected %#v", c1.Activity.Object, n)
	}
}
