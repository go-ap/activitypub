package activitypub

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/go-ap/errors"
	"github.com/google/go-cmp/cmp"
)

func TestItemCollection_Append(t *testing.T) {
	tests := []struct {
		name     string
		items    ItemCollection
		toAppend ItemCollection
		want     ItemCollection
		wantErr  error
	}{
		{
			name:     "empty_collection_nil_item",
			items:    ItemCollection{},
			toAppend: nil,
			want:     ItemCollection{},
		},
		{
			name:     "empty_collection_non_nil_item",
			items:    ItemCollection{},
			toAppend: ItemCollection{&Object{}},
			want:     ItemCollection{&Object{}},
		},
		{
			name:     "non_empty_collection_nil_item",
			items:    ItemCollection{&Object{ID: "test"}},
			toAppend: nil,
			want:     ItemCollection{&Object{ID: "test"}},
		},
		{
			name:     "non_empty_collection_non_contained_item_empty_ID",
			items:    ItemCollection{&Object{ID: "test"}},
			toAppend: ItemCollection{&Object{}},
			want:     ItemCollection{&Object{ID: "test"}, &Object{}},
		},
		{
			name:     "non_empty_collection_non_contained_item",
			items:    ItemCollection{&Object{ID: "test"}},
			toAppend: ItemCollection{&Object{ID: "test123"}},
			want:     ItemCollection{&Object{ID: "test"}, &Object{ID: "test123"}},
		},
		{
			name:     "non_empty_collection_just_contained_item",
			items:    ItemCollection{&Object{ID: "test"}},
			toAppend: ItemCollection{&Object{ID: "test"}},
			want:     ItemCollection{&Object{ID: "test"}},
		},
		{
			name:     "non_empty_collection_contained_item_first_pos",
			items:    ItemCollection{&Object{ID: "test"}, &Object{ID: "test123"}},
			toAppend: ItemCollection{&Object{ID: "test"}},
			want:     ItemCollection{&Object{ID: "test"}, &Object{ID: "test123"}},
		},
		{
			name: "non_empty_collection_contained_item_not_first_pos",
			items: ItemCollection{&Object{ID: "test123"}, &Object{ID: "test"}, &Object{ID: "test321"},
			},
			toAppend: ItemCollection{&Object{ID: "test"}},
			want:     ItemCollection{&Object{ID: "test123"}, &Object{ID: "test321"}, &Object{ID: "test"}},
		},
		{
			name:     "non_empty_collection_add_item",
			items:    ItemCollection{&Object{ID: "test123"}, &Object{ID: "test"}},
			toAppend: ItemCollection{&Object{ID: "test1234"}},
			want:     ItemCollection{&Object{ID: "test123"}, &Object{ID: "test"}, &Object{ID: "test1234"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.items.Append(tt.toAppend...)
			if !cmp.Equal(err, tt.wantErr) {
				t.Errorf("Append() error = %s", cmp.Diff(tt.wantErr, err))
			}

			if tt.want.Count() != tt.items.Count() {
				t.Errorf("Post Append() %T has count %d different than expected %d", tt.items, tt.items.Count(), tt.want.Count())
			}
			for _, it := range tt.toAppend {
				if !tt.items.Contains(it) {
					t.Errorf("Post Append() unable able to find %s in %T Items %v", it.GetLink(), tt.items, tt.items)
				}
			}
		})
	}
}

func TestItemCollection_Collection(t *testing.T) {
	tests := []struct {
		name string
		i    *ItemCollection
		want ItemCollection
	}{
		{
			name: "empty",
			i:    nil,
		},
		{
			name: "no-items",
			i:    &ItemCollection{},
		},
		{
			name: "one-nil",
			i:    &ItemCollection{nil},
			want: ItemCollection{nil},
		},
		{
			name: "two-nils",
			i:    &ItemCollection{nil, nil},
			want: ItemCollection{nil, nil},
		},
		{
			name: "one",
			i:    &ItemCollection{&Object{}},
			want: ItemCollection{&Object{}},
		},
		{
			name: "three",
			i:    &ItemCollection{&Object{}, IRI("test"), Item(nil)},
			want: ItemCollection{&Object{}, IRI("test"), Item(nil)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Collection(); !cmp.Equal(got, tt.want) {
				t.Errorf("Collection() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestItemCollection_GetID(t *testing.T) {
	tests := []struct {
		name string
		ic   ItemCollection
		want IRI
	}{
		{
			name: "empty",
		},
		{
			name: "not-empty",
			ic:   ItemCollection{IRI("http://example.com")},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ic.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemCollection_GetLink(t *testing.T) {
	tests := []struct {
		name string
		ic   ItemCollection
		want IRI
	}{
		{
			name: "empty",
		},
		{
			name: "not-empty",
			ic:   ItemCollection{IRI("http://example.com")},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ic.GetLink(); got != tt.want {
				t.Errorf("GetLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemCollection_GetType(t *testing.T) {
	tests := []struct {
		name string
		ic   ItemCollection
		want ActivityVocabularyType
	}{
		{
			name: "empty",
			want: CollectionOfItems,
		},
		{
			name: "not-empty",
			ic:   ItemCollection{IRI("http://example.com")},
			want: CollectionOfItems,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ic.GetType(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemCollection_First(t *testing.T) {
	tests := []struct {
		name string
		i    ItemCollection
		want Item
	}{
		{
			name: "empty",
			i:    nil,
			want: nil,
		},
		{
			name: "nil item",
			i:    ItemCollection{nil},
			want: nil,
		},
		{
			name: "nil item from two",
			i:    ItemCollection{nil, IRI("test")},
			want: nil,
		},
		{
			name: "iri from two",
			i:    ItemCollection{IRI("test"), Object{}},
			want: IRI("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.First(); !cmp.Equal(got, tt.want) {
				t.Errorf("First() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestItemCollection_Last(t *testing.T) {
	tests := []struct {
		name string
		i    ItemCollection
		want Item
	}{
		{
			name: "empty",
			i:    nil,
			want: nil,
		},
		{
			name: "nil item",
			i:    ItemCollection{nil},
			want: nil,
		},
		{
			name: "nil item from two",
			i:    ItemCollection{IRI("test"), nil},
			want: nil,
		},
		{
			name: "iri from two",
			i:    ItemCollection{Object{}, IRI("test")},
			want: IRI("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Last(); !cmp.Equal(got, tt.want) {
				t.Errorf("Last() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestItemCollection_Count(t *testing.T) {
	tests := []struct {
		name string
		i    *ItemCollection
		want uint
	}{
		{
			name: "empty",
			i:    nil,
			want: 0,
		},
		{
			name: "no-items",
			i:    &ItemCollection{},
			want: 0,
		},
		{
			name: "one-nil",
			i:    &ItemCollection{nil},
			want: 1,
		},
		{
			name: "two-nils",
			i:    &ItemCollection{nil, nil},
			want: 2,
		},
		{
			name: "one",
			i:    &ItemCollection{&Object{}},
			want: 1,
		},
		{
			name: "three",
			i:    &ItemCollection{&Object{}, IRI("test"), Item(nil)},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Count(); got != tt.want {
				t.Errorf("Count() = got %d, expected %d", got, tt.want)
			}
		})
	}
}

func TestItemCollection_Contains(t *testing.T) {
	t.Skipf("TODO")
}

func TestItemCollection_Remove(t *testing.T) {
	tests := []struct {
		name      string
		items     ItemCollection
		toRemove  ItemCollection
		remaining ItemCollection
	}{
		{
			name:      "empty_collection_nil_item",
			items:     ItemCollection{},
			toRemove:  nil,
			remaining: ItemCollection{},
		},
		{
			name:      "empty_collection_non_nil_item",
			items:     ItemCollection{},
			toRemove:  ItemCollection{&Object{}},
			remaining: ItemCollection{},
		},
		{
			name: "non_empty_collection_nil_item",
			items: ItemCollection{
				&Object{ID: "test"},
			},
			toRemove: nil,
			remaining: ItemCollection{
				&Object{ID: "test"},
			},
		},
		{
			name: "non_empty_collection_non_contained_item_empty_ID",
			items: ItemCollection{
				&Object{ID: "test"},
			},
			toRemove: ItemCollection{&Object{}},
			remaining: ItemCollection{
				&Object{ID: "test"},
			},
		},
		{
			name: "non_empty_collection_non_contained_item",
			items: ItemCollection{
				&Object{ID: "test"},
			},
			toRemove: ItemCollection{&Object{ID: "test123"}},
			remaining: ItemCollection{
				&Object{ID: "test"},
			},
		},
		{
			name: "non_empty_collection_just_contained_item",
			items: ItemCollection{
				&Object{ID: "test"},
			},
			toRemove:  ItemCollection{&Object{ID: "test"}},
			remaining: ItemCollection{},
		},
		{
			name: "non_empty_collection_contained_item_first_pos",
			items: ItemCollection{
				&Object{ID: "test"},
				&Object{ID: "test123"},
			},
			toRemove: ItemCollection{&Object{ID: "test"}},
			remaining: ItemCollection{
				&Object{ID: "test123"},
			},
		},
		{
			name: "non_empty_collection_contained_item_not_first_pos",
			items: ItemCollection{
				&Object{ID: "test123"},
				&Object{ID: "test"},
				&Object{ID: "test321"},
			},
			toRemove: ItemCollection{&Object{ID: "test"}},
			remaining: ItemCollection{
				&Object{ID: "test123"},
				&Object{ID: "test321"},
			},
		},
		{
			name: "non_empty_collection_contained_item_last_pos",
			items: ItemCollection{
				&Object{ID: "test123"},
				&Object{ID: "test"},
			},
			toRemove: ItemCollection{&Object{ID: "test"}},
			remaining: ItemCollection{
				&Object{ID: "test123"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.items.Remove(tt.toRemove...)

			if tt.remaining.Count() != tt.items.Count() {
				t.Errorf("Post Remove() %T has count %d different than expected %d", tt.items, tt.items.Count(), tt.remaining.Count())
			}
			for _, it := range tt.toRemove {
				if tt.items.Contains(it) {
					t.Errorf("Post Remove() was still able to find %s in %T Items %v", it.GetLink(), tt.items, tt.items)
				}
			}
			for _, it := range tt.remaining {
				if !tt.items.Contains(it) {
					t.Errorf("Post Remove() unable to find %s in %T Items %v", it.GetLink(), tt.items, tt.items)
				}
			}
		})
	}
}

func TestItemCollectionDeduplication(t *testing.T) {
	tests := []struct {
		name      string
		args      []*ItemCollection
		want      ItemCollection
		remaining []*ItemCollection
	}{
		{
			name: "empty",
		},
		{
			name: "no-overlap",
			args: []*ItemCollection{
				{
					IRI("https://example.com"),
					IRI("https://example.com/2"),
				},
				{
					IRI("https://example.com/1"),
				},
			},
			want: ItemCollection{
				IRI("https://example.com"),
				IRI("https://example.com/2"),
				IRI("https://example.com/1"),
			},
			remaining: []*ItemCollection{
				{
					IRI("https://example.com"),
					IRI("https://example.com/2"),
				},
				{
					IRI("https://example.com/1"),
				},
			},
		},
		{
			name: "some-overlap",
			args: []*ItemCollection{
				{
					IRI("https://example.com"),
					IRI("https://example.com/2"),
				},
				{
					IRI("https://example.com/1"),
					IRI("https://example.com/2"),
				},
			},
			want: ItemCollection{
				IRI("https://example.com"),
				IRI("https://example.com/2"),
				IRI("https://example.com/1"),
			},
			remaining: []*ItemCollection{
				{
					IRI("https://example.com"),
					IRI("https://example.com/2"),
				},
				{
					IRI("https://example.com/1"),
				},
			},
		},
		{
			name: "test from spammy",
			args: []*ItemCollection{
				{
					IRI("https://example.dev/a801139a-0d9a-4703-b0a5-9d14ae1438e4/followers"),
					IRI("https://www.w3.org/ns/activitystreams#Public"),
				},
				{
					IRI("https://example.dev/a801139a-0d9a-4703-b0a5-9d14ae1438e4"),
				},
				{
					IRI("https://example.dev"),
					IRI("https://example.dev/a801139a-0d9a-4703-b0a5-9d14ae1438e4"),
					IRI("https://www.w3.org/ns/activitystreams#Public"),
				},
			},
			want: ItemCollection{
				IRI("https://example.dev/a801139a-0d9a-4703-b0a5-9d14ae1438e4/followers"),
				IRI("https://www.w3.org/ns/activitystreams#Public"),
				IRI("https://example.dev/a801139a-0d9a-4703-b0a5-9d14ae1438e4"),
				IRI("https://example.dev"),
			},
			remaining: []*ItemCollection{
				{
					IRI("https://example.dev/a801139a-0d9a-4703-b0a5-9d14ae1438e4/followers"),
					IRI("https://www.w3.org/ns/activitystreams#Public"),
				},
				{
					IRI("https://example.dev/a801139a-0d9a-4703-b0a5-9d14ae1438e4"),
				},
				{
					IRI("https://example.dev"),
				},
			},
		},
		{
			name: "different order for spammy test",
			args: []*ItemCollection{
				{
					IRI("https://example.dev/a801139a-0d9a-4703-b0a5-9d14ae1438e4/followers"),
					IRI("https://www.w3.org/ns/activitystreams#Public"),
				},
				{
					IRI("https://example.dev"),
					IRI("https://example.dev/a801139a-0d9a-4703-b0a5-9d14ae1438e4"),
					IRI("https://www.w3.org/ns/activitystreams#Public"),
				},
				{
					IRI("https://example.dev/a801139a-0d9a-4703-b0a5-9d14ae1438e4"),
				},
			},
			want: ItemCollection{
				IRI("https://example.dev/a801139a-0d9a-4703-b0a5-9d14ae1438e4/followers"),
				IRI("https://www.w3.org/ns/activitystreams#Public"),
				IRI("https://example.dev"),
				IRI("https://example.dev/a801139a-0d9a-4703-b0a5-9d14ae1438e4"),
			},
			remaining: []*ItemCollection{
				{
					IRI("https://example.dev/a801139a-0d9a-4703-b0a5-9d14ae1438e4/followers"),
					IRI("https://www.w3.org/ns/activitystreams#Public"),
				},
				{
					IRI("https://example.dev"),
					IRI("https://example.dev/a801139a-0d9a-4703-b0a5-9d14ae1438e4"),
				},
				{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ItemCollectionDeduplication(tt.args...); !cmp.Equal(tt.want, got) {
				t.Errorf("ItemCollectionDeduplication() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestToItemCollection(t *testing.T) {
	iri := IRI("https://example.com")
	tests := []struct {
		name    string
		it      LinkOrIRI
		want    *ItemCollection
		wantErr error
	}{
		{
			name: "empty",
		},
		{
			name: "IRI to ItemCollection",
			it:   IRI("https://example.com"),
			want: &ItemCollection{IRI("https://example.com")},
		},
		{
			name: "*IRI to ItemCollection",
			it:   &iri,
			want: &ItemCollection{IRI("https://example.com")},
		},
		{
			name: "IRIs to ItemCollection",
			it:   IRIs{"https://example.com", "https://example.com/example"},
			want: &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
		},
		{
			name: "*IRIs to ItemCollection",
			it:   &IRIs{"https://example.com", "https://example.com/example"},
			want: &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
		},
		{
			name: "ItemCollection to ItemCollection",
			it:   ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
			want: &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
		},
		{
			name: "*ItemCollection to ItemCollection",
			it:   &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
			want: &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
		},
		{
			name: "*Collection to ItemCollection",
			it:   &Collection{Items: ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")}},
			want: &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
		},
		{
			name: "Collection to ItemCollection",
			it:   Collection{Items: ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")}},
			want: &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
		},
		{
			name: "*CollectionPage to ItemCollection",
			it:   &CollectionPage{Items: ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")}},
			want: &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
		},
		{
			name: "CollectionPage to ItemCollection",
			it:   CollectionPage{Items: ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")}},
			want: &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
		},
		{
			name: "*OrderedCollection to ItemCollection",
			it:   &OrderedCollection{OrderedItems: ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")}},
			want: &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
		},
		{
			name: "OrderedCollection to ItemCollection",
			it:   OrderedCollection{OrderedItems: ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")}},
			want: &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
		},
		{
			name: "*OrderedCollectionPage to ItemOrderedCollection",
			it:   &OrderedCollectionPage{OrderedItems: ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")}},
			want: &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
		},
		{
			name: "OrderedCollectionPage to ItemOrderedCollection",
			it:   OrderedCollectionPage{OrderedItems: ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")}},
			want: &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
		},
		{
			name: "no-items",
			it:   &ItemCollection{},
			want: &ItemCollection{},
		},
		{
			name: "one-nil",
			it:   &ItemCollection{nil},
			want: &ItemCollection{nil},
		},
		{
			name: "two-nils",
			it:   &ItemCollection{nil, nil},
			want: &ItemCollection{nil, nil},
		},
		{
			name: "one",
			it:   &ItemCollection{&Object{}},
			want: &ItemCollection{&Object{}},
		},
		{
			name: "three",
			it:   &ItemCollection{&Object{}, IRI("test"), Item(nil)},
			want: &ItemCollection{&Object{}, IRI("test"), Item(nil)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToItemCollection(tt.it)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("ToItemCollection() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ToItemCollection() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestToIRIs(t *testing.T) {
	iri := IRI("https://example.com")
	tests := []struct {
		name    string
		it      LinkOrIRI
		want    *IRIs
		wantErr error
	}{
		{
			name: "empty",
		},
		{
			name: "IRI to ItemCollection",
			it:   IRI("https://example.com"),
			want: &IRIs{"https://example.com"},
		},
		{
			name: "*IRI to ItemCollection",
			it:   &iri,
			want: &IRIs{"https://example.com"},
		},
		{
			name: "IRIs to ItemCollection",
			it:   IRIs{"https://example.com", "https://example.com/example"},
			want: &IRIs{"https://example.com", "https://example.com/example"},
		},
		{
			name: "*IRIs to ItemCollection",
			it:   &IRIs{"https://example.com", "https://example.com/example"},
			want: &IRIs{"https://example.com", "https://example.com/example"},
		},
		{
			name: "ItemCollection to ItemCollection",
			it:   ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
			want: &IRIs{"https://example.com", "https://example.com/example"},
		},
		{
			name: "*ItemCollection to ItemCollection",
			it:   &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
			want: &IRIs{"https://example.com", "https://example.com/example"},
		},
		{
			name: "no-items",
			it:   &ItemCollection{},
			want: &IRIs{},
		},
		{
			name: "just-nils",
			it:   &ItemCollection{nil, nil},
			want: &IRIs{},
		},
		{
			name: "just-nils in item collection",
			it:   ItemCollection{nil, nil},
			want: &IRIs{},
		},
		{
			name: "with nils",
			it:   &ItemCollection{IRI("test"), Item(nil)},
			want: &IRIs{IRI("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToIRIs(tt.it)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("ToIRIs() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ToIRIs() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestItemCollection_IRIs(t *testing.T) {
	tests := []struct {
		name string
		i    ItemCollection
		want IRIs
	}{
		{
			name: "empty",
			i:    nil,
			want: nil,
		},
		{
			name: "one item",
			i: ItemCollection{
				&Object{ID: "https://example.com"},
			},
			want: IRIs{"https://example.com"},
		},
		{
			name: "two items",
			i: ItemCollection{
				&Object{ID: "https://example.com"},
				&Actor{ID: "https://example.com/~jdoe"},
			},
			want: IRIs{"https://example.com", "https://example.com/~jdoe"},
		},
		{
			name: "mixed items",
			i: ItemCollection{
				&Object{ID: "https://example.com"},
				IRI("https://example.com/666"),
				&Actor{ID: "https://example.com/~jdoe"},
			},
			want: IRIs{"https://example.com", "https://example.com/666", "https://example.com/~jdoe"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.IRIs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IRIs() = %v, want %v", got, tt.want)
			}
		})
	}
}

var items = ItemCollection{}

func TestItemCollection_Equal(t *testing.T) {
	tests := []struct {
		name string
		i    ItemCollection
		with ItemCollection
		want bool
	}{
		{
			name: "empty",
			want: true,
		},
		{
			name: "empty collections",
			i:    ItemCollection{},
			with: ItemCollection{},
			want: true,
		},
		{
			name: "nil collection equal to empty collection",
			i:    nil,
			with: ItemCollection{},
			want: true,
		},
		{
			name: "empty collection equal to nil collection",
			i:    nil,
			with: ItemCollection{},
			want: true,
		},
		{
			name: "IRIs collections",
			i:    (&IRIs{"http://example.com", "http://social.example.com"}).Collection(),
			with: (&IRIs{"http://example.com", "http://social.example.com"}).Collection(),
			want: true,
		},
		{
			name: "Item Collection with same IRIs",
			i:    ItemCollection{IRI("http://example.com"), IRI("http://social.example.com")},
			with: ItemCollection{IRI("http://example.com"), IRI("http://social.example.com")},
			want: true,
		},
		{
			name: "Item Collection with same IRIs in different order",
			i:    ItemCollection{IRI("http://social.example.com"), IRI("http://example.com")},
			with: ItemCollection{IRI("http://example.com"), IRI("http://social.example.com")},
			want: false,
		},
		{
			name: "Item Collection with Object and IRI with same IRI",
			i:    ItemCollection{IRI("http://example.com"), Object{ID: "http://example.com"}},
			with: ItemCollection{IRI("http://example.com"), Object{ID: "http://example.com"}},
			want: true,
		},
		{
			name: "Item Collection with Object and IRI with same IRI in different order",
			i:    ItemCollection{Object{ID: "http://example.com"}, IRI("http://example.com")},
			with: ItemCollection{IRI("http://example.com"), Object{ID: "http://example.com"}},
			want: false,
		},
		{
			name: "Item Collection with Link and IRI with same IRI",
			i:    ItemCollection{IRI("http://example.com"), Link{ID: "http://example.com"}},
			with: ItemCollection{IRI("http://example.com"), Link{ID: "http://example.com"}},
			want: true,
		},
		{
			name: "Item Collection with Link and IRI with same IRI in different order",
			i:    ItemCollection{Link{ID: "http://example.com"}, IRI("http://example.com")},
			with: ItemCollection{IRI("http://example.com"), Link{ID: "http://example.com"}},
			want: false,
		},
		{
			name: "Item Collection with Link and IRI with same IRI",
			i:    ItemCollection{Object{ID: "http://example.com"}, IRI("http://example.com")},
			with: ItemCollection{IRI("http://example.com"), Link{ID: "http://example.com"}},
			want: false,
		},
		{
			name: "different sizes",
			i:    ItemCollection{IRI("http://example.com")},
			with: ItemCollection{IRI("http://example.com"), IRI("http://social.example.com")},
			want: false,
		},
		{
			name: "plausible objects",
			i:    items,
			with: items,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Equal(tt.with); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
			// NOTE(marius): check symmetry of the function
			if got := tt.with.Equal(tt.i); got != tt.want {
				t.Errorf("Revered Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleOnItemCollection() {
	// When operating on a known ItemCollection slice, we can use a direct
	// approach for accessing and modifying it, instead of relying on the
	// slice iteration functionality provided by the other OnXXX() functions.
	//
	// This allows more flexibility regarding the slice itself, as opposed to
	// its elements.

	many := ItemCollection{
		&Place{ID: "http://example.com/home"},
		&Tombstone{ID: "http://example.com/missing"},
	}

	// In a way that's perhaps confusing compared to the behaviour
	// of the other OnXXX() functions that allow for side effects
	// outside the closure based on whether the Item element is a
	// pointer value or not, this function's side effects are dependent
	// on the fact that a pointer value is passed as its first argument.
	_ = OnItemCollection(many, func(col *ItemCollection) error {
		// Here the appended Item will not be present outside
		// this closure.
		*col = append(*col, &Question{Type: QuestionType})
		return nil
	})
	fmt.Printf("Many: %v\n", many)

	_ = OnItemCollection(&many, func(col *ItemCollection) error {
		// But here it will.
		*col = append(*col, &Question{Type: QuestionType})
		return nil
	})
	fmt.Printf("    : %v\n\n", many)

	others := ItemCollection{
		&Person{ID: "http://example.com/~jdoe", Type: PersonType},
		&Group{ID: "http://example.com/the-samples", Type: GroupType},
	}
	_ = OnItemCollection(&others, func(col *ItemCollection) error {
		for i, it := range *col {
			_ = OnActor(it, func(act *Actor) error {
				act.Name = DefaultLangValue("Actor #" + strconv.Itoa(i))
				return nil
			})
		}
		return nil
	})
	fmt.Printf("Others: %v\n", others)

	// Output:
	// Many: [activitypub.Place { id: http://example.com/home } activitypub.Tombstone { id: http://example.com/missing }]
	//     : [activitypub.Place { id: http://example.com/home } activitypub.Tombstone { id: http://example.com/missing } activitypub.Question[Question] {  }]
	//
	// Others: [activitypub.Actor[Person] { id: http://example.com/~jdoe, name: Actor #0 } activitypub.Actor[Group] { id: http://example.com/the-samples, name: Actor #1 }]
}

func TestItemCollection_Normalize(t *testing.T) {
	tests := []struct {
		name string
		i    ItemCollection
		want Item
	}{
		{
			name: "empty",
			i:    nil,
			want: nil,
		},
		{
			name: "single IRI",
			i:    ItemCollection{IRI("http://example.com")},
			want: IRI("http://example.com"),
		},
		{
			name: "multiple IRIs",
			i:    ItemCollection{IRI("http://example.com/1"), IRI("http://example.com/2")},
			want: ItemCollection{IRI("http://example.com/1"), IRI("http://example.com/2")},
		},
		{
			name: "single Object",
			i:    ItemCollection{Object{ID: "http://example.com"}},
			want: Object{ID: "http://example.com"},
		},
		{
			name: "multiple Objects",
			i:    ItemCollection{Activity{ID: "http://example.com/a1"}, Place{ID: "http://example.com/home"}},
			want: ItemCollection{Activity{ID: "http://example.com/a1"}, Place{ID: "http://example.com/home"}},
		},
		{
			name: "mixed Objects and IRIs",
			i:    ItemCollection{Activity{ID: "http://example.com/a1"}, IRI("http://example.com/~jdoe"), Place{ID: "http://example.com/home"}, IRI("http://example.com")},
			want: ItemCollection{Activity{ID: "http://example.com/a1"}, IRI("http://example.com/~jdoe"), Place{ID: "http://example.com/home"}, IRI("http://example.com")},
		},
		{
			name: "same ID objects",
			i:    ItemCollection{Activity{ID: "http://example.com/a1"}, Object{ID: "http://example.com/a1"}},
			want: ItemCollection{Activity{ID: "http://example.com/a1"}, Object{ID: "http://example.com/a1"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Normalize(); !cmp.Equal(got, tt.want) {
				t.Errorf("Normalize() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestDerefItem(t *testing.T) {
	tests := []struct {
		name string
		arg  Item
		want ItemCollection
	}{
		{
			name: "empty",
		},
		{
			name: "simple object",
			arg:  &Object{ID: "https://example.com"},
			want: ItemCollection{&Object{ID: "https://example.com"}},
		},
		{
			name: "simple IRI",
			arg:  IRI("https://example.com"),
			want: ItemCollection{IRI("https://example.com")},
		},
		{
			name: "IRI collection",
			arg:  IRIs{IRI("https://example.com"), IRI("https://example.com/~jdoe")},
			want: ItemCollection{IRI("https://example.com"), IRI("https://example.com/~jdoe")},
		},
		{
			name: "Item collection",
			arg: ItemCollection{
				&Object{ID: "https://example.com"},
				&Actor{ID: "https://example.com/~jdoe"},
			},
			want: ItemCollection{
				&Object{ID: "https://example.com"},
				&Actor{ID: "https://example.com/~jdoe"},
			},
		},
		{
			name: "mixed item collection",
			arg: ItemCollection{
				&Object{ID: "https://example.com"},
				IRI("https://example.com/666"),
				&Actor{ID: "https://example.com/~jdoe"},
			},
			want: ItemCollection{
				&Object{ID: "https://example.com"},
				IRI("https://example.com/666"),
				&Actor{ID: "https://example.com/~jdoe"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DerefItem(tt.arg); !cmp.Equal(got, tt.want) {
				t.Errorf("DerefItem() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestFlattenItemCollection(t *testing.T) {
	tests := []struct {
		name string
		col  ItemCollection
		want ItemCollection
	}{
		{
			name: "nil",
			col:  nil,
			want: nil,
		},
		{
			name: "empty",
			col:  ItemCollection{},
			want: nil,
		},
		{
			name: "with iri",
			col:  ItemCollection{IRI("http://example.com")},
			want: ItemCollection{IRI("http://example.com")},
		},
		{
			name: "with object",
			col:  ItemCollection{&Object{ID: "http://example.com"}},
			want: ItemCollection{IRI("http://example.com")},
		},
		{
			name: "with objects",
			col:  ItemCollection{&Object{ID: "http://example.com"}, &Actor{ID: "http://example.com/~jdoe"}},
			want: ItemCollection{IRI("http://example.com"), IRI("http://example.com/~jdoe")},
		},
		{
			name: "with duplicates",
			col:  ItemCollection{&Object{ID: "http://example.com"}, &Actor{ID: "http://example.com/~jdoe"}, IRI("http://example.com")},
			want: ItemCollection{IRI("http://example.com"), IRI("http://example.com/~jdoe")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenItemCollection(tt.col); !cmp.Equal(got, tt.want) {
				t.Errorf("FlattenItemCollection() = %s", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestItemCollection_UnmarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    ItemCollection
		wantErr error
	}{
		{
			name:    "nil",
			data:    nil,
			wantErr: errors.NotImplementedf("Binary functionality not implemented"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := ItemCollection{}
			if err := l.UnmarshalBinary(tt.data); !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("UnmarshalBinary() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
			}
			if !cmp.Equal(l, tt.want) {
				t.Errorf("UnmarshalBinary() got = %s", cmp.Diff(tt.want, l))
			}
		})
	}
}

func TestItemCollection_MarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		sub     ItemCollection
		want    []byte
		wantErr error
	}{
		{
			name:    "nil",
			sub:     ItemCollection{},
			want:    nil,
			wantErr: errors.NotImplementedf("Binary functionality not implemented"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sub.MarshalBinary()
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("MarshalBinary() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("MarshalBinary() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestItemCollection_Match(t *testing.T) {
	tests := []struct {
		name string
		i    ItemCollection
		tt   ActivityVocabularyTypes
		want bool
	}{
		{
			name: "nil",
			i:    nil,
			tt:   nil,
			want: false,
		},
		{
			name: "empty",
			i:    ItemCollection{},
			tt:   nil,
			want: false,
		},
		{
			name: "match-collectionOfItems",
			i:    ItemCollection{},
			tt:   ActivityVocabularyTypes{CollectionOfItems},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Match(tt.tt...); got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemCollection_ItemsMatch(t *testing.T) {
	tests := []struct {
		name    string
		i       ItemCollection
		toCheck []Item
		want    bool
	}{
		{
			name:    "nil match nil",
			i:       nil,
			toCheck: nil,
			want:    true,
		},
		{
			name:    "empty match nil",
			i:       ItemCollection{},
			toCheck: nil,
			want:    true,
		},
		{
			name:    "not-empty match nil toCheck",
			i:       ItemCollection{IRI("http://example.com")},
			toCheck: nil,
			want:    true,
		},
		{
			name:    "item-collection, no match",
			i:       ItemCollection{IRI("http://example.com")},
			toCheck: ItemCollection{Object{}},
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.ItemsMatch(tt.toCheck...); got != tt.want {
				t.Errorf("ItemsMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemCollection_Equals(t *testing.T) {
	tests := []struct {
		name string
		i    ItemCollection
		with Item
		want bool
	}{
		{
			name: "nil",
			i:    nil,
			with: nil,
			want: false,
		},
		{
			name: "collection not equals non-collection",
			i:    ItemCollection{},
			with: &Object{},
			want: false,
		},
		{
			name: "empty equals empty",
			i:    ItemCollection{},
			with: ItemCollection{},
			want: true,
		},
		{
			name: "empty not equals not empty",
			i:    ItemCollection{},
			with: ItemCollection{nil},
			want: false,
		},
		{
			name: "equal with one IRI",
			i:    ItemCollection{IRI("http://example.com")},
			with: ItemCollection{IRI("http://example.com")},
			want: true,
		},
		{
			name: "equal with empty object",
			i:    ItemCollection{&Object{}},
			with: ItemCollection{&Object{}},
			want: true,
		},
		{
			name: "not equal with different counts",
			i:    ItemCollection{&Object{}, IRI("http://example.com")},
			with: ItemCollection{&Object{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Equals(tt.with); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemCollection_Recipients(t *testing.T) {
	tests := []struct {
		name string
		i    ItemCollection
		want ItemCollection
	}{
		{
			name: "nil",
			i:    nil,
			want: nil,
		},
		{
			name: "empty",
			i:    ItemCollection{},
			want: nil,
		},
		{
			name: "one object with Audience",
			i:    ItemCollection{&Object{Audience: ItemCollection{PublicNS}}},
			want: ItemCollection{PublicNS},
		},
		{
			name: "one object with To",
			i:    ItemCollection{&Object{To: ItemCollection{IRI("http://example.com")}}},
			want: ItemCollection{IRI("http://example.com")},
		},
		{
			name: "one object with CC",
			i:    ItemCollection{&Object{CC: ItemCollection{IRI("http://example.com")}}},
			want: ItemCollection{IRI("http://example.com")},
		},
		{
			name: "one object with Bto",
			i:    ItemCollection{&Object{Bto: ItemCollection{IRI("http://example.com"), &Object{ID: "http://example.com/ob"}}}},
			want: ItemCollection{IRI("http://example.com"), IRI("http://example.com/ob")},
		},
		{
			name: "one object with BCC",
			i:    ItemCollection{&Object{BCC: ItemCollection{IRI("http://example.com"), &Actor{ID: "http://example.com/~jdoe"}}}},
			want: ItemCollection{IRI("http://example.com"), IRI("http://example.com/~jdoe")},
		},
		//
		{
			name: "two objects",
			i: ItemCollection{
				&Object{
					Audience: ItemCollection{PublicNS},
				},
				&Actor{
					To: ItemCollection{PublicNS},
				},
				&Actor{
					Bto: ItemCollection{IRI("http://example.com/followers")},
				},
				&Object{
					Bto: ItemCollection{
						&Collection{ID: "http://example.com/hidden-to"},
					},
				},
				&Object{
					BCC: ItemCollection{
						IRI("http://example.com"),
						&Actor{ID: "http://example.com/~jdoe"},
					},
				},
				Link{ID: "http://example.com/no-recipients"},
			},
			want: ItemCollection{PublicNS, IRI("http://example.com/followers"), IRI("http://example.com/hidden-to"), IRI("http://example.com"), IRI("http://example.com/~jdoe")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Recipients(); !cmp.Equal(got, tt.want) {
				t.Errorf("Recipients() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
