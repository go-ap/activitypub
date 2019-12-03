package activitypub

import (
	"testing"
)

func TestLikedNew(t *testing.T) {
	l := LikedNew()

	id := ObjectID("liked")
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
