package activitypub

import "testing"

func TestCollectionsNew(t *testing.T) {
	var testValue = ObjectId("test")

	c := CollectionNew(testValue)

	if c.Id != testValue {
		t.Errorf("Object Id '%v' different than expected '%v'", c.Id, testValue)
	}
	if c.Type != CollectionType {
		t.Errorf("Object Type '%v' different than expected '%v'", c.Type, CollectionType)
	}
}

func TestOrderedCollectionsNew(t *testing.T) {
	var testValue = ObjectId("test")

	c := OrderedCollectionNew(testValue)

	if c.Id != testValue {
		t.Errorf("Object Id '%v' different than expected '%v'", c.Id, testValue)
	}
	if c.Type != OrderedCollectionType {
		t.Errorf("Object Type '%v' different than expected '%v'", c.Type, OrderedCollectionType)
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
