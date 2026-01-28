package activitypub

import "testing"

func TestFlatten(t *testing.T) {
	t.Skipf("TODO")
}

func TestItemsEqual(t *testing.T) {
	type args struct {
		it   Item
		with Item
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil_items_equal",
			args: args{nil, nil},
			want: true,
		},
		{
			name: "nil_item_with_object",
			args: args{nil, &Object{}},
			want: false,
		},
		{
			name: "nil_item_with_object#1",
			args: args{&Object{}, nil},
			want: false,
		},
		{
			name: "empty_objects",
			args: args{&Object{}, &Object{}},
			want: true,
		},
		{
			name: "empty_objects_different_alias_type",
			args: args{&Activity{}, &Object{}},
			want: true,
		},
		{
			name: "empty_objects_different_alias_type#1",
			args: args{&Actor{}, &Object{}},
			want: true,
		},
		{
			name: "same_id_object",
			args: args{&Object{ID: "test"}, &Object{ID: "test"}},
			want: true,
		},
		{
			name: "same_id_object_different_alias",
			args: args{&Activity{ID: "test"}, &Object{ID: "test"}},
			want: true,
		},
		{
			name: "same_id_object_different_alias#1",
			args: args{&Activity{ID: "test"}, &Actor{ID: "test"}},
			want: true,
		},
		{
			name: "different_id_objects",
			args: args{&Object{ID: "test1"}, &Object{ID: "test"}},
			want: false,
		},
		{
			name: "different_id_types",
			args: args{&Object{ID: "test", Type: NoteType}, &Object{ID: "test", Type: ArticleType}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ItemsEqual(tt.args.it, tt.args.with); got != tt.want {
				t.Errorf("ItemsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNil(t *testing.T) {
	type args struct {
		it Item
	}
	var (
		o      *Object
		col    *ItemCollection
		iris   *IRIs
		obNil  Item = o
		colNil Item = col
		itIRIs Item = iris
	)
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil is nil",
			args: args{
				it: nil,
			},
			want: true,
		},
		{
			name: "Item is nil",
			args: args{
				it: Item(nil),
			},
			want: true,
		},
		{
			name: "Object nil",
			args: args{
				it: obNil,
			},
			want: true,
		},
		{
			name: "IRIs nil",
			args: args{
				it: iris,
			},
			want: true,
		},
		{
			name: "IRIs as Item nil",
			args: args{
				it: itIRIs,
			},
			want: true,
		},
		{
			name: "IRIs not nil",
			args: args{
				it: IRIs{},
			},
			want: false,
		},
		{
			name: "IRIs as Item not nil",
			args: args{
				it: Item(IRIs{}),
			},
			want: false,
		},
		{
			name: "ItemCollection nil",
			args: args{
				it: col,
			},
			want: true,
		},
		{
			name: "ItemCollection as Item nil",
			args: args{
				it: colNil,
			},
			want: true,
		},
		{
			name: "ItemCollection not nil",
			args: args{
				it: ItemCollection{},
			},
			want: false,
		},
		{
			name: "object-not-nil",
			args: args{
				it: &Object{},
			},
			want: false,
		},
		{
			name: "place-not-nil",
			args: args{
				it: &Place{},
			},
			want: false,
		},
		{
			name: "tombstone-not-nil",
			args: args{
				it: &Tombstone{},
			},
			want: false,
		},
		{
			name: "collection-not-nil",
			args: args{
				it: &Collection{},
			},
			want: false,
		},
		{
			name: "activity-not-nil",
			args: args{
				it: &Activity{},
			},
			want: false,
		},
		{
			name: "intransitive-activity-not-nil",
			args: args{
				it: &IntransitiveActivity{},
			},
			want: false,
		},
		{
			name: "actor-not-nil",
			args: args{
				it: &Actor{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNil(tt.args.it); got != tt.want {
				t.Errorf("IsNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemsEqual1(t *testing.T) {
	type args struct {
		it   Item
		with Item
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil",
			args: args{},
			want: true,
		},
		{
			name: "equal empty items",
			args: args{
				it:   &Object{},
				with: &Actor{},
			},
			want: true,
		},
		{
			name: "equal same ID items",
			args: args{
				it:   &Object{ID: "example-1"},
				with: &Object{ID: "example-1"},
			},
			want: true,
		},
		{
			name: "different IDs",
			args: args{
				it:   &Object{ID: "example-1"},
				with: &Object{ID: "example-2"},
			},
			want: false,
		},
		{
			name: "different properties",
			args: args{
				it:   &Object{ID: "example-1"},
				with: &Object{Type: ArticleType},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ItemsEqual(tt.args.it, tt.args.with); got != tt.want {
				t.Errorf("ItemsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	nilIRI *IRI = nil

	nilLink *Link = nil

	nilObject       *Object       = nil
	nilTombstone    *Tombstone    = nil
	nilProfile      *Profile      = nil
	nilPlace        *Place        = nil
	nilRelationship *Relationship = nil

	nilActor *Actor = nil

	nilActivity             *Activity             = nil
	nilIntransitiveActivity *IntransitiveActivity = nil

	nilCollectionIntf CollectionInterface = nil

	nilCollection     *Collection     = nil
	nilCollectionPage *CollectionPage = nil

	nilOrderedCollection     *OrderedCollection     = nil
	nilOrderedCollectionPage *OrderedCollectionPage = nil
)

func TestIsObject(t *testing.T) {
	tests := []struct {
		name string
		it   Item
		want bool
	}{
		{
			name: "nil",
			want: false,
		},
		{
			name: "interface with nil value",
			it:   Item(nil),
			want: false,
		},
		{
			name: "empty object",
			it:   Object{},
			want: true,
		},
		{
			name: "pointer to empty object",
			it:   &Object{},
			want: true,
		},
		{
			name: "pointer to nil object",
			it:   nilObject,
			want: false,
		},
		{
			name: "pointer to nil tombstone",
			it:   nilTombstone,
			want: false,
		},
		{
			name: "pointer to nil profile",
			it:   nilProfile,
			want: false,
		},
		{
			name: "pointer to nil place",
			it:   nilPlace,
			want: false,
		},
		{
			name: "pointer to nil relationship",
			it:   nilRelationship,
			want: false,
		},
		{
			name: "pointer to nil actor",
			it:   nilActor,
			want: false,
		},
		{
			name: "pointer to nil activity",
			it:   nilActivity,
			want: false,
		},
		{
			name: "pointer to nil intransitive activity",
			it:   nilIntransitiveActivity,
			want: false,
		},
		{
			name: "pointer to nil collection interface",
			it:   nilCollectionIntf,
			want: false,
		},
		{
			name: "pointer to nil collection",
			it:   nilCollection,
			want: false,
		},
		{
			name: "pointer to nil collection page",
			it:   nilCollectionPage,
			want: false,
		},
		{
			name: "pointer to nil ordered collection",
			it:   nilOrderedCollection,
			want: false,
		},
		{
			name: "pointer to nil ordered collection page",
			it:   nilOrderedCollectionPage,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsObject(tt.it); got != tt.want {
				t.Errorf("IsObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemsEqual2(t *testing.T) {
	type args struct {
		it   Item
		with Item
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil vs nil",
			args: args{
				it:   nil,
				with: nil,
			},
			want: true,
		},
		{
			name: "nil vs object",
			args: args{
				it:   nil,
				with: Object{},
			},
			want: false,
		},
		{
			name: "object vs nil",
			args: args{
				it:   Object{},
				with: nil,
			},
			want: false,
		},
		{
			name: "empty object vs empty object",
			args: args{
				it:   Object{},
				with: Object{},
			},
			want: true,
		},
		{
			name: "object-id vs empty object",
			args: args{
				it:   Object{ID: "https://example.com"},
				with: Object{},
			},
			want: false,
		},
		{
			name: "empty object vs object-id",
			args: args{
				it:   Object{},
				with: Object{ID: "https://example.com"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ItemsEqual(tt.args.it, tt.args.with); got != tt.want {
				t.Errorf("ItemsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsItemCollection(t *testing.T) {
	tests := []struct {
		name string
		it   LinkOrIRI
		want bool
	}{
		{
			name: "empty",
			it:   nil,
			want: false,
		},
		{
			name: "object",
			it:   Object{},
			want: false,
		},
		{
			name: "nil item",
			it:   Item(nil),
			want: false,
		},
		{
			name: "nil item collection",
			it:   ItemCollection(nil),
			want: true,
		},
		{
			name: "item collection with nil item",
			it:   ItemCollection{nil},
			want: true,
		},
		{
			name: "item collection with one item",
			it:   ItemCollection{Object{}},
			want: true,
		},
		{
			name: "nil iris",
			it:   IRIs(nil),
			want: false,
		},
		{
			name: "iris with no items",
			it:   IRIs{},
			want: true,
		},
		{
			name: "iris with one item",
			it:   IRIs{""},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsItemCollection(tt.it); got != tt.want {
				t.Errorf("IsItemCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIRI(t *testing.T) {
	tests := []struct {
		name string
		it   LinkOrIRI
		want bool
	}{
		{
			name: "empty",
			it:   nil,
			want: false,
		},
		{
			name: "nil iri",
			it:   nilIRI,
			want: false,
		},
		{
			name: "empty string",
			it:   IRI(""),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIRI(tt.it); got != tt.want {
				t.Errorf("IsIRI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIRIs(t *testing.T) {
	tests := []struct {
		name string
		it   LinkOrIRI
		want bool
	}{
		{
			name: "empty",
			it:   nil,
			want: false,
		},
		{
			name: "object",
			it:   Object{},
			want: false,
		},
		{
			name: "nil iri",
			it:   nilIRI,
			want: false,
		},
		{
			name: "nil item",
			it:   Item(nil),
			want: false,
		},
		{
			name: "nil item collection",
			it:   ItemCollection(nil),
			want: false,
		},
		{
			name: "item collection with nil item",
			it:   ItemCollection{nil},
			want: false,
		},
		{
			name: "item collection with one item",
			it:   ItemCollection{Object{}},
			want: false,
		},
		{
			name: "nil iris",
			it:   IRIs(nil),
			want: false,
		},
		{
			name: "iris with no items",
			it:   IRIs{},
			want: true,
		},
		{
			name: "iris with one item",
			it:   IRIs{""},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIRIs(tt.it); got != tt.want {
				t.Errorf("IsIRIs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLink(t *testing.T) {
	tests := []struct {
		name string
		it   LinkOrIRI
		want bool
	}{
		{
			name: "empty",
			it:   nil,
			want: false,
		},
		{
			name: "nil link",
			it:   nilLink,
			want: false,
		},
		{
			name: "object",
			it:   Object{},
			want: false,
		},
		{
			name: "link",
			it:   Link{},
			want: true,
		},
		{
			name: "link pointer",
			it:   &Link{},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLink(tt.it); got != tt.want {
				t.Errorf("IsLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
