package activitypub

import (
	"reflect"
	"testing"
)

func TestLinkNew(t *testing.T) {
	testValue := ID("test")
	var testType ActivityVocabularyType

	l := LinkNew(testValue, testType)

	if l.ID != testValue {
		t.Errorf("APObject Id '%v' different than expected '%v'", l.ID, testValue)
	}
	if l.Type != LinkType {
		t.Errorf("APObject Type '%v' different than expected '%v'", l.Type, LinkType)
	}
}

func TestLink_IsLink(t *testing.T) {
	l := LinkNew("test", LinkType)
	if !l.IsLink() {
		t.Errorf("%#v should be a valid link", l.Type)
	}
	m := LinkNew("test", MentionType)
	if !m.IsLink() {
		t.Errorf("%#v should be a valid link", m.Type)
	}
}

func TestLink_IsObject(t *testing.T) {
	l := LinkNew("test", LinkType)
	if l.IsObject() {
		t.Errorf("%#v should not be a valid object", l.Type)
	}
	m := LinkNew("test", MentionType)
	if m.IsObject() {
		t.Errorf("%#v should not be a valid object", m.Type)
	}
}

func TestLink_GetID(t *testing.T) {
	t.Skipf("TODO")
}

func TestLink_GetLink(t *testing.T) {
	t.Skipf("TODO")
}

func TestLink_GetType(t *testing.T) {
	t.Skipf("TODO")
}

func TestLink_UnmarshalJSON(t *testing.T) {
	t.Skipf("TODO")
}

func TestMentionNew(t *testing.T) {
	t.Skipf("TODO")
}

func TestLink_IsCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestLink_GobEncode(t *testing.T) {
	type fields struct {
		ID        ID
		Type      ActivityVocabularyType
		Name      NaturalLanguageValues
		Rel       IRI
		MediaType MimeType
		Height    uint
		Width     uint
		Preview   Item
		Href      IRI
		HrefLang  LangRef
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Link{
				ID:        tt.fields.ID,
				Type:      tt.fields.Type,
				Name:      tt.fields.Name,
				Rel:       tt.fields.Rel,
				MediaType: tt.fields.MediaType,
				Height:    tt.fields.Height,
				Width:     tt.fields.Width,
				Preview:   tt.fields.Preview,
				Href:      tt.fields.Href,
				HrefLang:  tt.fields.HrefLang,
			}
			got, err := l.GobEncode()
			if (err != nil) != tt.wantErr {
				t.Errorf("GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GobEncode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLink_GobDecode(t *testing.T) {
	type fields struct {
		ID        ID
		Type      ActivityVocabularyType
		Name      NaturalLanguageValues
		Rel       IRI
		MediaType MimeType
		Height    uint
		Width     uint
		Preview   Item
		Href      IRI
		HrefLang  LangRef
	}
	tests := []struct {
		name    string
		fields  fields
		data    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			fields:  fields{},
			data:    []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Link{
				ID:        tt.fields.ID,
				Type:      tt.fields.Type,
				Name:      tt.fields.Name,
				Rel:       tt.fields.Rel,
				MediaType: tt.fields.MediaType,
				Height:    tt.fields.Height,
				Width:     tt.fields.Width,
				Preview:   tt.fields.Preview,
				Href:      tt.fields.Href,
				HrefLang:  tt.fields.HrefLang,
			}
			if err := l.GobDecode(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("GobDecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
