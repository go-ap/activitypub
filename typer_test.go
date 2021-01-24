package handlers

import (
	"github.com/go-ap/activitypub"
	"testing"
)

func TestPathTyper_Type(t *testing.T) {
	t.Skipf("TODO")
}

func TestValidActivityCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestValidCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestValidObjectCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestValidCollectionIRI(t *testing.T) {
	t.Skipf("TODO")
}

func TestSplit(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollectionType_AddTo(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollectionType_IRI(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollectionType_OfActor(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollectionTypes_Contains(t *testing.T) {
	t.Skipf("TODO")
}

func TestIRIf(t *testing.T) {
	type args struct {
		i activitypub.IRI
		t CollectionType
	}
	tests := []struct {
		name string
		args args
		want activitypub.IRI
	}{
		{
			name: "plain concat",
			args: args{
				i: "https://example.com",
				t: "inbox",
			},
			want: "https://example.com/inbox",
		},
		{
			name: "strip root from iri",
			args: args{
				i: "https://example.com/",
				t: "inbox",
			},
			want: "https://example.com/inbox",
		},
		{
			name: "invalid iri",
			args: args{
				i: "example.com",
				t: "test",
			},
			want: "example.com/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IRIf(tt.args.i, tt.args.t); got != tt.want {
				t.Errorf("IRIf() = %v, want %v", got, tt.want)
			}
		})
	}
}
