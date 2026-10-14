package build

import (
	"testing"
	"time"

	vocab "github.com/go-ap/activitypub"
	"github.com/google/go-cmp/cmp"
)

func TestObject(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Object
	}{
		{
			name: "empty",
			want: &vocab.Object{},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Object{ID: "http://example.com"},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Object{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
		{
			name: "WithName",
			args: []InitFn{WithName("Lorem ipsum")},
			want: &vocab.Object{Name: vocab.NaturalLanguageValues{vocab.DefaultLang: vocab.Content("Lorem ipsum")}},
		},
		{
			name: "WithSummary",
			args: []InitFn{WithSummary("Lorem ipsum")},
			want: &vocab.Object{Summary: vocab.NaturalLanguageValues{vocab.DefaultLang: vocab.Content("Lorem ipsum")}},
		},
		{
			name: "WithContent",
			args: []InitFn{WithContent("Lorem ipsum")},
			want: &vocab.Object{Content: vocab.NaturalLanguageValues{vocab.DefaultLang: vocab.Content("Lorem ipsum")}},
		},
		{
			name: "WithSourceContent",
			args: []InitFn{WithSourceContent("Lorem ipsum")},
			want: &vocab.Object{
				Source: vocab.Source{
					Content: vocab.NaturalLanguageValues{vocab.DefaultLang: vocab.Content("Lorem ipsum")},
				},
			},
		},
		{
			name: "WithSourceMediaType",
			args: []InitFn{WithSourceMediaType("test/test")},
			want: &vocab.Object{
				Source: vocab.Source{
					MediaType: vocab.MimeType("test/test"),
				},
			},
		},
		{
			name: "WithSource",
			args: []InitFn{WithSource(
				vocab.Source{
					Content:   vocab.NaturalLanguageValues{vocab.DefaultLang: vocab.Content("Lorem ipsum")},
					MediaType: "test/test",
				},
			)},
			want: &vocab.Object{
				Source: vocab.Source{
					Content:   vocab.NaturalLanguageValues{vocab.DefaultLang: vocab.Content("Lorem ipsum")},
					MediaType: vocab.MimeType("test/test"),
				},
			},
		},
		{
			name: "WithReplies",
			args: []InitFn{WithID("http://example.com/"), WithShares()},
			want: &vocab.Object{ID: "http://example.com/", Shares: vocab.Shares.IRI(vocab.IRI("http://example.com"))},
		},
		{
			name: "WithShares",
			args: []InitFn{WithShares(vocab.IRI("http://example.com/shares"))},
			want: &vocab.Object{Shares: vocab.IRI("http://example.com/shares")},
		},
		{
			name: "WithLikes",
			args: []InitFn{WithLikes(vocab.IRI("http://example.com/likes"))},
			want: &vocab.Object{Likes: vocab.IRI("http://example.com/likes")},
		},
		{
			name: "WithPublished",
			args: []InitFn{WithPublished("2001-01-01T06:06:06.123Z")},
			want: &vocab.Object{Published: time.Date(2001, time.January, 1, 6, 6, 6, 123000000, time.UTC)},
		},
		{
			name: "WithUpdated",
			args: []InitFn{WithUpdated("2001-01-01T06:06:06.123Z")},
			want: &vocab.Object{Updated: time.Date(2001, time.January, 1, 6, 6, 6, 123000000, time.UTC)},
		},
		{
			name: "WithTo",
			args: []InitFn{WithTo(vocab.IRI("http://example.com/to"))},
			want: &vocab.Object{To: vocab.ItemCollection{vocab.IRI("http://example.com/to")}},
		},
		{
			name: "WithCC",
			args: []InitFn{WithCC(&vocab.Object{ID: "http://example.com/cc-ob"})},
			want: &vocab.Object{CC: vocab.ItemCollection{&vocab.Object{ID: "http://example.com/cc-ob"}}},
		},
		{
			name: "WithBto",
			args: []InitFn{WithBto(vocab.IRI("http://exaple.com/bto"), &vocab.Object{ID: "http://example.com/bto-ob"})},
			want: &vocab.Object{Bto: vocab.ItemCollection{vocab.IRI("http://exaple.com/bto"), &vocab.Object{ID: "http://example.com/bto-ob"}}},
		},
		{
			name: "WithBCC",
			args: []InitFn{WithBCC(vocab.IRI("http://exaple.com/bto"), vocab.PublicNS)},
			want: &vocab.Object{BCC: vocab.ItemCollection{vocab.IRI("http://exaple.com/bto"), vocab.PublicNS}},
		},
		{
			name: "WithTag",
			args: []InitFn{WithTag(vocab.IRI("http://exaple.com/tag1"), vocab.IRI("http://exaple.com/tag2"))},
			want: &vocab.Object{Tag: vocab.ItemCollection{vocab.IRI("http://exaple.com/tag1"), vocab.IRI("http://exaple.com/tag2")}},
		},
		{
			name: "WithAttachment",
			args: []InitFn{WithAttachment(vocab.IRI("http://exaple.com/attachment"))},
			want: &vocab.Object{Attachment: vocab.IRI("http://exaple.com/attachment")},
		},
		{
			name: "WithAttributedTo",
			args: []InitFn{WithAttributedTo(vocab.IRI("http://exaple.com/~jdoe"), vocab.IRI("http://example.com/~alice"))},
			want: &vocab.Object{AttributedTo: vocab.ItemCollection{vocab.IRI("http://exaple.com/~jdoe"), vocab.IRI("http://example.com/~alice")}},
		},
		{
			name: "WithInReplyTo",
			args: []InitFn{WithInReplyTo(vocab.IRI("http://exaple.com/ob-1"), vocab.IRI("http://example.com/ob-2"))},
			want: &vocab.Object{InReplyTo: vocab.ItemCollection{vocab.IRI("http://exaple.com/ob-1"), vocab.IRI("http://example.com/ob-2")}},
		},
		{
			name: "WithURL",
			args: []InitFn{WithURL(vocab.IRI("http://exaple.com/url"), vocab.IRI("http://example.com"))},
			want: &vocab.Object{URL: vocab.ItemCollection{vocab.IRI("http://exaple.com/url"), vocab.IRI("http://example.com")}},
		},
		{
			name: "WithContext",
			args: []InitFn{WithContext(vocab.IRI("http://exaple.com"))},
			want: &vocab.Object{Context: vocab.IRI("http://exaple.com")},
		},
		{
			name: "WithGenerator",
			args: []InitFn{WithGenerator(vocab.IRI("http://exaple.com"))},
			want: &vocab.Object{Generator: vocab.IRI("http://exaple.com")},
		},
		{
			name: "WithAudience",
			args: []InitFn{WithAudience(vocab.IRI("http://exaple.com"))},
			want: &vocab.Object{Audience: vocab.ItemCollection{vocab.IRI("http://exaple.com")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Object(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Object() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestTombstone(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Tombstone
	}{
		{
			name: "empty",
			want: &vocab.Tombstone{Type: vocab.TombstoneType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Tombstone{ID: "http://example.com", Type: vocab.TombstoneType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Tombstone{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
		{
			name: "WithFormerType",
			args: []InitFn{WithFormerType("invalid-type")},
			want: &vocab.Tombstone{Type: vocab.TombstoneType, FormerType: vocab.ActivityVocabularyType("invalid-type")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tombstone(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Tombstone() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestPlace(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Place
	}{
		{
			name: "empty",
			want: &vocab.Place{Type: vocab.PlaceType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Place{ID: "http://example.com", Type: vocab.PlaceType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Place{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
		{
			name: "WithRadius",
			args: []InitFn{WithRadius(666)},
			want: &vocab.Place{Type: vocab.PlaceType, Radius: 666},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Place(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Place() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestProfile(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Profile
	}{
		{
			name: "empty",
			want: &vocab.Profile{Type: vocab.ProfileType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Profile{ID: "http://example.com", Type: vocab.ProfileType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Profile{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
		{
			name: "WithDescribes",
			args: []InitFn{WithDescribes(vocab.IRI("http://example.com"))},
			want: &vocab.Profile{Type: vocab.ProfileType, Describes: vocab.IRI("http://example.com")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Profile(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Profile() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestRelationship(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Relationship
	}{
		{
			name: "empty",
			want: &vocab.Relationship{Type: vocab.RelationshipType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Relationship{ID: "http://example.com", Type: vocab.RelationshipType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Relationship{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
		{
			name: "WithRelationship",
			args: []InitFn{WithRelationship(vocab.IRI("http://example.com/relationship"))},
			want: &vocab.Relationship{Type: vocab.RelationshipType, Relationship: vocab.IRI("http://example.com/relationship")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Relationship(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Relationship() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestVideo(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Video
	}{
		{
			name: "empty",
			want: &vocab.Video{Type: vocab.VideoType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Video{ID: "http://example.com", Type: vocab.VideoType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Video{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Video(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Video() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestAudio(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Audio
	}{
		{
			name: "empty",
			want: &vocab.Audio{Type: vocab.AudioType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Audio{ID: "http://example.com", Type: vocab.AudioType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Audio{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Audio(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Audio() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestImage(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Image
	}{
		{
			name: "empty",
			want: &vocab.Image{Type: vocab.ImageType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Image{ID: "http://example.com", Type: vocab.ImageType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Image{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Image(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Image() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestArticle(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Article
	}{
		{
			name: "empty",
			want: &vocab.Article{Type: vocab.ArticleType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Article{ID: "http://example.com", Type: vocab.ArticleType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Article{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Article(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Article() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestNote(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Note
	}{
		{
			name: "empty",
			want: &vocab.Note{Type: vocab.NoteType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Note{ID: "http://example.com", Type: vocab.NoteType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Note{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Note(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Note() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestDocument(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Document
	}{
		{
			name: "empty",
			want: &vocab.Document{Type: vocab.DocumentType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Document{ID: "http://example.com", Type: vocab.DocumentType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Document{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Document(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Document() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestPage(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Page
	}{
		{
			name: "empty",
			want: &vocab.Page{Type: vocab.PageType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Page{ID: "http://example.com", Type: vocab.PageType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Page{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Page(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Page() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestEvent(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Event
	}{
		{
			name: "empty",
			want: &vocab.Event{Type: vocab.EventType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Event{ID: "http://example.com", Type: vocab.EventType},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Event{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Event(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Event() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
