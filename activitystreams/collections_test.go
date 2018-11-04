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
		t.Errorf("Inbox collection of %q should have one element", *c.GetID())
	}
	if !reflect.DeepEqual(c.OrderedItems[0], val) {
		t.Errorf("First item in Inbox is does not match %q", val.ID)
	}
}

func TestCollection_Append(t *testing.T) {
	id := ObjectID("test")

	val := Object{ID: ObjectID("grrr")}

	c := CollectionNew(id)
	c.Append(val)

	if c.TotalItems != 1 {
		t.Errorf("Inbox collection of %q should have one element", *c.GetID())
	}
	if !reflect.DeepEqual(c.Items[0], val) {
		t.Errorf("First item in Inbox is does not match %q", val.ID)
	}
}

func TestCollectionPage_Append(t *testing.T) {
	id := ObjectID("test")

	val := Object{ID: ObjectID("grrr")}

	c := CollectionNew(id)

	p := CollectionPageNew(c)
	p.Append(val)

	if p.PartOf != c.GetLink() {
		t.Errorf("Collection page should point to collection %q", c.GetLink())
	}
	if p.TotalItems != 1 {
		t.Errorf("Collection page of %q should have exactly one element", *p.GetID())
	}
	if !reflect.DeepEqual(p.Items[0], val) {
		t.Errorf("First item in Inbox is does not match %q", val.ID)
	}
}

func TestCollection_Collection(t *testing.T) {
	id := ObjectID("test")

	c := CollectionNew(id)

	if c.Collection() != c {
		t.Errorf("Collection should return itself %q", *c.GetID())
	}
}

func TestCollection_GetID(t *testing.T) {
	id := ObjectID("test")

	c := CollectionNew(id)

	if *c.GetID() != id {
		t.Errorf("GetID should return %s, received %s", id, *c.GetID())
	}
}

func TestCollection_GetLink(t *testing.T) {
	id := ObjectID("test")
	link := IRI(id)

	c := CollectionNew(id)

	if c.GetLink() != link {
		t.Errorf("GetLink should return %q, received %q", link, c.GetLink())
	}
}

func TestCollection_GetType(t *testing.T) {
	id := ObjectID("test")

	c := CollectionNew(id)

	if c.GetType() != CollectionType {
		t.Errorf("Collection Type should be %q, received %q", CollectionType, c.GetType())
	}
}

func TestCollection_IsLink(t *testing.T) {
	id := ObjectID("test")

	c := CollectionNew(id)

	if c.IsLink() != false {
		t.Errorf("Collection should not be a link, received %t", c.IsLink())
	}
}

func TestCollection_IsObject(t *testing.T) {
	id := ObjectID("test")

	c := CollectionNew(id)

	if c.IsObject() != true {
		t.Errorf("Collection should be an object, received %t", c.IsObject())
	}
}

func TestCollection_UnmarshalJSON(t *testing.T) {

}

func TestCollectionPage_UnmarshalJSON(t *testing.T) {

}

func TestOrderedCollection_Append(t *testing.T) {
	id := ObjectID("test")

	val := Object{ID: ObjectID("grrr")}

	c := OrderedCollectionNew(id)

	p := OrderedCollectionPageNew(c)
	p.Append(val)

	if p.PartOf != c.GetLink() {
		t.Errorf("Ordereed collection page should point to ordered collection %q", c.GetLink())
	}
	if p.TotalItems != 1 {
		t.Errorf("Ordered collection page of %q should have exactly one element", *p.GetID())
	}
	if !reflect.DeepEqual(p.OrderedItems[0], val) {
		t.Errorf("First item in Inbox is does not match %q", val.ID)
	}
}

func TestOrderedCollection_Collection(t *testing.T) {
	id := ObjectID("test")

	c := OrderedCollectionNew(id)

	if c.Collection() != c {
		t.Errorf("Collection should return itself %q", *c.GetID())
	}
}

func TestOrderedCollection_GetID(t *testing.T) {
	id := ObjectID("test")

	c := OrderedCollectionNew(id)

	if *c.GetID() != id {
		t.Errorf("GetID should return %q, received %q", id, *c.GetID())
	}
}

func TestOrderedCollection_GetLink(t *testing.T) {
	id := ObjectID("test")
	link := IRI(id)

	c := OrderedCollectionNew(id)

	if c.GetLink() != link {
		t.Errorf("GetLink should return %q, received %q", link, c.GetLink())
	}
}

func TestOrderedCollection_GetType(t *testing.T) {
	id := ObjectID("test")

	c := OrderedCollectionNew(id)

	if c.GetType() != OrderedCollectionType {
		t.Errorf("OrderedCollection Type should be %q, received %q", OrderedCollectionType, c.GetType())
	}
}

func TestOrderedCollection_IsLink(t *testing.T) {
	id := ObjectID("test")

	c := OrderedCollectionNew(id)

	if c.IsLink() != false {
		t.Errorf("OrderedCollection should not be a link, received %t", c.IsLink())
	}
}

func TestOrderedCollection_IsObject(t *testing.T) {
	id := ObjectID("test")

	c := OrderedCollectionNew(id)

	if c.IsObject() != true {
		t.Errorf("OrderedCollection should be an object, received %t", c.IsObject())
	}
}

func TestOrderedCollection_UnmarshalJSON(t *testing.T) {

}

func TestOrderedCollectionPage_Append(t *testing.T) {

}

func TestOrderedCollectionPage_UnmarshalJSON(t *testing.T) {
	id := ObjectID("test")

	val := Object{ID: ObjectID("grrr")}

	c := OrderedCollectionNew(id)

	p := OrderedCollectionPageNew(c)
	p.Append(val)

	if p.PartOf != c.GetLink() {
		t.Errorf("OrderedCollection page should point to OrderedCollection %q", c.GetLink())
	}
	if p.TotalItems != 1 {
		t.Errorf("OrderedCollection page of %q should have exactly one element", *p.GetID())
	}
	if !reflect.DeepEqual(p.OrderedItems[0], val) {
		t.Errorf("First item in Inbox is does not match %q", val.ID)
	}
}
