package activitystreams

import (
	"reflect"
	"testing"
)

func TestCollectionNew(t *testing.T) {
	var testValue = ObjectID("test")

	c := CollectionNew(testValue)

	if c.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", c.ID, testValue)
	}
	if c.Type != CollectionType {
		t.Errorf("APObject Type '%v' different than expected '%v'", c.Type, CollectionType)
	}
}

func TestOrderedCollectionNew(t *testing.T) {
	var testValue = ObjectID("test")

	c := OrderedCollectionNew(testValue)

	if c.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", c.ID, testValue)
	}
	if c.Type != OrderedCollectionType {
		t.Errorf("APObject Type '%v' different than expected '%v'", c.Type, OrderedCollectionType)
	}
}

func TestCollectionPageNew(t *testing.T) {
	var testValue = ObjectID("test")

	c := CollectionNew(testValue)
	p := CollectionPageNew(c)
	if reflect.DeepEqual(p.Collection, c) {
		t.Errorf("Invalid collection parent '%v'", p.PartOf)
	}
	if p.PartOf != c.GetLink() {
		t.Errorf("Invalid collection '%v'", p.PartOf)
	}
}

func TestOrderedCollectionPageNew(t *testing.T) {
	var testValue = ObjectID("test")

	c := OrderedCollectionNew(testValue)
	p := OrderedCollectionPageNew(c)
	if reflect.DeepEqual(p.OrderedCollection, c) {
		t.Errorf("Invalid ordered collection parent '%v'", p.PartOf)
	}
	if p.PartOf != c.GetLink() {
		t.Errorf("Invalid collection '%v'", p.PartOf)
	}
}

func TestValidCollectionType(t *testing.T) {
	for _, validType := range validCollectionTypes {
		if !ValidCollectionType(validType) {
			t.Errorf("Generic Type '%#v' should be valid", validType)
		}
	}
}

func Test_OrderedCollection_Append(t *testing.T) {
	id := ObjectID("test")

	val := Object{ID: ObjectID("grrr")}

	c := OrderedCollectionNew(id)
	c.Append(val)

	if c.TotalItems != 1 {
		t.Errorf("Inbox collection of %q should have exactly an GetID", *c.GetID())
	}
	if !reflect.DeepEqual(c.OrderedItems[0], val) {
		t.Errorf("First item in Inbox is does not match %q", val.ID)
	}
}

func Test_Collection_Append(t *testing.T) {
	id := ObjectID("test")

	val := Object{ID: ObjectID("grrr")}

	c := CollectionNew(id)
	c.Append(val)

	if c.TotalItems != 1 {
		t.Errorf("Inbox collection of %q should have exactly an GetID", *c.GetID())
	}
	if !reflect.DeepEqual(c.Items[0], val) {
		t.Errorf("First item in Inbox is does not match %q", val.ID)
	}
}
