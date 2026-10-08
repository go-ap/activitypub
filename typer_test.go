package activitypub

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPathTyper_Type(t *testing.T) {
	t.Skipf("TODO")
}

func TestValidActivityCollection(t *testing.T) {
	tests := []struct {
		name string
		typ  CollectionPath
		want bool
	}{
		{
			name: "empty",
			typ:  "",
			want: false,
		},
		{
			name: "invalid",
			typ:  "invalid",
			want: false,
		},
		{
			name: "valid",
			typ:  "inbox",
			want: true,
		},
		{
			name: "valid-outbox",
			typ:  Outbox,
			want: true,
		},
		{
			name: "invalid-replies",
			typ:  Replies,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidActivityCollection(tt.typ); got != tt.want {
				t.Errorf("ValidActivityCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidCollection(t *testing.T) {
	tests := []struct {
		name string
		typ  CollectionPath
		want bool
	}{
		{
			name: "empty",
			typ:  "",
			want: false,
		},
		{
			name: "invalid",
			typ:  "invalid",
			want: false,
		},
		{
			name: "valid",
			typ:  "inbox",
			want: true,
		},
		{
			name: "valid-outbox",
			typ:  Outbox,
			want: true,
		},
		{
			name: "valid-replies",
			typ:  Replies,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidCollection(tt.typ); got != tt.want {
				t.Errorf("ValidCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidObjectCollection(t *testing.T) {
	tests := []struct {
		name string
		typ  CollectionPath
		want bool
	}{
		{
			name: "empty",
			typ:  "",
			want: false,
		},
		{
			name: "invalid",
			typ:  "invalid",
			want: false,
		},
		{
			name: "valid",
			typ:  "followers",
			want: true,
		},
		{
			name: "invalid-outbox",
			typ:  Outbox,
			want: false,
		},
		{
			name: "valid-replies",
			typ:  Replies,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidObjectCollection(tt.typ); got != tt.want {
				t.Errorf("ValidObjectCollection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidCollectionIRI(t *testing.T) {
	tests := []struct {
		name string
		i    IRI
		want bool
	}{
		{
			name: "empty",
			i:    "",
			want: false,
		},
		{
			name: "not-a-collection",
			i:    "http://example.com",
			want: false,
		},
		{
			name: "inbox",
			i:    "http://example.com/inbox",
			want: true,
		},
		{
			name: string(Shares),
			i:    Shares.IRI(IRI("http://example.com")),
			want: true,
		},
		{
			name: string(Followers),
			i:    Followers.IRI(IRI("http://example.com")),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidCollectionIRI(tt.i); got != tt.want {
				t.Errorf("ValidCollectionIRI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		name     string
		i        IRI
		wantIRI  IRI
		wantPath CollectionPath
	}{
		{
			name: "empty",
		},
		{
			name:     "no-path",
			i:        "http://example.com",
			wantIRI:  "http://example.com",
			wantPath: "",
		},
		{
			name:     "with-inbox",
			i:        "http://example.com/inbox",
			wantIRI:  "http://example.com",
			wantPath: "inbox",
		},
		{
			name:     string(Inbox),
			i:        Inbox.IRI(IRI("http://example.com")),
			wantIRI:  "http://example.com",
			wantPath: Inbox,
		},
		{
			name:     string(Outbox),
			i:        Outbox.IRI(IRI("http://example.com")),
			wantIRI:  "http://example.com",
			wantPath: Outbox,
		},
		{
			name:     string(Followers),
			i:        Followers.IRI(IRI("http://example.com")),
			wantIRI:  "http://example.com",
			wantPath: Followers,
		},
		{
			name:     string(Replies),
			i:        Replies.IRI(IRI("http://example.com")),
			wantIRI:  "http://example.com",
			wantPath: Replies,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIRI, gotPath := Split(tt.i)
			if gotIRI != tt.wantIRI {
				t.Errorf("Split() got = %v, want %v", gotIRI, tt.wantIRI)
			}
			if gotPath != tt.wantPath {
				t.Errorf("Split() got1 = %v, want %v", gotPath, tt.wantPath)
			}
		})
	}
}

func TestCollectionTypes_Of(t *testing.T) {
	type args struct {
		o Item
		t CollectionPath
	}
	tests := []struct {
		name string
		args args
		want Item
	}{
		{
			name: "nil from nil object",
			args: args{
				o: nil,
				t: "likes",
			},
			want: nil,
		},
		{
			name: "nil from invalid CollectionPath type",
			args: args{
				o: Object{
					Likes: IRI("test"),
				},
				t: "like",
			},
			want: nil,
		},
		{
			name: "nil from nil CollectionPath type",
			args: args{
				o: Object{
					Likes: nil,
				},
				t: "likes",
			},
			want: nil,
		},
		{
			name: "get likes iri",
			args: args{
				o: Object{
					Likes: IRI("test"),
				},
				t: "likes",
			},
			want: IRI("test"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if ob := test.args.t.Of(test.args.o); ob != test.want {
				t.Errorf("Object received %#v is different, expected #%v", ob, test.want)
			}
		})
	}
}

func TestCollectionType_IRI(t *testing.T) {
	type args struct {
		o Item
		t CollectionPath
	}
	tests := []struct {
		name string
		args args
		want IRI
	}{
		{
			name: "just path from nil object",
			args: args{
				o: nil,
				t: "likes",
			},
			want: IRI(""),
		},
		{
			name: "emptyIRI from invalid CollectionPath type",
			args: args{
				o: Object{
					Likes: IRI("test"),
				},
				t: "like",
			},
			want: "",
		},
		{
			name: "just path from object without ID",
			args: args{
				o: Object{},
				t: "likes",
			},
			want: IRI(""),
		},
		{
			name: "likes iri on object",
			args: args{
				o: Object{
					ID:    "http://example.com",
					Likes: IRI("test"),
				},
				t: "likes",
			},
			want: IRI("test"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if ob := test.args.t.IRI(test.args.o); ob != test.want {
				t.Errorf("IRI received '%s' is different, expected '%s'", ob, test.want)
			}
		})
	}
}

func TestCollectionType_OfActor(t *testing.T) {
	t.Skipf("TODO")
}

func TestCollectionTypes_Contains(t *testing.T) {
	t.Skipf("TODO")
}

func TestIRIf(t *testing.T) {
	type args struct {
		i IRI
		t CollectionPath
	}
	tests := []struct {
		name string
		args args
		want IRI
	}{
		{
			name: "nil iri",
			args: args{
				i: Object{}.ID,
				t: "inbox",
			},
			want: "",
		},
		{
			name: "empty iri",
			args: args{
				i: "",
				t: "inbox",
			},
			want: "",
		},
		{
			name: "plain concat",
			args: args{
				i: "https://example.com",
				t: "inbox",
			},
			want: "https://example.com/inbox",
		},
		{
			name: "strip root from iri",
			args: args{
				i: "https://example.com/",
				t: "inbox",
			},
			want: "https://example.com/inbox",
		},
		{
			name: "invalid iri",
			args: args{
				i: "example.com",
				t: "test",
			},
			want: "example.com/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IRIf(tt.args.i, tt.args.t); got != tt.want {
				t.Errorf("IRIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollectionPath_AddTo(t *testing.T) {
	tests := []struct {
		name        string
		t           CollectionPath
		i           Item
		wantIRI     IRI
		wantSuccess bool
	}{
		{
			name:        "empty",
			t:           "",
			i:           nil,
			wantIRI:     EmptyIRI,
			wantSuccess: false,
		},
		{
			name:        "simple",
			t:           "test",
			i:           &Object{ID: "http://example.com/addTo"},
			wantIRI:     "http://example.com/addTo/test",
			wantSuccess: true,
		},
		{
			name:        "on-nil-item",
			t:           "test",
			i:           Item(nil),
			wantIRI:     EmptyIRI,
			wantSuccess: false,
		},
		{
			name:        "on-nil",
			t:           "test",
			i:           nil,
			wantIRI:     EmptyIRI,
			wantSuccess: false,
		},
		{
			name:        "on-nil-object",
			t:           "test",
			i:           o,
			wantIRI:     EmptyIRI,
			wantSuccess: false,
		},
		{
			name:        "on-nil-item",
			t:           "test",
			i:           Item(nil),
			wantIRI:     EmptyIRI,
			wantSuccess: false,
		},
		{
			name:        "inbox-on-Actor",
			t:           Inbox,
			i:           &Actor{ID: "http://example.com/~jdoe"},
			wantIRI:     "http://example.com/~jdoe/inbox",
			wantSuccess: true,
		},
		{
			name:        "outbox-on-Actor",
			t:           Outbox,
			i:           &Actor{ID: "http://example.com/~jdoe"},
			wantIRI:     "http://example.com/~jdoe/outbox",
			wantSuccess: true,
		},
		{
			name:        "liked-on-Actor",
			t:           Liked,
			i:           &Actor{ID: "http://example.com/~jdoe"},
			wantIRI:     "http://example.com/~jdoe/liked",
			wantSuccess: true,
		},
		{
			name:        "likes-on-Actor",
			t:           Likes,
			i:           &Actor{ID: "http://example.com/~jdoe"},
			wantIRI:     "http://example.com/~jdoe/likes",
			wantSuccess: true,
		},
		{
			name:        "following-on-Actor",
			t:           Following,
			i:           &Actor{ID: "http://example.com/~jdoe"},
			wantIRI:     "http://example.com/~jdoe/following",
			wantSuccess: true,
		},
		{
			name:        "followers-on-Actor",
			t:           Followers,
			i:           &Actor{ID: "http://example.com/~jdoe"},
			wantIRI:     "http://example.com/~jdoe/followers",
			wantSuccess: true,
		},
		{
			name:        "likes-on-object",
			t:           Likes,
			i:           &Object{ID: "http://example.com/ob-1"},
			wantIRI:     "http://example.com/ob-1/likes",
			wantSuccess: true,
		},
		{
			name:        "shares-on-object",
			t:           Shares,
			i:           &Object{ID: "http://example.com/ob-1"},
			wantIRI:     "http://example.com/ob-1/shares",
			wantSuccess: true,
		},
		{
			name:        "replies-on-object",
			t:           Replies,
			i:           &Object{ID: "http://example.com/ob-1"},
			wantIRI:     "http://example.com/ob-1/replies",
			wantSuccess: true,
		},
		{
			name:        "random-path-on-Object",
			t:           "random-path",
			i:           &Object{ID: "http://example.com/ob-1"},
			wantIRI:     "http://example.com/ob-1/random-path",
			wantSuccess: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIRI, gotOK := tt.t.AddTo(tt.i)
			if gotIRI != tt.wantIRI {
				t.Errorf("AddTo() got IRI = %v, want %v", gotIRI, tt.wantIRI)
			}
			if gotOK != tt.wantSuccess {
				t.Errorf("AddTo() got  OK = %v, want %v", gotOK, tt.wantSuccess)
			}
		})
	}
}

func TestCollectionPaths_Split(t *testing.T) {
	tests := []struct {
		name       string
		t          CollectionPaths
		given      IRI
		maybeActor IRI
		maybeCol   CollectionPath
	}{
		{
			name:       "empty",
			t:          nil,
			given:      "",
			maybeActor: "",
			maybeCol:   "",
		},
		{
			name:       "nil with example.com",
			t:          nil,
			given:      "example.com",
			maybeActor: "example.com",
			maybeCol:   "",
		},
		{
			name:       "nil with https://example.com",
			t:          nil,
			given:      "https://example.com/",
			maybeActor: "https://example.com",
			maybeCol:   Unknown,
		},
		{
			name:       "outbox with https://example.com/outbox",
			t:          CollectionPaths{Outbox},
			given:      "https://example.com/outbox",
			maybeActor: "https://example.com",
			maybeCol:   Outbox,
		},
		{
			name:       "{outbox,inbox} with https://example.com/inbox",
			t:          CollectionPaths{Outbox, Inbox},
			given:      "https://example.com/inbox",
			maybeActor: "https://example.com",
			maybeCol:   Inbox,
		},
		{
			name:       "IRI with query params",
			t:          CollectionPaths{Outbox, Inbox},
			given:      "https://example.com/inbox?maxItems=100",
			maybeActor: "https://example.com",
			maybeCol:   Inbox,
		},
		{
			name:       "IRI with fragment",
			t:          CollectionPaths{Outbox, Inbox},
			given:      "https://example.com/outbox#top",
			maybeActor: "https://example.com",
			maybeCol:   Outbox,
		},
		{
			// TODO(marius): This feels wrong.
			name:       "outbox with https://example.com/inbox",
			t:          CollectionPaths{Outbox},
			given:      "https://example.com/inbox",
			maybeActor: "https://example.com",
			maybeCol:   Unknown,
		},
		{
			name:       "invalid url",
			t:          CollectionPaths{Inbox},
			given:      "127.0.0.1:666/inbox",
			maybeActor: "127.0.0.1:666",
			maybeCol:   Inbox,
		},
		{
			name:       "invalid url - collection doesn't match",
			t:          CollectionPaths{Outbox},
			given:      "127.0.0.1:666/inbox",
			maybeActor: "127.0.0.1:666/inbox",
			maybeCol:   Unknown,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ma, mc := tt.t.Split(tt.given)
			if ma != tt.maybeActor {
				t.Errorf("Split() got Actor = %q, want %q", ma, tt.maybeActor)
			}
			if mc != tt.maybeCol {
				t.Errorf("Split() got Colletion Path = %q, want %q", mc, tt.maybeCol)
			}
		})
	}
}

func TestCollectionPath_Of(t *testing.T) {
	tests := []struct {
		name string
		t    CollectionPath
		arg  Item
		want Item
	}{
		{
			name: "all-nil",
			t:    "",
		},
		{
			name: "inbox-nil",
			t:    Inbox,
		},
		{
			name: "outbox-nil",
			t:    Outbox,
		},
		{
			name: "followers-nil",
			t:    Followers,
		},
		{
			name: "following-nil",
			t:    Following,
		},
		{
			name: "liked-nil",
			t:    Liked,
		},
		{
			name: "likes-nil",
			t:    Likes,
		},
		{
			name: "shares-nil",
			t:    Shares,
		},
		{
			name: "replies-nil",
			t:    Replies,
		},
		{
			name: "inbox-empty",
			t:    Inbox,
			arg:  &Actor{},
		},
		{
			name: "outbox-empty",
			t:    Outbox,
			arg:  &Actor{},
		},
		{
			name: "followers-empty",
			t:    Followers,
			arg:  &Actor{},
		},
		{
			name: "following-empty",
			t:    Following,
			arg:  &Actor{},
		},
		{
			name: "liked-empty",
			t:    Liked,
			arg:  &Actor{},
		},
		{
			name: "likes-empty",
			t:    Likes,
			arg:  &Object{},
		},
		{
			name: "shares-empty",
			t:    Shares,
			arg:  &Object{},
		},
		{
			name: "replies-empty",
			t:    Replies,
			arg:  &Object{},
		},
		//
		{
			name: "inbox",
			t:    Inbox,
			arg: &Actor{
				Type:  PersonType,
				Inbox: IRI("https://example.com/inbox"),
			},
			want: IRI("https://example.com/inbox"),
		},
		{
			name: "outbox",
			t:    Outbox,
			arg: &Actor{
				Type:   PersonType,
				Outbox: IRI("https://example.com/outbox"),
			},
			want: IRI("https://example.com/outbox"),
		},
		{
			name: "followers",
			t:    Followers,
			arg: &Actor{
				Type:      GroupType,
				Followers: IRI("https://example.com/c132-333"),
			},
			want: IRI("https://example.com/c132-333"),
		},
		{
			name: "following",
			t:    Following,
			arg: &Actor{
				Type:      GroupType,
				Following: IRI("https://example.com/c666-333"),
			},
			want: IRI("https://example.com/c666-333"),
		},
		{
			name: "liked",
			t:    Liked,
			arg: &Actor{
				Type:  ApplicationType,
				Liked: IRI("https://example.com/l666"),
			},
			want: IRI("https://example.com/l666"),
		},
		{
			name: "likes",
			t:    Likes,
			arg: &Object{
				Type:  NoteType,
				Likes: IRI("https://example.com/l166"),
			},
			want: IRI("https://example.com/l166"),
		},
		{
			name: "shares",
			t:    Shares,
			arg: &Object{
				Type:   PageType,
				Shares: IRI("https://example.com/s266"),
			},
			want: IRI("https://example.com/s266"),
		},
		{
			name: "replies",
			t:    Replies,
			arg: &Object{
				Type:    ArticleType,
				Replies: IRI("https://example.com/r466"),
			},
			want: IRI("https://example.com/r466"),
		},
		{
			name: "replies-on-iri",
			t:    Replies,
			arg:  IRI("http://example.com"),
			want: IRI("http://example.com/replies"),
		},
		{
			name: "replies-on-item-collection",
			t:    Likes,
			arg: ItemCollection{
				Object{ID: "http://example.com/ob1", Likes: IRI("http://example.com/ob1/likes")},
				Actor{ID: "http://example.com/~jdoe", Likes: IRI("http://example.com/~jdoe/likes")},
			},
			want: ItemCollection{
				IRI("http://example.com/ob1/likes"),
				IRI("http://example.com/~jdoe/likes"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Of(tt.arg); !cmp.Equal(got, tt.want) {
				t.Errorf("Of() = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
