package activitypub

import (
	"fmt"
	"testing"

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
	t.Skipf("TODO")
}

func TestPlace_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestPlace_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestPlace_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestPlace_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestPlace_IsObject(t *testing.T) {
	t.Skipf("TODO")
}

func TestPlace_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
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
