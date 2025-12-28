package activitypub

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClone(t *testing.T) {
	tests := []struct {
		name string
		it   Item
		want Item
	}{
		{
			name: "empty",
		},
		{
			name: "*object empty",
			it:   &Object{},
			want: &Object{},
		},
		{
			name: "object empty",
			it:   Object{},
			want: &Object{},
		},
		{
			name: "*object with ID",
			it:   &Object{ID: "http://example.com"},
			want: &Object{ID: "http://example.com"},
		},
		{
			name: "object with ID",
			it:   Object{ID: "http://example.com"},
			want: &Object{ID: "http://example.com"},
		},
		{
			name: "*object with Type",
			it:   &Object{Type: NoteType},
			want: &Object{Type: NoteType},
		},
		{
			name: "object with Type",
			it:   Object{Type: NoteType},
			want: &Object{Type: NoteType},
		},
		{
			name: "*collection",
			it:   &Collection{Type: CollectionType},
			want: &Collection{Type: CollectionType},
		},
		{
			name: "collection",
			it:   Collection{Type: CollectionType},
			want: &Collection{Type: CollectionType},
		},
		{
			name: "*collection with items",
			it:   &Collection{Type: CollectionType, Items: ItemCollection{&Object{ID: "http://example.com"}, &Place{ID: "http//:example.com/home"}}},
			want: &Collection{Type: CollectionType, Items: ItemCollection{&Object{ID: "http://example.com"}, &Place{ID: "http//:example.com/home"}}},
		},
		{
			name: "collection with items",
			it:   Collection{Type: CollectionType, Items: ItemCollection{&Object{ID: "http://example.com"}, &Place{ID: "http//:example.com/home"}}},
			want: &Collection{Type: CollectionType, Items: ItemCollection{&Object{ID: "http://example.com"}, &Place{ID: "http//:example.com/home"}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Clone(tt.it)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Clone() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
