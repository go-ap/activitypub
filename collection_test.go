package activitypub

import (
	"reflect"
	"testing"

	"github.com/go-ap/errors"
	"github.com/google/go-cmp/cmp"
)

func mockCollection(items ...Item) Collection {
	cc := Collection{
		ID:   IRIf("https://example.com", Inbox),
		Type: CollectionType,
	}
	if len(items) == 0 {
		cc.Items = make(ItemCollection, 0)
	} else {
		cc.Items = items
		cc.TotalItems = uint(len(items))
	}
	return cc
}

func TestCollection_Append(t *testing.T) {
	tests := []struct {
		name    string
		col     Collection
		it      []Item
		wantErr error
	}{
		{
			name: "empty",
			col:  mockCollection(),
			it:   ItemCollection{},
		},
		{
			name: "add one item",
			col:  mockCollection(),
			it: ItemCollection{
				Object{ID: ID("grrr")},
			},
		},
		{
			name: "add multiple items",
			col:  mockCollection(),
			it: ItemCollection{
				Object{ID: ID("grrr")},
				Activity{ID: ID("one")},
				Actor{ID: ID("jdoe")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				tt.col.Items = tt.col.Items[:0]
				tt.col.TotalItems = 0
			}()
			if err := tt.col.Append(tt.it...); (err != nil) && errors.Is(err, tt.wantErr) {
				t.Errorf("Append() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.col.TotalItems != uint(len(tt.it)) {
				t.Errorf("Post Append() %T TotalItems %d different than added count %d", tt.col, tt.col.TotalItems, len(tt.it))
			}
			for _, it := range tt.it {
				if !tt.col.Items.Contains(it) {
					t.Errorf("Post Append() unable to find %s in %T Items %v", tt.col, it.GetLink(), tt.col.Items)
				}
			}
		})
	}
}

func TestCollection_Collection(t *testing.T) {
	id := ID("test")

	c := &Collection{Type: CollectionType, ID: id}

	if !reflect.DeepEqual(c.Collection(), c.Items) {
		t.Errorf("Collection items should be equal %v %v", c.Collection(), c.Items)
	}
}

func TestCollection_GetID(t *testing.T) {
	id := ID("test")

	c := &Collection{Type: CollectionType, ID: id}

	if c.GetID() != id {
		t.Errorf("GetID should return %s, received %s", id, c.GetID())
	}
}

func TestCollection_GetLink(t *testing.T) {
	id := ID("test")
	link := IRI(id)

	c := &Collection{Type: CollectionType, ID: id}

	if c.GetLink() != link {
		t.Errorf("GetLink should return %q, received %q", link, c.GetLink())
	}
}

func TestCollection_GetType(t *testing.T) {
	id := ID("test")

	c := &Collection{Type: CollectionType, ID: id}

	if !c.Match(CollectionType) {
		t.Errorf("Collection Type should be %q, received %q", CollectionType, c.GetType())
	}
}

func TestCollection_UnmarshalJSON(t *testing.T) {
	c := Collection{}

	dataEmpty := []byte("{}")
	c.UnmarshalJSON(dataEmpty)
	if c.ID != "" {
		t.Errorf("Unmarshaled object should have empty ID, received %q", c.ID)
	}
	if HasTypes(c) {
		t.Errorf("Unmarshaled object should have empty Type, received %q", c.GetType())
	}
	if c.AttributedTo != nil {
		t.Errorf("Unmarshaled object should have empty AttributedTo, received %q", c.AttributedTo)
	}
	if len(c.Name) != 0 {
		t.Errorf("Unmarshaled object should have empty Name, received %q", c.Name)
	}
	if len(c.Summary) != 0 {
		t.Errorf("Unmarshaled object should have empty Summary, received %q", c.Summary)
	}
	if len(c.Content) != 0 {
		t.Errorf("Unmarshaled object should have empty Content, received %q", c.Content)
	}
	if c.TotalItems != 0 {
		t.Errorf("Unmarshaled object should have empty TotalItems, received %d", c.TotalItems)
	}
	if len(c.Items) > 0 {
		t.Errorf("Unmarshaled object should have empty Items, received %v", c.Items)
	}
	if c.URL != nil {
		t.Errorf("Unmarshaled object should have empty URL, received %v", c.URL)
	}
	if !c.Published.IsZero() {
		t.Errorf("Unmarshaled object should have empty Published, received %q", c.Published)
	}
	if !c.StartTime.IsZero() {
		t.Errorf("Unmarshaled object should have empty StartTime, received %q", c.StartTime)
	}
	if !c.Updated.IsZero() {
		t.Errorf("Unmarshaled object should have empty Updated, received %q", c.Updated)
	}
}

func TestCollection_Count(t *testing.T) {
	id := ID("test")

	c := &Collection{Type: CollectionType, ID: id}

	if c.TotalItems != 0 {
		t.Errorf("Empty object should have empty TotalItems, received %d", c.TotalItems)
	}
	if len(c.Items) > 0 {
		t.Errorf("Empty object should have empty Items, received %v", c.Items)
	}
	if c.Count() != uint(len(c.Items)) {
		t.Errorf("%T.Count() returned %d, expected %d", c, c.Count(), len(c.Items))
	}

	_ = c.Append(IRI("test"))
	if c.TotalItems != 1 {
		t.Errorf("Object should have %d TotalItems, received %d", 1, c.TotalItems)
	}
	if c.Count() != uint(len(c.Items)) {
		t.Errorf("%T.Count() returned %d, expected %d", c, c.Count(), len(c.Items))
	}
}

func TestCollection_Contains(t *testing.T) {
	tests := []struct {
		name   string
		items  ItemCollection
		search Item
		want   bool
	}{
		{
			name: "nil",
			want: false,
		},
		{
			name:   "iri not found",
			items:  ItemCollection{IRI("http://example.com/1"), IRI("http://example.com/2"), Object{ID: "http://example.com"}},
			search: IRI("http://example.com"),
			want:   false,
		},
		{
			name:   "iri found",
			items:  ItemCollection{IRI("http://example.com")},
			search: IRI("http://example.com"),
			want:   true,
		},
		{
			name:   "object not found",
			items:  ItemCollection{Object{ID: "http://example.com/1", Type: NoteType}},
			search: Object{ID: "http://example.com/1", Type: VideoType},
			want:   false,
		},
		{
			name:   "object found",
			items:  ItemCollection{Object{ID: "http://example.com"}},
			search: Object{ID: "http://example.com"},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Collection{Items: tt.items}
			if got := o.Contains(tt.search); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollection_MarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollection_ItemMatches(t *testing.T) {
	t.Skipf("TODO")
}

func TestToCollection(t *testing.T) {
	tests := []struct {
		name    string
		it      LinkOrIRI
		want    *Collection
		wantErr error
	}{
		{
			name: "empty",
		},
		{
			name: "Valid Collection",
			it:   Collection{ID: "test", Type: CollectionType},
			want: &Collection{ID: "test", Type: CollectionType},
		},
		{
			name: "Valid *Collection",
			it:   &Collection{ID: "test", Type: CollectionType},
			want: &Collection{ID: "test", Type: CollectionType},
		},
		{
			name: "Valid CollectionPage",
			it:   CollectionPage{ID: "test", Type: CollectionPageType},
			want: &Collection{ID: "test", Type: CollectionPageType},
		},
		{
			name: "Valid *CollectionPage",
			it:   &CollectionPage{ID: "test", Type: CollectionPageType},
			want: &Collection{ID: "test", Type: CollectionPageType},
		},
		{
			name: "Valid OrderedCollection",
			it:   OrderedCollection{ID: "test", Type: OrderedCollectionType},
			want: &Collection{ID: "test", Type: OrderedCollectionType},
		},
		{
			name: "Valid *OrderedCollection",
			it:   &OrderedCollection{ID: "test", Type: OrderedCollectionType},
			want: &Collection{ID: "test", Type: OrderedCollectionType},
		},
		{
			name: "Valid OrderedCollectionPage",
			it:   OrderedCollectionPage{ID: "test", Type: OrderedCollectionPageType},
			want: &Collection{ID: "test", Type: OrderedCollectionPageType},
		},
		{
			name: "Valid *OrderedCollectionPage",
			it:   &OrderedCollectionPage{ID: "test", Type: OrderedCollectionPageType},
			want: &Collection{ID: "test", Type: OrderedCollectionPageType},
		},
		{
			name:    "IRI",
			it:      IRI("https://example.com"),
			wantErr: ErrorInvalidType[Collection](IRI("")),
		},
		{
			name:    "IRIs",
			it:      IRIs{IRI("https://example.com")},
			wantErr: ErrorInvalidType[Collection](IRIs{}),
		},
		{
			name:    "ItemCollection",
			it:      ItemCollection{},
			wantErr: ErrorInvalidType[Collection](ItemCollection{}),
		},
		{
			name:    "Object",
			it:      &Object{ID: "test", Type: ArticleType},
			wantErr: ErrorInvalidType[Collection](&Object{}),
		},
		{
			name:    "Activity",
			it:      &Activity{ID: "test", Type: CreateType},
			wantErr: ErrorInvalidType[Collection](&Activity{}),
		},
		{
			name:    "IntransitiveActivity",
			it:      &IntransitiveActivity{ID: "test", Type: ArriveType},
			wantErr: ErrorInvalidType[Collection](&IntransitiveActivity{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToCollection(tt.it)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("ToCollection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ToCollection() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestCollection_Equals(t *testing.T) {
	tests := []struct {
		name   string
		fields Collection
		item   Item
		want   bool
	}{
		{
			name: "collection with items in different order",
			fields: Collection{
				ID:    "https://example.com/1",
				Type:  CollectionType,
				First: IRI("https://example.com/1?first"),
				Items: ItemCollection{
					Object{ID: "https://example.com/1/1", Type: NoteType},
					Object{ID: "https://example.com/1/3", Type: ImageType},
				},
			},
			item: &Collection{
				ID:    "https://example.com/1",
				Type:  CollectionType,
				First: IRI("https://example.com/1?first"),
				Items: ItemCollection{
					Object{ID: "https://example.com/1/3", Type: ImageType},
					Object{ID: "https://example.com/1/1", Type: NoteType},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.Equals(tt.item); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_Remove(t *testing.T) {
	tests := []struct {
		name      string
		col       Collection
		toRemove  ItemCollection
		remaining ItemCollection
	}{
		{
			name: "empty",
			col:  mockCollection(),
		},
		{
			name:     "remove one item from empty collection",
			col:      mockCollection(),
			toRemove: ItemCollection{Object{ID: ID("grrr")}},
		},
		{
			name: "remove all from collection",
			col: mockCollection(
				Object{ID: ID("grrr")},
				Activity{ID: ID("one")},
				Actor{ID: ID("jdoe")},
			),
			toRemove: ItemCollection{
				Object{ID: ID("grrr")},
				Activity{ID: ID("one")},
				Actor{ID: ID("jdoe")},
			},
			remaining: ItemCollection{},
		},
		{
			name:      "empty_collection_non_nil_item",
			col:       mockCollection(),
			toRemove:  ItemCollection{&Object{}},
			remaining: ItemCollection{},
		},
		{
			name:      "non_empty_collection_nil_item",
			col:       mockCollection(&Object{ID: "test"}),
			toRemove:  nil,
			remaining: ItemCollection{&Object{ID: "test"}},
		},
		{
			name:     "non_empty_collection_non_contained_item_empty_ID",
			col:      mockCollection(&Object{ID: "test"}),
			toRemove: ItemCollection{&Object{}},
			remaining: ItemCollection{
				&Object{ID: "test"},
			},
		},
		{
			name:     "non_empty_collection_non_contained_item",
			col:      mockCollection(&Object{ID: "test"}),
			toRemove: ItemCollection{&Object{ID: "test123"}},
			remaining: ItemCollection{
				&Object{ID: "test"},
			},
		},
		{
			name:      "non_empty_collection_just_contained_item",
			col:       mockCollection(&Object{ID: "test"}),
			toRemove:  ItemCollection{&Object{ID: "test"}},
			remaining: ItemCollection{},
		},
		{
			name: "non_empty_collection_contained_item_first_pos",
			col: mockCollection(
				&Object{ID: "test"},
				&Object{ID: "test123"},
			),
			toRemove: ItemCollection{&Object{ID: "test"}},
			remaining: ItemCollection{
				&Object{ID: "test123"},
			},
		},
		{
			name: "non_empty_collection_contained_item_not_first_pos",
			col: mockCollection(
				&Object{ID: "test123"},
				&Object{ID: "test"},
				&Object{ID: "test321"},
			),
			toRemove: ItemCollection{&Object{ID: "test"}},
			remaining: ItemCollection{
				&Object{ID: "test123"},
				&Object{ID: "test321"},
			},
		},
		{
			name: "non_empty_collection_contained_item_last_pos",
			col: mockCollection(
				&Object{ID: "test123"},
				&Object{ID: "test"},
			),
			toRemove: ItemCollection{&Object{ID: "test"}},
			remaining: ItemCollection{
				&Object{ID: "test123"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.col.Remove(tt.toRemove...)

			if tt.col.TotalItems != uint(len(tt.remaining)) {
				t.Errorf("Post Remove() %T TotalItems %d different than expected %d", tt.col, tt.col.TotalItems, len(tt.remaining))
			}
			for _, it := range tt.remaining {
				if !tt.col.Items.Contains(it) {
					t.Errorf("Post Remove() unable to find %s in %T Items %v", it.GetLink(), tt.col, tt.col.Items)
				}
			}
			for _, it := range tt.toRemove {
				if tt.col.Items.Contains(it) {
					t.Errorf("Post Remove() was still able to find %s in %T Items %v", it.GetLink(), tt.col, tt.col.Items)
				}
			}
		})
	}
}

func TestCollection_Clean(t *testing.T) {
	tests := []struct {
		name string
		o    Collection
		want Item
	}{
		{
			name: "empty",
			o:    Collection{},
			want: &Collection{},
		},
		{
			name: "has Bto",
			o:    Collection{Type: CollectionType, Bto: ItemCollection{IRI("http://example.com")}},
			want: &Collection{Type: CollectionType},
		},
		{
			name: "has BCC",
			o:    Collection{Type: CollectionType, BCC: ItemCollection{IRI("http://example.com")}},
			want: &Collection{Type: CollectionType},
		},
		{
			name: "audience has BCC",
			o:    Collection{Type: CollectionType, Audience: ItemCollection{&Collection{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &Collection{Type: CollectionType, Audience: ItemCollection{&Collection{}}},
		},
		{
			name: "attachment has BCC",
			o:    Collection{Type: CollectionType, Attachment: &Collection{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Collection{Type: CollectionType, Attachment: &Collection{}},
		},
		{
			name: "icon has BCC",
			o:    Collection{Type: CollectionType, Icon: &Collection{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Collection{Type: CollectionType, Icon: &Collection{}},
		},
		{
			name: "image has BCC",
			o:    Collection{Type: CollectionType, Image: &Collection{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Collection{Type: CollectionType, Image: &Collection{}},
		},
		{
			name: "context has BCC",
			o:    Collection{Type: CollectionType, Context: &Collection{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Collection{Type: CollectionType, Context: &Collection{}},
		},
		{
			name: "generator has BCC",
			o:    Collection{Type: CollectionType, Generator: &Collection{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Collection{Type: CollectionType, Generator: &Collection{}},
		},
		{
			name: "attributedTo has BCC",
			o:    Collection{Type: CollectionType, AttributedTo: &Collection{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Collection{Type: CollectionType, AttributedTo: &Collection{}},
		},
		{
			name: "preview has BCC",
			o:    Collection{Type: CollectionType, Preview: &Collection{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Collection{Type: CollectionType, Preview: &Collection{}},
		},
		{
			name: "tag has BCC",
			o:    Collection{Type: CollectionType, Tag: ItemCollection{&Collection{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &Collection{Type: CollectionType, Tag: ItemCollection{&Collection{}}},
		},
		{
			name: "items has BCC",
			o:    Collection{Type: CollectionType, Items: ItemCollection{&Collection{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &Collection{Type: CollectionType, Items: ItemCollection{&Collection{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Clean(); !cmp.Equal(got, tt.want, EquateItems) {
				t.Errorf("Clean() = %s", cmp.Diff(tt.want, got, EquateItems))
			}
		})
	}
}

func TestCopyCollectionProperties(t *testing.T) {
	type args struct {
		to   *Collection
		from *Collection
	}
	tests := []struct {
		name    string
		args    args
		want    *Collection
		wantErr error
	}{
		{
			name:    "nil",
			args:    args{},
			want:    nil,
			wantErr: nil,
		},
		{
			name: "local properties",
			args: args{
				to: &Collection{},
				from: &Collection{
					Items:      ItemCollection{IRI("one")},
					TotalItems: 1,
					First:      IRI("http://example.com/first"),
					Last:       IRI("http://example.com/last"),
					Current:    IRI("http://example.com/current"),
				},
			},
			want: &Collection{
				Items:      ItemCollection{IRI("one")},
				TotalItems: 1,
				First:      IRI("http://example.com/first"),
				Last:       IRI("http://example.com/last"),
				Current:    IRI("http://example.com/current"),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CopyCollectionProperties(tt.args.to, tt.args.from)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("CopyCollectionProperties() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("CopyCollectionProperties() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestCollection_Recipients(t *testing.T) {
	type fields struct {
		Audience ItemCollection
		To       ItemCollection
		Bto      ItemCollection
		CC       ItemCollection
		BCC      ItemCollection
	}
	tests := []struct {
		name   string
		fields fields
		want   ItemCollection
	}{
		{
			name:   "nil",
			fields: fields{},
			want:   nil,
		},
		{
			name: "audience",
			fields: fields{
				Audience: ItemCollection{IRI("http://example.com/aud")},
			},
			want: ItemCollection{IRI("http://example.com/aud")},
		},
		{
			name: "to",
			fields: fields{
				To: ItemCollection{IRI("http://example.com/to")},
			},
			want: ItemCollection{IRI("http://example.com/to")},
		},
		{
			name: "bto",
			fields: fields{
				Bto: ItemCollection{IRI("http://example.com/bto")},
			},
			want: ItemCollection{IRI("http://example.com/bto")},
		},
		{
			name: "cc",
			fields: fields{
				CC: ItemCollection{IRI("http://example.com/cc")},
			},
			want: ItemCollection{IRI("http://example.com/cc")},
		},
		{
			name: "bcc",
			fields: fields{
				BCC: ItemCollection{IRI("http://example.com/bcc")},
			},
			want: ItemCollection{IRI("http://example.com/bcc")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Collection{
				Audience: tt.fields.Audience,
				To:       tt.fields.To,
				Bto:      tt.fields.Bto,
				CC:       tt.fields.CC,
				BCC:      tt.fields.BCC,
			}
			if got := c.Recipients(); !cmp.Equal(got, tt.want) {
				t.Errorf("Recipients() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestCollection_ItemsMatch(t *testing.T) {
	tests := []struct {
		name   string
		items  ItemCollection
		search ItemCollection
		want   bool
	}{
		{
			name: "nil",
			want: false,
		},
		{
			name:   "iri not found",
			items:  ItemCollection{IRI("http://example.com/1"), IRI("http://example.com/2"), Object{ID: "http://example.com"}},
			search: ItemCollection{IRI("http://example.com")},
			want:   false,
		},
		{
			name:   "iri found",
			items:  ItemCollection{IRI("http://example.com")},
			search: ItemCollection{IRI("http://example.com")},
			want:   true,
		},
		{
			name:   "object not found",
			items:  ItemCollection{Object{ID: "http://example.com/1", Type: NoteType}},
			search: ItemCollection{Object{ID: "http://example.com/1", Type: VideoType}},
			want:   false,
		},
		{
			name:   "object found",
			items:  ItemCollection{Object{ID: "http://example.com"}},
			search: ItemCollection{Object{ID: "http://example.com"}},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Collection{Items: tt.items}
			if got := o.ItemsMatch(tt.search...); got != tt.want {
				t.Errorf("ItemsMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
