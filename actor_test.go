package activitypub

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var (
	fnAct = func(_ *Actor) error { return nil }

	maybeActor Item = new(Actor)
)

func Benchmark_ToActor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToActor(maybeActor)
	}
}

func Benchmark_To_T_Actor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		To[Actor](maybeActor)
	}
}

func Benchmark_OnActor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OnActor(maybeActor, fnAct)
	}
}

func Benchmark_On_T_Actor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On[Actor](maybeActor, fnAct)
	}
}

func TestActor_Object(t *testing.T) {
	m := &Actor{Type: ActorType, ID: "test"}
	if reflect.DeepEqual(ID(""), m.GetID()) {
		t.Errorf("%#v should not be an empty activity pub object", m.GetID())
	}
}

func TestActor_Type(t *testing.T) {
	m := &Actor{Type: ActorType, ID: "test"}
	if !cmp.Equal(m.GetType(), ActorType) {
		t.Errorf("%#v should be an empty Link object", m.GetType())
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
	tests := []struct {
		name string
		a    Actor
		want Item
	}{
		{
			name: "empty",
			a:    Actor{},
			want: &Actor{},
		},
		{
			name: "has Bto",
			a:    Actor{Type: GroupType, Bto: ItemCollection{IRI("http://example.com")}},
			want: &Actor{Type: GroupType},
		},
		{
			name: "has BCC",
			a:    Actor{Type: ServiceType, BCC: ItemCollection{IRI("http://example.com")}},
			want: &Actor{Type: ServiceType},
		},
		{
			name: "audience has BCC",
			a:    Actor{Type: ApplicationType, Audience: ItemCollection{&Object{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &Actor{Type: ApplicationType, Audience: ItemCollection{&Object{}}},
		},
		{
			name: "attachment has BCC",
			a:    Actor{Type: PersonType, Attachment: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Actor{Type: PersonType, Attachment: &Object{}},
		},
		{
			name: "icon has BCC",
			a:    Actor{Type: PersonType, Icon: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Actor{Type: PersonType, Icon: &Object{}},
		},
		{
			name: "image has BCC",
			a:    Actor{Type: PersonType, Image: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Actor{Type: PersonType, Image: &Object{}},
		},
		{
			name: "context has BCC",
			a:    Actor{Type: PersonType, Context: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Actor{Type: PersonType, Context: &Object{}},
		},
		{
			name: "generator has BCC",
			a:    Actor{Type: PersonType, Generator: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Actor{Type: PersonType, Generator: &Object{}},
		},
		{
			name: "attributedTo has BCC",
			a:    Actor{Type: PersonType, AttributedTo: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Actor{Type: PersonType, AttributedTo: &Object{}},
		},
		{
			name: "preview has BCC",
			a:    Actor{Type: PersonType, Preview: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Actor{Type: PersonType, Preview: &Object{}},
		},
		{
			name: "tag has BCC",
			a:    Actor{Type: ApplicationType, Tag: ItemCollection{&Object{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &Actor{Type: ApplicationType, Tag: ItemCollection{&Object{}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Clean(); !cmp.Equal(got, tt.want, EquateItems) {
				t.Errorf("Clean() = %s", cmp.Diff(tt.want, got, EquateItems))
			}
		})
	}
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
		{
			name: "Tombstone",
			it:   &Tombstone{ID: "test", Type: TombstoneType, FormerType: PersonType},
			want: &Actor{ID: "test", Type: TombstoneType},
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
			want:   false,
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

func ExampleActor_initialization() {
	// actor1 is a struct literal which can be operated on directly.
	// For example, we can set the URL
	actor1 := Actor{}
	actor1.URL = IRI("http://example.com/1")
	fmt.Printf("Actor1: %v\n", actor1)

	// actor2 is wrapped in an Item interface.
	// It is probably the most common way of interacting with objects in the library,
	// because usually they get unmarshaled from an HTTP request, or another representation.
	var actor2 Item = &Actor{Type: ActorType}

	// That means we can't set any properties directly, so
	// if you uncomment the next line you will get a compiler error.
	// actor2.URL = IRI("http://example.com")

	_ = OnActor(actor2, func(actor *Actor) error {
		// In order to operate on it, we must wrap it in a call to OnActor
		actor.URL = IRI("http://example.com/~jdoe")
		actor.PreferredUsername = DefaultNaturalLanguage("jdoe")
		return nil
	})
	fmt.Printf("Actor2: %v\n", actor2)

	// Output:
	// Actor1: activitypub.Actor { url: http://example.com/1 }
	// Actor2: activitypub.Actor[Actor] { url: http://example.com/~jdoe, preferredUsername: jdoe }
}

func ExampleToActor() {
	// We can see here an initialization for the Application type which,
	// unlike the disjoint Place type that we've seen before,
	// is just an alias for Actor, but can be used to convey additional meaning.
	// There are additional aliases for all Actor types: Person, Group, and Service.
	//
	// However, there is no mechanism to coerce the Type property to the correct
	// Vocabulary Type value corresponding to the used alias, so it must be set manually.
	var actor1 Item = &Application{ID: "http://example.com/app-1", Type: ApplicationType}

	// As we've seen previously, we can't operate on actor1 as it's an Item instance, so
	// uncommenting the following line will trigger a compiler error:
	//actor1.PreferredUsername = DefaultNaturalLanguage("app")
	a1, _ := ToActor(actor1)
	a1.PreferredUsername = DefaultNaturalLanguage("app")
	fmt.Printf("Actor1: %v\n", actor1)
	fmt.Printf("      : %v\n\n", a1)

	// Here we see another valid, but maybe slightly misleading inline initialization.
	var actor2 Item = Actor{ID: "http://example.com/~jdoe", Type: PersonType, Name: DefaultNaturalLanguage("Jane Doe")}
	a2, _ := ToActor(actor2)
	a2.PreferredUsername = DefaultNaturalLanguage("jdoe")
	// We can still access the properties common with the Object type
	a2.Name = DefaultNaturalLanguage("Jane Doe Phd")
	// However since the Item interface is not wrapping a pointer to Actor,
	// these changes won't be reflected onto the actor2 value.
	fmt.Printf("Actor2: %v\n", actor2)
	// They persist though, on the "a2" value.
	fmt.Printf("      : %v\n", a2)

	// Another thing that we could do here is to assign back to actor2.
	// But, take care because stomping on the original value might not *always* be what you want.
	// I think that most times the correct thing to do is to ensure that the Item interface
	// wraps a pointer to the data you want.
	actor2 = a2
	fmt.Printf("Actor2: %v\n\n", actor2)

	// A special case in the library must be allowed for Tombstone types.
	// Tombstones are the objects left behind Delete activities, and they can replace
	// any of the other types, including Actor.
	//
	// Therefore, we support converting them to an Actor instance,
	// even though there is loss of data incurred, and the mechanism is slower as it
	// requires copying the common properties.
	maybeActor := Tombstone{ID: "http://example.com", Type: TombstoneType, FormerType: ServiceType}
	ma, _ := ToActor(maybeActor)
	fmt.Printf("MaybeActor: %v\n", maybeActor)
	fmt.Printf("          : %v\n\n", ma)

	// We should also consider that the hierarchy of types that can be converted is unidirectional.
	// As an example, an Object type can't be converted to an Actor type, and trying results in an error.
	//
	// We could use the same mechanism as for Tombstone objects, but we want to enforce the fact that
	// the types are actually disjoint in the Activity Vocabulary ontology.
	var notActor Item = &Object{ID: "http://example.com", Type: GroupType}
	na, err := ToActor(notActor)
	fmt.Printf("NotActor: %v\n", notActor)
	fmt.Printf("        : %v\n", na)
	fmt.Printf("Error   : %v\n\n", err)

	// Output:
	// Actor1: activitypub.Actor[Application] { id: http://example.com/app-1, preferredUsername: app }
	//       : activitypub.Actor[Application] { id: http://example.com/app-1, preferredUsername: app }
	//
	// Actor2: activitypub.Actor[Person] { id: http://example.com/~jdoe, name: Jane Doe }
	//       : activitypub.Actor[Person] { id: http://example.com/~jdoe, name: Jane Doe Phd, preferredUsername: jdoe }
	// Actor2: activitypub.Actor[Person] { id: http://example.com/~jdoe, name: Jane Doe Phd, preferredUsername: jdoe }
	//
	// MaybeActor: activitypub.Tombstone[Tombstone] { id: http://example.com, formerType: Service }
	//           : activitypub.Actor[Tombstone] { id: http://example.com }
	//
	// NotActor: activitypub.Object[Group] { id: http://example.com }
	//         : <nil>
	// Error   : unable to convert *activitypub.Object to *activitypub.Actor
}

func ExampleOnActor() {
	// In the ExampleToActor function, we saw how we can convert data types
	// to Actor pointer values and be allowed to use their properties in that way.
	//
	// Here we can see how this mechanism can be used to build specific logic when dealing
	// with opaque Item interface values.

	// As we've seen, you can not access this Actor's properties
	// because it's wrapped in the Item interface.
	var actor1 Item = &Actor{ID: "http://example.com/~jdoe", Type: PersonType}
	// Uncommenting this line will trigger a compilation error.
	//actor1.Name = DefaultNaturalLanguage("John Doe")
	_ = OnActor(actor1, func(act *Actor) error {
		// Instead we can wrap it in an OnActor() call in which
		// we can modify it, and the changes will be visible outside its scope.
		act.PreferredUsername = DefaultNaturalLanguage("jdoe")

		// Similarly, as we've seen in the ExampleToActor, we can also modify
		// the properties in common with the Object type, without needing
		// a call to OnObject/ToObject.
		act.Name = DefaultNaturalLanguage("John Doe")
		return nil
	})
	fmt.Printf("Actor1: %v\n", actor1)

	// Here's an example of handling a Tombstone actor that's a little more realistic.
	var deletedActor *Actor
	tombstone := Tombstone{ID: "http://example.com", Type: TombstoneType, FormerType: ServiceType}
	_ = OnActor(tombstone, func(act *Actor) error {
		// We can set the deletedActor's Type to the value contained in the
		// Tombstone.formerType property, as a way to ensure that future
		// type switches happening in the code, work as expected.
		act.Type = tombstone.FormerType
		deletedActor = act
		return nil
	})
	if ServiceType.Match(deletedActor.GetType()) {
		fmt.Printf("DeletedActor: %v\n", deletedActor)
	}

	// Output:
	// Actor1: activitypub.Actor[Person] { id: http://example.com/~jdoe, name: John Doe, preferredUsername: jdoe }
	// DeletedActor: activitypub.Actor[Service] { id: http://example.com }
}
