package activitypub

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-ap/errors"
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
			if got != nil && !got.Match(ProfileType) {
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
	tests := []struct {
		name    string
		data    []byte
		want    Profile
		wantErr error
	}{
		{
			name:    "nil",
			data:    nil,
			wantErr: errors.Newf(`cannot parse JSON: cannot parse empty string; unparsed tail: ""`),
		},
		{
			name: "just with ID",
			data: []byte(`{"id":"https://example.com/profile-1"}`),
			want: Profile{ID: "https://example.com/profile-1"},
		},
		{
			name: "a mocked profile",
			data: []byte(`{"id":"https://example.com/~root/profile/1","type":"Profile","attributedTo":"https://example.com/~root","audience":"https://www.w3.org/ns/activitystreams#Public","published":"2000-01-04T22:23:22.000Z","likes":"https://example.com/~root/profile/1/likes"}`),
			want: Profile{
				ID:           "https://example.com/~root/profile/1",
				Type:         ProfileType,
				AttributedTo: IRI("https://example.com/~root"),
				Audience:     ItemCollection{IRI("https://www.w3.org/ns/activitystreams#Public")},
				Published:    time.Date(2000, time.January, 4, 22, 23, 22, 0, time.UTC),
				Likes:        IRI("https://example.com/~root/profile/1/likes"),
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Profile{}
			if err := p.UnmarshalJSON(tt.data); !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("UnmarshalJSON() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
			}
			if !cmp.Equal(p, tt.want) {
				t.Errorf("UnmarshalJSON() got = %s", cmp.Diff(tt.want, p))
			}
		})
	}
}

func TestProfile_Clean(t *testing.T) {
	tests := []struct {
		name string
		a    Profile
		want Item
	}{
		{
			name: "empty",
			a:    Profile{},
			want: &Profile{},
		},
		{
			name: "has Bto",
			a:    Profile{Type: ProfileType, Bto: ItemCollection{IRI("http://example.com")}},
			want: &Profile{Type: ProfileType},
		},
		{
			name: "has BCC",
			a:    Profile{Type: ProfileType, BCC: ItemCollection{IRI("http://example.com")}},
			want: &Profile{Type: ProfileType},
		},
		{
			name: "audience has BCC",
			a:    Profile{Type: ProfileType, Audience: ItemCollection{&Object{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &Profile{Type: ProfileType, Audience: ItemCollection{&Object{}}},
		},
		{
			name: "attachment has BCC",
			a:    Profile{Type: ProfileType, Attachment: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Profile{Type: ProfileType, Attachment: &Object{}},
		},
		{
			name: "icon has BCC",
			a:    Profile{Type: ProfileType, Icon: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Profile{Type: ProfileType, Icon: &Object{}},
		},
		{
			name: "image has BCC",
			a:    Profile{Type: ProfileType, Image: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Profile{Type: ProfileType, Image: &Object{}},
		},
		{
			name: "context has BCC",
			a:    Profile{Type: ProfileType, Context: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Profile{Type: ProfileType, Context: &Object{}},
		},
		{
			name: "generator has BCC",
			a:    Profile{Type: ProfileType, Generator: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Profile{Type: ProfileType, Generator: &Object{}},
		},
		{
			name: "attributedTo has BCC",
			a:    Profile{Type: ProfileType, AttributedTo: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Profile{Type: ProfileType, AttributedTo: &Object{}},
		},
		{
			name: "preview has BCC",
			a:    Profile{Type: ProfileType, Preview: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Profile{Type: ProfileType, Preview: &Object{}},
		},
		{
			name: "tag has BCC",
			a:    Profile{Type: ProfileType, Tag: ItemCollection{&Object{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &Profile{Type: ProfileType, Tag: ItemCollection{&Object{}}},
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

func assertProfileWithTesting(fn logFn, expected *Profile) WithProfileFn {
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
		fn func(logFn, *Profile) WithProfileFn
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
			var logFn logFn
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

func TestProfile_UnmarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    Profile
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
			l := Profile{}
			if err := l.UnmarshalBinary(tt.data); !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("UnmarshalBinary() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
			}
			if !cmp.Equal(l, tt.want) {
				t.Errorf("UnmarshalBinary() got = %s", cmp.Diff(tt.want, l))
			}
		})
	}
}

func TestProfile_MarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		sub     Profile
		want    []byte
		wantErr error
	}{
		{
			name:    "nil",
			sub:     Profile{},
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
