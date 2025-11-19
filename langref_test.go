package activitypub

import (
	"reflect"
	"testing"

	"golang.org/x/text/language"
)

func TestLangRef_GobEncode(t *testing.T) {
	tests := []struct {
		name    string
		l       LangRef
		want    []byte
		wantErr bool
	}{
		{
			name:    "Nil lang ref",
			l:       NilLangRef,
			want:    gobValue([]byte(language.Und.String())),
			wantErr: false,
		},
		{
			name:    "invalid text",
			l:       MakeRef([]byte("ana are")),
			want:    gobValue([]byte(language.Und.String())),
			wantErr: false,
		},
		{
			name:    "valid English",
			l:       MakeRef([]byte("en")),
			want:    gobValue([]byte(language.English.String())),
			wantErr: false,
		},
		{
			name:    "valid French Canadian",
			l:       MakeRef([]byte("fr_ca")),
			want:    gobValue([]byte(language.CanadianFrench.String())),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.GobEncode()
			if (err != nil) != tt.wantErr {
				t.Errorf("GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GobEncode() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestLangRefValue_GobEncode(t *testing.T) {
	type fields struct {
		Ref   LangRef
		Value Content
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			fields:  fields{},
			want:    []byte{},
			wantErr: false,
		},
		{
			name: "some values",
			fields: fields{
				Ref:   MakeRef([]byte("ana")),
				Value: Content("are mere"),
			},
			want:    gobValue(kv{K: []byte("ana"), V: []byte("are mere")}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LangRefValue{
				Ref:   tt.fields.Ref,
				Value: tt.fields.Value,
			}
			got, err := l.GobEncode()
			if (err != nil) != tt.wantErr {
				t.Errorf("GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GobEncode() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}

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
