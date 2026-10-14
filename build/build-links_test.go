package build

import (
	"testing"

	vocab "github.com/go-ap/activitypub"
	"github.com/google/go-cmp/cmp"
)

func TestLink(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Link
	}{
		{
			name: "empty",
			want: &vocab.Link{},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Link{ID: "http://example.com"},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Link{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
		{
			name: "WithName",
			args: []InitFn{WithName("Lorem ipsum")},
			want: &vocab.Link{Name: vocab.NaturalLanguageValues{vocab.DefaultLang: vocab.Content("Lorem ipsum")}},
		},
		{
			name: "WithPreview",
			args: []InitFn{WithPreview(&vocab.Object{ID: "http://example.com/preview"})},
			want: &vocab.Link{Preview: &vocab.Object{ID: "http://example.com/preview"}},
		},
		{
			name: "WithMediaType",
			args: []InitFn{WithMediaType("test/test")},
			want: &vocab.Link{MediaType: "test/test"},
		},
		{
			name: "WithHref",
			args: []InitFn{WithHref("http://example.com/href")},
			want: &vocab.Link{Href: "http://example.com/href"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Link(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Link() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestMention(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Mention
	}{
		{
			name: "empty",
			want: &vocab.Mention{Type: vocab.MentionType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Mention{ID: "http://example.com", Type: vocab.MentionType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Mention{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mention(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Mention() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
