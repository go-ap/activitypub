package activitypub

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClone(t *testing.T) {
	tests := []struct {
		name    string
		it      Item
		want    Item
		wantErr error
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Clone(tt.it)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("Clone() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Clone() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
