package activitypub

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestProfile_Recipients(t *testing.T) {
	t.Skipf("TODO")
}

func TestToProfile(t *testing.T) {
	tests := []struct {
		name    string
		it      LinkOrIRI
		want    *Profile
		wantErr error
	}{
		{
			name: "empty",
		},
		{
			name: "Valid Profile",
			it:   Profile{ID: "test", Type: ProfileType},
			want: &Profile{ID: "test", Type: ProfileType},
		},
		{
			name: "Valid *Profile",
			it:   &Profile{ID: "test", Type: ProfileType},
			want: &Profile{ID: "test", Type: ProfileType},
		},
		{
			name:    "IRI",
			it:      IRI("https://example.com"),
			wantErr: ErrorInvalidType[Profile](IRI("")),
		},
		{
			name:    "IRIs",
			it:      IRIs{IRI("https://example.com")},
			wantErr: ErrorInvalidType[Profile](IRIs{}),
		},
		{
			name:    "ItemCollection",
			it:      ItemCollection{},
			wantErr: ErrorInvalidType[Profile](ItemCollection{}),
		},
		{
			name:    "Object",
			it:      &Object{ID: "test", Type: ArticleType},
			wantErr: ErrorInvalidType[Profile](&Object{}),
		},
		{
			name:    "Activity",
			it:      &Activity{ID: "test", Type: CreateType},
			wantErr: ErrorInvalidType[Profile](&Activity{}),
		},
		{
			name:    "IntransitiveActivity",
			it:      &IntransitiveActivity{ID: "test", Type: ArriveType},
			wantErr: ErrorInvalidType[Profile](&IntransitiveActivity{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToProfile(tt.it)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("ToProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ToProfile() got = %s", cmp.Diff(tt.want, got))
			}
			if got != nil && !got.Matches(ProfileType) {
				t.Errorf("ToProfile() expected to match Profile type, got = %v", got.GetType())
			}
		})
	}
}

func TestProfile_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestProfile_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestProfile_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestProfile_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestProfile_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestProfile_IsObject(t *testing.T) {
	t.Skipf("TODO")
}

func TestProfile_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestProfile_Clean(t *testing.T) {
	t.Skipf("TODO")
}

func assertProfileWithTesting(fn canErrorFunc, expected *Profile) withProfileFn {
	return func(p *Profile) error {
		if !assertDeepEquals(fn, p, expected) {
			return fmt.Errorf("not equal")
		}
		return nil
	}
}

func TestOnProfile(t *testing.T) {
	testProfile := Profile{
		ID: "https://example.com",
	}
	type args struct {
		it Item
		fn func(canErrorFunc, *Profile) withProfileFn
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "single",
			args:    args{testProfile, assertProfileWithTesting},
			wantErr: false,
		},
		{
			name:    "single fails",
			args:    args{&Profile{ID: "https://not-equal"}, assertProfileWithTesting},
			wantErr: true,
		},
		{
			name:    "collection of profiles",
			args:    args{ItemCollection{testProfile, testProfile}, assertProfileWithTesting},
			wantErr: false,
		},
		{
			name:    "collection of profiles fails",
			args:    args{ItemCollection{testProfile, &Profile{ID: "not-equal"}}, assertProfileWithTesting},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var logFn canErrorFunc
			if tt.wantErr {
				logFn = t.Logf
			} else {
				logFn = t.Errorf
			}
			if err := OnProfile(tt.args.it, tt.args.fn(logFn, &testProfile)); (err != nil) != tt.wantErr {
				t.Errorf("OnProfile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
