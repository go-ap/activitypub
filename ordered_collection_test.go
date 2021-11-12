package activitypub

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOrderedCollectionNew(t *testing.T) {
	var testValue = ID("test")

	c := OrderedCollectionNew(testValue)

	if c.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", c.ID, testValue)
	}
	if c.Type != OrderedCollectionType {
		t.Errorf("APObject Type '%v' different than expected '%v'", c.Type, OrderedCollectionType)
	}
}

func Test_OrderedCollection_Append(t *testing.T) {
	id := ID("test")

	val := Object{ID: ID("grrr")}

	c := OrderedCollectionNew(id)
	c.Append(val)

	if c.Count() != 1 {
		t.Errorf("Inbox collection of %q should have one element", c.GetID())
	}
	if !reflect.DeepEqual(c.OrderedItems[0], val) {
		t.Errorf("First item in Inbox is does not match %q", val.ID)
	}
}

func TestOrderedCollection_Append(t *testing.T) {
	id := ID("test")

	val := Object{ID: ID("grrr")}

	c := OrderedCollectionNew(id)

	p := OrderedCollectionPageNew(c)
	p.Append(val)

	if p.PartOf != c.GetLink() {
		t.Errorf("Ordereed collection page should point to ordered collection %q", c.GetLink())
	}
	if p.Count() != 1 {
		t.Errorf("Ordered collection page of %q should have exactly one element", p.GetID())
	}
	if !reflect.DeepEqual(p.OrderedItems[0], val) {
		t.Errorf("First item in Inbox is does not match %q", val.ID)
	}
}

func TestOrderedCollection_Collection(t *testing.T) {
	id := ID("test")

	o := OrderedCollectionNew(id)

	if !reflect.DeepEqual(o.Collection(), o.OrderedItems) {
		t.Errorf("Collection items should be equal %v %v", o.Collection(), o.OrderedItems)
	}
}

func TestOrderedCollection_GetID(t *testing.T) {
	id := ID("test")

	c := OrderedCollectionNew(id)

	if c.GetID() != id {
		t.Errorf("GetID should return %q, received %q", id, c.GetID())
	}
}

func TestOrderedCollection_GetLink(t *testing.T) {
	id := ID("test")
	link := IRI(id)

	c := OrderedCollectionNew(id)

	if c.GetLink() != link {
		t.Errorf("GetLink should return %q, received %q", link, c.GetLink())
	}
}

func TestOrderedCollection_GetType(t *testing.T) {
	id := ID("test")

	c := OrderedCollectionNew(id)

	if c.GetType() != OrderedCollectionType {
		t.Errorf("OrderedCollection Type should be %q, received %q", OrderedCollectionType, c.GetType())
	}
}

func TestOrderedCollection_IsLink(t *testing.T) {
	id := ID("test")

	c := OrderedCollectionNew(id)

	if c.IsLink() != false {
		t.Errorf("OrderedCollection should not be a link, received %t", c.IsLink())
	}
}

func TestOrderedCollection_IsObject(t *testing.T) {
	id := ID("test")

	c := OrderedCollectionNew(id)

	if c.IsObject() != true {
		t.Errorf("OrderedCollection should be an object, received %t", c.IsObject())
	}
}

func TestOrderedCollection_UnmarshalJSON(t *testing.T) {
	c := OrderedCollection{}

	dataEmpty := []byte("{}")
	c.UnmarshalJSON(dataEmpty)
	if c.ID != "" {
		t.Errorf("Unmarshaled object should have empty ID, received %q", c.ID)
	}
	if c.Type != "" {
		t.Errorf("Unmarshaled object should have empty Type, received %q", c.Type)
	}
	if c.AttributedTo != nil {
		t.Errorf("Unmarshaled object should have empty AttributedTo, received %q", c.AttributedTo)
	}
	if len(c.Name) != 0 {
		t.Errorf("Unmarshaled object should have empty Name, received %q", c.Name)
	}
	if len(c.Summary) != 0 {
		t.Errorf("Unmarshaled object should have empty Summary, received %q", c.Summary)
	}
	if len(c.Content) != 0 {
		t.Errorf("Unmarshaled object should have empty Content, received %q", c.Content)
	}
	if c.TotalItems != 0 {
		t.Errorf("Unmarshaled object should have empty TotalItems, received %d", c.TotalItems)
	}
	if len(c.OrderedItems) > 0 {
		t.Errorf("Unmarshaled object should have empty OrderedItems, received %v", c.OrderedItems)
	}
	if c.URL != nil {
		t.Errorf("Unmarshaled object should have empty URL, received %v", c.URL)
	}
	if !c.Published.IsZero() {
		t.Errorf("Unmarshaled object should have empty Published, received %q", c.Published)
	}
	if !c.StartTime.IsZero() {
		t.Errorf("Unmarshaled object should have empty StartTime, received %q", c.StartTime)
	}
	if !c.Updated.IsZero() {
		t.Errorf("Unmarshaled object should have empty Updated, received %q", c.Updated)
	}
}

func TestOrderedCollection_Count(t *testing.T) {
	id := ID("test")

	c := OrderedCollectionNew(id)

	if c.TotalItems != 0 {
		t.Errorf("Empty object should have empty TotalItems, received %d", c.TotalItems)
	}
	if len(c.OrderedItems) > 0 {
		t.Errorf("Empty object should have empty Items, received %v", c.OrderedItems)
	}
	if c.Count() != uint(len(c.OrderedItems)) {
		t.Errorf("%T.Count() returned %d, expected %d", c, c.Count(), len(c.OrderedItems))
	}

	c.Append(IRI("test"))
	if c.TotalItems != 0 {
		t.Errorf("Empty object should have empty TotalItems, received %d", c.TotalItems)
	}
	if c.Count() != uint(len(c.OrderedItems)) {
		t.Errorf("%T.Count() returned %d, expected %d", c, c.Count(), len(c.OrderedItems))
	}
}

func TestOrderedCollectionPage_Count(t *testing.T) {
	id := ID("test")

	c := OrderedCollectionNew(id)
	p := OrderedCollectionPageNew(c)

	if p.TotalItems != 0 {
		t.Errorf("Empty object should have empty TotalItems, received %d", p.TotalItems)
	}
	if len(p.OrderedItems) > 0 {
		t.Errorf("Empty object should have empty Items, received %v", p.OrderedItems)
	}
	if p.Count() != uint(len(p.OrderedItems)) {
		t.Errorf("%T.Count() returned %d, expected %d", c, p.Count(), len(p.OrderedItems))
	}

	p.Append(IRI("test"))
	if p.TotalItems != 0 {
		t.Errorf("Empty object should have empty TotalItems, received %d", p.TotalItems)
	}
	if p.Count() != uint(len(p.OrderedItems)) {
		t.Errorf("%T.Count() returned %d, expected %d", c, p.Count(), len(p.OrderedItems))
	}
}

func TestOnOrderedCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestToOrderedCollection(t *testing.T) {
	err := fmt.Errorf("unable to convert to ordered collection")
	tests := map[string]struct {
		it      Item
		want    *OrderedCollection
		wantErr error
	}{
		"OrderedCollection": {
			it:      new(OrderedCollection),
			want:    new(OrderedCollection),
			wantErr: nil,
		},
		"OrderedCollectionPage": {
			it:      new(OrderedCollectionPage),
			want:    new(OrderedCollection),
			wantErr: nil,
		},
		"Collection": {
			it:      new(Collection),
			want:    new(OrderedCollection),
			wantErr: err,
		},
		"CollectionPage": {
			it:      new(CollectionPage),
			want:    new(OrderedCollection),
			wantErr: err,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ToOrderedCollection(tt.it)
			if tt.wantErr != nil && err == nil {
				t.Errorf("ToOrderedCollection() no error returned, wanted error = [%T]%s", tt.wantErr, tt.wantErr)
				return
			}
			if err != nil {
				if tt.wantErr == nil {
					t.Errorf("ToOrderedCollection() returned unexpected error[%T]%s", err, err)
					return
				}
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("ToOrderedCollection() received error %v, wanted error %v", err, tt.wantErr)
					return
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToOrderedCollection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedCollection_Contains(t *testing.T) {
	t.Skipf("TODO")
}

func TestInboxNew(t *testing.T) {
	i := InboxNew()

	id := ID("inbox")
	if i.ID != id {
		t.Errorf("%T should be initialized with %q as %T", i, id, id)
	}
	if len(i.Name) != 0 {
		t.Errorf("%T should be initialized with 0 length Name", i)
	}
	if len(i.Content) != 0 {
		t.Errorf("%T should be initialized with 0 length Content", i)
	}
	if len(i.Summary) != 0 {
		t.Errorf("%T should be initialized with 0 length Summary", i)
	}
	if i.TotalItems != 0 {
		t.Errorf("%T should be initialized with 0 TotalItems", i)
	}
}

func TestLikedNew(t *testing.T) {
	l := LikedNew()

	id := ID("liked")
	if l.ID != id {
		t.Errorf("%T should be initialized with %q as %T", l, id, id)
	}
	if len(l.Name) != 0 {
		t.Errorf("%T should be initialized with 0 length Name", l)
	}
	if len(l.Content) != 0 {
		t.Errorf("%T should be initialized with 0 length Content", l)
	}
	if len(l.Summary) != 0 {
		t.Errorf("%T should be initialized with 0 length Summary", l)
	}
	if l.TotalItems != 0 {
		t.Errorf("%T should be initialized with 0 TotalItems", l)
	}
}

func TestLikesNew(t *testing.T) {
	t.Skipf("TODO")
}

func TestOutboxNew(t *testing.T) {
	o := OutboxNew()

	id := ID("outbox")
	if o.ID != id {
		t.Errorf("%T should be initialized with %q as %T", o, id, id)
	}
	if len(o.Name) != 0 {
		t.Errorf("%T should be initialized with 0 length Name", o)
	}
	if len(o.Content) != 0 {
		t.Errorf("%T should be initialized with 0 length Content", o)
	}
	if len(o.Summary) != 0 {
		t.Errorf("%T should be initialized with 0 length Summary", o)
	}
	if o.TotalItems != 0 {
		t.Errorf("%T should be initialized with 0 TotalItems", o)
	}
}

func TestSharesNew(t *testing.T) {
	t.Skipf("TODO")
}

func TestOrderedCollection_MarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestOrderedCollection_ItemMatches(t *testing.T) {
	t.Skipf("TODO")
}

func TestOrderedCollection_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}
