package activitypub

import "testing"

func TestCollectionNew(t *testing.T) {
	var testValue = ObjectId("test")

	c := CollectionNew(testValue)

	if c.Id != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", c.Id, testValue)
	}
	if c.Type != CollectionType {
		t.Errorf("APObject Type '%v' different than expected '%v'", c.Type, CollectionType)
	}
}

func TestOrderedCollectionNew(t *testing.T) {
	var testValue = ObjectId("test")

	c := OrderedCollectionNew(testValue)

	if c.Id != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", c.Id, testValue)
	}
	if c.Type != OrderedCollectionType {
		t.Errorf("APObject Type '%v' different than expected '%v'", c.Type, OrderedCollectionType)
	}
}

func TestCollectionPageNew(t *testing.T) {
	var testValue = ObjectId("test")

	c := CollectionNew(testValue)
	p := CollectionPageNew(c)
	if p.PartOf != c {
		t.Errorf("Invalid collection '%v'", p.PartOf)
	}
}

func TestOrderedCollectionPageNew(t *testing.T) {
	var testValue = ObjectId("test")

	c := OrderedCollectionNew(testValue)
	p := OrderedCollectionPageNew(c)
	if p.PartOf != c {
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
