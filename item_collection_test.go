package activitypub

import "testing"

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
		name string
		i    ItemCollection
		arg  Item
	}{
		{
			name: "empty_collection_nil_item",
			i:    ItemCollection{},
			arg:  nil,
		},
		{
			name: "empty_collection_non_nil_item",
			i:    ItemCollection{},
			arg:  &Object{},
		},
		{
			name: "non_empty_collection_nil_item",
			i: ItemCollection{
				&Object{ID: "test"},
			},
			arg: nil,
		},
		{
			name: "non_empty_collection_non_contained_item_empty_ID",
			i: ItemCollection{
				&Object{ID: "test"},
			},
			arg: &Object{},
		},
		{
			name: "non_empty_collection_non_contained_item",
			i: ItemCollection{
				&Object{ID: "test"},
			},
			arg: &Object{ID: "test123"},
		},
		{
			name: "non_empty_collection_just_contained_item",
			i: ItemCollection{
				&Object{ID: "test"},
			},
			arg: &Object{ID: "test"},
		},
		{
			name: "non_empty_collection_contained_item_first_pos",
			i: ItemCollection{
				&Object{ID: "test"},
				&Object{ID: "test123"},
			},
			arg: &Object{ID: "test"},
		},
		{
			name: "non_empty_collection_contained_item_not_first_pos",
			i: ItemCollection{
				&Object{ID: "test123"},
				&Object{ID: "test"},
				&Object{ID: "test321"},
			},
			arg: &Object{ID: "test"},
		},
		{
			name: "non_empty_collection_contained_item_last_pos",
			i: ItemCollection{
				&Object{ID: "test123"},
				&Object{ID: "test"},
			},
			arg: &Object{ID: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			origContains := tt.i.Contains(tt.arg)
			origLen := tt.i.Count()
			should := ""
			does := "n't"
			if origContains {
				should = "n't"
				does = ""
			}

			tt.i.Remove(tt.arg)
			if tt.i.Contains(tt.arg) {
				t.Errorf("%T should%s contain %T, but it does%s: %#v", tt.i, should, tt.arg, does, tt.i)
			}
			if origContains {
				if tt.i.Count() > origLen-1 {
					t.Errorf("%T should have a count lower than %d, got %d", tt.i, origLen, tt.i.Count())
				}
			} else {
				if tt.i.Count() != origLen {
					t.Errorf("%T should have a count equal to %d, got %d", tt.i, origLen, tt.i.Count())
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
