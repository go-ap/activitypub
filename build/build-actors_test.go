package build

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"testing"

	vocab "github.com/go-ap/activitypub"
	"github.com/google/go-cmp/cmp"
)

func TestActor(t *testing.T) {
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := pk.Public()
	pem, _ := pemPublicKey(pub)

	tests := []struct {
		name string
		args []InitFn
		want *vocab.Actor
	}{
		{
			name: "empty",
			want: &vocab.Actor{},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Actor{ID: "http://example.com"},
		},
		{
			name: "WithType",
			args: []InitFn{WithType("invalid-type")},
			want: &vocab.Actor{Type: vocab.ActivityVocabularyType("invalid-type")},
		},
		{
			name: "WithPreferredUsername",
			args: []InitFn{WithPreferredUsername("Jane Doe")},
			want: &vocab.Actor{PreferredUsername: vocab.NaturalLanguageValues{vocab.DefaultLang: vocab.Content("Jane Doe")}},
		},
		{
			name: "WithInbox_null",
			args: []InitFn{WithID("http://example.com/~jdoe"), WithInbox()},
			want: &vocab.Actor{ID: "http://example.com/~jdoe", Inbox: vocab.IRI("http://example.com/~jdoe/inbox")},
		},
		{
			name: "WithInbox",
			args: []InitFn{WithInbox(vocab.IRI("http://example.com/inbox"))},
			want: &vocab.Actor{Inbox: vocab.IRI("http://example.com/inbox")},
		},
		{
			name: "WithOutbox_null",
			args: []InitFn{WithID("http://example.com/~jdoe"), WithOutbox()},
			want: &vocab.Actor{ID: "http://example.com/~jdoe", Outbox: vocab.IRI("http://example.com/~jdoe/outbox")},
		},
		{
			name: "WithOutbox",
			args: []InitFn{WithOutbox(vocab.IRI("http://example.com/inbox"))},
			want: &vocab.Actor{Outbox: vocab.IRI("http://example.com/inbox")},
		},
		{
			name: "WithLiked_null",
			args: []InitFn{WithID("http://example.com/~jdoe"), WithLiked()},
			want: &vocab.Actor{ID: "http://example.com/~jdoe", Liked: vocab.IRI("http://example.com/~jdoe/liked")},
		},
		{
			name: "WithLiked",
			args: []InitFn{WithLiked(vocab.IRI("http://example.com/inbox"))},
			want: &vocab.Actor{Liked: vocab.IRI("http://example.com/inbox")},
		},
		{
			name: "WithFollowing_null",
			args: []InitFn{WithID("http://example.com/~jdoe"), WithFollowing()},
			want: &vocab.Actor{ID: "http://example.com/~jdoe", Following: vocab.IRI("http://example.com/~jdoe/following")},
		},
		{
			name: "WithFollowing",
			args: []InitFn{WithFollowing(vocab.IRI("http://example.com/inbox"))},
			want: &vocab.Actor{Following: vocab.IRI("http://example.com/inbox")},
		},
		{
			name: "WithFollowers_null",
			args: []InitFn{WithID("http://example.com/~jdoe"), WithFollowers()},
			want: &vocab.Actor{ID: "http://example.com/~jdoe", Followers: vocab.IRI("http://example.com/~jdoe/followers")},
		},
		{
			name: "WithFollowers",
			args: []InitFn{WithFollowers(vocab.IRI("http://example.com/inbox"))},
			want: &vocab.Actor{Followers: vocab.IRI("http://example.com/inbox")},
		},
		{
			name: "WithAuthEndpoint",
			args: []InitFn{WithAuthEndpoint("http://example.com/auth")},
			want: &vocab.Actor{Endpoints: &vocab.Endpoints{OauthAuthorizationEndpoint: vocab.IRI("http://example.com/auth")}},
		},
		{
			name: "WithTokenEndpoint",
			args: []InitFn{WithTokenEndpoint("http://example.com/token")},
			want: &vocab.Actor{Endpoints: &vocab.Endpoints{OauthTokenEndpoint: vocab.IRI("http://example.com/token")}},
		},
		{
			name: "WithProxyURL",
			args: []InitFn{WithProxyURL("http://example.com/proxyUrl")},
			want: &vocab.Actor{Endpoints: &vocab.Endpoints{ProxyURL: vocab.IRI("http://example.com/proxyUrl")}},
		},
		{
			name: "WithSharedInbox",
			args: []InitFn{WithSharedInbox("http://example.com/inbox")},
			want: &vocab.Actor{Endpoints: &vocab.Endpoints{SharedInbox: vocab.IRI("http://example.com/inbox")}},
		},
		{
			name: "WithStreams",
			args: []InitFn{WithStreams(vocab.IRI("http://example.com/s1"), vocab.IRI("http://example.com/s2"))},
			want: &vocab.Actor{Streams: vocab.ItemCollection{vocab.IRI("http://example.com/s1"), vocab.IRI("http://example.com/s2")}},
		},
		{
			name: "WithPublicKey, no id",
			args: []InitFn{WithPublicKey(PublicKey("http://example.com", pub))},
			want: &vocab.Actor{PublicKey: vocab.PublicKey{ID: "http://example.com", PublicKeyPem: pem}},
		},
		{
			name: "WithPublicKey",
			args: []InitFn{WithID("http://example.com/~jdoe"), WithPublicKey(PublicKey("", pub))},
			want: &vocab.Actor{ID: "http://example.com/~jdoe", PublicKey: vocab.PublicKey{ID: "http://example.com/~jdoe#main", Owner: "http://example.com/~jdoe", PublicKeyPem: pem}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Actor(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Actor() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestPerson(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Person
	}{
		{
			name: "empty",
			want: &vocab.Person{Type: vocab.PersonType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Person{ID: "http://example.com", Type: vocab.PersonType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Person(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Person() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestApplication(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Application
	}{
		{
			name: "empty",
			want: &vocab.Application{Type: vocab.ApplicationType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Application{ID: "http://example.com", Type: vocab.ApplicationType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Application(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Application() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestGroup(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Group
	}{
		{
			name: "empty",
			want: &vocab.Group{Type: vocab.GroupType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Group{ID: "http://example.com", Type: vocab.GroupType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Group(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Group() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestService(t *testing.T) {
	tests := []struct {
		name string
		args []InitFn
		want *vocab.Service
	}{
		{
			name: "empty",
			want: &vocab.Service{Type: vocab.ServiceType},
		},
		{
			name: "WithID",
			args: []InitFn{WithID("http://example.com")},
			want: &vocab.Service{ID: "http://example.com", Type: vocab.ServiceType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Service(tt.args...); !cmp.Equal(got, tt.want) {
				t.Errorf("Service() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}

func TestPublicKey(t *testing.T) {
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)
	pem, _ := pemPublicKey(pk.Public())

	type args struct {
		id vocab.IRI
		k  crypto.PublicKey
	}

	tests := []struct {
		name string
		args args
		want vocab.PublicKey
	}{
		{
			name: "empty",
			args: args{},
			want: vocab.PublicKey{},
		},
		{
			name: "withIRI",
			args: args{
				id: "http://example.com#main",
			},
			want: vocab.PublicKey{ID: "http://example.com#main"},
		},
		{
			name: "withPublicKey",
			args: args{k: &pk.PublicKey},
			want: vocab.PublicKey{PublicKeyPem: pem},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PublicKey(tt.args.id, tt.args.k); !cmp.Equal(got, tt.want) {
				t.Errorf("PublicKey() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
