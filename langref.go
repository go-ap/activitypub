package activitypub

import (
	"github.com/go-ap/errors"
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

// Equal checks if the language ref is equal to the "other"
func (l LangRef) Equal(other LangRef) bool {
	return (l.Valid() == other.Valid()) || l == other
}

// MakeRef
func MakeRef(raw []byte) LangRef {
	return LangRef(language.Make(string(raw)))
}

func (l LangRef) MarshalBinary() ([]byte, error) {
	panic(errors.NotImplementedf("Binary functionality not implemented"))
}

func (l *LangRef) UnmarshalBinary(data []byte) error {
	panic(errors.NotImplementedf("Binary functionality not implemented"))
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
