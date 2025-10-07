package activitypub

import (
	"reflect"
	"testing"
)

func TestItemCollection_Append(t *testing.T) {
	t.Skipf("TODO")
}

func TestItemCollection_Collection(t *testing.T) {
	t.Skipf("TODO")
}

func TestItemCollection_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestItemCollection_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestItemCollection_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestItemCollection_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestItemCollection_IsObject(t *testing.T) {
	t.Skipf("TODO")
}

func TestItemCollection_First(t *testing.T) {
	t.Skipf("TODO")
}

func TestItemCollection_Count(t *testing.T) {
	t.Skipf("TODO")
}

func TestItemCollection_Contains(t *testing.T) {
	t.Skipf("TODO")
}

func TestItemCollection_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestToItemCollection(t *testing.T) {
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
				IRI("https://example.dev/a801139a-0d9a-4703-b0a5-9d14ae1438e4"),
				IRI("https://example.dev"),
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
			if got := ItemCollectionDeduplication(tt.args...); !tt.want.Equals(got) {
				t.Errorf("ItemCollectionDeduplication() = %v, want %v", got, tt.want)
			}
			if len(tt.remaining) != len(tt.args) {
				t.Errorf("ItemCollectionDeduplication() arguments count %d, want %d", len(tt.args), len(tt.remaining))
			}
			for i, remArg := range tt.remaining {
				arg := tt.args[i]
				if !remArg.Equals(arg) {
					t.Errorf("ItemCollectionDeduplication() argument at pos %d = %v, want %v", i, arg, remArg)
				}
			}
		})
	}
}

func TestToItemCollection1(t *testing.T) {
	tests := []struct {
		name    string
		it      Item
		want    *ItemCollection
		wantErr bool
	}{
		{
			name: "empty",
		},
		{
			name:    "IRIs to ItemCollection",
			it:      IRIs{"https://example.com", "https://example.com/example"},
			want:    &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
			wantErr: false,
		},
		{
			name:    "ItemCollection to ItemCollection",
			it:      ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
			want:    &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
			wantErr: false,
		},
		{
			name:    "*ItemCollection to ItemCollection",
			it:      &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
			want:    &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
			wantErr: false,
		},
		{
			name:    "Collection to ItemCollection",
			it:      &Collection{Items: ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")}},
			want:    &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
			wantErr: false,
		},
		{
			name:    "CollectionPage to ItemCollection",
			it:      &CollectionPage{Items: ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")}},
			want:    &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
			wantErr: false,
		},
		{
			name:    "OrderedCollection to ItemCollection",
			it:      &OrderedCollection{OrderedItems: ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")}},
			want:    &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
			wantErr: false,
		},
		{
			name:    "OrderedCollectionPage to ItemOrderedCollection",
			it:      &OrderedCollectionPage{OrderedItems: ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")}},
			want:    &ItemCollection{IRI("https://example.com"), IRI("https://example.com/example")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToItemCollection(tt.it)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToItemCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToItemCollection() got = %v, want %v", got, tt.want)
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
