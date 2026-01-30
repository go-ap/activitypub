package activitypub

import (
	"testing"
)

func TestActivityVocabularyTypes_EmptyTypes(t *testing.T) {
	if !EmptyTypes(NilType) {
		t.Errorf("expected NilType to be empty")
	}
	if !EmptyTypes() {
		t.Errorf("expected nil ActivityVocabularyType collection to be empty")
	}
	if !EmptyTypes(nil...) {
		t.Errorf("expected nil ActivityVocabularyType collection to be empty")
	}
	if !EmptyTypes(ActivityVocabularyTypes{}...) {
		t.Errorf("expected default ActivityVocabularyTypes slice to be empty")
	}
	if !EmptyTypes(ActivityVocabularyTypes{NilType, ""}...) {
		t.Errorf("expected ActivityVocabularyTypes slice of all NilType to be empty")
	}
	if EmptyTypes(ActivityVocabularyTypes{NilType, ObjectType}...) {
		t.Errorf("expected ActivityVocabularyTypes slice with non-NilType to not be empty")
	}
	if EmptyTypes(ActivityVocabularyTypes{"X"}...) {
		t.Errorf("expected ActivityVocabularyTypes slice with non-NilType to not be empty")
	}
}

func TestActivityVocabularyTypes_Match(t *testing.T) {
	tests := []struct {
		name string
		t    ActivityVocabularyTypes
		arg  Typer
		want bool
	}{
		{
			name: "nil matches nil",
			want: true,
		},
		{
			name: "nil matches empty types",
			t:    ActivityVocabularyTypes{},
			want: true,
		},
		{
			name: "nil matches []NilType",
			arg:  ActivityVocabularyTypes{NilType},
			want: true,
		},
		{
			name: "nil matches NilType",
			arg:  NilType,
			want: true,
		},
		{
			name: "NilType matches nil",
			t:    ActivityVocabularyTypes{NilType},
			want: true,
		},
		{
			name: "empty matches []NilType",
			t:    ActivityVocabularyTypes{},
			arg:  ActivityVocabularyTypes{NilType},
			want: true,
		},
		{
			name: "empty matches NilType",
			t:    ActivityVocabularyTypes{},
			arg:  NilType,
			want: true,
		},
		{
			name: `empty matches ""`,
			t:    ActivityVocabularyTypes{},
			arg:  ActivityVocabularyTypes{""},
			want: true,
		},
		{
			name: `empty matches ""`,
			t:    ActivityVocabularyTypes{},
			arg:  ActivityVocabularyType(""),
			want: true,
		},
		{
			name: `empty does not match single type`,
			t:    ActivityVocabularyTypes{},
			arg:  ActivityVocabularyTypes{"X"},
			want: false,
		},
		{
			name: `empty does not match single type`,
			t:    ActivityVocabularyTypes{},
			arg:  ActivityVocabularyType("X"),
			want: false,
		},
		{
			name: `empty does not match multiple types`,
			t:    ActivityVocabularyTypes{},
			arg:  ActivityVocabularyTypes{"X", "Y"},
			want: false,
		},
		{
			name: "NilType matches empty types",
			t:    ActivityVocabularyTypes{NilType},
			arg:  ActivityVocabularyTypes{},
			want: true,
		},
		{
			name: "NilType matches NilType",
			t:    ActivityVocabularyTypes{NilType},
			arg:  ActivityVocabularyTypes{NilType},
			want: true,
		},
		{
			name: `"" matches nil`,
			t:    ActivityVocabularyTypes{""},
			want: true,
		},
		{
			name: `"" matches empty types`,
			t:    ActivityVocabularyTypes{""},
			arg:  ActivityVocabularyTypes{},
			want: true,
		},
		{
			name: `"" matches NilType`,
			t:    ActivityVocabularyTypes{""},
			arg:  ActivityVocabularyTypes{NilType},
			want: true,
		},
		{
			name: "nil does not match X",
			arg:  ActivityVocabularyTypes{"X"},
			want: false,
		},
		{
			name: `"" does not match X`,
			t:    ActivityVocabularyTypes{""},
			arg:  ActivityVocabularyTypes{"X"},
			want: false,
		},
		{
			name: "NilType does not match X",
			t:    ActivityVocabularyTypes{NilType},
			arg:  ActivityVocabularyTypes{"X"},
			want: false,
		},
		{
			name: "X matches X",
			t:    ActivityVocabularyTypes{"X"},
			arg:  ActivityVocabularyTypes{"X"},
			want: true,
		},
		{
			name: "X matches X,Y",
			t:    ActivityVocabularyTypes{"X"},
			arg:  ActivityVocabularyTypes{"X", "Y"},
			want: true,
		},
		{
			name: "X matches Y,X",
			t:    ActivityVocabularyTypes{"X"},
			arg:  ActivityVocabularyTypes{"Y", "X"},
			want: true,
		},
		{
			name: "X,Y matches []X",
			t:    ActivityVocabularyTypes{"X", "Y"},
			arg:  ActivityVocabularyTypes{"X"},
			want: true,
		},
		{
			name: "X,Y matches X",
			t:    ActivityVocabularyTypes{"X", "Y"},
			arg:  ActivityVocabularyType("X"),
			want: true,
		},
		{
			name: "X,Y matches []Y",
			t:    ActivityVocabularyTypes{"X", "Y"},
			arg:  ActivityVocabularyTypes{"Y"},
			want: true,
		},
		{
			name: "X,Y matches Y",
			t:    ActivityVocabularyTypes{"X", "Y"},
			arg:  ActivityVocabularyType("Y"),
			want: true,
		},
		{
			name: "X,Y matches X,Y",
			t:    ActivityVocabularyTypes{"X", "Y"},
			arg:  ActivityVocabularyTypes{"X", "Y"},
			want: true,
		},
		{
			name: "X,Y matches Y,X",
			t:    ActivityVocabularyTypes{"X", "Y"},
			arg:  ActivityVocabularyTypes{"Y", "X"},
			want: true,
		},
		{
			name: `XY does not match empty`,
			t:    ActivityVocabularyTypes{"X", "Y"},
			arg:  ActivityVocabularyTypes{},
			want: false,
		},
		{
			name: "NilType,X,Y to check, matches NilType ",
			t:    ActivityVocabularyTypes{NilType, "X", "Y"},
			arg:  ActivityVocabularyTypes{NilType},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Match(tt.arg); got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityVocabularyType_Match(t *testing.T) {
	tests := []struct {
		name      string
		a         ActivityVocabularyType
		arg       Typer
		wantMatch bool
	}{
		{
			name:      "nil matches nil",
			wantMatch: true,
		},
		{
			name:      "nil matches empty types",
			arg:       ActivityVocabularyTypes{},
			wantMatch: true,
		},
		{
			name:      "nil matches NilType",
			arg:       ActivityVocabularyTypes{NilType},
			wantMatch: true,
		},
		{
			name:      "NilType matches nil",
			a:         NilType,
			wantMatch: true,
		},
		{
			name:      "NilType matches empty types",
			a:         NilType,
			arg:       ActivityVocabularyTypes{},
			wantMatch: true,
		},
		{
			name:      "NilType matches NilType",
			a:         NilType,
			arg:       ActivityVocabularyTypes{NilType},
			wantMatch: true,
		},
		{
			name:      `"" matches nil`,
			a:         "",
			wantMatch: true,
		},
		{
			name:      `"" matches empty types`,
			a:         "",
			arg:       ActivityVocabularyTypes{},
			wantMatch: true,
		},
		{
			name:      `"" matches NilType`,
			a:         "",
			arg:       ActivityVocabularyTypes{NilType},
			wantMatch: true,
		},
		{
			name:      "nil does not match X",
			arg:       ActivityVocabularyTypes{"X"},
			wantMatch: false,
		},
		{
			name:      `"" does not match X`,
			a:         "",
			arg:       ActivityVocabularyTypes{"X"},
			wantMatch: false,
		},
		{
			name:      "NilType does not match X",
			a:         NilType,
			arg:       ActivityVocabularyTypes{"X"},
			wantMatch: false,
		},
		{
			name:      "X matches X",
			a:         ActivityVocabularyType("X"),
			arg:       ActivityVocabularyTypes{"X"},
			wantMatch: true,
		},
		{
			name:      "X matches X,Y",
			a:         ActivityVocabularyType("X"),
			arg:       ActivityVocabularyTypes{"X", "Y"},
			wantMatch: true,
		},
		{
			name:      "X matches Y,X",
			a:         ActivityVocabularyType("X"),
			arg:       ActivityVocabularyTypes{"Y", "X"},
			wantMatch: true,
		},
		{
			name:      "nil matches nil",
			wantMatch: true,
		},
		{
			name:      "nil matches empty types",
			arg:       ActivityVocabularyTypes{},
			wantMatch: true,
		},
		{
			name:      "nil matches NilType",
			arg:       ActivityVocabularyTypes{NilType},
			wantMatch: true,
		},
		{
			name:      "NilType matches nil",
			a:         NilType,
			wantMatch: true,
		},
		{
			name:      "NilType matches empty types",
			a:         NilType,
			arg:       ActivityVocabularyTypes{},
			wantMatch: true,
		},
		{
			name:      "NilType matches NilType",
			a:         NilType,
			arg:       ActivityVocabularyTypes{NilType},
			wantMatch: true,
		},
		{
			name:      `"" matches nil`,
			a:         "",
			wantMatch: true,
		},
		{
			name:      `"" matches empty types`,
			a:         "",
			arg:       ActivityVocabularyTypes{},
			wantMatch: true,
		},
		{
			name:      `"" matches NilType`,
			a:         "",
			arg:       ActivityVocabularyTypes{NilType},
			wantMatch: true,
		},
		{
			name:      "nil does not match X",
			arg:       ActivityVocabularyTypes{"X"},
			wantMatch: false,
		},
		{
			name:      `"" does not match X`,
			a:         "",
			arg:       ActivityVocabularyTypes{"X"},
			wantMatch: false,
		},
		{
			name:      "NilType does not match X",
			a:         NilType,
			arg:       ActivityVocabularyTypes{"X"},
			wantMatch: false,
		},
		{
			name:      "X matches X",
			a:         ActivityVocabularyType("X"),
			arg:       ActivityVocabularyTypes{"X"},
			wantMatch: true,
		},
		{
			name:      "X matches X,Y",
			a:         ActivityVocabularyType("X"),
			arg:       ActivityVocabularyTypes{"X", "Y"},
			wantMatch: true,
		},
		{
			name:      "X matches Y,X",
			a:         ActivityVocabularyType("X"),
			arg:       ActivityVocabularyTypes{"Y", "X"},
			wantMatch: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMatch := tt.a.Match(tt.arg); gotMatch != tt.wantMatch {
				t.Errorf("Match() = %v, want %v", gotMatch, tt.wantMatch)
			}
		})
	}
}

func TestHasTypes(t *testing.T) {
	tests := []struct {
		name string
		it   ActivityObject
		want bool
	}{
		{
			name: "empty",
			want: false,
		},
		{
			name: "no type",
			it:   &Object{},
			want: false,
		},
		{
			name: "NilType",
			it:   &Object{Type: ActivityVocabularyType("")},
			want: false,
		},
		{
			name: "NilType",
			it:   &Object{Type: NilType},
			want: false,
		},
		{
			name: "empty vocabulary types",
			it:   &Object{Type: ActivityVocabularyTypes{}},
			want: false,
		},
		{
			name: "only NilType in vocabulary types",
			it:   &Object{Type: ActivityVocabularyTypes{NilType}},
			want: false,
		},
		{
			name: `only "" in vocabulary types`,
			it:   &Object{Type: ActivityVocabularyTypes{ActivityVocabularyType("")}},
			want: false,
		},
		{
			name: "only empty vocabulary types",
			it:   &Object{Type: ActivityVocabularyTypes{ActivityVocabularyType(""), NilType}},
			want: false,
		},
		{
			name: "X type",
			it:   &Object{Type: ActivityVocabularyType("X")},
			want: true,
		},
		{
			name: "X in vocabulary types",
			it:   &Object{Type: ActivityVocabularyTypes{"X"}},
			want: true,
		},
		{
			name: "X, with empty in vocabulary types",
			it:   &Object{Type: ActivityVocabularyTypes{"X", NilType, ActivityVocabularyType("")}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasTypes(tt.it); got != tt.want {
				t.Errorf("HasTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllTypes(t *testing.T) {
	tests := []struct {
		name    string
		toMatch []ActivityVocabularyType
		toCheck []ActivityVocabularyType
		want    bool
	}{
		{
			name: "nil to match, matches nil to check",
			want: true,
		},
		{
			name:    "nil to match, matches empty to check",
			toCheck: ActivityVocabularyTypes{},
			want:    true,
		},
		{
			name:    "nil to match, matches NilType to check",
			toCheck: ActivityVocabularyTypes{NilType},
			want:    true,
		},
		{
			name:    "nil matches X",
			toCheck: ActivityVocabularyTypes{"X"},
			want:    true,
		},
		{
			name:    "empty to match, matches nil to check",
			toMatch: ActivityVocabularyTypes{},
			want:    true,
		},
		{
			name:    "empty to match, matches empty to check",
			toMatch: ActivityVocabularyTypes{},
			toCheck: ActivityVocabularyTypes{},
			want:    true,
		},
		{
			name:    "empty to match, matches NilType to check",
			toMatch: ActivityVocabularyTypes{},
			toCheck: ActivityVocabularyTypes{NilType},
			want:    true,
		},
		{
			name:    "empty matches X",
			toMatch: ActivityVocabularyTypes{},
			toCheck: ActivityVocabularyTypes{"X"},
			want:    true,
		},
		{
			name:    "empty matches X,Y",
			toMatch: ActivityVocabularyTypes{},
			toCheck: ActivityVocabularyTypes{"X", "Y"},
			want:    true,
		},
		{
			name:    "X matches X",
			toMatch: ActivityVocabularyTypes{"X"},
			toCheck: ActivityVocabularyTypes{"X"},
			want:    true,
		},
		{
			name:    "X matches X,Y",
			toMatch: ActivityVocabularyTypes{"X"},
			toCheck: ActivityVocabularyTypes{"X", "Y"},
			want:    true,
		},
		{
			name:    "X,Y does not match X",
			toMatch: ActivityVocabularyTypes{"X", "Y"},
			toCheck: ActivityVocabularyTypes{"X"},
			want:    false,
		},
		{
			name:    "X,Y matches X,Y",
			toMatch: ActivityVocabularyTypes{"X", "Y"},
			toCheck: ActivityVocabularyTypes{"X", "Y"},
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllTypes(tt.toMatch...).Match(tt.toCheck...)
			if got != tt.want {
				t.Errorf("AllTypes().Match() = %t, want %t", got, tt.want)
			}
		})
	}
}

func TestAnyTypes(t *testing.T) {
	tests := []struct {
		name    string
		toMatch []ActivityVocabularyType
		toCheck []ActivityVocabularyType
		want    bool
	}{
		{
			name: "nil to match, matches nil to check",
			want: true,
		},
		{
			name:    "nil to match, matches empty to check",
			toCheck: ActivityVocabularyTypes{},
			want:    true,
		},
		{
			name:    "nil to match, matches NilType to check",
			toCheck: ActivityVocabularyTypes{NilType},
			want:    true,
		},
		{
			name:    "nil does not match X",
			toCheck: ActivityVocabularyTypes{"X"},
			want:    false,
		},
		{
			name:    "nil does not match X,Y",
			toCheck: ActivityVocabularyTypes{"X", "Y"},
			want:    false,
		},
		{
			name:    "empty to match, matches nil to check",
			toMatch: ActivityVocabularyTypes{},
			want:    true,
		},
		{
			name:    "empty to match, matches empty to check",
			toMatch: ActivityVocabularyTypes{},
			toCheck: ActivityVocabularyTypes{},
			want:    true,
		},
		{
			name:    "empty to match, matches NilType to check",
			toMatch: ActivityVocabularyTypes{},
			toCheck: ActivityVocabularyTypes{NilType},
			want:    true,
		},
		{
			name:    "empty does not match X",
			toMatch: ActivityVocabularyTypes{},
			toCheck: ActivityVocabularyTypes{"X"},
			want:    false,
		},
		{
			name:    "empty does not match X,Y",
			toMatch: ActivityVocabularyTypes{},
			toCheck: ActivityVocabularyTypes{"X", "Y"},
			want:    false,
		},
		{
			name:    "X matches X",
			toMatch: ActivityVocabularyTypes{"X"},
			toCheck: ActivityVocabularyTypes{"X"},
			want:    true,
		},
		{
			name:    "X matches X,Y",
			toMatch: ActivityVocabularyTypes{"X"},
			toCheck: ActivityVocabularyTypes{"X", "Y"},
			want:    true,
		},
		{
			name:    "X,Y matches X",
			toMatch: ActivityVocabularyTypes{"X", "Y"},
			toCheck: ActivityVocabularyTypes{"X"},
			want:    true,
		},
		{
			name:    "X,Y matches X,Y",
			toMatch: ActivityVocabularyTypes{"X", "Y"},
			toCheck: ActivityVocabularyTypes{"X", "Y"},
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnyTypes(tt.toMatch...).Match(tt.toCheck...)
			if got != tt.want {
				t.Errorf("AnyTypes().Match() = %t, want %t", got, tt.want)
			}
		})
	}
}
