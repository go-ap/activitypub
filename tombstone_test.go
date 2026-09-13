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
	tests := []struct {
		name string
		a    Tombstone
		want Item
	}{
		{
			name: "empty",
			a:    Tombstone{},
			want: &Tombstone{},
		},
		{
			name: "has Bto",
			a:    Tombstone{Type: TombstoneType, Bto: ItemCollection{IRI("http://example.com")}},
			want: &Tombstone{Type: TombstoneType},
		},
		{
			name: "has BCC",
			a:    Tombstone{Type: TombstoneType, BCC: ItemCollection{IRI("http://example.com")}},
			want: &Tombstone{Type: TombstoneType},
		},
		{
			name: "audience has BCC",
			a:    Tombstone{Type: TombstoneType, Audience: ItemCollection{&Object{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &Tombstone{Type: TombstoneType, Audience: ItemCollection{&Object{}}},
		},
		{
			name: "attachment has BCC",
			a:    Tombstone{Type: TombstoneType, Attachment: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Tombstone{Type: TombstoneType, Attachment: &Object{}},
		},
		{
			name: "icon has BCC",
			a:    Tombstone{Type: TombstoneType, Icon: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Tombstone{Type: TombstoneType, Icon: &Object{}},
		},
		{
			name: "image has BCC",
			a:    Tombstone{Type: TombstoneType, Image: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Tombstone{Type: TombstoneType, Image: &Object{}},
		},
		{
			name: "context has BCC",
			a:    Tombstone{Type: TombstoneType, Context: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Tombstone{Type: TombstoneType, Context: &Object{}},
		},
		{
			name: "generator has BCC",
			a:    Tombstone{Type: TombstoneType, Generator: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Tombstone{Type: TombstoneType, Generator: &Object{}},
		},
		{
			name: "attributedTo has BCC",
			a:    Tombstone{Type: TombstoneType, AttributedTo: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Tombstone{Type: TombstoneType, AttributedTo: &Object{}},
		},
		{
			name: "preview has BCC",
			a:    Tombstone{Type: TombstoneType, Preview: &Object{BCC: ItemCollection{IRI("http://example.com")}}},
			want: &Tombstone{Type: TombstoneType, Preview: &Object{}},
		},
		{
			name: "tag has BCC",
			a:    Tombstone{Type: TombstoneType, Tag: ItemCollection{&Object{BCC: ItemCollection{IRI("http://example.com")}}}},
			want: &Tombstone{Type: TombstoneType, Tag: ItemCollection{&Object{}}},
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

func assertTombstoneWithTesting(fn logFn, expected *Tombstone) withTombstoneFn {
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
		fn func(logFn, *Tombstone) withTombstoneFn
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
			var logFn logFn
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
			it:   Tombstone{ID: "test", Type: TombstoneType},
			want: &Tombstone{ID: "test", Type: TombstoneType},
		},
		{
			name: "Valid *Tombstone",
			it:   &Tombstone{ID: "test", Type: TombstoneType},
			want: &Tombstone{ID: "test", Type: TombstoneType},
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
			it:      &Object{ID: "test", Type: ArticleType},
			wantErr: ErrorInvalidType[Tombstone](&Object{}),
		},
		{
			name:    "Activity",
			it:      &Activity{ID: "test", Type: CreateType},
			wantErr: ErrorInvalidType[Tombstone](&Activity{}),
		},
		{
			name:    "IntransitiveActivity",
			it:      &IntransitiveActivity{ID: "test", Type: ArriveType},
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
			if got != nil && !got.Match(TombstoneType) {
				t.Errorf("ToTombstone() expected to match Tombstone type, got = %v", got.GetType())
			}
		})
	}
}
