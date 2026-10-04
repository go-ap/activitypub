package activitypub

import "testing"

func TestContentOf(t *testing.T) {
	tests := []struct {
		name string
		arg  Item
		want string
	}{
		{
			name: "empty",
			arg:  nil,
			want: "",
		},
		{
			name: "object",
			arg:  &Object{Content: DefaultLangValue("test")},
			want: "test",
		},
		{
			name: "place",
			arg:  &Place{Content: DefaultLangValue("p1")},
			want: "p1",
		},
		{
			name: "profile",
			arg:  &Profile{Content: DefaultLangValue("p2")},
			want: "p2",
		},
		{
			name: "tombstone",
			arg:  &Tombstone{Content: DefaultLangValue("t1")},
			want: "t1",
		},
		{
			name: "relationship",
			arg:  &Relationship{Content: DefaultLangValue("r1")},
			want: "r1",
		},
		{
			name: "activity",
			arg:  &Activity{Content: DefaultLangValue("a1")},
			want: "a1",
		},
		{
			name: "intransitive-activity",
			arg:  &IntransitiveActivity{Content: DefaultLangValue("i1")},
			want: "i1",
		},
		{
			name: "Question",
			arg:  &Question{Content: DefaultLangValue("q1")},
			want: "q1",
		},
		{
			name: "actor",
			arg:  &Actor{Content: DefaultLangValue("p3"), PreferredUsername: DefaultLangValue("test123")},
			want: "p3",
		},
		{
			name: "link",
			arg:  &Link{Name: DefaultLangValue("L3")},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContentOf(tt.arg); got != tt.want {
				t.Errorf("ContentOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNameOf(t *testing.T) {
	tests := []struct {
		name string
		arg  Item
		want string
	}{
		{
			name: "empty",
			arg:  nil,
			want: "",
		},
		{
			name: "object",
			arg:  &Object{Name: DefaultLangValue("test")},
			want: "test",
		},
		{
			name: "place",
			arg:  &Place{Name: DefaultLangValue("p1")},
			want: "p1",
		},
		{
			name: "profile",
			arg:  &Profile{Name: DefaultLangValue("p2")},
			want: "p2",
		},
		{
			name: "tombstone",
			arg:  &Tombstone{Name: DefaultLangValue("t1")},
			want: "t1",
		},
		{
			name: "relationship",
			arg:  &Relationship{Name: DefaultLangValue("r1")},
			want: "r1",
		},
		{
			name: "activity",
			arg:  &Activity{Name: DefaultLangValue("a1")},
			want: "a1",
		},
		{
			name: "intransitive-activity",
			arg:  &IntransitiveActivity{Name: DefaultLangValue("i1")},
			want: "i1",
		},
		{
			name: "Question",
			arg:  &Question{Name: DefaultLangValue("q1")},
			want: "q1",
		},
		{
			name: "actor",
			arg:  &Actor{Name: DefaultLangValue("p3"), PreferredUsername: DefaultLangValue("test123")},
			want: "p3",
		},
		{
			name: "link",
			arg:  &Link{Name: DefaultLangValue("L3")},
			want: "L3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NameOf(tt.arg); got != tt.want {
				t.Errorf("NameOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPreferredNameOf(t *testing.T) {
	tests := []struct {
		name string
		arg  Item
		want string
	}{
		{
			name: "empty",
			arg:  nil,
			want: "",
		},
		{
			name: "object",
			arg:  &Object{Name: DefaultLangValue("test")},
			want: "",
		},
		{
			name: "place",
			arg:  &Place{Name: DefaultLangValue("p1")},
			want: "",
		},
		{
			name: "profile",
			arg:  &Profile{Name: DefaultLangValue("p2")},
			want: "",
		},
		{
			name: "tombstone",
			arg:  &Tombstone{Name: DefaultLangValue("t1")},
			want: "",
		},
		{
			name: "relationship",
			arg:  &Relationship{Name: DefaultLangValue("r1")},
			want: "",
		},
		{
			name: "activity",
			arg:  &Activity{Name: DefaultLangValue("a1")},
			want: "",
		},
		{
			name: "intransitive-activity",
			arg:  &IntransitiveActivity{Name: DefaultLangValue("i1")},
			want: "",
		},
		{
			name: "Question",
			arg:  &Question{Name: DefaultLangValue("q1")},
			want: "",
		},
		{
			name: "actor",
			arg:  &Actor{Name: DefaultLangValue("test"), PreferredUsername: DefaultLangValue("test123")},
			want: "test123",
		},
		{
			name: "link",
			arg:  &Link{Name: DefaultLangValue("L3")},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PreferredNameOf(tt.arg); got != tt.want {
				t.Errorf("PreferredNameOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSummaryOf(t *testing.T) {
	tests := []struct {
		name string
		arg  Item
		want string
	}{
		{
			name: "empty",
			arg:  nil,
			want: "",
		},
		{
			name: "object",
			arg:  &Object{Summary: DefaultLangValue("test")},
			want: "test",
		},
		{
			name: "place",
			arg:  &Place{Summary: DefaultLangValue("p1")},
			want: "p1",
		},
		{
			name: "profile",
			arg:  &Profile{Summary: DefaultLangValue("p2")},
			want: "p2",
		},
		{
			name: "tombstone",
			arg:  &Tombstone{Summary: DefaultLangValue("t1")},
			want: "t1",
		},
		{
			name: "relationship",
			arg:  &Relationship{Summary: DefaultLangValue("r1")},
			want: "r1",
		},
		{
			name: "activity",
			arg:  &Activity{Summary: DefaultLangValue("a1")},
			want: "a1",
		},
		{
			name: "intransitive-activity",
			arg:  &IntransitiveActivity{Summary: DefaultLangValue("i1")},
			want: "i1",
		},
		{
			name: "Question",
			arg:  &Question{Summary: DefaultLangValue("q1")},
			want: "q1",
		},
		{
			name: "actor",
			arg:  &Actor{Summary: DefaultLangValue("p3"), PreferredUsername: DefaultLangValue("test123")},
			want: "p3",
		},
		{
			name: "link",
			arg:  &Link{Name: DefaultLangValue("L3")},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SummaryOf(tt.arg); got != tt.want {
				t.Errorf("SummaryOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
