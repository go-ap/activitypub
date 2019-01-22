package activitypub

import (
	as "github.com/go-ap/activitystreams"
	"reflect"
	"testing"
)

func TestInboxNew(t *testing.T) {
	i := InboxNew()

	id := as.ObjectID("inbox")
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

func TestInboxStream_GetID(t *testing.T) {
	o := InboxStream{}
	if *o.GetID() != "" {
		t.Errorf("%T should be initialized with empty %T", o, o.GetID())
	}
	id := as.ObjectID("test_out_stream")
	o.ID = id
	if *o.GetID() != id {
		t.Errorf("%T should have %T as %q", o, id, id)
	}
}

func TestInboxStream_GetType(t *testing.T) {
	o := InboxStream{}

	if o.GetType() != "" {
		t.Errorf("%T should be initialized with empty %T", o, o.GetType())
	}

	o.Type = as.OrderedCollectionType
	if o.GetType() != as.OrderedCollectionType {
		t.Errorf("%T should have %T as %q", o, o.GetType(), as.OrderedCollectionType)
	}
}

func TestInboxStream_Append(t *testing.T) {
	o := InboxStream{}

	val := as.Object{ID: as.ObjectID("grrr")}

	o.Append(val)
	if o.TotalItems != 1 {
		t.Errorf("%T should have exactly an element, found %d", o, o.TotalItems)
	}
	if !reflect.DeepEqual(o.OrderedItems[0], val) {
		t.Errorf("First item in %T.%T does not match %q", o, o.OrderedItems, val.ID)
	}
}
func TestInbox_Append(t *testing.T) {
	i := InboxNew()

	val := as.Object{ID: as.ObjectID("grrr")}

	i.Append(val)
	if i.TotalItems != 1 {
		t.Errorf("%T should have exactly an element, found %d", i, i.TotalItems)
	}
	if !reflect.DeepEqual(i.OrderedItems[0], val) {
		t.Errorf("First item in %T.%T does not match %q", i, i.OrderedItems, val.ID)
	}
}
