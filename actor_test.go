package activitypub

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"
	"time"

	"github.com/go-ap/errors"
	"github.com/google/go-cmp/cmp"
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

func Benchmark_OnActor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OnActor(maybeActor, fnAct)
	}
}

func TestActor_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    Actor
		wantErr error
	}{
		{
			name:    "nil",
			data:    nil,
			wantErr: fmt.Errorf(`cannot parse JSON: cannot parse empty string; unparsed tail: ""`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Actor{}
			if err := l.UnmarshalJSON(tt.data); !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("UnmarshalJSON() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
			}
			if !cmp.Equal(l, tt.want) {
				t.Errorf("UnmarshalJSON() got = %s", cmp.Diff(tt.want, l))
			}
		})
	}
}

func TestActor_GetID(t *testing.T) {
	tests := []struct {
		name string
		ID   ID
		want ID
	}{
		{
			name: "empty",
			ID:   "",
			want: "",
		},
		{
			name: "not empty",
			ID:   "http://example.com",
			want: "http://example.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Actor{ID: tt.ID}
			if got := l.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActor_GetLink(t *testing.T) {
	tests := []struct {
		name string
		IRI  IRI
		want IRI
	}{
		{
			name: "empty",
			IRI:  "",
			want: "",
		},
		{
			name: "not empty",
			IRI:  "http://example.com",
			want: "http://example.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Actor{ID: tt.IRI}
			if got := l.GetLink(); got != tt.want {
				t.Errorf("GetLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActor_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestEndpoints_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    Endpoints
		wantErr error
	}{
		{
			name:    "nil",
			data:    nil,
			want:    Endpoints{},
			wantErr: fmt.Errorf(`cannot parse JSON: cannot parse empty string; unparsed tail: ""`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Endpoints{}
			if err := e.UnmarshalJSON(tt.data); !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("UnmarshalJSON() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
			}
			if !cmp.Equal(e, tt.want) {
				t.Errorf("UnmarshalJSON() got = %s", cmp.Diff(tt.want, e))
			}
		})
	}
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

func TestActor_Recipients(t *testing.T) {
	t.Skipf("TODO")
}

func TestPublicKey_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestActor_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		sub     Actor
		want    []byte
		wantErr error
	}{
		{
			name: "nil",
			sub:  Actor{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sub.MarshalJSON()
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("MarshalJSON() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("MarshalJSON() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestEndpoints_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		sub     Endpoints
		want    []byte
		wantErr error
	}{
		{
			name: "nil",
			sub:  Endpoints{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.sub.MarshalJSON()
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("MarshalJSON() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("MarshalJSON() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
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

func TestOnActor(t *testing.T) {
	type args struct {
		it Item
		fn func(*testing.T) WithActorFn
	}
	tests := []struct {
		name     string
		args     args
		expected Item
		wantErr  error
	}{
		{
			name: "empty",
			args: args{
				it: nil,
				fn: func(t *testing.T) WithActorFn {
					return nil
				},
			},
		},
		{
			name: "single",
			args: args{
				it: Actor{ID: "https://example.com"},
				fn: func(t *testing.T) WithActorFn {
					return func(actor *Actor) error {
						return nil
					}
				},
			},
			expected: &Actor{ID: "https://example.com"},
		},
		{
			name: "single fails",
			args: args{
				it: Person{ID: "https://not-equals"},
				fn: func(t *testing.T) WithActorFn {
					return func(actor *Actor) error {
						return nil
					}
				}},
			expected: &Actor{ID: "https://example.com"},
		},
		{
			name: "collectionOfPersons",
			args: args{
				it: ItemCollection{
					Actor{ID: "https://example.com"},
					Actor{ID: "https://example.com"},
				},
				fn: func(t *testing.T) WithActorFn {
					return func(actor *Actor) error {
						return nil
					}
				},
			},
			expected: &Actor{ID: "https://example.com"},
		},
		{
			name: "collectionOfPersons fails",
			args: args{
				it: ItemCollection{Person{}, Person{ID: "https://not-equals"}},
				fn: func(t *testing.T) WithActorFn {
					return func(actor *Actor) error {
						return nil
					}
				},
			},
			expected: &Actor{ID: "https://example.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := OnActor(tt.args.it, tt.args.fn(t))
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("OnPerson() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
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
		name string
		act  Actor
		arg  Item
		want bool
	}{
		{
			name: "equal-empty-actor",
			arg:  Actor{},
			want: true,
		},
		{
			name: "equal-actor-just-id",
			act:  Actor{ID: "test"},
			arg:  Actor{ID: "test"},
			want: true,
		},
		{
			name: "equal-actor-id",
			act:  Actor{ID: "test", URL: IRI("example.com")},
			arg:  Actor{ID: "test"},
			want: false,
		},
		{
			name: "equal-false-with-id-and-url",
			act:  Actor{ID: "test"},
			arg:  Actor{ID: "test", URL: IRI("example.com")},
			want: false,
		},
		{
			name: "not a valid actor",
			act:  Actor{ID: "http://example.com"},
			arg:  Activity{ID: "http://example.com"},
			want: false,
		},
		{
			name: "inbox different",
			act:  Actor{Inbox: CollectionPath("test").IRI(IRI("http://example.com"))},
			arg:  Actor{Inbox: IRI("http://example.com")},
			want: false,
		},
		{
			name: "inbox same",
			act:  Actor{Inbox: CollectionPath("test").IRI(IRI("http://example.com"))},
			arg:  Actor{Inbox: IRI("http://example.com/test")},
			want: true,
		},
		{
			name: "outbox different",
			act:  Actor{Outbox: CollectionPath("test").IRI(IRI("http://example.com"))},
			arg:  Actor{Outbox: IRI("http://example.com")},
			want: false,
		},
		{
			name: "outbox same",
			act:  Actor{Outbox: CollectionPath("test").IRI(IRI("http://example.com"))},
			arg:  Actor{Outbox: IRI("http://example.com/test")},
			want: true,
		},
		{
			name: "following different",
			act:  Actor{Following: CollectionPath("test").IRI(IRI("http://example.com"))},
			arg:  Actor{Following: IRI("http://example.com")},
			want: false,
		},
		{
			name: "following same",
			act:  Actor{Following: CollectionPath("test").IRI(IRI("http://example.com"))},
			arg:  Actor{Following: IRI("http://example.com/test")},
			want: true,
		},
		{
			name: "followers different",
			act:  Actor{Followers: CollectionPath("test").IRI(IRI("http://example.com"))},
			arg:  Actor{Followers: IRI("http://example.com")},
			want: false,
		},
		{
			name: "followers same",
			act:  Actor{Followers: CollectionPath("test").IRI(IRI("http://example.com"))},
			arg:  Actor{Followers: IRI("http://example.com/test")},
			want: true,
		},
		{
			name: "liked different",
			act:  Actor{Liked: CollectionPath("test").IRI(IRI("http://example.com"))},
			arg:  Actor{Liked: IRI("http://example.com")},
			want: false,
		},
		{
			name: "liked same",
			act:  Actor{Liked: CollectionPath("test").IRI(IRI("http://example.com"))},
			arg:  Actor{Liked: IRI("http://example.com/test")},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.act.Equals(tt.arg); got != tt.want {
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
		actor.PreferredUsername = DefaultLangValue("jdoe")
		return nil
	})
	fmt.Printf("Actor2: %v\n", actor2)

	// Output:
	// Actor1: activitypub.Actor { url: http://example.com/1 }
	// Actor2: activitypub.Actor[Actor] { url: http://example.com/~jdoe, preferredUsername: jdoe }
}

func ExampleToActor() {
	// The Actor types form a large part of the data handled by the library,
	// especially in the ActivityPub state machine module.

	// We can see here an initialization for the Application type which
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
	a1.PreferredUsername = DefaultLangValue("app")
	fmt.Printf("Actor1: %v\n", actor1)
	fmt.Printf("      : %v\n\n", a1)

	// Here we see another valid, but maybe slightly misleading inline initialization.
	var actor2 Item = Actor{ID: "http://example.com/~jdoe", Type: PersonType, Name: DefaultLangValue("Jane Doe")}
	a2, _ := ToActor(actor2)
	a2.PreferredUsername = DefaultLangValue("jdoe")
	// We can still access the properties common with the Object type
	a2.Name = DefaultLangValue("Jane Doe Phd")
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
	// In the ToActor() example function, we saw how we can convert data types
	// to Actor pointer values and be allowed to use their properties in that way.
	//
	// Here we can see how this mechanism can be used to build specific logic when dealing
	// with opaque Item interface values.

	// As we've seen, you can not access this Actor's properties
	// because it's wrapped in the Item interface.
	var actor1 Item = &Actor{ID: "http://example.com/~jdoe", Type: PersonType}
	// Uncommenting this line will trigger a compiler error.
	//actor1.Name = DefaultNaturalLanguage("John Doe")

	_ = OnActor(actor1, func(act *Actor) error {
		// Instead we can wrap it in an OnActor() call in which
		// we can modify it, and the changes will be visible outside its scope.
		act.PreferredUsername = DefaultLangValue("jdoe")

		// Similarly, as we've seen in the ToActor() example, we can also modify
		// the properties in common with the Object type, without needing
		// a call to OnObject/ToObject.
		act.Name = DefaultLangValue("John Doe")
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

func TestActor_UnmarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    Actor
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
			l := Actor{}
			if err := l.UnmarshalBinary(tt.data); !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("UnmarshalBinary() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
			}
			if !cmp.Equal(l, tt.want) {
				t.Errorf("UnmarshalBinary() got = %s", cmp.Diff(tt.want, l))
			}
		})
	}
}

func TestActor_MarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		sub     Actor
		want    []byte
		wantErr error
	}{
		{
			name:    "nil",
			sub:     Actor{},
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

func TestCopyActorProperties(t *testing.T) {
	type args struct {
		to   *Actor
		from *Actor
	}
	tests := []struct {
		name    string
		args    args
		want    *Actor
		wantErr error
	}{
		{
			name: "nil",
			args: args{},
			want: nil,
		},
		{
			name: string(Inbox),
			args: args{
				to:   &Actor{},
				from: &Actor{Inbox: CollectionPath("test").IRI(IRI("http://example.com"))},
			},
			want: &Actor{Inbox: IRI("http://example.com/test")},
		},
		{
			name: string(Outbox),
			args: args{
				to:   &Actor{},
				from: &Actor{Outbox: CollectionPath("test").IRI(IRI("http://example.com"))},
			},
			want: &Actor{Outbox: IRI("http://example.com/test")},
		},
		{
			name: string(Following),
			args: args{
				to:   &Actor{},
				from: &Actor{Following: CollectionPath("test").IRI(IRI("http://example.com"))},
			},
			want: &Actor{Following: IRI("http://example.com/test")},
		},
		{
			name: string(Followers),
			args: args{
				to:   &Actor{},
				from: &Actor{Followers: CollectionPath("test").IRI(IRI("http://example.com"))},
			},
			want: &Actor{Followers: IRI("http://example.com/test")},
		},
		{
			name: string(Liked),
			args: args{
				to:   &Actor{},
				from: &Actor{Liked: CollectionPath("test").IRI(IRI("http://example.com"))},
			},
			want: &Actor{Liked: IRI("http://example.com/test")},
		},
		{
			name: "preferredUsername",
			args: args{
				to:   &Actor{},
				from: &Actor{PreferredUsername: DefaultLangValue("jdoe")},
			},
			want: &Actor{PreferredUsername: DefaultLangValue("jdoe")},
		},
		{
			name: "endpoints",
			args: args{
				to: &Actor{},
				from: &Actor{Endpoints: &Endpoints{
					UploadMedia:                IRI("http://example.com/upload"),
					OauthAuthorizationEndpoint: IRI("http://example.com/auth-url"),
					OauthTokenEndpoint:         IRI("http://example.com/token-url"),
					SharedInbox:                IRI("http://example.com/shared-inbox"),
					ProxyURL:                   "http://example.com/proxy-url",
				}},
			},
			want: &Actor{Endpoints: &Endpoints{
				UploadMedia:                IRI("http://example.com/upload"),
				OauthAuthorizationEndpoint: IRI("http://example.com/auth-url"),
				OauthTokenEndpoint:         IRI("http://example.com/token-url"),
				SharedInbox:                IRI("http://example.com/shared-inbox"),
				ProxyURL:                   "http://example.com/proxy-url",
			}},
		},
		{
			name: "streams",
			args: args{
				to:   &Actor{},
				from: &Actor{Streams: ItemCollection{IRI("http://example.com/stream1"), IRI("stream2")}},
			},
			want: &Actor{Streams: ItemCollection{IRI("http://example.com/stream1"), IRI("stream2")}},
		},
		{
			name: "publickKey",
			args: args{
				to: &Actor{},
				from: &Actor{PublicKey: PublicKey{
					ID:           "http://example.com/public-key",
					Owner:        "http://example.com/~jdoe",
					PublicKeyPem: "--- NOT A PUBLIC KEY ---",
				}},
			},
			want: &Actor{PublicKey: PublicKey{
				ID:           "http://example.com/public-key",
				Owner:        "http://example.com/~jdoe",
				PublicKeyPem: "--- NOT A PUBLIC KEY ---",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CopyActorProperties(tt.args.to, tt.args.from)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("CopyActorProperties() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("CopyActorProperties() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestFlattenActorProperties(t *testing.T) {
	tests := []struct {
		name string
		a    *Actor
		want *Actor
	}{
		{
			name: "nil",
			a:    nil,
			want: nil,
		},
		{
			name: string(Inbox),
			a: &Actor{
				Inbox: Object{ID: CollectionPath("test").IRI(IRI("http://example.com"))},
			},
			want: &Actor{Inbox: IRI("http://example.com/test")},
		},
		{
			name: string(Outbox),
			a: &Actor{
				Outbox: Object{ID: CollectionPath("test").IRI(IRI("http://example.com"))},
			},
			want: &Actor{Outbox: IRI("http://example.com/test")},
		},
		{
			name: string(Following),
			a: &Actor{
				Following: Object{ID: CollectionPath("test").IRI(IRI("http://example.com"))},
			},
			want: &Actor{Following: IRI("http://example.com/test")},
		},
		{
			name: string(Followers),
			a: &Actor{
				Followers: Object{ID: CollectionPath("test").IRI(IRI("http://example.com"))},
			},
			want: &Actor{Followers: IRI("http://example.com/test")},
		},
		{
			name: string(Liked),
			a: &Actor{
				Liked: Object{ID: CollectionPath("test").IRI(IRI("http://example.com"))},
			},
			want: &Actor{Liked: IRI("http://example.com/test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlattenActorProperties(tt.a); !cmp.Equal(got, tt.want) {
				t.Errorf("FlattenActorProperties() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

var (
	// Example ECC P-256 Test Key
	// https://www.rfc-editor.org/rfc/rfc9421.html#name-example-ecc-p-256-test-key
	prvECDSA = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIFKbhfNZfpDsW43+0+JjUr9K+bTeuxopu653+hBaXGA7oAoGCCqGSM49
AwEHoUQDQgAEqIVYZVLCrPZHGHjP17CTW0/+D9Lfw0EkjqF7xB4FivAxzic30tMM
4GF+hR6Dxh71Z50VGGdldkkDXZCnTNnoXQ==
-----END EC PRIVATE KEY-----`

	pubECDSA = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEqIVYZVLCrPZHGHjP17CTW0/+D9Lf
w0EkjqF7xB4FivAxzic30tMM4GF+hR6Dxh71Z50VGGdldkkDXZCnTNnoXQ==
-----END PUBLIC KEY-----`

	prvKeyECDSA = func() *ecdsa.PrivateKey {
		prvBlockECDSA, _ := pem.Decode([]byte(prvECDSA))
		k, _ := x509.ParseECPrivateKey(prvBlockECDSA.Bytes)
		return k
	}()

	// Example Ed25519 Test Key
	// https://www.rfc-editor.org/rfc/rfc9421.html#name-example-ed25519-test-key
	prvEd25519 = `-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIJ+DYvh6SEqVTm50DFtMDoQikTmiCqirVv9mWG9qfSnF
-----END PRIVATE KEY-----`

	pubEd25519 = `-----BEGIN PUBLIC KEY-----
MCowBQYDK2VwAyEAJrQLj5P/89iXES9+vFgrIy29clF9CC/oPPsw3c5D0bs=
-----END PUBLIC KEY-----`

	prvKeyEd25519 = func() ed25519.PrivateKey {
		prvBlockEd25519, _ := pem.Decode([]byte(prvEd25519))
		k, _ := x509.ParsePKCS8PrivateKey(prvBlockEd25519.Bytes)
		return k.(ed25519.PrivateKey)
	}()

	prvRSA = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDEZNd5f+5jjw7Y
vzhwniZgFOiz80cOWAJtMGtmorjkjaQPE2cmrgWvEHiCYqQ0jnbCSJrMPZXUlXUm
vsdHaczpbHKlnPKUgC35QpXs3NikWvoZFBhJR99lbGuZilsQj/lMi7Ht7lzmZRDR
/ZeapRi+otxjSKNe3FH1ONIaXZdBEfRHKfRW2FV9W3L76gX9jsH/2s26r6LYlyLM
lnQkt2dM+GwYSG4pv/Kl/KE2i/UdJ/o/tealiO5usyZwK3U2vZCJaseMWDbluHTM
Q1RPPV8SeI5pBqREa2XrSwbcZUI+TaB+xiPvIAjrTboxLY5XyIDwjag+a/aMvzRx
JAKRmRqhAgMBAAECggEAGYiEzSqZSz9Xrk1aIKYnFhnR0UeBRvehRSHk7MCeKjTS
DhW3NPuuCH8rM8RwVdbp0MOQwJoHJ07RHtrx3LKALh7n3ulDTpRFpeEGzfc+gUvE
tUr8B1b9T9njOWCYC1S0lEObO/RgBqJAKBUAx13MlEhnP887UkNxsmCTTFM7rX1W
AVgHGZo71M6IebHjoLEFmYAXtFgY3+W29J3JOAUsWWCIZYntnOtNAslRYjzyuTRh
bDTKayKMuRuPo4/jgQhsHS6qsRonK0dQRKLuBX8iHyoXFQP7GVea9liHnm9lqwQ1
Ve4zvD4IUbV71NVZ6xsY3Nfr7/f8ehatMhSZ/H6BKwKBgQD3Eq7yhOmzGk8Woa2b
7NBB9UwCoGHIWh/lXj731a/3zD85sS7VXdvyXOwhilMpx/lNmTBhq2ko5qf122xu
JyQ1DfsKvHs0sHqL1ZVHovpe0lvwxu1Nj/HfsEqlt//qFvXa4Gp5uw/hqcLcxfSa
Yo8z6m/+QPESMPxwD3FaUmKuiwKBgQDLfWYv/t6m+4v0ROSY+oGW/r6HbYkFWodw
UU+S2THwOttddJDEOznEZxmJhqAXPMChTkmXL8kn1aRSmn/AHJDVywCxv0xNiIcR
kjfFmmzdig/HcBoWHPc5aAoZorxDEIb6NWEJt/vEDfnOOLItfufKCy9aKdQ7pwjP
FcPH7TwNAwKBgFTPNP5KYW35Oeyq0s0THOmHKfA83VPIm+o/z52C3ERS9+D10P2s
mjM3claRBLryycC5NMJR9Gb1xfG+wBmPlf4gLmwhBqmvamFVj0hnyUmDK8wafJqD
LqN6ACWiY1YXS402O1ZNv8XWX+0ohi34Zu+LKaY85IM6DWzp4B8A6J7BAoGBALkK
SAE/B6LavXKbrzA5I9x1vDYUcfQPVXfaSLzlipbEPrRmCjqXDLm/cyZu6GcZFKXa
NesoRghWKv3+hkrg7weqePApX65lh0WAK/0hpvtxz1VxaBdRsbJfHEghhoaJoeQm
5B3dUzD98HoJbmUWsJo2v5GC1f6Erur5BLZp0SCXAoGAPFS3336Z1pK0Ne+Qw83x
V9dcGszH2HYfU4AVOyxIRK6DDsw0k/r4zi/0AfRjnPlP89qffTZfkEgW+/jXQYzr
B5MKfJ21HzAim4dyNzhbViNqzQtRZhP7/ESJJVF1CUbyIr3xqwww0rVdWxgTPrra
CymanYHXIAFVDKU/1A099Uw=
-----END PRIVATE KEY-----`

	pubRSA = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxGTXeX/uY48O2L84cJ4m
YBTos/NHDlgCbTBrZqK45I2kDxNnJq4FrxB4gmKkNI52wkiazD2V1JV1Jr7HR2nM
6WxypZzylIAt+UKV7NzYpFr6GRQYSUffZWxrmYpbEI/5TIux7e5c5mUQ0f2XmqUY
vqLcY0ijXtxR9TjSGl2XQRH0Ryn0VthVfVty++oF/Y7B/9rNuq+i2JcizJZ0JLdn
TPhsGEhuKb/ypfyhNov1HSf6P7XmpYjubrMmcCt1Nr2QiWrHjFg25bh0zENUTz1f
EniOaQakRGtl60sG3GVCPk2gfsYj7yAI6026MS2OV8iA8I2oPmv2jL80cSQCkZka
oQIDAQAB
-----END PUBLIC KEY-----`

	prvKeyRSA = func() *rsa.PrivateKey {
		prvBlockRSA, _ := pem.Decode([]byte(prvRSA))
		k, err := x509.ParsePKCS8PrivateKey(prvBlockRSA.Bytes)
		if err != nil {
			panic(err)
		}
		return k.(*rsa.PrivateKey)
	}()

	// NOTE(marius): values from the Example RSA key test cases in the HTTP Signatures RFC9421
	//  https://www.rfc-editor.org/rfc/rfc9421.html#appendix-B.1.1
	prvRSA1 = `-----BEGIN RSA PRIVATE KEY-----
MIIEqAIBAAKCAQEAhAKYdtoeoy8zcAcR874L8cnZxKzAGwd7v36APp7Pv6Q2jdsP
BRrwWEBnez6d0UDKDwGbc6nxfEXAy5mbhgajzrw3MOEt8uA5txSKobBpKDeBLOsd
JKFqMGmXCQvEG7YemcxDTRPxAleIAgYYRjTSd/QBwVW9OwNFhekro3RtlinV0a75
jfZgkne/YiktSvLG34lw2zqXBDTC5NHROUqGTlML4PlNZS5Ri2U4aCNx2rUPRcKI
lE0PuKxI4T+HIaFpv8+rdV6eUgOrB2xeI1dSFFn/nnv5OoZJEIB+VmuKn3DCUcCZ
SFlQPSXSfBDiUGhwOw76WuSSsf1D4b/vLoJ10wIDAQABAoIBAG/JZuSWdoVHbi56
vjgCgkjg3lkO1KrO3nrdm6nrgA9P9qaPjxuKoWaKO1cBQlE1pSWp/cKncYgD5WxE
CpAnRUXG2pG4zdkzCYzAh1i+c34L6oZoHsirK6oNcEnHveydfzJL5934egm6p8DW
+m1RQ70yUt4uRc0YSor+q1LGJvGQHReF0WmJBZHrhz5e63Pq7lE0gIwuBqL8SMaA
yRXtK+JGxZpImTq+NHvEWWCu09SCq0r838ceQI55SvzmTkwqtC+8AT2zFviMZkKR
Qo6SPsrqItxZWRty2izawTF0Bf5S2VAx7O+6t3wBsQ1sLptoSgX3QblELY5asI0J
YFz7LJECgYkAsqeUJmqXE3LP8tYoIjMIAKiTm9o6psPlc8CrLI9CH0UbuaA2JCOM
cCNq8SyYbTqgnWlB9ZfcAm/cFpA8tYci9m5vYK8HNxQr+8FS3Qo8N9RJ8d0U5Csw
DzMYfRghAfUGwmlWj5hp1pQzAuhwbOXFtxKHVsMPhz1IBtF9Y8jvgqgYHLbmyiu1
mwJ5AL0pYF0G7x81prlARURwHo0Yf52kEw1dxpx+JXER7hQRWQki5/NsUEtv+8RT
qn2m6qte5DXLyn83b1qRscSdnCCwKtKWUug5q2ZbwVOCJCtmRwmnP131lWRYfj67
B/xJ1ZA6X3GEf4sNReNAtaucPEelgR2nsN0gKQKBiGoqHWbK1qYvBxX2X3kbPDkv
9C+celgZd2PW7aGYLCHq7nPbmfDV0yHcWjOhXZ8jRMjmANVR/eLQ2EfsRLdW69bn
f3ZD7JS1fwGnO3exGmHO3HZG+6AvberKYVYNHahNFEw5TsAcQWDLRpkGybBcxqZo
81YCqlqidwfeO5YtlO7etx1xLyqa2NsCeG9A86UjG+aeNnXEIDk1PDK+EuiThIUa
/2IxKzJKWl1BKr2d4xAfR0ZnEYuRrbeDQYgTImOlfW6/GuYIxKYgEKCFHFqJATAG
IxHrq1PDOiSwXd2GmVVYyEmhZnbcp8CxaEMQoevxAta0ssMK3w6UsDtvUvYvF22m
qQKBiD5GwESzsFPy3Ga0MvZpn3D6EJQLgsnrtUPZx+z2Ep2x0xc5orneB5fGyF1P
WtP+fG5Q6Dpdz3LRfm+KwBCWFKQjg7uTxcjerhBWEYPmEMKYwTJF5PBG9/ddvHLQ
EQeNC8fHGg4UXU8mhHnSBt3EA10qQJfRDs15M38eG2cYwB1PZpDHScDnDA0=
-----END RSA PRIVATE KEY-----`

	pubRSA1 = `-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAhAKYdtoeoy8zcAcR874L8cnZxKzAGwd7v36APp7Pv6Q2jdsPBRrw
WEBnez6d0UDKDwGbc6nxfEXAy5mbhgajzrw3MOEt8uA5txSKobBpKDeBLOsdJKFq
MGmXCQvEG7YemcxDTRPxAleIAgYYRjTSd/QBwVW9OwNFhekro3RtlinV0a75jfZg
kne/YiktSvLG34lw2zqXBDTC5NHROUqGTlML4PlNZS5Ri2U4aCNx2rUPRcKIlE0P
uKxI4T+HIaFpv8+rdV6eUgOrB2xeI1dSFFn/nnv5OoZJEIB+VmuKn3DCUcCZSFlQ
PSXSfBDiUGhwOw76WuSSsf1D4b/vLoJ10wIDAQAB
-----END RSA PUBLIC KEY-----
`

	prvKeyRSA1 = func() *rsa.PrivateKey {
		prvBlockRSA, _ := pem.Decode([]byte(prvRSA1))
		k, _ := x509.ParsePKCS1PrivateKey(prvBlockRSA.Bytes)
		return k
	}()
)

func actorKeyMock(pem string) PublicKey {
	return PublicKey{
		ID:           "https://example.com/~johndoe#main",
		Owner:        "https://example.com/~johndoe",
		PublicKeyPem: pem,
	}
}

func TestPublicKey_CryptoValue(t *testing.T) {
	tests := []struct {
		name    string
		key     PublicKey
		want    crypto.PublicKey
		wantErr error
	}{
		{
			name:    "empty",
			key:     PublicKey{},
			wantErr: errors.Newf("unable to decode PEM payload for public key"),
		},
		{
			name: "rsa x509",
			key:  actorKeyMock(pubRSA),
			want: prvKeyRSA.Public(),
		},
		{
			name: "rsa pkcs#1",
			key:  actorKeyMock(pubRSA1),
			want: prvKeyRSA1.Public(),
		},
		{
			name: "ecdsa",
			key:  actorKeyMock(pubECDSA),
			want: prvKeyECDSA.Public(),
		},
		{
			name: "ed25519",
			key:  actorKeyMock(pubEd25519),
			want: prvKeyEd25519.Public(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.key.CryptoValue()
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("CryptoValue() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
				return
			}
			if !cmp.Equal(got, tt.want, EquatePublicKeys) {
				t.Errorf("CryptoValue() got = %s", cmp.Diff(tt.want, got, EquatePublicKeys))
			}
		})
	}
}

func arePubKeys(a, b any) bool {
	_, ok1 := a.(cryptoPubKey)
	_, ok2 := b.(cryptoPubKey)
	return ok1 && ok2
}

type cryptoPubKey interface {
	Equal(key crypto.PublicKey) bool
}

func comparePubKeys(x, y any) bool {
	if rcmp, ok := x.(cryptoPubKey); ok {
		return rcmp.Equal(y)
	}
	return false
}

var EquatePublicKeys = cmp.FilterValues(arePubKeys, cmp.Comparer(comparePubKeys))
