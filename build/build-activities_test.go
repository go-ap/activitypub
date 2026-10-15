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
		{
			name: "WithObject",
			args: []InitFn{WithObject(vocab.IRI("http://example.com/o1"), vocab.IRI("http://example.com/o2"))},
			want: &vocab.Activity{Object: vocab.ItemCollection{vocab.IRI("http://example.com/o1"), vocab.IRI("http://example.com/o2")}},
		},
		{
			name: "WithActor",
			args: []InitFn{WithActor(vocab.IRI("http://example.com/~jdoe"))},
			want: &vocab.Activity{Actor: vocab.IRI("http://example.com/~jdoe")},
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
		{
			name: "WithActor",
			args: []InitFn{WithActor(vocab.IRI("http://example.com/~jdoe"))},
			want: &vocab.IntransitiveActivity{Actor: vocab.IRI("http://example.com/~jdoe")},
		},
		{
			name: "WithTarget",
			args: []InitFn{WithTarget(vocab.IRI("http://example.com"))},
			want: &vocab.IntransitiveActivity{Target: vocab.IRI("http://example.com")},
		},
		{
			name: "WithOrigin",
			args: []InitFn{WithOrigin(vocab.IRI("http://example.com"))},
			want: &vocab.IntransitiveActivity{Origin: vocab.IRI("http://example.com")},
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
		{
			name: "WithActor",
			args: []InitFn{WithActor(vocab.IRI("http://example.com/~jdoe"))},
			want: &vocab.Question{Type: vocab.QuestionType, Actor: vocab.IRI("http://example.com/~jdoe")},
		},
		{
			name: "WithAnyOf",
			args: []InitFn{WithAnyOf(vocab.IRI("http://example.com/a1"), vocab.IRI("http://example.com/a2"))},
			want: &vocab.Question{Type: vocab.QuestionType, AnyOf: vocab.ItemCollection{vocab.IRI("http://example.com/a1"), vocab.IRI("http://example.com/a2")}},
		},
		{
			name: "WithOneOf",
			args: []InitFn{WithOneOf(vocab.IRI("http://example.com/a1"), vocab.IRI("http://example.com/a2"))},
			want: &vocab.Question{Type: vocab.QuestionType, OneOf: vocab.ItemCollection{vocab.IRI("http://example.com/a1"), vocab.IRI("http://example.com/a2")}},
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

func TestMove(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Move
	}{
		{
			name: "empty",
			want: &vocab.Move{Type: vocab.MoveType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Move{ID: "http://example.com", Type: vocab.MoveType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Move(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Move() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Remove
	}{
		{
			name: "empty",
			want: &vocab.Remove{Type: vocab.RemoveType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Remove{ID: "http://example.com", Type: vocab.RemoveType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Remove() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Add
	}{
		{
			name: "empty",
			want: &vocab.Add{Type: vocab.AddType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Add{ID: "http://example.com", Type: vocab.AddType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Add() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestCreate(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Create
	}{
		{
			name: "empty",
			want: &vocab.Create{Type: vocab.CreateType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Create{ID: "http://example.com", Type: vocab.CreateType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Create(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Create() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Update
	}{
		{
			name: "empty",
			want: &vocab.Update{Type: vocab.UpdateType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Update{ID: "http://example.com", Type: vocab.UpdateType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Update(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Update() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Delete
	}{
		{
			name: "empty",
			want: &vocab.Delete{Type: vocab.DeleteType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Delete{ID: "http://example.com", Type: vocab.DeleteType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Delete(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Delete() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestLike(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Like
	}{
		{
			name: "empty",
			want: &vocab.Like{Type: vocab.LikeType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Like{ID: "http://example.com", Type: vocab.LikeType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Like(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Like() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestDislike(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Dislike
	}{
		{
			name: "empty",
			want: &vocab.Dislike{Type: vocab.DislikeType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Dislike{ID: "http://example.com", Type: vocab.DislikeType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dislike(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Dislike() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestIgnore(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Ignore
	}{
		{
			name: "empty",
			want: &vocab.Ignore{Type: vocab.IgnoreType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Ignore{ID: "http://example.com", Type: vocab.IgnoreType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ignore(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Ignore() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestUndo(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Undo
	}{
		{
			name: "empty",
			want: &vocab.Undo{Type: vocab.UndoType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Undo{ID: "http://example.com", Type: vocab.UndoType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Undo(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Undo() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestBlock(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Block
	}{
		{
			name: "empty",
			want: &vocab.Block{Type: vocab.BlockType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Block{ID: "http://example.com", Type: vocab.BlockType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Block(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Block() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestFlag(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Flag
	}{
		{
			name: "empty",
			want: &vocab.Flag{Type: vocab.FlagType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Flag{ID: "http://example.com", Type: vocab.FlagType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Flag(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Flag() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestAnnounce(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Announce
	}{
		{
			name: "empty",
			want: &vocab.Announce{Type: vocab.AnnounceType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Announce{ID: "http://example.com", Type: vocab.AnnounceType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Announce(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Announce() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestFollow(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Follow
	}{
		{
			name: "empty",
			want: &vocab.Follow{Type: vocab.FollowType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Follow{ID: "http://example.com", Type: vocab.FollowType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Follow(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Follow() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestAccept(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Accept
	}{
		{
			name: "empty",
			want: &vocab.Accept{Type: vocab.AcceptType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Accept{ID: "http://example.com", Type: vocab.AcceptType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Accept(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Accept() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestReject(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Reject
	}{
		{
			name: "empty",
			want: &vocab.Reject{Type: vocab.RejectType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Reject{ID: "http://example.com", Type: vocab.RejectType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reject(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Reject() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
