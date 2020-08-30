package activitypub

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOrderedCollectionPageNew(t *testing.T) {
	var testValue = ID("test")

	c := OrderedCollectionNew(testValue)
	p := OrderedCollectionPageNew(c)
	if reflect.DeepEqual(p, c) {
		t.Errorf("Invalid ordered collection parent '%v'", p.PartOf)
	}
	if p.PartOf != c.GetLink() {
		t.Errorf("Invalid collection '%v'", p.PartOf)
	}
}

func TestOrderedCollectionPage_UnmarshalJSON(t *testing.T) {
	p := OrderedCollectionPage{}

	dataEmpty := []byte("{}")
	p.UnmarshalJSON(dataEmpty)
	if p.ID != "" {
		t.Errorf("Unmarshaled object should have empty ID, received %q", p.ID)
	}
	if p.Type != "" {
		t.Errorf("Unmarshaled object should have empty Type, received %q", p.Type)
	}
	if p.AttributedTo != nil {
		t.Errorf("Unmarshaled object should have empty AttributedTo, received %q", p.AttributedTo)
	}
	if len(p.Name) != 0 {
		t.Errorf("Unmarshaled object should have empty Name, received %q", p.Name)
	}
	if len(p.Summary) != 0 {
		t.Errorf("Unmarshaled object should have empty Summary, received %q", p.Summary)
	}
	if len(p.Content) != 0 {
		t.Errorf("Unmarshaled object should have empty Content, received %q", p.Content)
	}
	if p.TotalItems != 0 {
		t.Errorf("Unmarshaled object should have empty TotalItems, received %d", p.TotalItems)
	}
	if len(p.OrderedItems) > 0 {
		t.Errorf("Unmarshaled object should have empty OrderedItems, received %v", p.OrderedItems)
	}
	if p.URL != nil {
		t.Errorf("Unmarshaled object should have empty URL, received %v", p.URL)
	}
	if !p.Published.IsZero() {
		t.Errorf("Unmarshaled object should have empty Published, received %q", p.Published)
	}
	if !p.StartTime.IsZero() {
		t.Errorf("Unmarshaled object should have empty StartTime, received %q", p.StartTime)
	}
	if !p.Updated.IsZero() {
		t.Errorf("Unmarshaled object should have empty Updated, received %q", p.Updated)
	}
	if p.PartOf != nil {
		t.Errorf("Unmarshaled object should have empty PartOf, received %q", p.PartOf)
	}
	if p.Current != nil {
		t.Errorf("Unmarshaled object should have empty Current, received %q", p.Current)
	}
	if p.First != nil {
		t.Errorf("Unmarshaled object should have empty First, received %q", p.First)
	}
	if p.Last != nil {
		t.Errorf("Unmarshaled object should have empty Last, received %q", p.Last)
	}
	if p.Next != nil {
		t.Errorf("Unmarshaled object should have empty Next, received %q", p.Next)
	}
	if p.Prev != nil {
		t.Errorf("Unmarshaled object should have empty Prev, received %q", p.Prev)
	}
}

func TestOrderedCollectionPage_Append(t *testing.T) {
	id := ID("test")

	val := Object{ID: ID("grrr")}

	c := OrderedCollectionNew(id)

	p := OrderedCollectionPageNew(c)
	p.Append(val)

	if p.PartOf != c.GetLink() {
		t.Errorf("OrderedCollection page should point to OrderedCollection %q", c.GetLink())
	}
	if p.Count() != 1 {
		t.Errorf("OrderedCollection page of %q should have exactly one element", p.GetID())
	}
	if !reflect.DeepEqual(p.OrderedItems[0], val) {
		t.Errorf("First item in Inbox is does not match %q", val.ID)
	}
}

func TestOrderedCollectionPage_Collection(t *testing.T) {
	id := ID("test")

	c := OrderedCollectionNew(id)
	p := OrderedCollectionPageNew(c)

	if !reflect.DeepEqual(p.Collection(), p.OrderedItems) {
		t.Errorf("Collection items should be equal %v %v", p.Collection(), p.OrderedItems)
	}
}

func TestOrderedCollectionPage_Contains(t *testing.T) {
	t.Skipf("TODO")
}

func TestToOrderedCollectionPage(t *testing.T) {
	err := fmt.Errorf("unable to convert to ordered collection page")
	tests := map[string]struct {
		it      Item
		want    *OrderedCollectionPage
		wantErr error
	}{
		"OrderedCollectionPage": {
			it: new(OrderedCollectionPage),
			want: new(OrderedCollectionPage),
			wantErr: nil,
		},
		"OrderedCollection": {
			it: new(OrderedCollection),
			want: new(OrderedCollectionPage),
			wantErr: err,
		},
		"Collection": {
			it: new(Collection),
			want: new(OrderedCollectionPage),
			wantErr: err,
		},
		"CollectionPage": {
			it: new(CollectionPage),
			want: new(OrderedCollectionPage),
			wantErr: err,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ToOrderedCollectionPage(tt.it)
			if tt.wantErr != nil && err == nil {
				t.Errorf("ToOrderedCollectionPage() no error returned, wanted error = [%T]%s", tt.wantErr, tt.wantErr)
				return
			} 
			if err != nil {
				if tt.wantErr == nil {
					t.Errorf("ToOrderedCollectionPage() returned unexpected error[%T]%s", err, err)
					return
				}
				if !reflect.DeepEqual(err, tt.wantErr) {
					t.Errorf("ToOrderedCollectionPage() received error %v, wanted error %v", err, tt.wantErr)
					return
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToOrderedCollectionPage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
