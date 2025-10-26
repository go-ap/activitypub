package activitypub

import (
	"bytes"
	"encoding/gob"

	"golang.org/x/text/language"
)

// LangRef is the type for a language reference code, should be an ISO639-1 language specifier.
type LangRef language.Tag

// NilLangRef represents a convention for a nil language reference.
// It is used for LangRefValue objects without an explicit language key.
var NilLangRef = und

// DefaultLang represents the default language reference used when using the convenience content generation.
var DefaultLang = English

// Valid
func (l LangRef) Valid() bool {
	return len(l.String()) > 0 && l != LangRef(language.Und)
}

// Equal
func (l LangRef) Equal(other LangRef) bool {
	return l.Valid() && other.Valid() && l == other
}

// MakeRef
func MakeRef(raw []byte) LangRef {
	return LangRef(language.Make(string(raw)))
}

func (l LangRef) GobEncode() ([]byte, error) {
	if len(l.String()) == 0 {
		return []byte{}, nil
	}
	b := new(bytes.Buffer)
	gg := gob.NewEncoder(b)
	if err := gobEncodeStringLikeType(gg, []byte(l.String())); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (l *LangRef) GobDecode(data []byte) error {
	if len(data) == 0 {
		// NOTE(marius): this behaviour diverges from vanilla gob package
		return nil
	}
	var bb []byte
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&bb); err != nil {
		return err
	}
	*l = MakeRef(bb)
	return nil
}

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (l *LangRef) UnmarshalJSON(data []byte) error {
	return l.UnmarshalText(data)
}

// UnmarshalText implements the TextEncoder interface
func (l *LangRef) UnmarshalText(data []byte) error {
	*l = NilLangRef
	if len(data) == 0 {
		return nil
	}
	if len(data) > 2 {
		if data[0] == '"' && data[len(data)-1] == '"' {
			*l = MakeRef(data[1 : len(data)-1])
		}
	} else {
		*l = MakeRef(data)
	}
	return nil
}

func (l LangRef) String() string {
	return language.Tag(l).String()
}
