package activitypub

import (
	"fmt"
	"testing"
)

func TestPlace_Recipients(t *testing.T) {
	t.Skipf("TODO")
}

func TestToPlace(t *testing.T) {
	t.Skipf("TODO")
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
	t.Skipf("TODO")
}

func assertPlaceWithTesting(fn canErrorFunc, expected *Place) withPlaceFn {
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
		fn func(canErrorFunc, *Place) withPlaceFn
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
		var logFn canErrorFunc
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
