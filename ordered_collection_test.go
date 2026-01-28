package activitypub

import (
	"reflect"
	"testing"

	"github.com/go-ap/errors"
	"github.com/google/go-cmp/cmp"
)

func mockOrderedCollection(items ...Item) OrderedCollection {
	cc := OrderedCollection{
		ID:   IRIf("https://example.com", Inbox),
		Type: OrderedCollectionType,
	}
	if len(items) == 0 {
		cc.OrderedItems = make(ItemCollection, 0)
	} else {
		cc.OrderedItems = items
		cc.TotalItems = uint(len(items))
	}
	return cc
}

func TestOrderedCollectionNew(t *testing.T) {
	testValue := ID("test")

	c := OrderedCollectionNew(testValue)

	if c.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", c.ID, testValue)
	}
	if !c.Matches(OrderedCollectionType) {
		t.Errorf("APObject Type '%v' different than expected '%v'", c.GetType(), OrderedCollectionType)
	}
}

func TestOrderedCollection_Append(t *testing.T) {
	tests := []struct {
		name    string
		col     OrderedCollection
		it      []Item
		wantErr error
	}{
		{
			name: "empty",
			col:  mockOrderedCollection(),
			it:   ItemCollection{},
		},
		{
			name: "add one item",
			col:  mockOrderedCollection(),
			it: ItemCollection{
				Object{ID: ID("grrr")},
			},
		},
		{
			name: "add multiple items",
			col:  mockOrderedCollection(),
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

	if !c.Matches(OrderedCollectionType) {
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
	if HasTypes(c) {
		t.Errorf("Unmarshaled object should have empty Type, received %q", c.GetType())
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

	_ = c.Append(IRI("test"))
	if c.TotalItems != 1 {
		t.Errorf("Object should have %d TotalItems, received %d", 1, c.TotalItems)
	}
	if c.Count() != uint(len(c.OrderedItems)) {
		t.Errorf("%T.Count() returned %d, expected %d", c, c.Count(), len(c.OrderedItems))
	}
}

func TestOnOrderedCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestToOrderedCollection(t *testing.T) {
	tests := []struct {
		name    string
		it      LinkOrIRI
		want    *OrderedCollection
		wantErr error
	}{
		{
			name: "empty",
		},
		{
			name: "Valid OrderedCollection",
			it:   OrderedCollection{ID: "test", Type: OrderedCollectionType},
			want: &OrderedCollection{ID: "test", Type: OrderedCollectionType},
		},
		{
			name: "Valid *OrderedCollection",
			it:   &OrderedCollection{ID: "test", Type: OrderedCollectionType},
			want: &OrderedCollection{ID: "test", Type: OrderedCollectionType},
		},
		{
			name: "Valid OrderedCollectionPage",
			it:   OrderedCollectionPage{ID: "test", Type: OrderedCollectionPageType},
			want: &OrderedCollection{ID: "test", Type: OrderedCollectionPageType},
		},
		{
			name: "Valid *OrderedCollectionPage",
			it:   &OrderedCollectionPage{ID: "test", Type: OrderedCollectionPageType},
			want: &OrderedCollection{ID: "test", Type: OrderedCollectionPageType},
		},
		{
			name: "Valid OrderedCollection",
			it:   OrderedCollection{ID: "test", Type: OrderedCollectionType},
			want: &OrderedCollection{ID: "test", Type: OrderedCollectionType},
		},
		{
			name: "Valid *OrderedCollection",
			it:   &OrderedCollection{ID: "test", Type: OrderedCollectionType},
			want: &OrderedCollection{ID: "test", Type: OrderedCollectionType},
		},
		{
			name: "Valid OrderedCollectionPage",
			it:   OrderedCollectionPage{ID: "test", Type: OrderedCollectionPageType},
			want: &OrderedCollection{ID: "test", Type: OrderedCollectionPageType},
		},
		{
			name: "Valid *OrderedCollectionPage",
			it:   &OrderedCollectionPage{ID: "test", Type: OrderedCollectionPageType},
			want: &OrderedCollection{ID: "test", Type: OrderedCollectionPageType},
		},
		{
			name:    "IRI",
			it:      IRI("https://example.com"),
			wantErr: ErrorInvalidType[OrderedCollection](IRI("")),
		},
		{
			name:    "IRIs",
			it:      IRIs{IRI("https://example.com")},
			wantErr: ErrorInvalidType[OrderedCollection](IRIs{}),
		},
		{
			name:    "ItemCollection",
			it:      ItemCollection{},
			wantErr: ErrorInvalidType[OrderedCollection](ItemCollection{}),
		},
		{
			name:    "Object",
			it:      &Object{ID: "test", Type: ArticleType},
			wantErr: ErrorInvalidType[OrderedCollection](&Object{}),
		},
		{
			name:    "Activity",
			it:      &Activity{ID: "test", Type: CreateType},
			wantErr: ErrorInvalidType[OrderedCollection](&Activity{}),
		},
		{
			name:    "IntransitiveActivity",
			it:      &IntransitiveActivity{ID: "test", Type: ArriveType},
			wantErr: ErrorInvalidType[OrderedCollection](&IntransitiveActivity{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToOrderedCollection(tt.it)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("ToOrderedCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ToOrderedCollection() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestOrderedCollection_Contains(t *testing.T) {
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

func TestOrderedCollection_Remove(t *testing.T) {
	tests := []struct {
		name      string
		col       OrderedCollection
		toRemove  ItemCollection
		remaining ItemCollection
	}{
		{
			name: "empty",
			col:  mockOrderedCollection(),
		},
		{
			name:     "remove one item from empty collection",
			col:      mockOrderedCollection(),
			toRemove: ItemCollection{Object{ID: ID("grrr")}},
		},
		{
			name: "remove all from collection",
			col: mockOrderedCollection(
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
			col:       mockOrderedCollection(),
			toRemove:  ItemCollection{&Object{}},
			remaining: ItemCollection{},
		},
		{
			name:      "non_empty_collection_nil_item",
			col:       mockOrderedCollection(&Object{ID: "test"}),
			toRemove:  nil,
			remaining: ItemCollection{&Object{ID: "test"}},
		},
		{
			name:     "non_empty_collection_non_contained_item_empty_ID",
			col:      mockOrderedCollection(&Object{ID: "test"}),
			toRemove: ItemCollection{&Object{}},
			remaining: ItemCollection{
				&Object{ID: "test"},
			},
		},
		{
			name:     "non_empty_collection_non_contained_item",
			col:      mockOrderedCollection(&Object{ID: "test"}),
			toRemove: ItemCollection{&Object{ID: "test123"}},
			remaining: ItemCollection{
				&Object{ID: "test"},
			},
		},
		{
			name:      "non_empty_collection_just_contained_item",
			col:       mockOrderedCollection(&Object{ID: "test"}),
			toRemove:  ItemCollection{&Object{ID: "test"}},
			remaining: ItemCollection{},
		},
		{
			name: "non_empty_collection_contained_item_first_pos",
			col: mockOrderedCollection(
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
			col: mockOrderedCollection(
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
			col: mockOrderedCollection(
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
