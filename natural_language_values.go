package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"
	"strings"

	"github.com/valyala/fastjson"
)

// NilLangRef represents a convention for a nil language reference.
// It is used for LangRefValue objects without an explicit language key.
const NilLangRef LangRef = "-"

type (
	// LangRef is the type for a language reference code, should be an ISO639-1 language specifier.
	LangRef string
	Content []byte

	// LangRefValue is a type for storing per language values
	LangRefValue struct {
		Ref   LangRef
		Value Content
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

func (n NaturalLanguageValues) Get(ref LangRef) Content {
	for _, val := range n {
		if val.Ref == ref {
			return val.Value
		}
	}
	return nil
}

// Set sets a language, value pair in a NaturalLanguageValues array
func (n *NaturalLanguageValues) Set(ref LangRef, v Content) error {
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

// MarshalJSON encodes the receiver object to a JSON document.
func (n NaturalLanguageValues) MarshalJSON() ([]byte, error) {
	l := len(n)
	if l <= 0 {
		return nil, nil
	}

	b := bytes.Buffer{}
	if l == 1 {
		v := n[0]
		if len(v.Value) > 0 {
			v.Value = Content(unescape([]byte(v.Value)))
			ll, err := b.WriteString(strconv.Quote(v.Value.String()))
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
func (n *NaturalLanguageValues) Append(lang LangRef, value Content) error {
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
	if n == nil {
		return 0
	}
	return uint(len(*n))
}

// String adds support for Stringer interface. It returns the Value[LangRef] text or just Value if LangRef is NIL
func (l LangRefValue) String() string {
	if l.Ref == NilLangRef {
		return l.Value.String()
	}
	return fmt.Sprintf("%s[%s]", l.Value, l.Ref)
}

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (l *LangRefValue) UnmarshalJSON(data []byte) error {
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		l.Ref = NilLangRef
		l.Value = unescape(data)
		return nil
	}
	switch val.Type() {
	case fastjson.TypeObject:
		o, _ := val.Object()
		o.Visit(func(key []byte, v *fastjson.Value) {
			l.Ref = LangRef(key)
			l.Value = unescape(v.GetStringBytes())
		})
	case fastjson.TypeString:
		l.Ref = NilLangRef
		l.Value = unescape(val.GetStringBytes())
	}

	return nil
}

// UnmarshalText implements the TextEncoder interface
func (l *LangRefValue) UnmarshalText(data []byte) error {
	l.Ref = NilLangRef
	l.Value = unescape(data)
	return nil
}

func (l LangRef) GobEncode() ([]byte, error) {
	if len(l) == 0 {
		return []byte{}, nil
	}
	b := new(bytes.Buffer)
	gg := gob.NewEncoder(b)
	if err := gobEncodeStringLikeType(gg, []byte(l)); err != nil {
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
	*l = LangRef(bb)
	return nil
}

// MarshalJSON encodes the receiver object to a JSON document.
func (l LangRefValue) MarshalJSON() ([]byte, error) {
	buf := bytes.Buffer{}
	if l.Ref != NilLangRef && len(l.Ref) > 0 {
		if l.Value.Equals(Content("")) {
			return nil, nil
		}
		buf.WriteByte('"')
		buf.WriteString(l.Ref.String())
		buf.Write([]byte{'"', ':'})
	}
	l.Value = Content(unescape([]byte(l.Value)))
	buf.WriteString(strconv.Quote(l.Value.String()))
	return buf.Bytes(), nil
}

// MarshalText serializes the LangRefValue into JSON
func (l LangRefValue) MarshalText() ([]byte, error) {
	if l.Ref != NilLangRef && l.Value.Equals(Content("")) {
		return nil, nil
	}
	buf := bytes.Buffer{}
	buf.WriteString(l.Value.String())
	if l.Ref != NilLangRef {
		buf.WriteByte('[')
		buf.WriteString(l.Ref.String())
		buf.WriteByte(']')
	}
	return buf.Bytes(), nil
}

type kv struct {
	K []byte
	V []byte
}

func (l LangRefValue) GobEncode() ([]byte, error) {
	if len(l.Value) == 0 && len(l.Ref) == 0 {
		return []byte{}, nil
	}
	b := new(bytes.Buffer)
	gg := gob.NewEncoder(b)
	mm := kv{
		K: []byte(l.Ref),
		V: []byte(l.Value),
	}
	if err := gg.Encode(mm); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (l *LangRefValue) GobDecode(data []byte) error {
	if len(data) == 0 {
		// NOTE(marius): this behaviour diverges from vanilla gob package
		return nil
	}
	mm := kv{}
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&mm); err != nil {
		return err
	}
	l.Ref = LangRef(mm.K)
	l.Value = mm.V
	return nil
}

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (l *LangRef) UnmarshalJSON(data []byte) error {
	return l.UnmarshalText(data)
}

// UnmarshalText implements the TextEncoder interface
func (l *LangRef) UnmarshalText(data []byte) error {
	*l = ""
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

func (l LangRef) String() string {
	return string(l)
}

func (l LangRefValue) Equals(other LangRefValue) bool {
	return l.Ref == other.Ref && l.Value.Equals(other.Value)
}

func (c *Content) UnmarshalJSON(data []byte) error {
	return c.UnmarshalText(data)
}

func (c *Content) UnmarshalText(data []byte) error {
	*c = Content{}
	if len(data) == 0 {
		return nil
	}
	if len(data) > 2 {
		if data[0] == '"' && data[len(data)-1] == '"' {
			*c = Content(data[1 : len(data)-1])
		}
	} else {
		*c = Content(data)
	}
	return nil
}

func (c Content) GobEncode() ([]byte, error) {
	if len(c) == 0 {
		return []byte{}, nil
	}
	b := new(bytes.Buffer)
	gg := gob.NewEncoder(b)
	if err := gobEncodeStringLikeType(gg, c); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (c *Content) GobDecode(data []byte) error {
	if len(data) == 0 {
		// NOTE(marius): this behaviour diverges from vanilla gob package
		return nil
	}
	bb := make([]byte, 0)
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&bb); err != nil {
		return err
	}
	*c = bb
	return nil
}

func (c Content) String() string {
	return string(c)
}

func (c Content) Equals(other Content) bool {
	return bytes.Equal(c, other)
}

func unescape(b []byte) []byte {
	// FIXME(marius): I feel like I'm missing something really obvious about encoding/decoding from Json regarding
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

// UnmarshalJSON decodes an incoming JSON document into the receiver object.
func (n *NaturalLanguageValues) UnmarshalJSON(data []byte) error {
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		// try our luck if data contains an unquoted string
		n.Append(NilLangRef, unescape(data))
		return nil
	}
	switch val.Type() {
	case fastjson.TypeObject:
		ob, _ := val.Object()
		ob.Visit(func(key []byte, v *fastjson.Value) {
			if dat := v.GetStringBytes(); len(dat) > 0 {
				n.Append(LangRef(key), unescape(dat))
			}
		})
	case fastjson.TypeString:
		if dat := val.GetStringBytes(); len(dat) > 0 {
			n.Append(NilLangRef, unescape(dat))
		}
	case fastjson.TypeArray:
		for _, v := range val.GetArray() {
			l := LangRefValue{}
			l.UnmarshalJSON([]byte(v.String()))
			if len(l.Value) > 0 {
				n.Append(l.Ref, l.Value)
			}
		}
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
		n.Append(LangRef(NilLangRef), Content(data[1:len(data)-1]))
	}
	return nil
}

func (n NaturalLanguageValues) GobEncode() ([]byte, error) {
	if len(n) == 0 {
		return []byte{}, nil
	}
	b := new(bytes.Buffer)
	gg := gob.NewEncoder(b)
	mm := make([]kv, len(n))
	for i, l := range n {
		mm[i] = kv{K: []byte(l.Ref), V: l.Value}
	}
	if err := gg.Encode(mm); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (n *NaturalLanguageValues) GobDecode(data []byte) error {
	if len(data) == 0 {
		// NOTE(marius): this behaviour diverges from vanilla gob package
		return nil
	}
	mm := make([]kv, 0)
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&mm); err != nil {
		return err
	}
	for _, m := range mm {
		*n = append(*n, LangRefValue{Ref: LangRef(m.K), Value: m.V})
	}
	return nil
}

// Equals
func (n NaturalLanguageValues) Equals(with NaturalLanguageValues) bool {
	if n.Count() != with.Count() {
		return false
	}
	for _, wv := range with {
		for _, nv := range n {
			if nv.Equals(wv) {
				continue
			}
			return false
		}
	}
	return true
}
