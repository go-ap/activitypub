package activitypub

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTombstone_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestTombstone_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestTombstone_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestTombstone_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestTombstone_IsLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestTombstone_IsObject(t *testing.T) {
	t.Skipf("TODO")
}

func TestTombstone_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestTombstone_Clean(t *testing.T) {
	t.Skipf("TODO")
}

func assertTombstoneWithTesting(fn canErrorFunc, expected *Tombstone) withTombstoneFn {
	return func(p *Tombstone) error {
		if !assertDeepEquals(fn, p, expected) {
			return fmt.Errorf("not equal")
		}
		return nil
	}
}

func TestOnTombstone(t *testing.T) {
	testTombstone := Tombstone{
		ID: "https://example.com",
	}
	type args struct {
		it Item
		fn func(canErrorFunc, *Tombstone) withTombstoneFn
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "single",
			args:    args{testTombstone, assertTombstoneWithTesting},
			wantErr: false,
		},
		{
			name:    "single fails",
			args:    args{&Tombstone{ID: "https://not-equal"}, assertTombstoneWithTesting},
			wantErr: true,
		},
		{
			name:    "collection of profiles",
			args:    args{ItemCollection{testTombstone, testTombstone}, assertTombstoneWithTesting},
			wantErr: false,
		},
		{
			name:    "collection of profiles fails",
			args:    args{ItemCollection{testTombstone, &Tombstone{ID: "not-equal"}}, assertTombstoneWithTesting},
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
			if err := OnTombstone(tt.args.it, tt.args.fn(logFn, &testTombstone)); (err != nil) != tt.wantErr {
				t.Errorf("OnTombstone() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestToTombstone(t *testing.T) {
	tests := []struct {
		name    string
		it      LinkOrIRI
		want    *Tombstone
		wantErr error
	}{
		{
			name: "empty",
		},
		{
			name: "Valid Tombstone",
			it:   Tombstone{ID: "test", Type: TombstoneType.ToTypes()},
			want: &Tombstone{ID: "test", Type: TombstoneType.ToTypes()},
		},
		{
			name: "Valid *Tombstone",
			it:   &Tombstone{ID: "test", Type: TombstoneType.ToTypes()},
			want: &Tombstone{ID: "test", Type: TombstoneType.ToTypes()},
		},
		{
			name:    "IRI",
			it:      IRI("https://example.com"),
			wantErr: ErrorInvalidType[Tombstone](IRI("")),
		},
		{
			name:    "IRIs",
			it:      IRIs{IRI("https://example.com")},
			wantErr: ErrorInvalidType[Tombstone](IRIs{}),
		},
		{
			name:    "ItemCollection",
			it:      ItemCollection{},
			wantErr: ErrorInvalidType[Tombstone](ItemCollection{}),
		},
		{
			name:    "Object",
			it:      &Object{ID: "test", Type: ArticleType.ToTypes()},
			wantErr: ErrorInvalidType[Tombstone](&Object{}),
		},
		{
			name:    "Activity",
			it:      &Activity{ID: "test", Type: CreateType.ToTypes()},
			wantErr: ErrorInvalidType[Tombstone](&Activity{}),
		},
		{
			name:    "IntransitiveActivity",
			it:      &IntransitiveActivity{ID: "test", Type: ArriveType.ToTypes()},
			wantErr: ErrorInvalidType[Tombstone](&IntransitiveActivity{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToTombstone(tt.it)
			if !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("ToTombstone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("ToTombstone() got = %s", cmp.Diff(tt.want, got))
			}
		})
	}
}
