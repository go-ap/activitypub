package activitypub

import (
	"bytes"
	"fmt"
	"github.com/buger/jsonparser"
	"strconv"
	"strings"
)

const (
	NilLangRef LangRef = "-"
)

type (
	// LangRef is the type for a language reference code, should be an ISO639-1 language specifier.
	LangRef string

	// LangRefValue is a type for storing per language values
	LangRefValue struct {
		Ref   LangRef
		Value string
	}
	// NaturalLanguageValues is a mapping for multiple language values
	NaturalLanguageValues []LangRefValue
)

func NaturalLanguageValuesNew() NaturalLanguageValues {
	return make(NaturalLanguageValues, 0)
}

func (n NaturalLanguageValues) String() string {
	cnt := len(n)
	if cnt == 1 {
		return n[0].String()
	}
	s := strings.Builder{}
	s.Write([]byte{'['})
	for k, v := range n {
		s.WriteString(v.String())
		if k != cnt-1 {
			s.Write([]byte{','})
		}
	}
	s.Write([]byte{']'})
	return s.String()
}

func (n NaturalLanguageValues) Get(ref LangRef) string {
	for _, val := range n {
		if val.Ref == ref {
			return val.Value
		}
	}
	return ""
}

// Set sets a language, value pair in a NaturalLanguageValues array
func (n *NaturalLanguageValues) Set(ref LangRef, v string) error {
	found := false
	for k, vv := range *n {
		if vv.Ref == ref {
			(*n)[k] = LangRefValue{ref, v}
			found = true
		}
	}
	if !found {
		n.Append(ref, v)
	}
	return nil
}

// MarshalJSON serializes the NaturalLanguageValues into JSON
func (n NaturalLanguageValues) MarshalJSON() ([]byte, error) {
	l := len(n)
	if l <= 0 {
		return nil, nil
	}
	
	b := bytes.Buffer{}
	if l == 1 {
		v := n[0]
		if len(v.Value) > 0 {
			v.Value = string(unescape([]byte(v.Value)))
			ll, err := b.WriteString(strconv.Quote(v.Value))
			if err != nil {
				return nil, err
			}
			if ll <= 0 {
				return nil, nil
			}
			return b.Bytes(), nil
		}
	}
	b.Write([]byte{'{'})
	empty := true
	for _, val := range n {
		if len(val.Ref) == 0 || len(val.Value) == 0 {
			continue
		}
		if !empty {
			b.Write([]byte{','})
		}
		if v, err := val.MarshalJSON(); err == nil && len(v) > 0 {
			l, err := b.Write(v)
			if err == nil && l > 0 {
				empty = false
			}
		}
	}
	b.Write([]byte{'}'})
	if !empty {
		return b.Bytes(), nil
	}
	return nil, nil
}

// First returns the first element in the array
func (n NaturalLanguageValues) First() LangRefValue {
	for _, v := range n {
		return v
	}
	return LangRefValue{}
}

// MarshalText serializes the NaturalLanguageValues into Text
func (n NaturalLanguageValues) MarshalText() ([]byte, error) {
	for _, v := range n {
		return []byte(fmt.Sprintf("%q", v)), nil
	}
	return nil, nil
}

// Append is syntactic sugar for resizing the NaturalLanguageValues map
// and appending an element
func (n *NaturalLanguageValues) Append(lang LangRef, value string) error {
	var t NaturalLanguageValues
	if len(*n) == 0 {
		t = make(NaturalLanguageValues, 0)
	} else {
		t = *n
	}
	t = append(*n, LangRefValue{lang, value})
	*n = t

	return nil
}

// Count returns the length of Items in the item collection
func (n *NaturalLanguageValues) Count() uint {
	return uint(len(*n))
}

// String adds support for Stringer interface. It returns the Value[LangRef] text or just Value if LangRef is NIL
func (l LangRefValue) String() string {
	if l.Ref == NilLangRef {
		return l.Value
	}
	return fmt.Sprintf("%s[%s]", l.Value, l.Ref)
}

// UnmarshalJSON implements the JsonEncoder interface
func (l *LangRefValue) UnmarshalJSON(data []byte) error {
	_, typ, _, err := jsonparser.Get(data)
	if err != nil {
		l.Ref = NilLangRef
		l.Value = string(unescape(data))
		return nil
	}
	switch typ {
	case jsonparser.Object:
		jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			l.Ref = LangRef(key)
			l.Value = string(unescape(value))
			return err
		})
	case jsonparser.String:
		l.Ref = NilLangRef
		l.Value = string(unescape(data))
	}

	return nil
}

// UnmarshalText implements the TextEncoder interface
func (l *LangRefValue) UnmarshalText(data []byte) error {
	return nil
}

// MarshalJSON serializes the LangRefValue into JSON
func (l LangRefValue) MarshalJSON() ([]byte, error) {
	buf := bytes.Buffer{}
	if l.Ref != NilLangRef && len(l.Ref) > 0{
		if l.Value == "" {
			return nil, nil
		}
		buf.WriteByte('"')
		buf.WriteString(string(l.Ref))
		buf.Write([]byte{'"', ':'})
	}
	l.Value = string(unescape([]byte(l.Value)))
	buf.WriteString(strconv.Quote(l.Value))
	return buf.Bytes(), nil
}

// MarshalText serializes the LangRefValue into JSON
func (l LangRefValue) MarshalText() ([]byte, error) {
	if l.Ref != NilLangRef && l.Value == "" {
		return nil, nil
	}
	buf := bytes.Buffer{}
	buf.WriteString(l.Value)
	if l.Ref != NilLangRef {
		buf.WriteByte('[')
		buf.WriteString(string(l.Ref))
		buf.WriteByte(']')
	}
	return buf.Bytes(), nil
}

// UnmarshalJSON implements the JsonEncoder interface
func (l *LangRef) UnmarshalJSON(data []byte) error {
	return l.UnmarshalText(data)
}

// UnmarshalText implements the TextEncoder interface
func (l *LangRef) UnmarshalText(data []byte) error {
	*l = LangRef("")
	if len(data) == 0 {
		return nil
	}
	if len(data) > 2 {
		if data[0] == '"' && data[len(data)-1] == '"' {
			*l = LangRef(data[1 : len(data)-1])
		}
	} else {
		*l = LangRef(data)
	}
	return nil
}

func unescape(b []byte) []byte {
	// FIMXE(marius): I feel like I'm missing something really obvious about encoding/decoding from Json regarding
	//    escape characters, and that this function is just a hack. Be better future Marius, find the real problem!
	b = bytes.ReplaceAll(b, []byte{'\\', 'a'}, []byte{'\a'})
	b = bytes.ReplaceAll(b, []byte{'\\', 'f'}, []byte{'\f'})
	b = bytes.ReplaceAll(b, []byte{'\\', 'n'}, []byte{'\n'})
	b = bytes.ReplaceAll(b, []byte{'\\', 'r'}, []byte{'\r'})
	b = bytes.ReplaceAll(b, []byte{'\\', 't'}, []byte{'\t'})
	b = bytes.ReplaceAll(b, []byte{'\\', 'v'}, []byte{'\v'})
	b = bytes.ReplaceAll(b, []byte{'\\', '"'}, []byte{'"'})
	b = bytes.ReplaceAll(b, []byte{'\\', '\\'}, []byte{'\\'}) // this should cover the case of \\u -> \u
	return b
}

// UnmarshalJSON tries to load the NaturalLanguage array from the incoming json value
func (n *NaturalLanguageValues) UnmarshalJSON(data []byte) error {
	val, typ, _, err := jsonparser.Get(data)
	if err != nil {
		// try our luck if data contains an unquoted string
		n.Append(NilLangRef, string(unescape(data)))
		return nil
	}
	switch typ {
	case jsonparser.Object:
		jsonparser.ObjectEach(data, func(key []byte, val []byte, dataType jsonparser.ValueType, offset int) error {
			n.Append(LangRef(key), string(unescape(val)))
			return err
		})
	case jsonparser.String:
		n.Append(NilLangRef, string(unescape(val)))
	case jsonparser.Array:
		jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			l := LangRefValue{}
			l.UnmarshalJSON(value)
			n.Append(l.Ref, l.Value)
		})
	}

	return nil
}

// UnmarshalText tries to load the NaturalLanguage array from the incoming Text value
func (n *NaturalLanguageValues) UnmarshalText(data []byte) error {
	if data[0] == '"' {
		// a quoted string - loading it to c.URL
		if data[len(data)-1] != '"' {
			return fmt.Errorf("invalid string value when unmarshaling %T value", n)
		}
		n.Append(LangRef(NilLangRef), string(data[1:len(data)-1]))
	}
	return nil
}
