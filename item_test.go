package activitypub

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func areItems(a, b any) bool {
	_, ok1 := a.(Item)
	_, ok2 := b.(Item)
	return ok1 && ok2
}

func compareItems(x, y any) bool {
	var i1 Item
	var i2 Item
	if ic1, ok := x.(Item); ok {
		i1 = ic1
	}
	if ic2, ok := y.(Item); ok {
		i2 = ic2
	}
	return ItemsEqual(i1, i2)
}

var EquateItems = cmp.FilterValues(areItems, cmp.Comparer(compareItems))

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
			want: false,
		},
		{
			name: "empty_objects_different_alias_type#1",
			args: args{&Actor{}, &Object{}},
			want: false,
		},
		{
			name: "same_id_object",
			args: args{&Object{ID: "test"}, &Object{ID: "test"}},
			want: true,
		},
		{
			name: "same_id_object_different_alias",
			args: args{&Activity{ID: "test"}, &Object{ID: "test"}},
			want: false,
		},
		{
			name: "same_id_object_different_alias#1",
			args: args{&Activity{ID: "test"}, &Actor{ID: "test"}},
			want: false,
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
		{
			name: "Link different than Object",
			args: args{&Object{ID: "test", Type: NoteType}, &Link{ID: "test", Type: MentionType}},
			want: false,
		},
		{
			name: "matching Link",
			args: args{&Link{ID: "test", Type: MentionType}, &Link{ID: "test", Type: MentionType}},
			want: true,
		},
		{
			name: "different id Links",
			args: args{&Link{ID: "test1", Type: MentionType}, &Link{ID: "test", Type: MentionType}},
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
			want: true,
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

func TestIsCollection(t *testing.T) {
	tests := []struct {
		name string
		it   LinkOrIRI
		want bool
	}{
		{
			name: "nil",
			it:   nil,
			want: false,
		},
		{
			name: "nil Item",
			it:   Item(nil),
			want: false,
		},
		{
			name: "nil IRIs",
			it:   IRIs(nil),
			want: false,
		},
		{
			name: "nil ItemCollection",
			it:   ItemCollection(nil),
			want: true,
		},
		{
			name: "nil *OrderedCollection",
			it:   (*OrderedCollection)(nil),
			want: true,
		},
		{
			name: "nil *OrderedCollectionPage",
			it:   (*OrderedCollectionPage)(nil),
			want: true,
		},
		{
			name: "nil *Collection",
			it:   (*Collection)(nil),
			want: true,
		},
		{
			name: "nil *CollectionPage",
			it:   (*CollectionPage)(nil),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCollection(tt.it); got != tt.want {
				t.Errorf("IsCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleOnItem() {
	// In the OnObject() examples we have seen how for the cases of Item values
	// wrapping a slice of elements, the call is done for each of them.
	//
	// This can work most of the time, as generally ItemCollections will contain
	// the same type of objects. But that is not always the case, there are cases
	// when collections contain heterogeneous types, eg (a collection of objects
	// that contains Object, Place, Relationship and Tombstone elements,
	// a collection of actors that contains Actor and Tombstone elements, or,
	// a collection of activities that contains both Activity and
	// IntransitiveActivity elements.
	//
	// When we want to apply similar logic to these types of slices, we can use
	// the OnItem() function.

	var heterogeneous Item = ItemCollection{
		// We haven't yet explored how an IRI string satisfies the Item interface
		// but here is a short example.
		IRI("http://example.com"),
		&Object{ID: "http://example.com/1", Type: ImageType},
		&Actor{ID: "http://example.com/~jdoe", Type: PersonType},
		&Tombstone{ID: "http://example.com/2", Type: TombstoneType, FormerType: NoteType},
	}

	cnt := 0
	_ = OnItem(heterogeneous, func(it Item) error {
		fmt.Printf("%d: %v\n", cnt, it)

		// NOTE(marius): we can operate on the items using the other OnXXX functions
		// and the changes will be reflected outside our scope if they are pointers.
		_ = OnObject(it, func(ob *Object) error {
			ob.ID = IRI("http://example.com/" + strconv.Itoa(cnt))
			return nil
		})
		cnt++
		return nil
	})
	// For a more idiomatic way about handling an ItemCollection slice,
	// see the OnItemCollection() example.

	// NOTE(marius): we can print again the modified objects.
	// We see that first element which is an IRI was not affected,
	// because it wasn't convertable to an Object type,
	// but it was still iterated over, as shown by the ID counter on the others.
	_ = OnItem(heterogeneous, func(it Item) error {
		fmt.Printf("%d: %v\n", cnt, it)
		cnt++
		return nil
	})

	// Output:
	// 0: http://example.com
	// 1: activitypub.Object[Image] { id: http://example.com/1 }
	// 2: activitypub.Actor[Person] { id: http://example.com/~jdoe }
	// 3: activitypub.Tombstone[Tombstone] { id: http://example.com/2, formerType: Note }
	// 4: http://example.com
	// 5: activitypub.Object[Image] { id: http://example.com/1 }
	// 6: activitypub.Actor[Person] { id: http://example.com/2 }
	// 7: activitypub.Tombstone[Tombstone] { id: http://example.com/3, formerType: Note }
}

func Test_typesEqual(t *testing.T) {
	type args struct {
		t1 Typer
		t2 Typer
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{},
			want: true,
		},
		{
			name: "left empty",
			args: args{t1: NoteType},
			want: false,
		},
		{
			name: "right empty",
			args: args{t2: NoteType},
			want: false,
		},
		{
			name: "equal single type",
			args: args{t1: NoteType, t2: NoteType},
			want: true,
		},
		{
			name: "not equal single type",
			args: args{t1: ArticleType, t2: NoteType},
			want: false,
		},
		{
			name: "equal multiple types",
			args: args{t1: ActivityVocabularyTypes{NoteType, TombstoneType}, t2: ActivityVocabularyTypes{TombstoneType, NoteType}},
			want: true,
		},
		{
			name: "not equal multiple types",
			args: args{t1: ActivityVocabularyTypes{ArticleType, NoteType, TombstoneType}, t2: ActivityVocabularyTypes{TombstoneType, NoteType}},
			want: false,
		},
		{
			name: "not equal multiple types",
			args: args{t1: ActivityVocabularyTypes{ArticleType, TombstoneType}, t2: ActivityVocabularyTypes{TombstoneType, NoteType}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := typesEqual(tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("typesEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareByType(t *testing.T) {
	type args struct {
		typ  Typer
		it   ObjectOrLink
		with ObjectOrLink
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{},
			want: true,
		},
		{
			name: "equal objects",
			args: args{
				typ:  ImageType,
				it:   &Object{ID: "http://example.com", Type: ImageType},
				with: &Object{ID: "http://example.com", Type: ImageType},
			},
			want: true,
		},
		{
			name: "objects w/o types not equal",
			args: args{
				typ:  ImageType,
				it:   &Object{ID: "http://example.com"},
				with: &Object{ID: "http://example.com"},
			},
			want: false,
		},
		//
		{
			name: "equal Tombstones",
			args: args{
				typ:  TombstoneType,
				it:   &Tombstone{ID: "http://example.com", Type: TombstoneType, FormerType: NoteType},
				with: &Tombstone{ID: "http://example.com", Type: TombstoneType, FormerType: NoteType},
			},
			want: true,
		},
		{
			name: "Tombstones different Summary",
			args: args{
				typ:  TombstoneType,
				it:   &Tombstone{ID: "http://example.com", Type: TombstoneType, Summary: DefaultLangValue("test")},
				with: &Tombstone{ID: "http://example.com", Type: TombstoneType, Summary: DefaultLangValue("different")},
			},
			want: false,
		},
		{
			name: "Tombstones different formerType",
			args: args{
				typ:  TombstoneType,
				it:   &Tombstone{ID: "http://example.com", Type: TombstoneType, FormerType: NoteType},
				with: &Tombstone{ID: "http://example.com", Type: TombstoneType, FormerType: ArticleType},
			},
			want: false,
		},
		{
			name: "Tombstones w/o types not equal",
			args: args{
				typ:  TombstoneType,
				it:   &Tombstone{ID: "http://example.com", FormerType: NoteType},
				with: &Tombstone{ID: "http://example.com", FormerType: NoteType},
			},
			want: false,
		},
		//
		{
			name: "equal Profiles",
			args: args{
				typ:  ProfileType,
				it:   &Profile{ID: "http://example.com", Type: ProfileType, Describes: IRI("http://example.com/1")},
				with: &Profile{ID: "http://example.com", Type: ProfileType, Describes: IRI("http://example.com/1")},
			},
			want: true,
		},
		{
			name: "Profiles different Summary",
			args: args{
				typ:  ProfileType,
				it:   &Profile{ID: "http://example.com", Type: ProfileType, Summary: DefaultLangValue("test")},
				with: &Profile{ID: "http://example.com", Type: ProfileType, Summary: DefaultLangValue("different")},
			},
			want: false,
		},
		{
			name: "Profiles different describes",
			args: args{
				typ:  ProfileType,
				it:   &Profile{ID: "http://example.com", Type: ProfileType, Describes: IRI("http://example.com/1")},
				with: &Profile{ID: "http://example.com", Type: ProfileType, Describes: IRI("http://example.com/2")},
			},
			want: false,
		},
		{
			name: "Profiles w/o types not equal",
			args: args{
				typ:  ProfileType,
				it:   &Profile{ID: "http://example.com", Describes: IRI("http://example.com/1")},
				with: &Profile{ID: "http://example.com", Describes: IRI("http://example.com/1")},
			},
			want: false,
		},
		//
		{
			name: "equal Relationships",
			args: args{
				typ:  RelationshipType,
				it:   &Relationship{ID: "http://example.com", Type: RelationshipType, Relationship: IRI("http://example.com/1")},
				with: &Relationship{ID: "http://example.com", Type: RelationshipType, Relationship: IRI("http://example.com/1")},
			},
			want: true,
		},
		{
			name: "Relationships different Summary",
			args: args{
				typ:  RelationshipType,
				it:   &Relationship{ID: "http://example.com", Type: RelationshipType, Summary: DefaultLangValue("test")},
				with: &Relationship{ID: "http://example.com", Type: RelationshipType, Summary: DefaultLangValue("different")},
			},
			want: false,
		},
		{
			name: "Relationships different relationship",
			args: args{
				typ:  RelationshipType,
				it:   &Relationship{ID: "http://example.com", Type: RelationshipType, Relationship: IRI("http://example.com/1")},
				with: &Relationship{ID: "http://example.com", Type: RelationshipType, Relationship: IRI("http://example.com/2")},
			},
			want: false,
		},
		{
			name: "Relationships w/o types not equal",
			args: args{
				typ:  RelationshipType,
				it:   &Relationship{ID: "http://example.com", Relationship: IRI("http://example.com/1")},
				with: &Relationship{ID: "http://example.com", Relationship: IRI("http://example.com/1")},
			},
			want: false,
		},
		//
		{
			name: "equal Places",
			args: args{
				typ:  PlaceType,
				it:   &Place{ID: "http://example.com", Type: PlaceType, Altitude: 200.1, Longitude: 0.666},
				with: &Place{ID: "http://example.com", Type: PlaceType, Altitude: 200.1, Longitude: 0.666},
			},
			want: true,
		},
		{
			name: "Places different Summary",
			args: args{
				typ:  PlaceType,
				it:   &Place{ID: "http://example.com", Type: PlaceType, Summary: DefaultLangValue("test")},
				with: &Place{ID: "http://example.com", Type: PlaceType, Summary: DefaultLangValue("different")},
			},
			want: false,
		},
		{
			name: "Places different latitude",
			args: args{
				typ:  PlaceType,
				it:   &Place{ID: "http://example.com", Type: PlaceType, Latitude: 6.66},
				with: &Place{ID: "http://example.com", Type: PlaceType, Latitude: 6.67},
			},
			want: false,
		},
		{
			name: "Places w/o types not equal",
			args: args{
				typ:  PlaceType,
				it:   &Place{ID: "http://example.com", Radius: 100},
				with: &Place{ID: "http://example.com", Radius: 100},
			},
			want: false,
		},
		//
		{
			// NOTE(marius): This might be wrong
			name: "Actors equal for person type",
			args: args{
				typ:  PersonType,
				it:   &Tombstone{ID: "http://example.com", Type: PersonType},
				with: &Actor{ID: "http://example.com", Type: PersonType, PreferredUsername: DefaultLangValue("jdoe")},
			},
			want: false,
		},
		//
		{
			name: "questions equal",
			args: args{
				typ:  QuestionType,
				it:   &Question{ID: "http://example.com", Type: QuestionType},
				with: &Question{ID: "http://example.com", Type: QuestionType},
			},
			want: true,
		},
		{
			name: "questions anyOf not equal",
			args: args{
				typ:  QuestionType,
				it:   &Question{ID: "http://example.com", Type: QuestionType, AnyOf: IRIs{"https://example.com/a-1"}},
				with: &Question{ID: "http://example.com", Type: QuestionType, AnyOf: IRIs{"https://example.com/a-2"}},
			},
			want: false,
		},
		{
			name: "questions oneOf nil",
			args: args{
				typ:  QuestionType,
				it:   &Question{ID: "http://example.com", Type: QuestionType, OneOf: IRIs{"https://example.com/a-1"}},
				with: &Question{ID: "http://example.com", Type: QuestionType},
			},
			want: false,
		},
		{
			name: "questions name different",
			args: args{
				typ:  QuestionType,
				it:   &Question{ID: "http://example.com", Type: QuestionType, Name: DefaultLangValue("Q1")},
				with: &Question{ID: "http://example.com", Type: QuestionType, Name: DefaultLangValue("Q2")},
			},
			want: false,
		},
		//
		{
			name: "intransitiveActivities equal",
			args: args{
				typ:  ArriveType,
				it:   &IntransitiveActivity{ID: "http://example.com", Type: ArriveType},
				with: &IntransitiveActivity{ID: "http://example.com", Type: ArriveType},
			},
			want: true,
		},
		{
			name: "intransitiveActivities actors not equal",
			args: args{
				typ:  TravelType,
				it:   &IntransitiveActivity{ID: "http://example.com", Type: TravelType, Actor: IRIs{"https://example.com/jdoe"}},
				with: &IntransitiveActivity{ID: "http://example.com", Type: TravelType, Actor: IRIs{"https://example.com/alice"}},
			},
			want: false,
		},
		{
			name: "intransitiveActivities target nil",
			args: args{
				typ:  TravelType,
				it:   &IntransitiveActivity{ID: "http://example.com", Type: TravelType, Target: IRIs{"https://example.com/a-1"}},
				with: &IntransitiveActivity{ID: "http://example.com", Type: TravelType},
			},
			want: false,
		},
		{
			name: "intransitiveActivities name different",
			args: args{
				typ:  TravelType,
				it:   &IntransitiveActivity{ID: "http://example.com", Type: TravelType, Name: DefaultLangValue("Q1")},
				with: &IntransitiveActivity{ID: "http://example.com", Type: TravelType, Name: DefaultLangValue("Q2")},
			},
			want: false,
		},
		//
		{
			name: "activities equal",
			args: args{
				typ:  CreateType,
				it:   &Activity{ID: "http://example.com", Type: CreateType},
				with: &Activity{ID: "http://example.com", Type: CreateType},
			},
			want: true,
		},
		{
			name: "Activities objects not equal",
			args: args{
				typ:  UpdateType,
				it:   &Activity{ID: "http://example.com", Type: UpdateType, Object: IRIs{"https://example.com/o1"}},
				with: &Activity{ID: "http://example.com", Type: UpdateType, Object: IRIs{"https://example.com/o2"}},
			},
			want: false,
		},
		{
			name: "Activities actors not equal",
			args: args{
				typ:  DeleteType,
				it:   &Activity{ID: "http://example.com", Type: DeleteType, Actor: IRIs{"https://example.com/jdoe"}},
				with: &Activity{ID: "http://example.com", Type: DeleteType, Actor: IRIs{"https://example.com/alice"}},
			},
			want: false,
		},
		{
			name: "Activities target nil",
			args: args{
				typ:  BlockType,
				it:   &Activity{ID: "http://example.com", Type: BlockType, Target: IRIs{"https://example.com/a-1"}},
				with: &Activity{ID: "http://example.com", Type: BlockType},
			},
			want: false,
		},
		{
			name: "Activities name different",
			args: args{
				typ:  LikeType,
				it:   &Activity{ID: "http://example.com", Type: LikeType, Content: DefaultLangValue("Q1")},
				with: &Activity{ID: "http://example.com", Type: LikeType, Content: DefaultLangValue("Q2")},
			},
			want: false,
		},
		//
		{
			name: "Actors equal",
			args: args{
				typ:  PersonType,
				it:   &Actor{ID: "http://example.com", Type: PersonType},
				with: &Actor{ID: "http://example.com", Type: PersonType},
			},
			want: true,
		},
		{
			name: "Actors replies not equal",
			args: args{
				typ:  ApplicationType,
				it:   &Actor{ID: "http://example.com", Type: ApplicationType, Replies: IRIs{"https://example.com/replies"}},
				with: &Actor{ID: "http://example.com", Type: ApplicationType, Replies: IRIs{"https://example.com/rr"}},
			},
			want: false,
		},
		{
			name: "Actors inbox nil",
			args: args{
				typ:  ServiceType,
				it:   &Actor{ID: "http://example.com", Type: ServiceType, Inbox: IRIs{"https://example.com/a-1"}},
				with: &Actor{ID: "http://example.com", Type: ServiceType},
			},
			want: false,
		},
		{
			name: "Actors preferredUsername different",
			args: args{
				typ:  GroupType,
				it:   &Actor{ID: "http://example.com", Type: GroupType, PreferredUsername: DefaultLangValue("Q1")},
				with: &Actor{ID: "http://example.com", Type: GroupType, PreferredUsername: DefaultLangValue("Q2")},
			},
			want: false,
		},
		// Collection
		{
			name: "collection equals",
			args: args{
				typ:  CollectionType,
				it:   &Collection{Type: CollectionType, First: IRI("http://example.com/1st"), TotalItems: 1, Items: ItemCollection{IRI("http://example.com")}},
				with: &Collection{Type: CollectionType, First: IRI("http://example.com/1st"), TotalItems: 1, Items: ItemCollection{IRI("http://example.com")}},
			},
			want: true,
		},
		// OrderedCollection
		{
			name: "ordered-collection equals",
			args: args{
				typ:  OrderedCollectionType,
				it:   &OrderedCollection{Type: OrderedCollectionType, First: IRI("http://example.com/1st"), TotalItems: 1, OrderedItems: ItemCollection{IRI("http://example.com")}},
				with: &OrderedCollection{Type: OrderedCollectionType, First: IRI("http://example.com/1st"), TotalItems: 1, OrderedItems: ItemCollection{IRI("http://example.com")}},
			},
			want: true,
		},
		// CollectionPage
		{
			name: "collection-page equals",
			args: args{
				typ:  CollectionPageType,
				it:   &CollectionPage{Type: CollectionPageType, First: IRI("http://example.com/1st"), TotalItems: 1, Items: ItemCollection{IRI("http://example.com")}},
				with: &CollectionPage{Type: CollectionPageType, First: IRI("http://example.com/1st"), TotalItems: 1, Items: ItemCollection{IRI("http://example.com")}},
			},
			want: true,
		},
		// OrderedCollectionPage
		{
			name: "ordered-collection-page equals",
			args: args{
				typ:  OrderedCollectionPageType,
				it:   &OrderedCollectionPage{Type: OrderedCollectionPageType, First: IRI("http://example.com/1st"), TotalItems: 1, OrderedItems: ItemCollection{IRI("http://example.com")}},
				with: &OrderedCollectionPage{Type: OrderedCollectionPageType, First: IRI("http://example.com/1st"), TotalItems: 1, OrderedItems: ItemCollection{IRI("http://example.com")}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareByType(tt.args.typ, tt.args.it, tt.args.with); got != tt.want {
				t.Errorf("compareByType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_typedObjectsEqual(t *testing.T) {
	type args struct {
		it   ObjectOrLink
		with ObjectOrLink
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{},
			want: true,
		},
		{
			name: "tombstone not equal to object",
			args: args{
				it:   &Object{Type: TombstoneType},
				with: &Tombstone{Type: TombstoneType},
			},
			want: false,
		},
		{
			name: "tombstones with different ids",
			args: args{
				it:   &Tombstone{ID: "http://example.com", Type: TombstoneType},
				with: &Tombstone{ID: "http://example.com/1", Type: TombstoneType},
			},
			want: false,
		},
		{
			name: "object with different types not equal",
			args: args{
				it:   &Object{Type: NoteType},
				with: &Object{Type: ArticleType},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := typedObjectsEqual(tt.args.it, tt.args.with); got != tt.want {
				t.Errorf("typedObjectsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
