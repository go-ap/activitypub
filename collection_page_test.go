package activitypub

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-ap/errors"
	"github.com/google/go-cmp/cmp"
)

func mockCollectionPage(items ...Item) CollectionPage {
	cc := CollectionPage{
		ID:   IRIf("https://example.com", Inbox),
		Type: CollectionPageType,
	}
	if len(items) == 0 {
		cc.Items = make(ItemCollection, 0)
	} else {
		cc.Items = items
		cc.TotalItems = uint(len(items))
	}
	return cc
}

func TestCollectionPage_Append(t *testing.T) {
	tests := []struct {
		name    string
		col     CollectionPage
		it      []Item
		wantErr error
	}{
		{
			name: "empty",
			col:  mockCollectionPage(),
			it:   ItemCollection{},
		},
		{
			name: "add one item",
			col:  mockCollectionPage(),
			it: ItemCollection{
				Object{ID: ID("grrr")},
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

func TestCollectionPage_UnmarshalJSON(t *testing.T) {
	p := CollectionPage{}

	dataEmpty := []byte("{}")
	p.UnmarshalJSON(dataEmpty)
	if p.ID != "" {
		t.Errorf("Unmarshaled object should have empty ID, received %q", p.ID)
	}
	if HasTypes(p) {
		t.Errorf("Unmarshaled object should have empty Type, received %q", p.GetType())
	}
	if p.AttributedTo != nil {
		t.Errorf("Unmarshaled object should have empty AttributedTo, received %q", p.AttributedTo)
	}
	if len(p.Name) != 0 {
		t.Errorf("Unmarshaled object should have empty Name, received %q", p.Name)
	}
	if len(p.Summary) != 0 {
		t.Errorf("Unmarshaled object should have empty Summary, received %q", p.Summary)
	}
	if len(p.Content) != 0 {
		t.Errorf("Unmarshaled object should have empty Content, received %q", p.Content)
	}
	if p.TotalItems != 0 {
		t.Errorf("Unmarshaled object should have empty TotalItems, received %d", p.TotalItems)
	}
	if len(p.Items) > 0 {
		t.Errorf("Unmarshaled object should have empty Items, received %v", p.Items)
	}
	if p.URL != nil {
		t.Errorf("Unmarshaled object should have empty URL, received %v", p.URL)
	}
	if !p.Published.IsZero() {
		t.Errorf("Unmarshaled object should have empty Published, received %q", p.Published)
	}
	if !p.StartTime.IsZero() {
		t.Errorf("Unmarshaled object should have empty StartTime, received %q", p.StartTime)
	}
	if !p.Updated.IsZero() {
		t.Errorf("Unmarshaled object should have empty Updated, received %q", p.Updated)
	}
	if p.PartOf != nil {
		t.Errorf("Unmarshaled object should have empty PartOf, received %q", p.PartOf)
	}
	if p.Current != nil {
		t.Errorf("Unmarshaled object should have empty Current, received %q", p.Current)
	}
	if p.First != nil {
		t.Errorf("Unmarshaled object should have empty First, received %q", p.First)
	}
	if p.Last != nil {
		t.Errorf("Unmarshaled object should have empty Last, received %q", p.Last)
	}
	if p.Next != nil {
		t.Errorf("Unmarshaled object should have empty Next, received %q", p.Next)
	}
	if p.Prev != nil {
		t.Errorf("Unmarshaled object should have empty Prev, received %q", p.Prev)
	}
}

func TestCollectionPage_Collection(t *testing.T) {
	id := ID("test")

	p := &CollectionPage{Type: CollectionPageType, ID: id}

	if !reflect.DeepEqual(p.Collection(), p.Items) {
		t.Errorf("Collection items should be equal %v %v", p.Collection(), p.Items)
	}
}

func TestCollectionPage_Count(t *testing.T) {
	id := ID("test")

	c := &Collection{Type: CollectionType, ID: id}
	p := &CollectionPage{Type: CollectionPageType, ID: id}

	if p.TotalItems != 0 {
		t.Errorf("Object should have empty TotalItems, received %d", p.TotalItems)
	}
	if len(p.Items) > 0 {
		t.Errorf("Empty object should have empty Items, received %v", p.Items)
	}
	if p.Count() != uint(len(p.Items)) {
		t.Errorf("%T.Count() returned %d, expected %d", c, p.Count(), len(p.Items))
	}

	_ = p.Append(IRI("test"))
	if p.TotalItems != 1 {
		t.Errorf("Empty object should have %d TotalItems, received %d", 1, p.TotalItems)
	}
	if p.Count() != uint(len(p.Items)) {
		t.Errorf("%T.Count() returned %d, expected %d", c, p.Count(), len(p.Items))
	}
}

func TestToCollectionPage(t *testing.T) {
	tests := []struct {
		name    string
		it      LinkOrIRI
		want    *CollectionPage
		wantErr error
	}{
		{
			name: "empty",
		},
		{
			name: "Valid CollectionPage",
			it:   CollectionPage{ID: "test", Type: CollectionPageType},
			want: &CollectionPage{ID: "test", Type: CollectionPageType},
		},
		{
			name: "Valid *CollectionPage",
			it:   &CollectionPage{ID: "test", Type: CollectionPageType},
			want: &CollectionPage{ID: "test", Type: CollectionPageType},
		},
		{
			name:    "Valid Collection",
			it:      Collection{ID: "test", Type: CollectionType},
			wantErr: ErrorInvalidType[CollectionPage](Collection{}),
		},
		{
			name:    "Valid *Collection",
			it:      &Collection{ID: "test", Type: CollectionType},
			wantErr: ErrorInvalidType[CollectionPage](new(Collection)),
		},
		{
			name:    "Valid OrderedCollection",
			it:      OrderedCollection{ID: "test", Type: OrderedCollectionType},
			wantErr: ErrorInvalidType[CollectionPage](OrderedCollection{}),
		},
		{
			name:    "Valid *OrderedCollection",
			it:      &OrderedCollection{ID: "test", Type: OrderedCollectionType},
			wantErr: ErrorInvalidType[CollectionPage](new(OrderedCollection)),
		},
		{
			name: "Valid OrderedCollectionPage",
			it:   OrderedCollectionPage{ID: "test", Type: OrderedCollectionPageType},
			want: &CollectionPage{ID: "test", Type: OrderedCollectionPageType},
		},
		{
			name: "Valid *OrderedCollectionPage",
			it:   &OrderedCollectionPage{ID: "test", Type: OrderedCollectionPageType},
			want: &CollectionPage{ID: "test", Type: OrderedCollectionPageType},
		},
		{
			name:    "IRI",
			it:      IRI("https://example.com"),
			wantErr: ErrorInvalidType[CollectionPage](IRI("")),
		},
		{
			name:    "IRIs",
			it:      IRIs{IRI("https://example.com")},
			wantErr: ErrorInvalidType[CollectionPage](IRIs{}),
		},
		{
			name:    "ItemCollection",
			it:      ItemCollection{},
			wantErr: ErrorInvalidType[CollectionPage](ItemCollection{}),
		},
		{
			name:    "Object",
			it:      &Object{ID: "test", Type: ArticleType},
			wantErr: ErrorInvalidType[CollectionPage](&Object{}),
		},
		{
			name:    "Activity",
			it:      &Activity{ID: "test", Type: CreateType},
			wantErr: ErrorInvalidType[CollectionPage](&Activity{}),
		},
		{
			name:    "IntransitiveActivity",
			it:      &IntransitiveActivity{ID: "test", Type: ArriveType},
			wantErr: ErrorInvalidType[CollectionPage](&IntransitiveActivity{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToCollectionPage(tt.it)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("ToCollectionPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ToCollectionPage() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestCollectionPage_Contains(t *testing.T) {
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
			o := CollectionPage{Items: tt.items}
			if got := o.Contains(tt.search); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollectionPage_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollectionPage_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollectionPage_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollectionPage_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollectionPage_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollectionPage_IsObject(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollectionPage_MarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollectionPage_ItemMatches(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollectionPage_Remove(t *testing.T) {
	tests := []struct {
		name      string
		col       CollectionPage
		toRemove  ItemCollection
		remaining ItemCollection
	}{
		{
			name: "empty",
			col:  mockCollectionPage(),
		},
		{
			name:     "remove one item from empty collection",
			col:      mockCollectionPage(),
			toRemove: ItemCollection{Object{ID: ID("grrr")}},
		},
		{
			name: "remove all from collection",
			col: mockCollectionPage(
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
			col:       mockCollectionPage(),
			toRemove:  ItemCollection{&Object{}},
			remaining: ItemCollection{},
		},
		{
			name:      "non_empty_collection_nil_item",
			col:       mockCollectionPage(&Object{ID: "test"}),
			toRemove:  nil,
			remaining: ItemCollection{&Object{ID: "test"}},
		},
		{
			name:     "non_empty_collection_non_contained_item_empty_ID",
			col:      mockCollectionPage(&Object{ID: "test"}),
			toRemove: ItemCollection{&Object{}},
			remaining: ItemCollection{
				&Object{ID: "test"},
			},
		},
		{
			name:     "non_empty_collection_non_contained_item",
			col:      mockCollectionPage(&Object{ID: "test"}),
			toRemove: ItemCollection{&Object{ID: "test123"}},
			remaining: ItemCollection{
				&Object{ID: "test"},
			},
		},
		{
			name:      "non_empty_collection_just_contained_item",
			col:       mockCollectionPage(&Object{ID: "test"}),
			toRemove:  ItemCollection{&Object{ID: "test"}},
			remaining: ItemCollection{},
		},
		{
			name: "non_empty_collection_contained_item_first_pos",
			col: mockCollectionPage(
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
			col: mockCollectionPage(
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
			col: mockCollectionPage(
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

func TestCollectionPage_Clean(t *testing.T) {
	tests := []struct {
		name string
		o    CollectionPage
		want Item
	}{
		{
			name: "empty",
			o:    CollectionPage{},
			want: &CollectionPage{},
		},
		{
			name: "has Bto",
			o:    CollectionPage{Type: CollectionPageType, Bto: ItemCollection{IRI("http://example.com")}},
			want: &CollectionPage{Type: CollectionPageType},
		},
		{
			name: "has BCC",
			o:    CollectionPage{Type: CollectionPageType, BCC: ItemCollection{IRI("http://example.com")}},
			want: &CollectionPage{Type: CollectionPageType},
		},
		{
			name: "audience has BCC",
			o:    CollectionPage{Type: CollectionPageType, Audience: ItemCollection{&CollectionPage{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &CollectionPage{Type: CollectionPageType, Audience: ItemCollection{&CollectionPage{}}},
		},
		{
			name: "attachment has BCC",
			o:    CollectionPage{Type: CollectionPageType, Attachment: &CollectionPage{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &CollectionPage{Type: CollectionPageType, Attachment: &CollectionPage{}},
		},
		{
			name: "icon has BCC",
			o:    CollectionPage{Type: CollectionPageType, Icon: &CollectionPage{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &CollectionPage{Type: CollectionPageType, Icon: &CollectionPage{}},
		},
		{
			name: "image has BCC",
			o:    CollectionPage{Type: CollectionPageType, Image: &CollectionPage{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &CollectionPage{Type: CollectionPageType, Image: &CollectionPage{}},
		},
		{
			name: "context has BCC",
			o:    CollectionPage{Type: CollectionPageType, Context: &CollectionPage{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &CollectionPage{Type: CollectionPageType, Context: &CollectionPage{}},
		},
		{
			name: "generator has BCC",
			o:    CollectionPage{Type: CollectionPageType, Generator: &CollectionPage{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &CollectionPage{Type: CollectionPageType, Generator: &CollectionPage{}},
		},
		{
			name: "attributedTo has BCC",
			o:    CollectionPage{Type: CollectionPageType, AttributedTo: &CollectionPage{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &CollectionPage{Type: CollectionPageType, AttributedTo: &CollectionPage{}},
		},
		{
			name: "preview has BCC",
			o:    CollectionPage{Type: CollectionPageType, Preview: &CollectionPage{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &CollectionPage{Type: CollectionPageType, Preview: &CollectionPage{}},
		},
		{
			name: "tag has BCC",
			o:    CollectionPage{Type: CollectionPageType, Tag: ItemCollection{&CollectionPage{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &CollectionPage{Type: CollectionPageType, Tag: ItemCollection{&CollectionPage{}}},
		},
		{
			name: "items has BCC",
			o:    CollectionPage{Type: CollectionPageType, Items: ItemCollection{&CollectionPage{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &CollectionPage{Type: CollectionPageType, Items: ItemCollection{&CollectionPage{}}},
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

func TestCollectionPage_Equals(t *testing.T) {
	type fields struct {
		ID           ID
		Type         Typer
		Name         NaturalLanguageValues
		Attachment   Item
		AttributedTo Item
		Audience     ItemCollection
		Content      NaturalLanguageValues
		Context      Item
		MediaType    MimeType
		EndTime      time.Time
		Generator    Item
		Icon         Item
		Image        Item
		InReplyTo    Item
		Location     Item
		Preview      Item
		Published    time.Time
		Replies      Item
		StartTime    time.Time
		Summary      NaturalLanguageValues
		Tag          Item
		Updated      time.Time
		URL          Item
		To           ItemCollection
		Bto          ItemCollection
		CC           ItemCollection
		BCC          ItemCollection
		Duration     time.Duration
		Likes        Item
		Shares       Item
		Source       Source
		Current      Item
		First        Item
		Last         Item
		TotalItems   uint
		Items        ItemCollection
		PartOf       Item
		Next         Item
		Prev         Item
		StartIndex   uint
	}

	tests := []struct {
		name   string
		fields fields
		with   Item
		want   bool
	}{
		{
			name:   "nil",
			fields: fields{},
			with:   Item(nil),
			want:   false,
		},
		{
			name:   "not-a-CollectionPage",
			fields: fields{ID: "http://example.com"},
			with:   &Object{ID: "http://example.com"},
			want:   false,
		},
		{
			name:   "CollectionPage-same-ID",
			fields: fields{ID: "http://example.com"},
			with:   &CollectionPage{ID: "http://example.com"},
			want:   true,
		},
		{
			name:   "CollectionPage-different-ID",
			fields: fields{ID: "http://example.com"},
			with:   &CollectionPage{ID: "http://example.com/1"},
			want:   false,
		},
		{
			name:   "CollectionPage-same-First",
			fields: fields{First: IRI("http://example.com")},
			with:   &CollectionPage{First: IRI("http://example.com")},
			want:   true,
		},
		{
			name:   "CollectionPage-different-First",
			fields: fields{First: IRI("http://example.com")},
			with:   &CollectionPage{First: IRI("http://example.com/1")},
			want:   false,
		},
		{
			name:   "CollectionPage-nil-First",
			fields: fields{First: IRI("http://example.com")},
			with:   &CollectionPage{},
			want:   false,
		},
		{
			name:   "CollectionPage-same-Last",
			fields: fields{Last: IRI("http://example.com")},
			with:   &CollectionPage{Last: IRI("http://example.com")},
			want:   true,
		},
		{
			name:   "CollectionPage-different-Last",
			fields: fields{Last: IRI("http://example.com")},
			with:   &CollectionPage{Last: IRI("http://example.com/2")},
			want:   false,
		},
		{
			name:   "CollectionPage-nil-Last",
			fields: fields{Last: IRI("http://example.com")},
			with:   &CollectionPage{},
			want:   false,
		},
		{
			name:   "CollectionPage-same-Next",
			fields: fields{Next: IRI("http://example.com")},
			with:   &CollectionPage{Next: IRI("http://example.com")},
			want:   true,
		},
		{
			name:   "CollectionPage-diff-Next",
			fields: fields{Next: IRI("http://example.com")},
			with:   &CollectionPage{Next: IRI("http://example.com/2")},
			want:   false,
		},
		{
			name:   "CollectionPage-nil-Next",
			fields: fields{Next: IRI("http://example.com")},
			with:   &CollectionPage{},
			want:   false,
		},
		{
			name:   "CollectionPage-same-Prev",
			fields: fields{Prev: IRI("http://example.com")},
			with:   &CollectionPage{Prev: IRI("http://example.com")},
			want:   true,
		},
		{
			name:   "CollectionPage-diff-Prev",
			fields: fields{Prev: IRI("http://example.com")},
			with:   &CollectionPage{Prev: IRI("http://example.com/2")},
			want:   false,
		},
		{
			name:   "CollectionPage-nil-Prev",
			fields: fields{Prev: IRI("http://example.com")},
			with:   &CollectionPage{},
			want:   false,
		},
		{
			name:   "CollectionPage-same-PartOf",
			fields: fields{PartOf: IRI("http://example.com")},
			with:   &CollectionPage{PartOf: IRI("http://example.com")},
			want:   true,
		},
		{
			name:   "CollectionPage-diff-PartOf",
			fields: fields{PartOf: IRI("http://example.com")},
			with:   &CollectionPage{PartOf: IRI("http://example.com/2")},
			want:   false,
		},
		{
			name:   "CollectionPage-nil-PartOf",
			fields: fields{PartOf: IRI("http://example.com")},
			with:   &CollectionPage{},
			want:   false,
		},
		{
			name:   "CollectionPage-same-Current",
			fields: fields{Current: IRI("http://example.com")},
			with:   &CollectionPage{Current: IRI("http://example.com")},
			want:   true,
		},
		{
			name:   "CollectionPage-diff-Current",
			fields: fields{Current: IRI("http://example.com")},
			with:   &CollectionPage{Current: IRI("http://example.com/2")},
			want:   false,
		},
		{
			name:   "CollectionPage-nil-Current",
			fields: fields{Current: IRI("http://example.com")},
			with:   &CollectionPage{},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := CollectionPage{
				ID:           tt.fields.ID,
				Type:         tt.fields.Type,
				Name:         tt.fields.Name,
				Attachment:   tt.fields.Attachment,
				AttributedTo: tt.fields.AttributedTo,
				Audience:     tt.fields.Audience,
				Content:      tt.fields.Content,
				Context:      tt.fields.Context,
				MediaType:    tt.fields.MediaType,
				EndTime:      tt.fields.EndTime,
				Generator:    tt.fields.Generator,
				Icon:         tt.fields.Icon,
				Image:        tt.fields.Image,
				InReplyTo:    tt.fields.InReplyTo,
				Location:     tt.fields.Location,
				Preview:      tt.fields.Preview,
				Published:    tt.fields.Published,
				Replies:      tt.fields.Replies,
				StartTime:    tt.fields.StartTime,
				Summary:      tt.fields.Summary,
				Tag:          tt.fields.Tag,
				Updated:      tt.fields.Updated,
				URL:          tt.fields.URL,
				To:           tt.fields.To,
				Bto:          tt.fields.Bto,
				CC:           tt.fields.CC,
				BCC:          tt.fields.BCC,
				Duration:     tt.fields.Duration,
				Likes:        tt.fields.Likes,
				Shares:       tt.fields.Shares,
				Source:       tt.fields.Source,
				Current:      tt.fields.Current,
				First:        tt.fields.First,
				Last:         tt.fields.Last,
				TotalItems:   tt.fields.TotalItems,
				Items:        tt.fields.Items,
				PartOf:       tt.fields.PartOf,
				Next:         tt.fields.Next,
				Prev:         tt.fields.Prev,
			}
			if got := o.Equals(tt.with); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollectionPage_ItemsMatch(t *testing.T) {
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
			o := CollectionPage{Items: tt.items}
			if got := o.ItemsMatch(tt.search...); got != tt.want {
				t.Errorf("ItemsMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyCollectionPageProperties(t *testing.T) {
	type args struct {
		to   *CollectionPage
		from *CollectionPage
	}
	tests := []struct {
		name    string
		args    args
		want    *CollectionPage
		wantErr error
	}{
		{
			name: "nil",
			args: args{},
			want: nil,
		},
		{
			name: "local properties",
			args: args{
				to: &CollectionPage{},
				from: &CollectionPage{
					Items:      ItemCollection{IRI("one")},
					TotalItems: 1,
					PartOf:     IRI("http://example.com/partOf"),
					First:      IRI("http://example.com/first"),
					Last:       IRI("http://example.com/last"),
					Current:    IRI("http://example.com/current"),
					Next:       IRI("http://example.com/next"),
					Prev:       IRI("http://example.com/prev"),
				},
			},
			want: &CollectionPage{
				Items:      ItemCollection{IRI("one")},
				TotalItems: 1,
				PartOf:     IRI("http://example.com/partOf"),
				First:      IRI("http://example.com/first"),
				Last:       IRI("http://example.com/last"),
				Current:    IRI("http://example.com/current"),
				Next:       IRI("http://example.com/next"),
				Prev:       IRI("http://example.com/prev"),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CopyCollectionPageProperties(tt.args.to, tt.args.from)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("CopyCollectionPageProperties() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("CopyCollectionPageProperties() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestCollectionPage_Recipients(t *testing.T) {
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
			c := &CollectionPage{
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

func TestCollectionPage_UnmarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    CollectionPage
		wantErr error
	}{
		{
			name:    "nil",
			data:    nil,
			wantErr: errors.NotImplementedf("Binary functionality not implemented"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := CollectionPage{}
			if err := l.UnmarshalBinary(tt.data); !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("UnmarshalBinary() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
			}
			if !cmp.Equal(l, tt.want) {
				t.Errorf("UnmarshalBinary() got = %s", cmp.Diff(tt.want, l))
			}
		})
	}
}

func TestCollectionPage_MarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		sub     CollectionPage
		want    []byte
		wantErr error
	}{
		{
			name:    "nil",
			sub:     CollectionPage{},
			want:    nil,
			wantErr: errors.NotImplementedf("Binary functionality not implemented"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sub.MarshalBinary()
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("MarshalBinary() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("MarshalBinary() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
