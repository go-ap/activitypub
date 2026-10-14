package build

import (
	"testing"

	vocab "github.com/go-ap/activitypub"
	"github.com/google/go-cmp/cmp"
)

func TestActivity(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Activity
	}{
		{
			name: "empty",
			want: &vocab.Activity{},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Activity{ID: "http://example.com"},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Activity{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Activity(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Activity() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestIntransitiveActivity(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.IntransitiveActivity
	}{
		{
			name: "empty",
			want: &vocab.IntransitiveActivity{},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.IntransitiveActivity{ID: "http://example.com"},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.IntransitiveActivity{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntransitiveActivity(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("IntransitiveActivity() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestQuestion(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Question
	}{
		{
			name: "empty",
			want: &vocab.Question{Type: vocab.QuestionType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Question{ID: "http://example.com", Type: vocab.QuestionType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Question{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Question(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Question() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
