package activitypub

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestItemCollection_Append(t *testing.T) {
	d := make(ItemCollection, 0)

	val := Object{ID: ID("grrr")}

	_ = d.Append(val)

	if len(d) != 1 {
		t.Errorf("Objects array should have exactly an element")
	}
	if !reflect.DeepEqual(d[0], val) {
		t.Errorf("First item in object array does not match %q", val.ID)
	}
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
