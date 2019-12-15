package activitypub

import (
	"bytes"
	"strings"
)

// ID designates an unique global identifier.
// All Objects in [ActivityStreams] should have unique global identifiers.
// ActivityPub extends this requirement; all objects distributed by the ActivityPub protocol MUST
// have unique global identifiers, unless they are intentionally transient
// (short lived activities that are not intended to be able to be looked up,
// such as some kinds of chat messages or game notifications).
// These identifiers must fall into one of the following groups:
//
// 1. Publicly dereferenceable URIs, such as HTTPS URIs, with their authority belonging
// to that of their originating server. (Publicly facing content SHOULD use HTTPS URIs).
// 2. An ID explicitly specified as the JSON null object, which implies an anonymous object
// (a part of its parent context)
type ID IRI

// UnmarshalJSON
func (i *ID) UnmarshalJSON(data []byte) error {
	*i = ID(strings.Trim(string(data), "\""))
	return nil
}

// MarshalJSON
func (i ID) MarshalJSON() ([]byte, error) {
	if len(i) == 0 {
		return nil, nil
	}
	b := bytes.Buffer{}
	b.Write([]byte{'"'})
	b.WriteString(string(i))
	b.Write([]byte{'"'})
	return b.Bytes(), nil
}

func (i *ID) IsValid() bool {
	return i != nil && len(*i) > 0
}
