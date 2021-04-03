package activitypub

import (
	"reflect"
	"testing"
)

func TestFlattenPersonProperties(t *testing.T) {
	t.Skipf("TODO")
}

func TestFlattenProperties(t *testing.T) {
	t.Skipf("TODO")
}

func TestFlattenItemCollection(t *testing.T) {
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
