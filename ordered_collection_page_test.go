package activitypub

import (
	"reflect"
	"testing"

	"github.com/go-ap/errors"
)

func mockOrderedCollectionPage(items ...Item) OrderedCollectionPage {
	cc := OrderedCollectionPage{
		ID:   IRIf("https://example.com", Inbox),
		Type: OrderedCollectionPageType,
	}
	if len(items) == 0 {
		cc.OrderedItems = make(ItemCollection, 0)
	} else {
		cc.OrderedItems = items
		cc.TotalItems = uint(len(items))
	}
	return cc
}

func TestOrderedCollectionPageNew(t *testing.T) {
	testValue := ID("test")

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
	tests := []struct {
		name    string
		col     OrderedCollectionPage
		it      []Item
		wantErr error
	}{
		{
			name: "empty",
			col:  mockOrderedCollectionPage(),
			it:   ItemCollection{},
		},
		{
			name: "add one item",
			col:  mockOrderedCollectionPage(),
			it: ItemCollection{
				Object{ID: ID("grrr")},
			},
		},
		{
			name: "add multiple items",
			col:  mockOrderedCollectionPage(),
			it: ItemCollection{
				Object{ID: ID("grrr")},
				Activity{ID: ID("one")},
				Actor{ID: ID("jdoe")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				tt.col.OrderedItems = tt.col.OrderedItems[:0]
				tt.col.TotalItems = 0
			}()
			if err := tt.col.Append(tt.it...); (err != nil) && errors.Is(err, tt.wantErr) {
				t.Errorf("Append() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.col.TotalItems != uint(len(tt.it)) {
				t.Errorf("Post Append() %T TotalItems %d different than added count %d", tt.col, tt.col.TotalItems, len(tt.it))
			}
			for _, it := range tt.it {
				if !tt.col.OrderedItems.Contains(it) {
					t.Errorf("Post Append() unable to find %s in %T Items %v", tt.col, it.GetLink(), tt.col.OrderedItems)
				}
			}
		})
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

	_ = p.Append(IRI("test"))
	if p.TotalItems != 1 {
		t.Errorf("Object should have %d TotalItems, received %d", 1, p.TotalItems)
	}
	if p.Count() != uint(len(p.OrderedItems)) {
		t.Errorf("%T.Count() returned %d, expected %d", c, p.Count(), len(p.OrderedItems))
	}
}

func TestToOrderedCollectionPage(t *testing.T) {
	err := func(it Item) error { return ErrorInvalidType[OrderedCollectionPage](it) }
	tests := map[string]struct {
		it      Item
		want    *OrderedCollectionPage
		wantErr error
	}{
		"OrderedCollectionPage": {
			it:      new(OrderedCollectionPage),
			want:    new(OrderedCollectionPage),
			wantErr: nil,
		},
		"OrderedCollection": {
			it:      new(OrderedCollection),
			want:    new(OrderedCollectionPage),
			wantErr: err(new(OrderedCollection)),
		},
		"Collection": {
			it:      new(Collection),
			want:    new(OrderedCollectionPage),
			wantErr: err(new(Collection)),
		},
		"CollectionPage": {
			it:      new(CollectionPage),
			want:    new(OrderedCollectionPage),
			wantErr: nil,
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

func TestOrderedCollectionPage_Remove(t *testing.T) {
	tests := []struct {
		name      string
		col       OrderedCollectionPage
		toRemove  ItemCollection
		remaining ItemCollection
	}{
		{
			name: "empty",
			col:  mockOrderedCollectionPage(),
		},
		{
			name:     "remove one item from empty collection",
			col:      mockOrderedCollectionPage(),
			toRemove: ItemCollection{Object{ID: ID("grrr")}},
		},
		{
			name: "remove all from collection",
			col: mockOrderedCollectionPage(
				Object{ID: ID("grrr")},
				Activity{ID: ID("one")},
				Actor{ID: ID("jdoe")},
			),
			toRemove: ItemCollection{
				Object{ID: ID("grrr")},
				Activity{ID: ID("one")},
				Actor{ID: ID("jdoe")},
			},
			remaining: ItemCollection{},
		},
		{
			name:      "empty_collection_non_nil_item",
			col:       mockOrderedCollectionPage(),
			toRemove:  ItemCollection{&Object{}},
			remaining: ItemCollection{},
		},
		{
			name:      "non_empty_collection_nil_item",
			col:       mockOrderedCollectionPage(&Object{ID: "test"}),
			toRemove:  nil,
			remaining: ItemCollection{&Object{ID: "test"}},
		},
		{
			name:     "non_empty_collection_non_contained_item_empty_ID",
			col:      mockOrderedCollectionPage(&Object{ID: "test"}),
			toRemove: ItemCollection{&Object{}},
			remaining: ItemCollection{
				&Object{ID: "test"},
			},
		},
		{
			name:     "non_empty_collection_non_contained_item",
			col:      mockOrderedCollectionPage(&Object{ID: "test"}),
			toRemove: ItemCollection{&Object{ID: "test123"}},
			remaining: ItemCollection{
				&Object{ID: "test"},
			},
		},
		{
			name:      "non_empty_collection_just_contained_item",
			col:       mockOrderedCollectionPage(&Object{ID: "test"}),
			toRemove:  ItemCollection{&Object{ID: "test"}},
			remaining: ItemCollection{},
		},
		{
			name: "non_empty_collection_contained_item_first_pos",
			col: mockOrderedCollectionPage(
				&Object{ID: "test"},
				&Object{ID: "test123"},
			),
			toRemove: ItemCollection{&Object{ID: "test"}},
			remaining: ItemCollection{
				&Object{ID: "test123"},
			},
		},
		{
			name: "non_empty_collection_contained_item_not_first_pos",
			col: mockOrderedCollectionPage(
				&Object{ID: "test123"},
				&Object{ID: "test"},
				&Object{ID: "test321"},
			),
			toRemove: ItemCollection{&Object{ID: "test"}},
			remaining: ItemCollection{
				&Object{ID: "test123"},
				&Object{ID: "test321"},
			},
		},
		{
			name: "non_empty_collection_contained_item_last_pos",
			col: mockOrderedCollectionPage(
				&Object{ID: "test123"},
				&Object{ID: "test"},
			),
			toRemove: ItemCollection{&Object{ID: "test"}},
			remaining: ItemCollection{
				&Object{ID: "test123"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.col.Remove(tt.toRemove...)

			if tt.col.TotalItems != uint(len(tt.remaining)) {
				t.Errorf("Post Remove() %T TotalItems %d different than expected %d", tt.col, tt.col.TotalItems, len(tt.remaining))
			}
			for _, it := range tt.remaining {
				if !tt.col.OrderedItems.Contains(it) {
					t.Errorf("Post Remove() unable to find %s in %T Items %v", it.GetLink(), tt.col, tt.col.OrderedItems)
				}
			}
			for _, it := range tt.toRemove {
				if tt.col.OrderedItems.Contains(it) {
					t.Errorf("Post Remove() was still able to find %s in %T Items %v", it.GetLink(), tt.col, tt.col.OrderedItems)
				}
			}
		})
	}
}
