package activitypub

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestActorNew(t *testing.T) {
	testValue := ID("test")
	testType := ApplicationType

	o := ActorNew(testValue, testType)

	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if !testType.Match(o.GetType()) {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.GetType(), testType)
	}

	n := ActorNew(testValue, "")
	if n.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", n.ID, testValue)
	}
	if !cmp.Equal(n.GetType(), ActorType) {
		t.Errorf("APObject Type '%v' different than expected '%v'", n.GetType(), ActorType)
	}
}

func TestPersonNew(t *testing.T) {
	testValue := ID("test")

	o := PersonNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if !cmp.Equal(o.GetType(), PersonType) {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, PersonType)
	}
}

func TestApplicationNew(t *testing.T) {
	testValue := ID("test")

	o := ApplicationNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if !cmp.Equal(o.GetType(), ApplicationType) {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, ApplicationType)
	}
}

func TestGroupNew(t *testing.T) {
	testValue := ID("test")

	o := GroupNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if !cmp.Equal(o.GetType(), GroupType) {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, GroupType)
	}
}

func TestOrganizationNew(t *testing.T) {
	testValue := ID("test")

	o := OrganizationNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if !cmp.Equal(o.GetType(), OrganizationType) {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, OrganizationType)
	}
}

func TestServiceNew(t *testing.T) {
	testValue := ID("test")

	o := ServiceNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if !cmp.Equal(o.GetType(), ServiceType) {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, ServiceType)
	}
}

func TestActor_IsLink(t *testing.T) {
	m := ActorNew("test", ActorType)
	if m.IsLink() {
		t.Errorf("%#v should not be a valid Link", m.Type)
	}
}

func TestActor_IsObject(t *testing.T) {
	m := ActorNew("test", ActorType)
	if !m.IsObject() {
		t.Errorf("%#v should be a valid object", m.Type)
	}
}

func TestActor_Object(t *testing.T) {
	m := ActorNew("test", ActorType)
	if reflect.DeepEqual(ID(""), m.GetID()) {
		t.Errorf("%#v should not be an empty activity pub object", m.GetID())
	}
}

func TestActor_Type(t *testing.T) {
	m := ActorNew("test", ActorType)
	if !cmp.Equal(m.GetType(), ActorType) {
		t.Errorf("%#v should be an empty Link object", m.GetType())
	}
}

func TestPerson_IsLink(t *testing.T) {
	m := PersonNew("test")
	if m.IsLink() {
		t.Errorf("%T should not be a valid Link", m)
	}
}

func TestPerson_IsObject(t *testing.T) {
	m := PersonNew("test")
	if !m.IsObject() {
		t.Errorf("%T should be a valid object", m)
	}
}

func TestActor_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestActor_GetActor(t *testing.T) {
	t.Skipf("TODO")
}

func TestActor_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestActor_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestActor_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestApplication_GetActor(t *testing.T) {
	t.Skipf("TODO")
}

func TestApplication_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestApplication_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestApplication_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestApplication_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestApplication_IsObject(t *testing.T) {
	t.Skipf("TODO")
}

func TestGroup_GetActor(t *testing.T) {
	t.Skipf("TODO")
}

func TestGroup_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestGroup_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestGroup_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestGroup_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestGroup_IsObject(t *testing.T) {
	t.Skipf("TODO")
}

func TestOrganization_GetActor(t *testing.T) {
	t.Skipf("TODO")
}

func TestOrganization_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestOrganization_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestOrganization_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestOrganization_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestOrganization_IsObject(t *testing.T) {
	t.Skipf("TODO")
}

func TestPerson_GetActor(t *testing.T) {
	t.Skipf("TODO")
}

func TestPerson_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestPerson_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestPerson_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func validateEmptyPerson(p Person, t *testing.T) {
	if p.ID != "" {
		t.Errorf("Unmarshaled object %T should have empty ID, received %q", p, p.ID)
	}
	if HasTypes(p) {
		t.Errorf("Unmarshaled object "+
			"%T should have empty Type, received %s", p, p.GetType())
	}
	if p.AttributedTo != nil {
		t.Errorf("Unmarshaled object %T should have empty AttributedTo, received %q", p, p.AttributedTo)
	}
	if len(p.Name) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Name, received %q", p, p.Name)
	}
	if len(p.Summary) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Summary, received %q", p, p.Summary)
	}
	if len(p.Content) != 0 {
		t.Errorf("Unmarshaled object %T should have empty Content, received %q", p, p.Content)
	}
	if p.URL != nil {
		t.Errorf("Unmarshaled object %T should have empty URL, received %v", p, p.URL)
	}
	if !p.Published.IsZero() {
		t.Errorf("Unmarshaled object %T should have empty Published, received %q", p, p.Published)
	}
	if !p.StartTime.IsZero() {
		t.Errorf("Unmarshaled object %T  should have empty StartTime, received %q", p, p.StartTime)
	}
	if !p.Updated.IsZero() {
		t.Errorf("Unmarshaled object %T  should have empty Updated, received %q", p, p.Updated)
	}
}

func TestPerson_UnmarshalJSON(t *testing.T) {
	p := Person{}

	dataEmpty := []byte("{}")
	_ = p.UnmarshalJSON(dataEmpty)
	validateEmptyPerson(p, t)
}

func TestApplication_UnmarshalJSON(t *testing.T) {
	a := Application{}

	dataEmpty := []byte("{}")
	a.UnmarshalJSON(dataEmpty)
	validateEmptyPerson(Person(a), t)
}

func TestGroup_UnmarshalJSON(t *testing.T) {
	g := Group{}

	dataEmpty := []byte("{}")
	g.UnmarshalJSON(dataEmpty)
	validateEmptyPerson(Person(g), t)
}

func TestOrganization_UnmarshalJSON(t *testing.T) {
	o := Organization{}

	dataEmpty := []byte("{}")
	o.UnmarshalJSON(dataEmpty)
	validateEmptyPerson(Person(o), t)
}

func TestService_UnmarshalJSON(t *testing.T) {
	s := Service{}

	dataEmpty := []byte("{}")
	s.UnmarshalJSON(dataEmpty)
	validateEmptyPerson(Person(s), t)
}

func TestService_GetActor(t *testing.T) {
	t.Skipf("TODO")
}

func TestService_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestService_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestService_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestService_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestService_IsObject(t *testing.T) {
	t.Skipf("TODO")
}

func TestToPerson(t *testing.T) {
	t.Skipf("TODO")
}

func TestEndpoints_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestActor_Clean(t *testing.T) {
	t.Skipf("TODO")
}

func TestToActor(t *testing.T) {
	tests := []struct {
		name    string
		it      LinkOrIRI
		want    *Actor
		wantErr error
	}{
		{
			name: "empty",
		},
		{
			name: "Valid Actor",
			it:   Actor{ID: "test", Type: UpdateType},
			want: &Actor{ID: "test", Type: UpdateType},
		},
		{
			name: "Valid *Actor",
			it:   ActorNew("test", CreateType),
			want: ActorNew("test", CreateType),
		},
		{
			name:    "IRI",
			it:      IRI("https://example.com"),
			wantErr: ErrorInvalidType[Actor](IRI("")),
		},
		{
			name:    "IRIs",
			it:      IRIs{IRI("https://example.com")},
			wantErr: ErrorInvalidType[Actor](IRIs{}),
		},
		{
			name:    "ItemCollection",
			it:      ItemCollection{},
			wantErr: ErrorInvalidType[Actor](ItemCollection{}),
		},
		{
			name:    "Object",
			it:      &Object{ID: "test", Type: ArticleType},
			wantErr: ErrorInvalidType[Actor](&Object{}),
		},
		{
			name:    "Activity",
			it:      &Activity{ID: "test", Type: CreateType},
			wantErr: ErrorInvalidType[Actor](&Activity{}),
		},
		{
			name:    "IntransitiveActivity",
			it:      &IntransitiveActivity{ID: "test", Type: ArriveType},
			wantErr: ErrorInvalidType[Actor](&IntransitiveActivity{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToActor(tt.it)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("ToActor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ToActor(): %s got = %s", tt.name, cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestActor_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestActor_Recipients(t *testing.T) {
	t.Skipf("TODO")
}

func TestPublicKey_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestActor_MarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestEndpoints_MarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestPublicKey_MarshalJSON(t *testing.T) {
	type fields struct {
		ID           ID
		Owner        IRI
		PublicKeyPem string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr error
	}{
		{
			name:   "empty",
			fields: fields{},
			want:   nil,
		},
		{
			name: "just id",
			fields: fields{
				ID: "https://example.com",
			},
			want:    []byte(`{"id":"https://example.com"}`),
			wantErr: nil,
		},
		{
			name: "just owner",
			fields: fields{
				Owner: "https://example.com/~jdoe",
			},
			want:    []byte(`{"owner":"https://example.com/~jdoe"}`),
			wantErr: nil,
		},
		{
			name: "just PEM",
			fields: fields{
				PublicKeyPem: "-----BEGIN PUBLIC KEY-----\ntest\n-----END PUBLIC KEY-----",
			},
			want:    []byte(`{"publicKeyPem":"-----BEGIN PUBLIC KEY-----\ntest\n-----END PUBLIC KEY-----"}`),
			wantErr: nil,
		},
		{
			name: "id and pem",
			fields: fields{
				ID:           "https://example.com",
				PublicKeyPem: "-----BEGIN PUBLIC KEY-----\nid_and_pem\n-----END PUBLIC KEY-----",
			},
			want:    []byte(`{"id":"https://example.com","publicKeyPem":"-----BEGIN PUBLIC KEY-----\nid_and_pem\n-----END PUBLIC KEY-----"}`),
			wantErr: nil,
		},
		{
			name: "owner and pem",
			fields: fields{
				Owner:        "https://example.com/~jdoe",
				PublicKeyPem: "-----BEGIN PUBLIC KEY-----\nowner_and_pem\n-----END PUBLIC KEY-----",
			},
			want:    []byte(`{"owner":"https://example.com/~jdoe","publicKeyPem":"-----BEGIN PUBLIC KEY-----\nowner_and_pem\n-----END PUBLIC KEY-----"}`),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PublicKey{
				ID:           tt.fields.ID,
				Owner:        tt.fields.Owner,
				PublicKeyPem: tt.fields.PublicKeyPem,
			}
			got, err := p.MarshalJSON()
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("MarshalJSON() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
				if err != nil {
					return
				}
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("MarshalJSON() got = %s, want %s", got, tt.want)
			}
		})
	}
}

func personNilFn(t *testing.T, expected Item) WithActorFn {
	return func(_ *Actor) error {
		return nil
	}
}

func personIsNotEqual(t *testing.T, expected Item) WithActorFn {
	return func(p *Person) error {
		if cmp.Equal(p, expected) {
			t.Errorf("Person equal assert failed %s", cmp.Diff(expected, p))
		}
		return nil
	}
}

func personIsEqual(t *testing.T, expected Item) WithActorFn {
	return func(p *Person) error {
		if !cmp.Equal(p, expected) {
			t.Errorf("Person equal assert failed %s", cmp.Diff(expected, p))
		}
		return nil
	}
}

func TestOnActor(t *testing.T) {
	testPerson := Actor{
		ID: "https://example.com",
	}
	type args struct {
		it Item
		fn func(*testing.T, Item) WithActorFn
	}
	tests := []struct {
		name     string
		args     args
		expected Item
		wantErr  error
	}{
		{
			name: "empty",
			args: args{nil, personNilFn},
		},
		{
			name:     "single",
			args:     args{testPerson, personNilFn},
			expected: &testPerson,
		},
		{
			name:     "single fails",
			args:     args{Person{ID: "https://not-equals"}, personIsNotEqual},
			expected: &testPerson,
		},
		{
			name:     "collectionOfPersons",
			args:     args{ItemCollection{testPerson, testPerson}, personIsEqual},
			expected: &testPerson,
		},
		{
			name:     "collectionOfPersons fails",
			args:     args{ItemCollection{Person{}, Person{ID: "https://not-equals"}}, personIsNotEqual},
			expected: &testPerson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := OnActor(tt.args.it, tt.args.fn(t, tt.expected))
			if !cmp.Equal(err, tt.wantErr, cmpopts.EquateErrors()) {
				t.Errorf("OnPerson() error = %s", cmp.Diff(tt.wantErr, err, cmpopts.EquateErrors()))
			}
		})
	}
}

func TestActor_Equals(t *testing.T) {
	type fields struct {
		ID                ID
		Type              Typer
		Name              NaturalLanguageValues
		Attachment        Item
		AttributedTo      Item
		Audience          ItemCollection
		Content           NaturalLanguageValues
		Context           Item
		MediaType         MimeType
		EndTime           time.Time
		Generator         Item
		Icon              Item
		Image             Item
		InReplyTo         Item
		Location          Item
		Preview           Item
		Published         time.Time
		Replies           Item
		StartTime         time.Time
		Summary           NaturalLanguageValues
		Tag               ItemCollection
		Updated           time.Time
		URL               Item
		To                ItemCollection
		Bto               ItemCollection
		CC                ItemCollection
		BCC               ItemCollection
		Duration          time.Duration
		Likes             Item
		Shares            Item
		Source            Source
		Inbox             Item
		Outbox            Item
		Following         Item
		Followers         Item
		Liked             Item
		PreferredUsername NaturalLanguageValues
		Endpoints         *Endpoints
		Streams           ItemCollection
		PublicKey         PublicKey
	}
	tests := []struct {
		name   string
		fields fields
		arg    Item
		want   bool
	}{
		{
			name:   "equal-empty-actor",
			fields: fields{},
			arg:    Actor{},
			want:   true,
		},
		{
			name:   "equal-actor-just-id",
			fields: fields{ID: "test"},
			arg:    Actor{ID: "test"},
			want:   true,
		},
		{
			name:   "equal-actor-id",
			fields: fields{ID: "test", URL: IRI("example.com")},
			arg:    Actor{ID: "test"},
			want:   true,
		},
		{
			name:   "equal-false-with-id-and-url",
			fields: fields{ID: "test"},
			arg:    Actor{ID: "test", URL: IRI("example.com")},
			want:   false,
		},
		{
			name:   "not a valid actor",
			fields: fields{ID: "http://example.com"},
			arg:    Activity{ID: "http://example.com"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Actor{
				ID:                tt.fields.ID,
				Type:              tt.fields.Type,
				Name:              tt.fields.Name,
				Attachment:        tt.fields.Attachment,
				AttributedTo:      tt.fields.AttributedTo,
				Audience:          tt.fields.Audience,
				Content:           tt.fields.Content,
				Context:           tt.fields.Context,
				MediaType:         tt.fields.MediaType,
				EndTime:           tt.fields.EndTime,
				Generator:         tt.fields.Generator,
				Icon:              tt.fields.Icon,
				Image:             tt.fields.Image,
				InReplyTo:         tt.fields.InReplyTo,
				Location:          tt.fields.Location,
				Preview:           tt.fields.Preview,
				Published:         tt.fields.Published,
				Replies:           tt.fields.Replies,
				StartTime:         tt.fields.StartTime,
				Summary:           tt.fields.Summary,
				Tag:               tt.fields.Tag,
				Updated:           tt.fields.Updated,
				URL:               tt.fields.URL,
				To:                tt.fields.To,
				Bto:               tt.fields.Bto,
				CC:                tt.fields.CC,
				BCC:               tt.fields.BCC,
				Duration:          tt.fields.Duration,
				Likes:             tt.fields.Likes,
				Shares:            tt.fields.Shares,
				Source:            tt.fields.Source,
				Inbox:             tt.fields.Inbox,
				Outbox:            tt.fields.Outbox,
				Following:         tt.fields.Following,
				Followers:         tt.fields.Followers,
				Liked:             tt.fields.Liked,
				PreferredUsername: tt.fields.PreferredUsername,
				Endpoints:         tt.fields.Endpoints,
				Streams:           tt.fields.Streams,
				PublicKey:         tt.fields.PublicKey,
			}
			if got := a.Equals(tt.arg); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
