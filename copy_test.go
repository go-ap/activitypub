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
			it:   &Object{Type: NoteType.ToTypes()},
			want: &Object{Type: NoteType.ToTypes()},
		},
		{
			name: "object with Type",
			it:   Object{Type: NoteType.ToTypes()},
			want: &Object{Type: NoteType.ToTypes()},
		},
		{
			name: "*collection",
			it:   &Collection{Type: CollectionType.ToTypes()},
			want: &Collection{Type: CollectionType.ToTypes()},
		},
		{
			name: "collection",
			it:   Collection{Type: CollectionType.ToTypes()},
			want: &Collection{Type: CollectionType.ToTypes()},
		},
		{
			name: "*collection with items",
			it:   &Collection{Type: CollectionType.ToTypes(), Items: ItemCollection{&Object{ID: "http://example.com"}, &Place{ID: "http//:example.com/home"}}},
			want: &Collection{Type: CollectionType.ToTypes(), Items: ItemCollection{&Object{ID: "http://example.com"}, &Place{ID: "http//:example.com/home"}}},
		},
		{
			name: "collection with items",
			it:   Collection{Type: CollectionType.ToTypes(), Items: ItemCollection{&Object{ID: "http://example.com"}, &Place{ID: "http//:example.com/home"}}},
			want: &Collection{Type: CollectionType.ToTypes(), Items: ItemCollection{&Object{ID: "http://example.com"}, &Place{ID: "http//:example.com/home"}}},
		},
		{
			name: "empty item collection",
			it:   ItemCollection{},
			want: &ItemCollection{},
		},
		{
			name: "empty *item collection",
			it:   &ItemCollection{},
			want: &ItemCollection{},
		},
		{
			name: "not empty item collection",
			it:   ItemCollection{&Object{ID: "http://example.com"}, &Place{ID: "http//:example.com/home"}},
			want: &ItemCollection{&Object{ID: "http://example.com"}, &Place{ID: "http//:example.com/home"}},
		},
		{
			name: "not empty *item collection",
			it:   &ItemCollection{&Actor{ID: "http://example.com/~jdoe"}, Link{ID: "https://example.com/test"}},
			want: &ItemCollection{&Actor{ID: "http://example.com/~jdoe"}, Link{ID: "https://example.com/test"}},
		},
		{
			name: "empty IRIs",
			it:   IRIs{},
			want: &IRIs{},
		},
		{
			name: "empty *IRIs",
			it:   &IRIs{},
			want: &IRIs{},
		},
		{
			name: "not empty IRIs",
			it:   IRIs{"http://example.com", "http//:example.com/home"},
			want: &IRIs{"http://example.com", "http//:example.com/home"},
		},
		{
			name: "not empty *IRIs",
			it:   &IRIs{"http://example.com/~jdoe", "https://example.com/test"},
			want: &IRIs{"http://example.com/~jdoe", "https://example.com/test"},
		},
		{
			name: "empty IRI",
			it:   IRI(""),
			want: IRI(""),
		},
		{
			name: "not empty IRI",
			it:   IRI("http://example.com"),
			want: IRI("http://example.com"),
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
