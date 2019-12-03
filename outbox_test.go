package activitypub

import (
	"reflect"
	"testing"
)

func TestOutboxNew(t *testing.T) {
	o := OutboxNew()

	id := ObjectID("outbox")
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

func TestOutboxStream_GetID(t *testing.T) {
	o := OutboxStream{}
	if o.GetID() != "" {
		t.Errorf("%T should be initialized with empty %T", o, o.GetID())
	}
	id := ObjectID("test_out_stream")
	o.ID = id
	if o.GetID() != id {
		t.Errorf("%T should have %T as %q", o, id, id)
	}
}

func TestOutboxStream_GetType(t *testing.T) {
	o := OutboxStream{}

	if o.GetType() != "" {
		t.Errorf("%T should be initialized with empty %T", o, o.GetType())
	}

	o.Type = OrderedCollectionType
	if o.GetType() != OrderedCollectionType {
		t.Errorf("%T should have %T as %q", o, o.GetType(), OrderedCollectionType)
	}
}

func TestOutboxStream_Append(t *testing.T) {
	o := OutboxStream{}

	val := Object{ID: ObjectID("grrr")}

	o.Append(val)
	if !reflect.DeepEqual(o.OrderedItems[0], val) {
		t.Errorf("First item in %T.%T does not match %q", o, o.OrderedItems, val.ID)
	}
}

func TestOutbox_Append(t *testing.T) {
	o := OutboxNew()

	val := Object{ID: ObjectID("grrr")}

	o.Append(val)
	if !reflect.DeepEqual(o.OrderedItems[0], val) {
		t.Errorf("First item in %T.%T does not match %q", o, o.OrderedItems, val.ID)
	}
}

func TestOutbox_Collection(t *testing.T) {
	t.Skipf("TODO")
}

func TestOutbox_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestOutbox_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestOutbox_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestOutbox_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestOutbox_IsObject(t *testing.T) {
	t.Skipf("TODO")
}

func TestOutbox_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}
