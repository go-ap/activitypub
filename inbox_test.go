package activitypub

import (
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
