package activitypub

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/valyala/fastjson"
)

func TestObjectNew(t *testing.T) {
	testValue := ID("test")
	testType := ArticleType

	o := ObjectNew(testType)
	o.ID = testValue

	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if !o.Matches(testType) {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.GetType(), testType)
	}

	n := ObjectNew(NilType)
	n.ID = testValue
	if n.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", n.ID, testValue)
	}
	if !n.Matches(ObjectType) {
		t.Errorf("APObject Type '%v' different than expected '%v'", n.GetType(), ObjectType)
	}
}

func TestObject_IsLink(t *testing.T) {
	o := ObjectNew(ObjectType)
	o.ID = "test"
	if o.IsLink() {
		t.Errorf("%#v should not be a valid link", o.Type)
	}
	m := ObjectNew(AcceptType)
	m.ID = "test"
	if m.IsLink() {
		t.Errorf("%#v should not be a valid link", m.Type)
	}
}

func TestObject_IsObject(t *testing.T) {
	o := ObjectNew(ObjectType)
	o.ID = "test"
	if !o.IsObject() {
		t.Errorf("%#v should be a valid object", o.Type)
	}
	m := ObjectNew(AcceptType)
	m.ID = "test"
	if !m.IsObject() {
		t.Errorf("%#v should be a valid object", m.Type)
	}
}

func TestObjectsArr_Append(t *testing.T) {
	d := make(ItemCollection, 0)

	val := Object{ID: ID("grrr")}

	_ = d.Append(val)

	if len(d) != 1 {
		t.Errorf("Objects array should have exactly an element")
	}
	if !reflect.DeepEqual(d[0], val) {
		t.Errorf("First item in object array does not match %q", val.ID)
	}
}

func TestRecipients(t *testing.T) {
	bob := PersonNew("bob")
	alice := PersonNew("alice")
	foo := OrganizationNew("foo")
	bar := GroupNew("bar")

	first := make(ItemCollection, 0)
	if len(first) != 0 {
		t.Errorf("Objects array should have exactly an element")
	}

	_ = first.Append(bob)
	_ = first.Append(alice)
	_ = first.Append(foo)
	_ = first.Append(bar)
	if len(first) != 4 {
		t.Errorf("Objects array should have exactly 4(four) elements, not %d", len(first))
	}

	_ = first.Append(bar)
	_ = first.Append(alice)
	_ = first.Append(foo)
	_ = first.Append(bob)
	if len(first) != 4 {
		t.Errorf("Objects array should have exactly 4(eight) elements, not %d", len(first))
	}

	ItemCollectionDeduplication(&first)
	if len(first) != 4 {
		t.Errorf("Objects array should have exactly 4(four) elements, not %d", len(first))
	}

	second := make(ItemCollection, 0)
	_ = second.Append(bar)
	_ = second.Append(foo)

	ItemCollectionDeduplication(&first, &second)
	if len(first) != 4 {
		t.Errorf("First Objects array should have exactly 8(eight) elements, not %d", len(first))
	}
	if len(second) != 0 {
		t.Errorf("Second Objects array should have exactly 0(zero) elements, not %d", len(second))
	}
}

func validateEmptyObject(o Object, t *testing.T) {
	if o.ID != "" {
		t.Errorf("Unmarshaled object %T should have empty ID, received %q", o, o.ID)
	}
	if HasTypes(o) {
		t.Errorf("Unmarshaled object %T should have empty Type, received %q", o, o.GetType())
	}
	if o.Matches(ActivityVocabularyTypes{NilType, NilType, ObjectType}...) {
		t.Errorf("Unmarshaled object %T should have empty Type, received %q", o, o.GetType())
	}
	if o.AttributedTo != nil {
		t.Errorf("Unmarshaled object %T should have empty AttributedTo, received %q", o, o.AttributedTo)
	}
	if len(o.Name) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Name, received %q", o, o.Name)
	}
	if len(o.Summary) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Summary, received %q", o, o.Summary)
	}
	if len(o.Content) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Content, received %q", o, o.Content)
	}
	if o.URL != nil {
		t.Errorf("Unmarshaled object %T should have empty URL, received %v", o, o.URL)
	}
	if o.Icon != nil {
		t.Errorf("Unmarshaled object %T should have empty Icon, received %v", o, o.Icon)
	}
	if o.Image != nil {
		t.Errorf("Unmarshaled object %T should have empty Image, received %v", o, o.Image)
	}
	if !o.Published.IsZero() {
		t.Errorf("Unmarshaled object %T should have empty Published, received %q", o, o.Published)
	}
	if !o.StartTime.IsZero() {
		t.Errorf("Unmarshaled object %T should have empty StartTime, received %q", o, o.StartTime)
	}
	if !o.Updated.IsZero() {
		t.Errorf("Unmarshaled object %T should have empty Updated, received %q", o, o.Updated)
	}
	if !o.EndTime.IsZero() {
		t.Errorf("Unmarshaled object %T should have empty EndTime, received %q", o, o.EndTime)
	}
	if o.Duration != 0 {
		t.Errorf("Unmarshaled object %T should have empty Duration, received %q", o, o.Duration)
	}
	if len(o.To) > 0 {
		t.Errorf("Unmarshaled object %T should have empty To, received %q", o, o.To)
	}
	if len(o.Bto) > 0 {
		t.Errorf("Unmarshaled object %T should have empty Bto, received %q", o, o.Bto)
	}
	if len(o.CC) > 0 {
		t.Errorf("Unmarshaled object %T should have empty CC, received %q", o, o.CC)
	}
	if len(o.BCC) > 0 {
		t.Errorf("Unmarshaled object %T should have empty BCC, received %q", o, o.BCC)
	}
	validateEmptySource(o.Source, t)
}

func validateEmptySource(s Source, t *testing.T) {
	if s.MediaType != "" {
		t.Errorf("Unmarshalled object %T should have empty Source.MediaType, received %q", s, s.MediaType)
	}
	if s.Content != nil {
		t.Errorf("Unmarshalled object %T should have empty Source.Content, received %q", s, s.Content)
	}
}

func TestObject_UnmarshalJSON(t *testing.T) {
	o := Object{}

	dataEmpty := []byte("{}")
	_ = o.UnmarshalJSON(dataEmpty)
	validateEmptyObject(o, t)
}

func TestMimeType_UnmarshalJSON(t *testing.T) {
	m := MimeType("")
	dataEmpty := []byte("")

	_ = m.UnmarshalJSON(dataEmpty)
	if m != "" {
		t.Errorf("Unmarshaled object %T should be an empty string, received %q", m, m)
	}
}

func TestMimeType_GobDecode(t *testing.T) {
	tests := []struct {
		name    string
		m       MimeType
		data    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			m:       "",
			data:    []byte{},
			wantErr: false,
		},
		{
			name:    "some mime-type",
			m:       "application/json",
			data:    gobValue([]byte("application/json")),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.GobDecode(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("GobDecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMimeType_GobEncode(t *testing.T) {
	tests := []struct {
		name    string
		m       MimeType
		want    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			m:       "",
			want:    []byte{},
			wantErr: false,
		},
		{
			name:    "some mime-type",
			m:       "application/json",
			want:    gobValue([]byte("application/json")),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.GobEncode()
			if (err != nil) != tt.wantErr {
				t.Errorf("GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GobEncode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLangRefValue_UnmarshalText(t *testing.T) {
	t.Skipf("TODO")
}

func TestLangRef_UnmarshalText(t *testing.T) {
	l := NilLangRef
	dataEmpty := []byte("")

	if _ = l.UnmarshalText(dataEmpty); l != NilLangRef {
		t.Errorf("Unmarshaled object %T should be an empty string, received %q", l, l)
	}
}

func TestObject_GetID(t *testing.T) {
	a := Object{}
	testVal := "crash$"
	a.ID = ID(testVal)
	if string(a.GetID()) != testVal {
		t.Errorf("%T should return %q, Received %q", a.GetID, testVal, a.GetID())
	}
}

func TestObject_GetLink(t *testing.T) {
	a := Object{}
	testVal := "crash$"
	a.ID = ID(testVal)
	if string(a.GetLink()) != testVal {
		t.Errorf("%T should return %q, Received %q", a.GetLink, testVal, a.GetLink())
	}
}

func TestObject_GetType(t *testing.T) {
	a := Object{}
	a.Type = ActorType
	if !a.Matches(ActorType) {
		t.Errorf("%T should return %q, Received %q", a.GetType(), ActorType, a.GetType())
	}
}

func TestToObject(t *testing.T) {
	tests := []struct {
		name    string
		it      LinkOrIRI
		want    *Object
		wantErr error
	}{
		{
			name: "empty",
		},
		{
			name: "Valid Object",
			it:   Object{ID: "test", Type: UpdateType},
			want: &Object{ID: "test", Type: UpdateType},
		},
		{
			name: "Valid *Object",
			it:   &Object{ID: "test", Type: CreateType},
			want: &Object{ID: "test", Type: CreateType},
		},
		{
			name: "Valid Place",
			it:   Place{ID: "test", Type: PlaceType},
			want: &Object{ID: "test", Type: PlaceType},
		},
		{
			name: "Valid *Place",
			it:   &Place{ID: "test", Type: PlaceType},
			want: &Object{ID: "test", Type: PlaceType},
		},
		{
			name: "Valid Profile",
			it:   Profile{ID: "test", Type: ProfileType},
			want: &Object{ID: "test", Type: ProfileType},
		},
		{
			name: "Valid *Profile",
			it:   &Profile{ID: "test", Type: ProfileType},
			want: &Object{ID: "test", Type: ProfileType},
		},
		{
			name: "Valid Relationship",
			it:   Relationship{ID: "test", Type: RelationshipType},
			want: &Object{ID: "test", Type: RelationshipType},
		},
		{
			name: "Valid *Relationship",
			it:   &Relationship{ID: "test", Type: RelationshipType},
			want: &Object{ID: "test", Type: RelationshipType},
		},
		{
			name: "Valid Tombstone",
			it:   Tombstone{ID: "test", Type: TombstoneType},
			want: &Object{ID: "test", Type: TombstoneType},
		},
		{
			name: "Valid *Tombstone",
			it:   &Tombstone{ID: "test", Type: TombstoneType},
			want: &Object{ID: "test", Type: TombstoneType},
		},
		{
			name: "Valid Activity",
			it:   &Activity{ID: "test", Type: CreateType},
			want: &Object{ID: "test", Type: CreateType},
		},
		{
			name: "Valid IntransitiveActivity",
			it:   &IntransitiveActivity{ID: "test", Type: ArriveType},
			want: &Object{ID: "test", Type: ArriveType},
		},
		{
			name: "Valid Question",
			it:   &Question{ID: "test", Type: QuestionType},
			want: &Object{ID: "test", Type: QuestionType},
		},
		{
			name: "Valid OrderedCollection",
			it:   OrderedCollection{ID: "test", Type: OrderedCollectionType},
			want: &Object{ID: "test", Type: OrderedCollectionType},
		},
		{
			name: "Valid *OrderedCollection",
			it:   &OrderedCollection{ID: "test", Type: OrderedCollectionType},
			want: &Object{ID: "test", Type: OrderedCollectionType},
		},
		{
			name: "Valid OrderedCollectionPage",
			it:   OrderedCollectionPage{ID: "test", Type: OrderedCollectionPageType},
			want: &Object{ID: "test", Type: OrderedCollectionPageType},
		},
		{
			name: "Valid *OrderedCollectionPage",
			it:   &OrderedCollectionPage{ID: "test", Type: OrderedCollectionPageType},
			want: &Object{ID: "test", Type: OrderedCollectionPageType},
		},
		{
			name: "Valid OrderedCollection",
			it:   OrderedCollection{ID: "test", Type: OrderedCollectionType},
			want: &Object{ID: "test", Type: OrderedCollectionType},
		},
		{
			name: "Valid *OrderedCollection",
			it:   &OrderedCollection{ID: "test", Type: OrderedCollectionType},
			want: &Object{ID: "test", Type: OrderedCollectionType},
		},
		{
			name: "Valid OrderedCollectionPage",
			it:   OrderedCollectionPage{ID: "test", Type: OrderedCollectionPageType},
			want: &Object{ID: "test", Type: OrderedCollectionPageType},
		},
		{
			name: "Valid *OrderedCollectionPage",
			it:   &OrderedCollectionPage{ID: "test", Type: OrderedCollectionPageType},
			want: &Object{ID: "test", Type: OrderedCollectionPageType},
		},
		{
			name:    "IRI",
			it:      IRI("https://example.com"),
			wantErr: ErrorInvalidType[Object](IRI("")),
		},
		{
			name:    "IRIs",
			it:      IRIs{IRI("https://example.com")},
			wantErr: ErrorInvalidType[Object](IRIs{}),
		},
		{
			name:    "ItemCollection",
			it:      ItemCollection{},
			wantErr: ErrorInvalidType[Object](ItemCollection{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToObject(tt.it)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("ToObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ToObject() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestToObject1(t *testing.T) {
	tests := []struct {
		name    string
		arg     Item
		want    Item
		wantErr bool
	}{
		{
			name: "Actor with ID, type, w/o extra properties",
			arg:  &Actor{ID: "https://example.com", Type: PersonType},
			want: &Object{ID: "https://example.com", Type: PersonType},
		},
		{
			name: "Actor with ID, type, w/ extra properties",
			arg: &Actor{ID: "https://example.com", Type: PersonType, Endpoints: &Endpoints{
				OauthAuthorizationEndpoint: IRI("https://example.com/oauth"),
			}},
			want: &Object{ID: "https://example.com", Type: PersonType},
		},
		{
			name: "Place w/o extra properties",
			arg:  &Place{ID: "https://example.com", Type: PlaceType},
			want: &Object{ID: "https://example.com", Type: PlaceType},
		},
		{
			name: "Place w/ extra properties",
			arg:  &Place{ID: "https://example.com", Type: PlaceType, Accuracy: 0.22, Altitude: 66.6},
			want: &Object{ID: "https://example.com", Type: PlaceType},
		},
		{
			name: "Profile w/o extra properties",
			arg:  &Profile{ID: "https://example.com", Type: ProfileType},
			want: &Object{ID: "https://example.com", Type: ProfileType},
		},
		{
			name: "Profile w/ extra properties",
			arg:  &Profile{ID: "https://example.com", Type: ProfileType, Describes: IRI("https://alt.example.com/")},
			want: &Object{ID: "https://example.com", Type: ProfileType},
		},
		{
			name: "Tombstone w/o extra properties",
			arg:  &Tombstone{ID: "https://example.com", Type: TombstoneType},
			want: &Object{ID: "https://example.com", Type: TombstoneType},
		},
		{
			name: "Tombstone w/ extra properties",
			arg:  &Tombstone{ID: "https://example.com", Type: TombstoneType, FormerType: GroupType, Deleted: time.Now()},
			want: &Object{ID: "https://example.com", Type: TombstoneType},
		},
		{
			name: "Create w/o extra properties",
			arg:  &Create{ID: "https://example.com", Type: CreateType},
			want: &Object{ID: "https://example.com", Type: CreateType},
		},
		{
			name: "Create w/ extra properties",
			arg:  &Create{ID: "https://example.com", Type: CreateType, Actor: IRI("https://example.com/1")},
			want: &Object{ID: "https://example.com", Type: CreateType},
		},
		{
			name: "Question w/o extra properties",
			arg:  &Question{ID: "https://example.com", Type: QuestionType},
			want: &Object{ID: "https://example.com", Type: QuestionType},
		},
		{
			name: "Question w/ extra properties",
			arg:  &Question{ID: "https://example.com", Type: QuestionType, AnyOf: ItemCollection{IRI("https://example.com")}},
			want: &Object{ID: "https://example.com", Type: QuestionType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := tt.arg.(Item)
			got, err := ToObject(arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !ItemsEqual(tt.want, got) {
				t.Errorf("ToObject() got = %v, expected %v", got, tt.want)
			}
		})
	}
}

func TestFlattenObjectProperties(t *testing.T) {
	t.Skipf("TODO")
}

func TestObject_Recipients(t *testing.T) {
	t.Skipf("TODO")
}

func TestRelationship_Recipients(t *testing.T) {
	t.Skipf("TODO")
}

func TestTombstone_Recipients(t *testing.T) {
	t.Skipf("TODO")
}

func TestSource_UnmarshalJSON(t *testing.T) {
	s := Source{}

	dataEmpty := []byte("{}")
	_ = s.UnmarshalJSON(dataEmpty)
	validateEmptySource(s, t)
}

func TestGetAPSource(t *testing.T) {
	data := []byte(`{"source": {"content": "test", "mediaType": "text/plain" }}`)

	par := fastjson.Parser{}
	val, _ := par.ParseBytes(data)
	a := GetAPSource(val)

	if a.Content.First().String() != "test" {
		t.Errorf("Content didn't match test value. Received %q, expecting %q", a.Content, "test")
	}
	if a.MediaType != "text/plain" {
		t.Errorf("Content didn't match test value. Received %q, expecting %q", a.MediaType, "text/plain")
	}
}

func TestObject_Clean(t *testing.T) {
	t.Skip("TODO")
}

func TestObject_IsCollection(t *testing.T) {
	t.Skip("TODO")
}

func TestActivityVocabularyType_MarshalJSON(t *testing.T) {
	t.Skip("TODO")
}

func TestActivityVocabularyType_GobDecode(t *testing.T) {
	tests := []struct {
		name    string
		t       ActivityVocabularyType
		data    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			t:       "",
			data:    []byte{},
			wantErr: false,
		},
		{
			name:    "some activity type",
			t:       PersonType,
			data:    gobValue([]byte("Person")),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.t.GobDecode(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("GobDecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestActivityVocabularyType_GobEncode(t *testing.T) {
	tests := []struct {
		name    string
		t       ActivityVocabularyType
		want    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			t:       "",
			want:    []byte{},
			wantErr: false,
		},
		{
			name:    "some activity type",
			t:       ActivityType,
			want:    []byte("Activity"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.t.GobEncode()
			if (err != nil) != tt.wantErr {
				t.Errorf("GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GobEncode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObject_MarshalJSON(t *testing.T) {
	type fields struct {
		ID           ID
		Type         TypeMatcher
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
		Tag          ItemCollection
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
	}
	tests := []struct {
		name    string
		fields  fields
		want    [][]byte
		wantErr bool
	}{
		{
			name:    "Empty",
			fields:  fields{},
			want:    nil,
			wantErr: false,
		},
		{
			name: "JustID",
			fields: fields{
				ID: ID("example.com"),
			},
			want:    [][]byte{[]byte(`{"id":"example.com"}`)},
			wantErr: false,
		},
		{
			name: "JustType",
			fields: fields{
				Type: ActivityVocabularyType("myType"),
			},
			want:    [][]byte{[]byte(`{"type":"myType"}`)},
			wantErr: false,
		},
		{
			name: "JustOneName",
			fields: fields{
				Name: NaturalLanguageValues{
					NilLangRef: Content("ana"),
				},
			},
			want:    [][]byte{[]byte(`{"name":"ana"}`)},
			wantErr: false,
		},
		{
			name: "MoreNames",
			fields: fields{
				Name: NaturalLanguageValues{
					English: Content("anna"),
					French:  Content("anne"),
				},
			},
			want: [][]byte{
				[]byte(`{"nameMap":{"en":"anna","fr":"anne"}}`),
				[]byte(`{"nameMap":{"fr":"anne","en":"anna"}}`),
			},
			wantErr: false,
		},
		{
			name: "JustOneSummary",
			fields: fields{
				Summary: NaturalLanguageValues{
					NilLangRef: Content("test summary"),
				},
			},
			want:    [][]byte{[]byte(`{"summary":"test summary"}`)},
			wantErr: false,
		},
		{
			name: "MoreSummaryEntries",
			fields: fields{
				Summary: NaturalLanguageValues{
					English: Content("test summary"),
					French:  Content("teste summary"),
				},
			},
			want: [][]byte{
				[]byte(`{"summaryMap":{"en":"test summary","fr":"teste summary"}}`),
				[]byte(`{"summaryMap":{"fr":"teste summary","en":"test summary"}}`),
			},
			wantErr: false,
		},
		{
			name: "JustOneContent",
			fields: fields{
				Content: NaturalLanguageValues{
					NilLangRef: Content("test content"),
				},
			},
			want:    [][]byte{[]byte(`{"content":"test content"}`)},
			wantErr: false,
		},
		{
			name: "MoreContentEntries",
			fields: fields{
				Content: NaturalLanguageValues{
					English: Content("test content"),
					French:  Content("teste content"),
				},
			},
			want: [][]byte{
				[]byte(`{"contentMap":{"en":"test content","fr":"teste content"}}`),
				[]byte(`{"contentMap":{"fr":"teste content","en":"test content"}}`),
			},
			wantErr: false,
		},
		{
			name: "MediaType",
			fields: fields{
				MediaType: MimeType("text/stupid"),
			},
			want:    [][]byte{[]byte(`{"mediaType":"text/stupid"}`)},
			wantErr: false,
		},
		{
			name: "Attachment",
			fields: fields{
				Attachment: &Object{
					ID:   "some example",
					Type: VideoType,
				},
			},
			want:    [][]byte{[]byte(`{"attachment":{"id":"some example","type":"Video"}}`)},
			wantErr: false,
		},
		{
			name: "AttributedTo",
			fields: fields{
				AttributedTo: &Actor{
					ID:   "http://example.com/ana",
					Type: PersonType,
				},
			},
			want:    [][]byte{[]byte(`{"attributedTo":{"id":"http://example.com/ana","type":"Person"}}`)},
			wantErr: false,
		},
		{
			name: "AttributedToDouble",
			fields: fields{
				AttributedTo: ItemCollection{
					&Actor{
						ID:   "http://example.com/ana",
						Type: PersonType,
					},
					&Actor{
						ID:   "http://example.com/GGG",
						Type: GroupType,
					},
				},
			},
			want:    [][]byte{[]byte(`{"attributedTo":[{"id":"http://example.com/ana","type":"Person"},{"id":"http://example.com/GGG","type":"Group"}]}`)},
			wantErr: false,
		},
		{
			name: "Source",
			fields: fields{
				Source: Source{
					MediaType: MimeType("text/plain"),
					Content:   NaturalLanguageValues{},
				},
			},
			want:    [][]byte{[]byte(`{"source":{"mediaType":"text/plain"}}`)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Object{
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
			}
			got, err := o.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			found := got == nil
			for _, wantBytes := range tt.want {
				if bytes.Equal(got, wantBytes) {
					found = true
				}
			}
			if !found {
				t.Errorf("MarshalJSON() got = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestSource_MarshalJSON(t *testing.T) {
	type fields struct {
		Content   NaturalLanguageValues
		MediaType MimeType
	}
	tests := []struct {
		name    string
		fields  fields
		want    [][]byte
		wantErr bool
	}{
		{
			name:    "Empty",
			fields:  fields{},
			want:    nil,
			wantErr: false,
		},
		{
			name: "MediaType",
			fields: fields{
				MediaType: MimeType("blank"),
			},
			want:    [][]byte{[]byte(`{"mediaType":"blank"}`)},
			wantErr: false,
		},
		{
			name: "OneContentValue",
			fields: fields{
				Content: NaturalLanguageValues{
					Und: Content("test"),
				},
			},
			want:    [][]byte{[]byte(`{"content":"test"}`)},
			wantErr: false,
		},
		{
			name: "MultipleContentValues",
			fields: fields{
				Content: NaturalLanguageValues{
					English: Content("test"),
					French:  Content("teste"),
				},
			},
			want: [][]byte{
				[]byte(`{"contentMap":{"en":"test","fr":"teste"}}`),
				[]byte(`{"contentMap":{"fr":"teste","en":"test"}}`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Source{
				Content:   tt.fields.Content,
				MediaType: tt.fields.MediaType,
			}
			got, err := s.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			found := got == nil
			for _, wantBytes := range tt.want {
				if bytes.Equal(got, wantBytes) {
					found = true
				}
			}
			if !found {
				t.Errorf("MarshalJSON() got = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestObject_Equals(t *testing.T) {
	type fields struct {
		ID           ID
		Type         TypeMatcher
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
		Tag          ItemCollection
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
	}
	tests := []struct {
		name   string
		fields fields
		arg    Item
		want   bool
	}{
		{
			name:   "equal-empty-object",
			fields: fields{},
			arg:    Object{},
			want:   true,
		},
		{
			name:   "equal-object-just-id",
			fields: fields{ID: "test"},
			arg:    Object{ID: "test"},
			want:   true,
		},
		{
			name:   "equal-object-id",
			fields: fields{ID: "test", URL: IRI("example.com")},
			arg:    Object{ID: "test"},
			want:   true,
		},
		{
			name:   "equal-false-with-id-and-url",
			fields: fields{ID: "test"},
			arg:    Object{ID: "test", URL: IRI("example.com")},
			want:   false,
		},
		{
			name:   "not a valid object",
			fields: fields{ID: "http://example.com"},
			arg:    Link{ID: "http://example.com"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Object{
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
			}
			if got := o.Equals(tt.arg); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObject_GobEncode(t *testing.T) {
	type fields struct {
		ID           ID
		Type         TypeMatcher
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
		Tag          ItemCollection
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
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "empty",
			fields:  fields{},
			wantErr: false,
		},
		{
			name:    "with ID",
			fields:  fields{ID: ID("https://example.com")},
			wantErr: false,
		},
		{
			name:    "with ID, type",
			fields:  fields{ID: ID("https://example.com"), Type: ObjectType},
			wantErr: false,
		},
		{
			name:    "with ID, type, name",
			fields:  fields{ID: ID("https://example.com"), Type: ObjectType, Name: NaturalLanguageValues{English: Content("ana")}},
			wantErr: false,
		},
		{
			name:    "with Source",
			fields:  fields{Source: Source{MediaType: "image/svg+xml", Content: NaturalLanguageValues{NilLangRef: Content("data:image/svg+xml,%3csvg%3e %3c/svg%3e")}}},
			wantErr: false,
		},
		{
			name:    "with IRI AttributedTo",
			fields:  fields{AttributedTo: IRI("https://example.com/1")},
			wantErr: false,
		},
		{
			name:    "with multiple IRIs AttributedTo",
			fields:  fields{AttributedTo: ItemCollection{IRI("https://example.com/1"), IRI("https://example.com/2")}},
			wantErr: false,
		},
		{
			name:    "with single object AttributedTo",
			fields:  fields{AttributedTo: Object{ID: "https://example.com/1"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Object{
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
			}
			got, err := o.GobEncode()
			if (err != nil) != tt.wantErr {
				t.Errorf("GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			ob := Object{}
			if err = ob.GobDecode(got); (err != nil) != tt.wantErr {
				t.Errorf("GobDecode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !ItemsEqual(ob, o) {
				t.Errorf("GobEncode() got = %s", cmp.Diff(ob, o))
			}
		})
	}
}

type reflectTest[T Objects | Links] struct {
	name    string
	arg     any
	want    *T
	wantErr bool
}

func Test_reflectedItemByType_Object(t *testing.T) {
	tests := []reflectTest[Object]{
		{
			name: "empty object",
			arg:  &Object{},
			want: &Object{},
		},
		{
			name: "object with ID",
			arg:  &Object{ID: "https://example.com"},
			want: &Object{ID: "https://example.com"},
		},
		{
			name: "object with ID, type",
			arg:  &Object{ID: "https://example.com", Type: ArticleType},
			want: &Object{ID: "https://example.com", Type: ArticleType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := tt.arg.(Item)
			got, err := reflectItemToType[Object](arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("reflectItemToType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !ItemsEqual(tt.want, got) {
				t.Errorf("reflectItemToType() got = %v, expected %v", got, tt.want)
			}
		})
	}
}
