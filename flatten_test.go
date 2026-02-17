package activitypub

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFlattenPersonProperties(t *testing.T) {
	t.Skipf("TODO")
}

func TestFlattenProperties(t *testing.T) {
	t.Skipf("TODO")
}

func TestFlattenCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestFlattenOrderedCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestFlattenIntransitiveActivityProperties(t *testing.T) {
	type args struct {
		act *IntransitiveActivity
	}
	tests := []struct {
		name string
		args args
		want *IntransitiveActivity
	}{
		{
			name: "blank",
			args: args{&IntransitiveActivity{}},
			want: &IntransitiveActivity{},
		},
		{
			name: "flatten-actor",
			args: args{&IntransitiveActivity{Actor: &Actor{ID: "example-actor-iri"}}},
			want: &IntransitiveActivity{Actor: IRI("example-actor-iri")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenIntransitiveActivityProperties(tt.args.act); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlattenIntransitiveActivityProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenActivityProperties(t *testing.T) {
	type args struct {
		act *Activity
	}
	tests := []struct {
		name string
		args args
		want *Activity
	}{
		{
			name: "blank",
			args: args{&Activity{}},
			want: &Activity{},
		},
		{
			name: "flatten-actor",
			args: args{&Activity{Actor: &Actor{ID: "example-actor-iri"}}},
			want: &Activity{Actor: IRI("example-actor-iri")},
		},
		{
			name: "flatten-object",
			args: args{&Activity{Object: &Object{ID: "example-actor-iri"}}},
			want: &Activity{Object: IRI("example-actor-iri")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenActivityProperties(tt.args.act); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlattenActivityProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		name string
		it   Item
		want Item
	}{
		{
			name: "nil",
			it:   nil,
			want: nil,
		},
		{
			name: "iri",
			it:   IRI("http://example.com"),
			want: IRI("http://example.com"),
		},
		{
			name: "object",
			it:   &Object{ID: IRI("http://example.com")},
			want: IRI("http://example.com"),
		},
		{
			name: "items",
			it:   ItemCollection{IRI("http://example.com"), IRI("http://jdoe.example.com")},
			want: ItemCollection{IRI("http://example.com"), IRI("http://jdoe.example.com")},
		},
		{
			name: "ordered collection",
			it:   &OrderedCollection{ID: "http://example.com"},
			want: IRI("http://example.com"),
		},
		{
			name: "ordered collection page",
			it:   &OrderedCollectionPage{ID: "http://example.com"},
			want: IRI("http://example.com"),
		},
		{
			name: "collection",
			it:   &Collection{ID: "http://example.com"},
			want: IRI("http://example.com"),
		},
		{
			name: "collection page",
			it:   &CollectionPage{ID: "http://example.com"},
			want: IRI("http://example.com"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Flatten(tt.it); !cmp.Equal(got, tt.want) {
				t.Errorf("Flatten() = %s", cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestFlattenItemCollection(t *testing.T) {
	tests := []struct {
		name string
		col  ItemCollection
		want ItemCollection
	}{
		{
			name: "nil",
			col:  nil,
			want: nil,
		},
		{
			name: "empty",
			col:  ItemCollection{},
			want: nil,
		},
		{
			name: "with iri",
			col:  ItemCollection{IRI("http://example.com")},
			want: ItemCollection{IRI("http://example.com")},
		},
		{
			name: "with object",
			col:  ItemCollection{&Object{ID: "http://example.com"}},
			want: ItemCollection{IRI("http://example.com")},
		},
		{
			name: "with objects",
			col:  ItemCollection{&Object{ID: "http://example.com"}, &Actor{ID: "http://example.com/~jdoe"}},
			want: ItemCollection{IRI("http://example.com"), IRI("http://example.com/~jdoe")},
		},
		{
			name: "with duplicates",
			col:  ItemCollection{&Object{ID: "http://example.com"}, &Actor{ID: "http://example.com/~jdoe"}, IRI("http://example.com")},
			want: ItemCollection{IRI("http://example.com"), IRI("http://example.com/~jdoe")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenItemCollection(tt.col); !cmp.Equal(got, tt.want) {
				t.Errorf("FlattenItemCollection() = %s", cmp.Diff(got, tt.want))
			}
		})
	}
}
