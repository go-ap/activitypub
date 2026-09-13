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
	tests := []struct {
		name string
		a    Relationship
		want Item
	}{
		{
			name: "empty",
			a:    Relationship{},
			want: &Relationship{},
		},
		{
			name: "has Bto",
			a:    Relationship{Type: RelationshipType, Bto: ItemCollection{IRI("http://example.com")}},
			want: &Relationship{Type: RelationshipType},
		},
		{
			name: "has BCC",
			a:    Relationship{Type: RelationshipType, BCC: ItemCollection{IRI("http://example.com")}},
			want: &Relationship{Type: RelationshipType},
		},
		{
			name: "audience has BCC",
			a:    Relationship{Type: RelationshipType, Audience: ItemCollection{&Object{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &Relationship{Type: RelationshipType, Audience: ItemCollection{&Object{}}},
		},
		{
			name: "attachment has BCC",
			a:    Relationship{Type: RelationshipType, Attachment: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Relationship{Type: RelationshipType, Attachment: &Object{}},
		},
		{
			name: "icon has BCC",
			a:    Relationship{Type: RelationshipType, Icon: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Relationship{Type: RelationshipType, Icon: &Object{}},
		},
		{
			name: "image has BCC",
			a:    Relationship{Type: RelationshipType, Image: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Relationship{Type: RelationshipType, Image: &Object{}},
		},
		{
			name: "context has BCC",
			a:    Relationship{Type: RelationshipType, Context: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Relationship{Type: RelationshipType, Context: &Object{}},
		},
		{
			name: "generator has BCC",
			a:    Relationship{Type: RelationshipType, Generator: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Relationship{Type: RelationshipType, Generator: &Object{}},
		},
		{
			name: "attributedTo has BCC",
			a:    Relationship{Type: RelationshipType, AttributedTo: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Relationship{Type: RelationshipType, AttributedTo: &Object{}},
		},
		{
			name: "preview has BCC",
			a:    Relationship{Type: RelationshipType, Preview: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Relationship{Type: RelationshipType, Preview: &Object{}},
		},
		{
			name: "tag has BCC",
			a:    Relationship{Type: RelationshipType, Tag: ItemCollection{&Object{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &Relationship{Type: RelationshipType, Tag: ItemCollection{&Object{}}},
		},
		{
			name: "relationship has BCC",
			a:    Relationship{Type: RelationshipType, Relationship: ItemCollection{&Object{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &Relationship{Type: RelationshipType, Relationship: ItemCollection{&Object{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Clean(); !cmp.Equal(got, tt.want, EquateItems) {
				t.Errorf("Clean() = %s", cmp.Diff(tt.want, got, EquateItems))
			}
		})
	}
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
		{
			name: "Tombstone",
			it:   &Tombstone{ID: "test", Type: TombstoneType, FormerType: PersonType},
			want: &Relationship{ID: "test", Type: TombstoneType},
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
			validTypes := ActivityVocabularyTypes{RelationshipType, TombstoneType}
			if got != nil && !validTypes.Match(got.Type) {
				t.Errorf("ToRelationship() expected to match %v types, got = %v", validTypes, got.Type)
			}
		})
	}
}
