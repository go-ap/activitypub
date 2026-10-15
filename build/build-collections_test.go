package build

import (
	"testing"

	vocab "github.com/go-ap/activitypub"
	"github.com/google/go-cmp/cmp"
)

func TestCollection(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Collection
	}{
		{
			name: "empty",
			want: &vocab.Collection{Type: vocab.CollectionType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Collection{ID: "http://example.com", Type: vocab.CollectionType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Collection{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
		{
			name: "WithFirst",
			args: []InitFn{WithFirst(vocab.IRI("http://example.com/first"))},
			want: &vocab.Collection{Type: vocab.CollectionType, First: vocab.IRI("http://example.com/first")},
		},
		{
			name: "WithLast",
			args: []InitFn{WithLast(vocab.IRI("http://example.com/last"))},
			want: &vocab.Collection{Type: vocab.CollectionType, Last: vocab.IRI("http://example.com/last")},
		},
		{
			name: "WithTotalItems",
			args: []InitFn{WithTotalItems(23456)},
			want: &vocab.Collection{Type: vocab.CollectionType, TotalItems: 23456},
		},
		{
			name: "WithItems",
			args: []InitFn{WithItems(vocab.IRI("http://example.com/1"), vocab.IRI("http://example.com/2"), vocab.IRI("http://example.com/3"))},
			want: &vocab.Collection{Type: vocab.CollectionType, Items: vocab.ItemCollection{vocab.IRI("http://example.com/1"), vocab.IRI("http://example.com/2"), vocab.IRI("http://example.com/3")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Collection(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Collection() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestCollectionPage(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.CollectionPage
	}{
		{
			name: "empty",
			want: &vocab.CollectionPage{Type: vocab.CollectionPageType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.CollectionPage{ID: "http://example.com", Type: vocab.CollectionPageType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.CollectionPage{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
		{
			name: "WithTotalItems",
			args: []InitFn{WithTotalItems(23456)},
			want: &vocab.CollectionPage{Type: vocab.CollectionPageType, TotalItems: 23456},
		},

		{
			name: "WithPrev",
			args: []InitFn{WithPrev(vocab.IRI("http://example.com/prev"))},
			want: &vocab.CollectionPage{Type: vocab.CollectionPageType, Prev: vocab.IRI("http://example.com/prev")},
		},
		{
			name: "WithNext",
			args: []InitFn{WithNext(vocab.IRI("http://example.com/next"))},
			want: &vocab.CollectionPage{Type: vocab.CollectionPageType, Next: vocab.IRI("http://example.com/next")},
		},
		{
			name: "WithCurrent",
			args: []InitFn{WithCurrent(vocab.IRI("http://example.com/cur"))},
			want: &vocab.CollectionPage{Type: vocab.CollectionPageType, Current: vocab.IRI("http://example.com/cur")},
		},
		{
			name: "WithPartOf",
			args: []InitFn{WithPartOf(vocab.IRI("http://example.com"))},
			want: &vocab.CollectionPage{Type: vocab.CollectionPageType, PartOf: vocab.IRI("http://example.com")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CollectionPage(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("CollectionPage() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestOrderedCollection(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.OrderedCollection
	}{
		{
			name: "empty",
			want: &vocab.OrderedCollection{Type: vocab.OrderedCollectionType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.OrderedCollection{ID: "http://example.com", Type: vocab.OrderedCollectionType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.OrderedCollection{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
		{
			name: "WithFirst",
			args: []InitFn{WithFirst(vocab.IRI("http://example.com/first"))},
			want: &vocab.OrderedCollection{Type: vocab.OrderedCollectionType, First: vocab.IRI("http://example.com/first")},
		},
		{
			name: "WithLast",
			args: []InitFn{WithLast(vocab.IRI("http://example.com/last"))},
			want: &vocab.OrderedCollection{Type: vocab.OrderedCollectionType, Last: vocab.IRI("http://example.com/last")},
		},
		{
			name: "WithTotalItems",
			args: []InitFn{WithTotalItems(23456)},
			want: &vocab.OrderedCollection{Type: vocab.OrderedCollectionType, TotalItems: 23456},
		},
		{
			name: "WithOrderedItems",
			args: []InitFn{WithOrderedItems(vocab.IRI("http://example.com/1"), vocab.IRI("http://example.com/2"), vocab.IRI("http://example.com/3"))},
			want: &vocab.OrderedCollection{Type: vocab.OrderedCollectionType, OrderedItems: vocab.ItemCollection{vocab.IRI("http://example.com/1"), vocab.IRI("http://example.com/2"), vocab.IRI("http://example.com/3")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrderedCollection(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("OrderedCollection() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestOrderedCollectionPage(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.OrderedCollectionPage
	}{
		{
			name: "empty",
			want: &vocab.OrderedCollectionPage{Type: vocab.OrderedCollectionPageType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.OrderedCollectionPage{ID: "http://example.com", Type: vocab.OrderedCollectionPageType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.OrderedCollectionPage{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
		{
			name: "WithTotalItems",
			args: []InitFn{WithTotalItems(23456)},
			want: &vocab.OrderedCollectionPage{Type: vocab.OrderedCollectionPageType, TotalItems: 23456},
		},
		{
			name: "WithOrderedItems",
			args: []InitFn{WithOrderedItems(vocab.IRI("http://example.com/1"), vocab.IRI("http://example.com/2"), vocab.IRI("http://example.com/3"))},
			want: &vocab.OrderedCollectionPage{Type: vocab.OrderedCollectionPageType, OrderedItems: vocab.ItemCollection{vocab.IRI("http://example.com/1"), vocab.IRI("http://example.com/2"), vocab.IRI("http://example.com/3")}},
		},
		{
			name: "WithPartOf",
			args: []InitFn{WithPartOf(vocab.IRI("http://example.com"))},
			want: &vocab.OrderedCollectionPage{Type: vocab.OrderedCollectionPageType, PartOf: vocab.IRI("http://example.com")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrderedCollectionPage(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("OrderedCollectionPage() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
