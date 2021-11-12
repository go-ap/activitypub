package activitypub

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCollectionNew(t *testing.T) {
	var testValue = ID("test")

	c := CollectionNew(testValue)

	if c.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", c.ID, testValue)
	}
	if c.Type != CollectionType {
		t.Errorf("APObject Type '%v' different than expected '%v'", c.Type, CollectionType)
	}
}

func TestCollection_Append(t *testing.T) {
	id := ID("test")

	val := Object{ID: ID("grrr")}

	c := CollectionNew(id)
	c.Append(val)

	if c.Count() != 1 {
		t.Errorf("Inbox collection of %q should have one element", c.GetID())
	}
	if !reflect.DeepEqual(c.Items[0], val) {
		t.Errorf("First item in Inbox is does not match %q", val.ID)
	}
}

func TestCollection_Collection(t *testing.T) {
	id := ID("test")

	c := CollectionNew(id)

	if !reflect.DeepEqual(c.Collection(), c.Items) {
		t.Errorf("Collection items should be equal %v %v", c.Collection(), c.Items)
	}
}

func TestCollection_GetID(t *testing.T) {
	id := ID("test")

	c := CollectionNew(id)

	if c.GetID() != id {
		t.Errorf("GetID should return %s, received %s", id, c.GetID())
	}
}

func TestCollection_GetLink(t *testing.T) {
	id := ID("test")
	link := IRI(id)

	c := CollectionNew(id)

	if c.GetLink() != link {
		t.Errorf("GetLink should return %q, received %q", link, c.GetLink())
	}
}

func TestCollection_GetType(t *testing.T) {
	id := ID("test")

	c := CollectionNew(id)

	if c.GetType() != CollectionType {
		t.Errorf("Collection Type should be %q, received %q", CollectionType, c.GetType())
	}
}

func TestCollection_IsLink(t *testing.T) {
	id := ID("test")

	c := CollectionNew(id)

	if c.IsLink() != false {
		t.Errorf("Collection should not be a link, received %t", c.IsLink())
	}
}

func TestCollection_IsObject(t *testing.T) {
	id := ID("test")

	c := CollectionNew(id)

	if c.IsObject() != true {
		t.Errorf("Collection should be an object, received %t", c.IsObject())
	}
}

func TestCollection_UnmarshalJSON(t *testing.T) {
	c := Collection{}

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
	if len(c.Items) > 0 {
		t.Errorf("Unmarshaled object should have empty Items, received %v", c.Items)
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

func TestCollection_Count(t *testing.T) {
	id := ID("test")

	c := CollectionNew(id)

	if c.TotalItems != 0 {
		t.Errorf("Empty object should have empty TotalItems, received %d", c.TotalItems)
	}
	if len(c.Items) > 0 {
		t.Errorf("Empty object should have empty Items, received %v", c.Items)
	}
	if c.Count() != uint(len(c.Items)) {
		t.Errorf("%T.Count() returned %d, expected %d", c, c.Count(), len(c.Items))
	}

	c.Append(IRI("test"))
	if c.TotalItems != 0 {
		t.Errorf("Empty object should have empty TotalItems, received %d", c.TotalItems)
	}
	if c.Count() != uint(len(c.Items)) {
		t.Errorf("%T.Count() returned %d, expected %d", c, c.Count(), len(c.Items))
	}
}

func TestCollection_Contains(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollection_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestFollowersNew(t *testing.T) {
	t.Skipf("TODO")
}

func TestFollowingNew(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollection_MarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollection_ItemMatches(t *testing.T) {
	t.Skipf("TODO")
}

func TestToCollection(t *testing.T) {
	err := fmt.Errorf("unable to convert to collection")
	tests := map[string]struct {
		it      Item
		want    *Collection
		wantErr error
	}{
		"Collection": {
			it:      new(Collection),
			want:    new(Collection),
			wantErr: nil,
		},
		"CollectionPage": {
			it:      new(CollectionPage),
			want:    new(Collection),
			wantErr: nil,
		},
		"OrderedCollectionPage": {
			it:      new(OrderedCollectionPage),
			want:    new(Collection),
			wantErr: err,
		},
		"OrderedCollection": {
			it:      new(OrderedCollection),
			want:    new(Collection),
			wantErr: err,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ToCollection(tt.it)
			if tt.wantErr != nil && err == nil {
				t.Errorf("ToCollection() no error returned, wanted error = [%T]%s", tt.wantErr, tt.wantErr)
				return
			}
			if err != nil {
				if tt.wantErr == nil {
					t.Errorf("ToCollection() returned unexpected error[%T]%s", err, err)
					return
				}
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("ToCollection() received error %v, wanted error %v", err, tt.wantErr)
					return
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToCollection() got = %v, want %v", got, tt.want)
			}
		})
	}
}
