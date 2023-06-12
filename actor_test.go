package activitypub

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestActorNew(t *testing.T) {
	testValue := ID("test")
	testType := ApplicationType

	o := ActorNew(testValue, testType)

	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != testType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, testType)
	}

	n := ActorNew(testValue, "")
	if n.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", n.ID, testValue)
	}
	if n.Type != ActorType {
		t.Errorf("APObject Type '%v' different than expected '%v'", n.Type, ActorType)
	}
}

func TestPersonNew(t *testing.T) {
	testValue := ID("test")

	o := PersonNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != PersonType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, PersonType)
	}
}

func TestApplicationNew(t *testing.T) {
	testValue := ID("test")

	o := ApplicationNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != ApplicationType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, ApplicationType)
	}
}

func TestGroupNew(t *testing.T) {
	testValue := ID("test")

	o := GroupNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != GroupType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, GroupType)
	}
}

func TestOrganizationNew(t *testing.T) {
	testValue := ID("test")

	o := OrganizationNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != OrganizationType {
		t.Errorf("APObject Type '%v' different than expected '%v'", o.Type, OrganizationType)
	}
}

func TestServiceNew(t *testing.T) {
	testValue := ID("test")

	o := ServiceNew(testValue)
	if o.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", o.ID, testValue)
	}
	if o.Type != ServiceType {
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
	if m.GetType() != ActorType {
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
	if p.Type != "" {
		t.Errorf("Unmarshaled object %T should have empty Type, received %q", p, p.Type)
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
	p.UnmarshalJSON(dataEmpty)
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
	t.Skipf("TODO")
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
	t.Skipf("TODO")
}

func assertPersonWithTesting(fn canErrorFunc, expected Item) WithActorFn {
	return func(p *Person) error {
		if !assertDeepEquals(fn, p, expected) {
			return fmt.Errorf("not equal")
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
		fn func(canErrorFunc, Item) WithActorFn
	}
	tests := []struct {
		name     string
		args     args
		expected Item
		wantErr  bool
	}{
		{
			name:     "single",
			args:     args{testPerson, assertPersonWithTesting},
			expected: &testPerson,
			wantErr:  false,
		},
		{
			name:     "single fails",
			args:     args{Person{ID: "https://not-equals"}, assertPersonWithTesting},
			expected: &testPerson,
			wantErr:  true,
		},
		{
			name:     "collectionOfPersons",
			args:     args{ItemCollection{testPerson, testPerson}, assertPersonWithTesting},
			expected: &testPerson,
			wantErr:  false,
		},
		{
			name:     "collectionOfPersons fails",
			args:     args{ItemCollection{testPerson, Person{ID: "https://not-equals"}}, assertPersonWithTesting},
			expected: &testPerson,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		var logFn canErrorFunc
		if tt.wantErr {
			logFn = t.Logf
		} else {
			logFn = t.Errorf
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := OnActor(tt.args.it, tt.args.fn(logFn, tt.expected)); (err != nil) != tt.wantErr {
				t.Errorf("OnPerson() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestActor_Equals(t *testing.T) {
	type fields struct {
		ID                ID
		Type              ActivityVocabularyType
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
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
