package activitypub

import (
	"testing"

	"github.com/go-ap/errors"
	"github.com/google/go-cmp/cmp"
)

func TestLangRef_UnmarshalJSON(t *testing.T) {
	lang := "en-US"
	rawJson := `"` + lang + `"`

	var a LangRef
	_ = a.UnmarshalJSON([]byte(rawJson))

	if a.String() != lang {
		t.Errorf("Invalid json unmarshal for %T. Expected %q, found %q", lang, lang, a)
	}
}

func TestLangRef_Equal(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name  string
		l     LangRef
		other LangRef
		want  bool
	}{
		{
			name:  "empty",
			l:     LangRef{},
			other: LangRef{},
			want:  true,
		},
		{
			name:  "und is zero",
			l:     und,
			other: LangRef{},
			want:  true,
		},
		{
			name:  "und",
			l:     und,
			other: und,
			want:  true,
		},
		{
			name:  "und vs en",
			l:     und,
			other: English,
			want:  false,
		},
		{
			name:  "en",
			l:     English,
			other: English,
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Equal(tt.other); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLangRef_UnmarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    LangRef
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
			l := LangRef{}
			if err := l.UnmarshalBinary(tt.data); !cmp.Equal(err, tt.wantErr, EquateWeakErrors) {
				t.Errorf("UnmarshalBinary() error = %s", cmp.Diff(tt.wantErr, err, EquateWeakErrors))
			}
			if !cmp.Equal(l, tt.want) {
				t.Errorf("UnmarshalBinary() got = %s", cmp.Diff(tt.want, l))
			}
		})
	}
}

func TestLangRef_MarshalBinary(t *testing.T) {
	tests := []struct {
		name    string
		sub     LangRef
		want    []byte
		wantErr error
	}{
		{
			name:    "nil",
			sub:     LangRef{},
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
