package activitypub

import (
	"fmt"
	"testing"

	"github.com/go-ap/errors"
	"github.com/google/go-cmp/cmp"
)

func TestTombstone_GetID(t *testing.T) {
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
			l := Tombstone{ID: tt.ID}
			if got := l.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTombstone_GetLink(t *testing.T) {
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
			l := Tombstone{ID: tt.IRI}
			if got := l.GetLink(); got != tt.want {
				t.Errorf("GetLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTombstone_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestTombstone_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    Tombstone
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
			l := Tombstone{}
			if err := l.UnmarshalJSON(tt.data); !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("UnmarshalJSON() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
			}
			if !cmp.Equal(l, tt.want) {
				t.Errorf("UnmarshalJSON() got = %s", cmp.Diff(tt.want, l))
			}
		})
	}
}

func TestTombstone_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		sub     Tombstone
		want    []byte
		wantErr error
	}{
		{
			name: "nil",
			sub:  Tombstone{},
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

func TestTombstone_UnmarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    Tombstone
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
			l := Tombstone{}
			if err := l.UnmarshalBinary(tt.data); !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("UnmarshalBinary() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
			}
			if !cmp.Equal(l, tt.want) {
				t.Errorf("UnmarshalBinary() got = %s", cmp.Diff(tt.want, l))
			}
		})
	}
}

func TestTombstone_MarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		sub     Tombstone
		want    []byte
		wantErr error
	}{
		{
			name:    "nil",
			sub:     Tombstone{},
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
