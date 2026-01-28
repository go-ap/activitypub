package activitypub

import "testing"

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

func TestActivityVocabularyTypes_Matches(t *testing.T) {
	tests := []struct {
		name string
		t    ActivityVocabularyTypes
		args []ActivityVocabularyType
		want bool
	}{
		{
			name: "nil matches nil",
			want: true,
		},
		{
			name: "empty matches nil",
			t:    ActivityVocabularyTypes{},
			want: true,
		},
		{
			name: "empty matches NilType",
			t:    ActivityVocabularyTypes{},
			args: ActivityVocabularyTypes{NilType},
			want: true,
		},
		{
			name: `empty matches ""`,
			t:    ActivityVocabularyTypes{},
			args: ActivityVocabularyTypes{""},
			want: true,
		},
		{
			name: `empty does not match single type`,
			t:    ActivityVocabularyTypes{},
			args: ActivityVocabularyTypes{"X"},
			want: false,
		},
		{
			name: `empty does not match multiple types`,
			t:    ActivityVocabularyTypes{},
			args: ActivityVocabularyTypes{"X", "Y"},
			want: false,
		},
		{
			name: `XY matches X`,
			t:    ActivityVocabularyTypes{"X", "Y"},
			args: ActivityVocabularyTypes{"X"},
			want: true,
		},
		{
			name: `XY matches Y`,
			t:    ActivityVocabularyTypes{"X", "Y"},
			args: ActivityVocabularyTypes{"Y"},
			want: true,
		},
		{
			name: `XY does not match empty`,
			t:    ActivityVocabularyTypes{"X", "Y"},
			args: ActivityVocabularyTypes{},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.Matches(tt.args...); got != tt.want {
				t.Errorf("Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypesMatch(t *testing.T) {
	type args struct {
		m1 TypeMatcher
		m2 TypeMatcher
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "nil matches nil",
			args: args{},
			want: true,
		},
		{
			name: "empty matches nil",
			args: args{m1: NilType},
			want: true,
		},
		{
			name: "nil matches empty",
			args: args{m2: NilType},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TypesMatch(tt.args.m1, tt.args.m2); got != tt.want {
				t.Errorf("TypesMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityVocabularyType_Matches(t *testing.T) {
	tests := []struct {
		name      string
		a         ActivityVocabularyType
		args      []ActivityVocabularyType
		wantMatch bool
	}{
		{
			name:      "nil matches nil",
			wantMatch: true,
		},
		{
			name:      "nil matches empty types",
			args:      ActivityVocabularyTypes{},
			wantMatch: true,
		},
		{
			name:      "nil matches NilType",
			args:      ActivityVocabularyTypes{NilType},
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
			args:      ActivityVocabularyTypes{},
			wantMatch: true,
		},
		{
			name:      "NilType matches NilType",
			a:         NilType,
			args:      ActivityVocabularyTypes{NilType},
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
			args:      ActivityVocabularyTypes{},
			wantMatch: true,
		},
		{
			name:      `"" matches NilType`,
			a:         "",
			args:      ActivityVocabularyTypes{NilType},
			wantMatch: true,
		},
		{
			name:      "nil does not match X",
			args:      ActivityVocabularyTypes{"X"},
			wantMatch: false,
		},
		{
			name:      `"" does not match X`,
			a:         "",
			args:      ActivityVocabularyTypes{"X"},
			wantMatch: false,
		},
		{
			name:      "NilType does not match X",
			a:         NilType,
			args:      ActivityVocabularyTypes{"X"},
			wantMatch: false,
		},
		{
			name:      "X matches X",
			a:         ActivityVocabularyType("X"),
			args:      ActivityVocabularyTypes{"X"},
			wantMatch: true,
		},
		{
			name:      "X matches X,Y",
			a:         ActivityVocabularyType("X"),
			args:      ActivityVocabularyTypes{"X", "Y"},
			wantMatch: true,
		},
		{
			name:      "X matches Y,X",
			a:         ActivityVocabularyType("X"),
			args:      ActivityVocabularyTypes{"Y", "X"},
			wantMatch: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMatch := tt.a.Matches(tt.args...); gotMatch != tt.wantMatch {
				t.Errorf("Matches() = %v, want %v", gotMatch, tt.wantMatch)
			}
		})
	}
}

func TestActivityVocabularyType_MatchOther(t *testing.T) {
	tests := []struct {
		name string
		a    ActivityVocabularyType
		arg  TypeMatcher
		want bool
	}{
		{
			name: "nil matches nil",
			want: true,
		},
		{
			name: "nil matches empty types",
			arg:  ActivityVocabularyTypes{},
			want: true,
		},
		{
			name: "nil matches NilType",
			arg:  ActivityVocabularyTypes{NilType},
			want: true,
		},
		{
			name: "NilType matches nil",
			a:    NilType,
			want: true,
		},
		{
			name: "NilType matches empty types",
			a:    NilType,
			arg:  ActivityVocabularyTypes{},
			want: true,
		},
		{
			name: "NilType matches NilType",
			a:    NilType,
			arg:  ActivityVocabularyTypes{NilType},
			want: true,
		},
		{
			name: `"" matches nil`,
			a:    "",
			want: true,
		},
		{
			name: `"" matches empty types`,
			a:    "",
			arg:  ActivityVocabularyTypes{},
			want: true,
		},
		{
			name: `"" matches NilType`,
			a:    "",
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
			a:    "",
			arg:  ActivityVocabularyTypes{"X"},
			want: false,
		},
		{
			name: "NilType does not match X",
			a:    NilType,
			arg:  ActivityVocabularyTypes{"X"},
			want: false,
		},
		{
			name: "X matches X",
			a:    ActivityVocabularyType("X"),
			arg:  ActivityVocabularyTypes{"X"},
			want: true,
		},
		{
			name: "X matches X,Y",
			a:    ActivityVocabularyType("X"),
			arg:  ActivityVocabularyTypes{"X", "Y"},
			want: true,
		},
		{
			name: "X matches Y,X",
			a:    ActivityVocabularyType("X"),
			arg:  ActivityVocabularyTypes{"Y", "X"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.MatchOther(tt.arg); got != tt.want {
				t.Errorf("MatchOther() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestActivityVocabularyTypes_MatchOther(t *testing.T) {
	tests := []struct {
		name string
		a    ActivityVocabularyTypes
		arg  TypeMatcher
		want bool
	}{
		{
			name: "nil matches nil",
			want: true,
		},
		{
			name: "nil matches empty types",
			arg:  ActivityVocabularyTypes{},
			want: true,
		},
		{
			name: "nil matches NilType",
			arg:  ActivityVocabularyTypes{NilType},
			want: true,
		},
		{
			name: "NilType matches nil",
			a:    ActivityVocabularyTypes{NilType},
			want: true,
		},
		{
			name: "NilType matches empty types",
			a:    ActivityVocabularyTypes{NilType},
			arg:  ActivityVocabularyTypes{},
			want: true,
		},
		{
			name: "NilType matches NilType",
			a:    ActivityVocabularyTypes{NilType},
			arg:  ActivityVocabularyTypes{NilType},
			want: true,
		},
		{
			name: `"" matches nil`,
			a:    ActivityVocabularyTypes{""},
			want: true,
		},
		{
			name: `"" matches empty types`,
			a:    ActivityVocabularyTypes{""},
			arg:  ActivityVocabularyTypes{},
			want: true,
		},
		{
			name: `"" matches NilType`,
			a:    ActivityVocabularyTypes{""},
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
			a:    ActivityVocabularyTypes{""},
			arg:  ActivityVocabularyTypes{"X"},
			want: false,
		},
		{
			name: "NilType does not match X",
			a:    ActivityVocabularyTypes{NilType},
			arg:  ActivityVocabularyTypes{"X"},
			want: false,
		},
		{
			name: "X matches X",
			a:    ActivityVocabularyTypes{"X"},
			arg:  ActivityVocabularyTypes{"X"},
			want: true,
		},
		{
			name: "X matches X,Y",
			a:    ActivityVocabularyTypes{"X"},
			arg:  ActivityVocabularyTypes{"X", "Y"},
			want: true,
		},
		{
			name: "X matches Y,X",
			a:    ActivityVocabularyTypes{"X"},
			arg:  ActivityVocabularyTypes{"Y", "X"},
			want: true,
		},
		{
			name: "X,Y does not match X",
			a:    ActivityVocabularyTypes{"X", "Y"},
			arg:  ActivityVocabularyTypes{"X"},
			want: false,
		},
		{
			name: "X,Y does not match Y",
			a:    ActivityVocabularyTypes{"X", "Y"},
			arg:  ActivityVocabularyTypes{"Y"},
			want: false,
		},
		{
			name: "X,Y matches X,Y",
			a:    ActivityVocabularyTypes{"X", "Y"},
			arg:  ActivityVocabularyTypes{"X", "Y"},
			want: true,
		},
		{
			name: "X,Y matches Y,X",
			a:    ActivityVocabularyTypes{"X", "Y"},
			arg:  ActivityVocabularyTypes{"Y", "X"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.MatchOther(tt.arg); got != tt.want {
				t.Errorf("MatchOther() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActivityVocabularyTypes_ContainsOld(t *testing.T) {
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ActivityTypes.Contains(ActivityType) {
			t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
		}
		for _, inValidType := range ObjectTypes {
			if ActivityTypes.Contains(inValidType) {
				t.Errorf("APObject Type '%v' should be invalid", inValidType)
			}
		}
		if ActivityTypes.Contains(invalidType) {
			t.Errorf("Activity Type '%v' should not be valid", invalidType)
		}
		for _, validType := range ActivityTypes {
			if !ActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should be valid", validType)
			}
		}
	}
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if IntransitiveActivityTypes.Contains(ActivityType) {
			t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
		}
		for _, inValidType := range ActivityTypes {
			if IntransitiveActivityTypes.Contains(inValidType) {
				t.Errorf("APObject Type '%v' should be invalid", inValidType)
			}
		}
		if IntransitiveActivityTypes.Contains(invalidType) {
			t.Errorf("Activity Type '%v' should not be valid", invalidType)
		}
		for _, validType := range IntransitiveActivityTypes {
			if !IntransitiveActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should be valid", validType)
			}
		}
	}
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ActivityTypes.Contains(ActivityType) {
			t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
		}
		for _, inValidType := range CollectionManagementActivityTypes {
			if !CollectionManagementActivityTypes.Contains(inValidType) {
				t.Errorf("APObject Type '%v' should be valid", inValidType)
			}
		}
		if CollectionManagementActivityTypes.Contains(invalidType) {
			t.Errorf("Activity Type '%v' should not be valid", invalidType)
		}
		for _, validType := range ContentManagementActivityTypes {
			if CollectionManagementActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
		for _, validType := range ReactionsActivityTypes {
			if CollectionManagementActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
	}

	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ActivityTypes.Contains(ActivityType) {
			t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
		}
		for _, inValidType := range ContentManagementActivityTypes {
			if !ContentManagementActivityTypes.Contains(inValidType) {
				t.Errorf("APObject Type '%v' should be valid", inValidType)
			}
		}
		if ContentManagementActivityTypes.Contains(invalidType) {
			t.Errorf("Activity Type '%v' should not be valid", invalidType)
		}
		for _, validType := range CollectionManagementActivityTypes {
			if ContentManagementActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
		for _, validType := range ReactionsActivityTypes {
			if ContentManagementActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
	}
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ReactionsActivityTypes.Contains(ActivityType) {
			t.Errorf("Generic Activity Type '%v' should not be valid", ActivityType)
		}
		for _, inValidType := range ReactionsActivityTypes {
			if !ReactionsActivityTypes.Contains(inValidType) {
				t.Errorf("APObject Type '%v' should be valid", inValidType)
			}
		}
		if ReactionsActivityTypes.Contains(invalidType) {
			t.Errorf("Activity Type '%v' should not be valid", invalidType)
		}
		for _, validType := range CollectionManagementActivityTypes {
			if ReactionsActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
		for _, validType := range ContentManagementActivityTypes {
			if ReactionsActivityTypes.Contains(validType) {
				t.Errorf("Activity Type '%v' should not be valid", validType)
			}
		}
	}
	{
		for _, validType := range CollectionTypes {
			if !CollectionTypes.Contains(validType) {
				t.Errorf("Generic Type '%#v' should be valid", validType)
			}
		}
	}
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ActorTypes.Contains(invalidType) {
			t.Errorf("APObject Type '%v' should not be valid", invalidType)
		}
		for _, validType := range ActorTypes {
			if !ActorTypes.Contains(validType) {
				t.Errorf("APObject Type '%v' should be valid", validType)
			}
		}
	}
	{
		for _, validType := range GenericTypes {
			if !GenericTypes.Contains(validType) {
				t.Errorf("Generic Type '%v' should be valid", validType)
			}
		}
	}
	{
		var invalidType ActivityVocabularyType = "RandomType"

		if ObjectTypes.Contains(invalidType) {
			t.Errorf("APObject Type '%v' should not be valid", invalidType)
		}
		for _, validType := range ObjectTypes {
			if !ObjectTypes.Contains(validType) {
				t.Errorf("APObject Type '%v' should be valid", validType)
			}
		}
	}
}

func TestActivityVocabularyTypes_Contains(t *testing.T) {
	tests := []struct {
		name string
		a    ActivityVocabularyTypes
		typ  ActivityVocabularyType
		want bool
	}{
		{
			name: "nil",
			a:    nil,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Contains(tt.typ); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOnType(t *testing.T) {
	type args struct {
		t  TypeMatcher
		fn func(typ ...ActivityVocabularyType) error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "empty",
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := OnType(tt.args.t, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("OnType() error = %v, wantErr %v", err, tt.wantErr)
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
