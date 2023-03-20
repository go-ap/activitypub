package activitypub

import (
	"testing"
)

func TestPathTyper_Type(t *testing.T) {
	t.Skipf("TODO")
}

func TestValidActivityCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestValidCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestValidObjectCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestValidCollectionIRI(t *testing.T) {
	t.Skipf("TODO")
}

func TestSplit(t *testing.T) {
	t.Skipf("TODO")
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
			want: IRI("/likes"),
		},
		{
			name: "emptyIRI from invalid CollectionPath type",
			args: args{
				o: Object{
					Likes: IRI("test"),
				},
				t: "like",
			},
			want: "/like",
		},
		{
			name: "just path from object without ID",
			args: args{
				o: Object{},
				t: "likes",
			},
			want: IRI("/likes"),
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
				t.Errorf("IRI received %q is different, expected %q", ob, test.want)
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
			name: "empty iri",
			args: args{
				i: "",
				t: "inbox",
			},
			want: "/inbox",
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

func TestCollectionType_AddTo(t *testing.T) {
	type args struct {
		i Item
	}
	var i Item
	var o *Object
	tests := []struct {
		name  string
		t     CollectionPath
		args  args
		want  IRI
		want1 bool
	}{
		{
			name: "simple",
			t:    "test",
			args: args{
				i: &Object{ID: "http://example.com/addTo"},
			},
			want:  "http://example.com/addTo/test",
			want1: false, // this seems to always be false
		},
		{
			name: "on-nil-item",
			t:    "test",
			args: args{
				i: i,
			},
			want:  NilIRI,
			want1: false,
		},
		{
			name: "on-nil",
			t:    "test",
			args: args{
				i: nil,
			},
			want:  NilIRI,
			want1: false,
		},
		{
			name: "on-nil-object",
			t:    "test",
			args: args{
				i: o,
			},
			want:  NilIRI,
			want1: false,
		},
		{
			name: "on-nil-item",
			t:    "test",
			args: args{
				i: i,
			},
			want:  NilIRI,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.t.AddTo(tt.args.i)
			if got != tt.want {
				t.Errorf("AddTo() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AddTo() got1 = %v, want %v", got1, tt.want1)
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
			maybeCol:   "",
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
			// TODO(marius): This feels wrong.
			name:       "outbox with https://example.com/inbox",
			t:          CollectionPaths{Outbox},
			given:      "https://example.com/inbox",
			maybeActor: "https://example.com",
			maybeCol:   Unknown,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maybeActor, maybeCol := tt.t.Split(tt.given)
			if maybeActor != tt.maybeActor {
				t.Errorf("Split() got = %v, want %v", maybeActor, tt.maybeActor)
			}
			if maybeCol != tt.maybeCol {
				t.Errorf("Split() got1 = %v, want %v", maybeCol, tt.maybeCol)
			}
		})
	}
}
