package activitypub

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRelationship_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestRelationship_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestRelationship_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestRelationship_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestRelationship_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestRelationship_IsObject(t *testing.T) {
	t.Skipf("TODO")
}

func TestRelationship_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestRelationship_Clean(t *testing.T) {
	t.Skipf("TODO")
}

func TestToRelationship(t *testing.T) {
	tests := []struct {
		name    string
		it      LinkOrIRI
		want    *Relationship
		wantErr error
	}{
		{
			name: "empty",
		},
		{
			name: "Valid Relationship",
			it:   Relationship{ID: "test", Type: RelationshipType},
			want: &Relationship{ID: "test", Type: RelationshipType},
		},
		{
			name: "Valid *Relationship",
			it:   &Relationship{ID: "test", Type: RelationshipType},
			want: &Relationship{ID: "test", Type: RelationshipType},
		},
		{
			name:    "IRI",
			it:      IRI("https://example.com"),
			wantErr: ErrorInvalidType[Relationship](IRI("")),
		},
		{
			name:    "IRIs",
			it:      IRIs{IRI("https://example.com")},
			wantErr: ErrorInvalidType[Relationship](IRIs{}),
		},
		{
			name:    "ItemCollection",
			it:      ItemCollection{},
			wantErr: ErrorInvalidType[Relationship](ItemCollection{}),
		},
		{
			name:    "Object",
			it:      &Object{ID: "test", Type: ArticleType},
			wantErr: ErrorInvalidType[Relationship](&Object{}),
		},
		{
			name:    "Activity",
			it:      &Activity{ID: "test", Type: CreateType},
			wantErr: ErrorInvalidType[Relationship](&Activity{}),
		},
		{
			name:    "IntransitiveActivity",
			it:      &IntransitiveActivity{ID: "test", Type: ArriveType},
			wantErr: ErrorInvalidType[Relationship](&IntransitiveActivity{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToRelationship(tt.it)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("ToRelationship() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ToRelationship() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
