package activitypub

import (
	"fmt"
	"testing"

	"github.com/go-ap/errors"
	"github.com/google/go-cmp/cmp"
)

func TestPlace_Recipients(t *testing.T) {
	t.Skipf("TODO")
}

func TestToPlace(t *testing.T) {
	tests := []struct {
		name    string
		it      LinkOrIRI
		want    *Place
		wantErr error
	}{
		{
			name: "empty",
		},
		{
			name: "Valid Place",
			it:   Place{ID: "test", Type: PlaceType},
			want: &Place{ID: "test", Type: PlaceType},
		},
		{
			name: "Valid *Place",
			it:   &Place{ID: "test", Type: PlaceType},
			want: &Place{ID: "test", Type: PlaceType},
		},
		{
			name:    "IRI",
			it:      IRI("https://example.com"),
			wantErr: ErrorInvalidType[Place](IRI("")),
		},
		{
			name:    "IRIs",
			it:      IRIs{IRI("https://example.com")},
			wantErr: ErrorInvalidType[Place](IRIs{}),
		},
		{
			name:    "ItemCollection",
			it:      ItemCollection{},
			wantErr: ErrorInvalidType[Place](ItemCollection{}),
		},
		{
			name:    "Object",
			it:      &Object{ID: "test", Type: ArticleType},
			wantErr: ErrorInvalidType[Place](&Object{}),
		},
		{
			name:    "Activity",
			it:      &Activity{ID: "test", Type: CreateType},
			wantErr: ErrorInvalidType[Place](&Activity{}),
		},
		{
			name:    "IntransitiveActivity",
			it:      &IntransitiveActivity{ID: "test", Type: ArriveType},
			wantErr: ErrorInvalidType[Place](&IntransitiveActivity{}),
		},
		{
			name: "Tombstone",
			it:   &Tombstone{ID: "test", Type: TombstoneType, FormerType: PersonType},
			want: &Place{ID: "test", Type: TombstoneType},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToPlace(tt.it)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("ToPlace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ToPlace() got = %s", cmp.Diff(tt.want, got))
			}
			validTypes := ActivityVocabularyTypes{PlaceType, TombstoneType}
			if got != nil && !validTypes.Match(got.Type) {
				t.Errorf("ToPlace() expected to match %v types, got = %v", validTypes, got.Type)
			}
		})
	}
}

func TestPlace_GetID(t *testing.T) {
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
			l := Place{ID: tt.ID}
			if got := l.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlace_GetLink(t *testing.T) {
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
			l := Place{ID: tt.IRI}
			if got := l.GetLink(); got != tt.want {
				t.Errorf("GetLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlace_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestPlace_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    Place
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
			l := Place{}
			if err := l.UnmarshalJSON(tt.data); !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("UnmarshalJSON() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
			}
			if !cmp.Equal(l, tt.want) {
				t.Errorf("UnmarshalJSON() got = %s", cmp.Diff(tt.want, l))
			}
		})
	}
}

func TestPlace_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		sub     Place
		want    []byte
		wantErr error
	}{
		{
			name: "nil",
			sub:  Place{},
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

func TestPlace_Clean(t *testing.T) {
	tests := []struct {
		name string
		a    Place
		want Item
	}{
		{
			name: "empty",
			a:    Place{},
			want: &Place{},
		},
		{
			name: "has Bto",
			a:    Place{Type: PlaceType, Bto: ItemCollection{IRI("http://example.com")}},
			want: &Place{Type: PlaceType},
		},
		{
			name: "has BCC",
			a:    Place{Type: PlaceType, BCC: ItemCollection{IRI("http://example.com")}},
			want: &Place{Type: PlaceType},
		},
		{
			name: "audience has BCC",
			a:    Place{Type: PlaceType, Audience: ItemCollection{&Object{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &Place{Type: PlaceType, Audience: ItemCollection{&Object{}}},
		},
		{
			name: "attachment has BCC",
			a:    Place{Type: PlaceType, Attachment: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Place{Type: PlaceType, Attachment: &Object{}},
		},
		{
			name: "icon has BCC",
			a:    Place{Type: PlaceType, Icon: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Place{Type: PlaceType, Icon: &Object{}},
		},
		{
			name: "image has BCC",
			a:    Place{Type: PlaceType, Image: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Place{Type: PlaceType, Image: &Object{}},
		},
		{
			name: "context has BCC",
			a:    Place{Type: PlaceType, Context: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Place{Type: PlaceType, Context: &Object{}},
		},
		{
			name: "generator has BCC",
			a:    Place{Type: PlaceType, Generator: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Place{Type: PlaceType, Generator: &Object{}},
		},
		{
			name: "attributedTo has BCC",
			a:    Place{Type: PlaceType, AttributedTo: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Place{Type: PlaceType, AttributedTo: &Object{}},
		},
		{
			name: "preview has BCC",
			a:    Place{Type: PlaceType, Preview: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Place{Type: PlaceType, Preview: &Object{}},
		},
		{
			name: "tag has BCC",
			a:    Place{Type: PlaceType, Tag: ItemCollection{&Object{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &Place{Type: PlaceType, Tag: ItemCollection{&Object{}}},
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

func assertPlaceWithTesting(fn logFn, expected *Place) WithPlaceFn {
	return func(p *Place) error {
		if !assertDeepEquals(fn, p, expected) {
			return fmt.Errorf("not equal")
		}
		return nil
	}
}

func TestOnPlace(t *testing.T) {
	testPlace := Place{
		ID: "https://example.com",
	}
	type args struct {
		it Item
		fn func(logFn, *Place) WithPlaceFn
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "single",
			args:    args{testPlace, assertPlaceWithTesting},
			wantErr: false,
		},
		{
			name:    "single fails",
			args:    args{Place{ID: "https://not-equals"}, assertPlaceWithTesting},
			wantErr: true,
		},
		{
			name:    "collectionOfPlaces",
			args:    args{ItemCollection{testPlace, testPlace}, assertPlaceWithTesting},
			wantErr: false,
		},
		{
			name:    "collectionOfPlaces fails",
			args:    args{ItemCollection{testPlace, Place{ID: "https://not-equals"}}, assertPlaceWithTesting},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		var logFn logFn
		if tt.wantErr {
			logFn = t.Logf
		} else {
			logFn = t.Errorf
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := OnPlace(tt.args.it, tt.args.fn(logFn, &testPlace)); (err != nil) != tt.wantErr {
				t.Errorf("OnPlace() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPlace_UnmarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    Place
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
			l := Place{}
			if err := l.UnmarshalBinary(tt.data); !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("UnmarshalBinary() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
			}
			if !cmp.Equal(l, tt.want) {
				t.Errorf("UnmarshalBinary() got = %s", cmp.Diff(tt.want, l))
			}
		})
	}
}

func TestPlace_MarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		sub     Place
		want    []byte
		wantErr error
	}{
		{
			name:    "nil",
			sub:     Place{},
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
