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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OrderedCollectionPage(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("OrderedCollectionPage() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
