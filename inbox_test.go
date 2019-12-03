package activitypub

import (
	"reflect"
	"testing"
)

func TestInboxNew(t *testing.T) {
	i := InboxNew()

	id := ObjectID("inbox")
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
	if o.GetID() != "" {
		t.Errorf("%T should be initialized with empty %T", o, o.GetID())
	}
	id := ObjectID("test_out_stream")
	o.ID = id
	if o.GetID() != id {
		t.Errorf("%T should have %T as %q", o, id, id)
	}
}

func TestInboxStream_GetType(t *testing.T) {
	o := InboxStream{}

	if o.GetType() != "" {
		t.Errorf("%T should be initialized with empty %T", o, o.GetType())
	}

	o.Type = OrderedCollectionType
	if o.GetType() != OrderedCollectionType {
		t.Errorf("%T should have %T as %q", o, o.GetType(), OrderedCollectionType)
	}
}

func TestInboxStream_Append(t *testing.T) {
	o := InboxStream{}

	val := Object{ID: ObjectID("grrr")}

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

	val := Object{ID: ObjectID("grrr")}

	i.Append(val)
	if i.TotalItems != 0 {
		t.Errorf("%T should have exactly an element, found %d", i, i.TotalItems)
	}
	if !reflect.DeepEqual(i.OrderedItems[0], val) {
		t.Errorf("First item in %T.%T does not match %q", i, i.OrderedItems, val.ID)
	}
}

func TestInbox_Collection(t *testing.T) {
	t.Skipf("TODO")
}

func TestInbox_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestInbox_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestInbox_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestInbox_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestInbox_IsObject(t *testing.T) {
	t.Skipf("TODO")
}

func TestInbox_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}
