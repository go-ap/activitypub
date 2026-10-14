package build

import (
	"testing"

	vocab "github.com/go-ap/activitypub"
	"github.com/google/go-cmp/cmp"
)

func TestActor(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Actor
	}{
		{
			name: "empty",
			want: &vocab.Actor{},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Actor{ID: "http://example.com"},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Actor{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
		{
			name: "WithPreferredUsername",
			args: []InitFn{WithPreferredUsername("Jane Doe")},
			want: &vocab.Actor{PreferredUsername: vocab.NaturalLanguageValues{vocab.DefaultLang: vocab.Content("Jane Doe")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Actor(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Actor() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
