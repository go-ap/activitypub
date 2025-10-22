package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/go-ap/errors"
	"github.com/valyala/fastjson"
)

type (
	Content []byte

	// LangRefValue is a type for storing per language values
	LangRefValue struct {
		Ref   LangRef
		Value Content
	}
	// NaturalLanguageValues is a mapping for multiple language values
	NaturalLanguageValues map[LangRef]Content
)

func NaturalLanguageValuesNew(values ...LangRefValue) NaturalLanguageValues {
	n := make(NaturalLanguageValues, len(values))
	for _, val := range values {
		n[val.Ref] = val.Value
	}
	return n
}

func DefaultNaturalLanguageValue(content string) NaturalLanguageValues {
	return NaturalLanguageValuesNew(DefaultLangRef(content))
}

func (n NaturalLanguageValues) String() string {
	cnt := len(n)
	if cnt == 1 {
		return n.First().String()
	}
	first := true
	s := strings.Builder{}
	s.Write([]byte{'['})
	for k, v := range n {
		if !first {
			s.Write([]byte{','})
		}
		s.WriteString(v.String())
		if k != NilLangRef {
			s.WriteString("[")
			s.WriteString(k.String())
			s.WriteString("]")
		}
		first = false
	}
	s.Write([]byte{']'})
	return s.String()
}

func (n NaturalLanguageValues) Get(ref LangRef) Content {
	for k, val := range n {
		if k == ref {
			return val
		}
	}
	return nil
}

// Set sets a language, value pair in a NaturalLanguageValues array
func (n *NaturalLanguageValues) Set(ref LangRef, v Content) error {
	(*n)[ref] = v
	return nil
}

var hex = "0123456789abcdef"

// safeSet holds the value true if the ASCII character with the given array
// position can be represented inside a JSON string without any further
// escaping.
//
// All values are true except for the ASCII control characters (0-31), the
// double quote ("), and the backslash character ("\").
var safeSet = [utf8.RuneSelf]bool{
	' ':      true,
	'!':      true,
	'"':      false,
	'#':      true,
	'$':      true,
	'%':      true,
	'&':      true,
	'\'':     true,
	'(':      true,
	')':      true,
	'*':      true,
	'+':      true,
	',':      true,
	'-':      true,
	'.':      true,
	'/':      true,
	'0':      true,
	'1':      true,
	'2':      true,
	'3':      true,
	'4':      true,
	'5':      true,
	'6':      true,
	'7':      true,
	'8':      true,
	'9':      true,
	':':      true,
	';':      true,
	'<':      true,
	'=':      true,
	'>':      true,
	'?':      true,
	'@':      true,
	'A':      true,
	'B':      true,
	'C':      true,
	'D':      true,
	'E':      true,
	'F':      true,
	'G':      true,
	'H':      true,
	'I':      true,
	'J':      true,
	'K':      true,
	'L':      true,
	'M':      true,
	'N':      true,
	'O':      true,
	'P':      true,
	'Q':      true,
	'R':      true,
	'S':      true,
	'T':      true,
	'U':      true,
	'V':      true,
	'W':      true,
	'X':      true,
	'Y':      true,
	'Z':      true,
	'[':      true,
	'\\':     false,
	']':      true,
	'^':      true,
	'_':      true,
	'`':      true,
	'a':      true,
	'b':      true,
	'c':      true,
	'd':      true,
	'e':      true,
	'f':      true,
	'g':      true,
	'h':      true,
	'i':      true,
	'j':      true,
	'k':      true,
	'l':      true,
	'm':      true,
	'n':      true,
	'o':      true,
	'p':      true,
	'q':      true,
	'r':      true,
	's':      true,
	't':      true,
	'u':      true,
	'v':      true,
	'w':      true,
	'x':      true,
	'y':      true,
	'z':      true,
	'{':      true,
	'|':      true,
	'}':      true,
	'~':      true,
	'\u007f': true,
}

// htmlSafeSet holds the value true if the ASCII character with the given
// array position can be safely represented inside a JSON string, embedded
// inside of HTML <script> tags, without any additional escaping.
//
// All values are true except for the ASCII control characters (0-31), the
// double quote ("), the backslash character ("\"), HTML opening and closing
// tags ("<" and ">"), and the ampersand ("&").
var htmlSafeSet = [utf8.RuneSelf]bool{
	' ':      true,
	'!':      true,
	'"':      false,
	'#':      true,
	'$':      true,
	'%':      true,
	'&':      false,
	'\'':     true,
	'(':      true,
	')':      true,
	'*':      true,
	'+':      true,
	',':      true,
	'-':      true,
	'.':      true,
	'/':      true,
	'0':      true,
	'1':      true,
	'2':      true,
	'3':      true,
	'4':      true,
	'5':      true,
	'6':      true,
	'7':      true,
	'8':      true,
	'9':      true,
	':':      true,
	';':      true,
	'<':      false,
	'=':      true,
	'>':      false,
	'?':      true,
	'@':      true,
	'A':      true,
	'B':      true,
	'C':      true,
	'D':      true,
	'E':      true,
	'F':      true,
	'G':      true,
	'H':      true,
	'I':      true,
	'J':      true,
	'K':      true,
	'L':      true,
	'M':      true,
	'N':      true,
	'O':      true,
	'P':      true,
	'Q':      true,
	'R':      true,
	'S':      true,
	'T':      true,
	'U':      true,
	'V':      true,
	'W':      true,
	'X':      true,
	'Y':      true,
	'Z':      true,
	'[':      true,
	'\\':     false,
	']':      true,
	'^':      true,
	'_':      true,
	'`':      true,
	'a':      true,
	'b':      true,
	'c':      true,
	'd':      true,
	'e':      true,
	'f':      true,
	'g':      true,
	'h':      true,
	'i':      true,
	'j':      true,
	'k':      true,
	'l':      true,
	'm':      true,
	'n':      true,
	'o':      true,
	'p':      true,
	'q':      true,
	'r':      true,
	's':      true,
	't':      true,
	'u':      true,
	'v':      true,
	'w':      true,
	'x':      true,
	'y':      true,
	'z':      true,
	'{':      true,
	'|':      true,
	'}':      true,
	'~':      true,
	'\u007f': true,
}

// NOTE: keep in sync with string above.
func stringBytes(e *bytes.Buffer, s []byte, escapeHTML bool) {
	e.WriteRune('"')
	start := 0
	for i := 0; i < len(s); {
		if b := s[i]; b < utf8.RuneSelf {
			if htmlSafeSet[b] || (!escapeHTML && safeSet[b]) {
				i++
				continue
			}
			if start < i {
				e.Write(s[start:i])
			}
			e.WriteRune('\\')
			switch b {
			case '\\', '"':
				e.WriteRune(rune(b))
			case '\n':
				e.WriteRune('n')
			case '\r':
				e.WriteRune('r')
			case '\t':
				e.WriteRune('t')
			default:
				// This encodes bytes < 0x20 except for \t, \n and \r.
				// If escapeHTML is set, it also escapes <, >, and &
				// because they can lead to security holes when
				// user-controlled strings are rendered into JSON
				// and served to some browsers.
				e.WriteString(`u00`)
				e.WriteByte(hex[b>>4])
				e.WriteByte(hex[b&0xF])
			}
			i++
			start = i
			continue
		}
		c, size := utf8.DecodeRune(s[i:])
		if c == utf8.RuneError && size == 1 {
			if start < i {
				e.Write(s[start:i])
			}
			e.WriteString(`\ufffd`)
			i += size
			start = i
			continue
		}
		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if c == '\u2028' || c == '\u2029' {
			if start < i {
				e.Write(s[start:i])
			}
			e.WriteString(`\u202`)
			e.WriteByte(hex[c&0xF])
			i += size
			start = i
			continue
		}
		i += size
	}
	if start < len(s) {
		e.Write(s[start:])
	}
	e.WriteByte('"')
}

// MarshalJSON encodes the receiver object to a JSON document.
func (n NaturalLanguageValues) MarshalJSON() ([]byte, error) {
	l := len(n)
	if l <= 0 {
		return nil, nil
	}

	b := bytes.Buffer{}
	if l == 1 {
		val := n.First()
		if len(val) > 0 {
			val = unescape(val)
			stringBytes(&b, val, false)
			return b.Bytes(), nil
		}
	}
	b.Write([]byte{'{'})
	empty := true
	for ref, val := range n {
		if len(val) == 0 {
			continue
		}
		if !empty {
			b.Write([]byte{','})
		}
		if ref.Valid() {
			stringBytes(&b, []byte(ref.String()), false)
			b.Write([]byte{':'})
		}
		stringBytes(&b, val, false)
		empty = len(val) == 0
	}
	b.Write([]byte{'}'})
	if !empty {
		return b.Bytes(), nil
	}
	return nil, nil
}

// First returns the first element in the map
func (n NaturalLanguageValues) First() Content {
	for _, v := range n {
		return v
	}
	return nil
}

// Default returns the default NaturalLanguageValue
func (n NaturalLanguageValues) Default() Content {
	v, _ := n[DefaultLang]
	return v
}

// MarshalText serializes the NaturalLanguageValues into Text
func (n NaturalLanguageValues) MarshalText() ([]byte, error) {
	bb := bytes.Buffer{}
	first := true
	for ref, val := range n {
		if !first {
			bb.WriteString(",")
		}
		bb.Write(val)
		if ref.Valid() {
			bb.WriteString("[")
			bb.WriteString(ref.String())
			bb.WriteString("]")
		}
		first = false
	}
	return bb.Bytes(), nil
}

func (n NaturalLanguageValues) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'q':
		_, _ = io.WriteString(s, "[")
		for _, nn := range n {
			nn.Format(s, verb)
		}
		_, _ = io.WriteString(s, "]")
	case 'v':
		_, _ = io.WriteString(s, "[")
		for _, nn := range n {
			nn.Format(s, verb)
		}
		_, _ = io.WriteString(s, "]")
	}
}

// Append is syntactic sugar for resizing the NaturalLanguageValues map
// and appending an element
func (n *NaturalLanguageValues) Append(lang LangRef, value Content) error {
	//*n = append(*n, LangRefValue{lang, value})
	(*n)[lang] = value
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
	if l.Ref == NilLangRef || !l.Ref.Valid() {
		return l.Value.String()
	}
	return fmt.Sprintf("%s[%s]", l.Value, l.Ref)
}

func DefaultLangRef(value string) LangRefValue {
	return LangRefValue{Ref: DefaultLang, Value: Content(value)}
}

func LangRefValueNew(lang LangRef, value string) LangRefValue {
	return LangRefValue{Ref: lang, Value: Content(value)}
}

func (l LangRefValue) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'q':
		if l.Ref == NilLangRef {
			_, _ = io.WriteString(s, string(l.Value))
		} else {
			_, _ = fmt.Fprintf(s, "%q[%s]", l.Value, l.Ref)
		}
	case 'v':
		if l.Ref == NilLangRef {
			_, _ = fmt.Fprintf(s, "%q", string(l.Value))
		} else {
			_, _ = fmt.Fprintf(s, "%q[%s]", string(l.Value), l.Ref)
		}
	}
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
			l.Ref = MakeRef(key)
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

// MarshalJSON encodes the receiver object to a JSON document.
func (l LangRefValue) MarshalJSON() ([]byte, error) {
	buf := bytes.Buffer{}
	if l.Ref != NilLangRef && l.Ref.Valid() {
		if l.Value.Equals(Content("")) {
			return nil, nil
		}
		stringBytes(&buf, []byte(l.Ref.String()), false)
		buf.Write([]byte{':'})
	}
	stringBytes(&buf, l.Value, false)
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
	if len(l.Value) == 0 && !l.Ref.Valid() {
		return []byte{}, nil
	}
	b := new(bytes.Buffer)
	gg := gob.NewEncoder(b)
	mm := kv{
		K: []byte(l.Ref.String()),
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
	l.Ref = MakeRef(mm.K)
	l.Value = mm.V
	return nil
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

func (c Content) Format(s fmt.State, verb rune) {
	switch verb {
	case 's', 'q':
		_, _ = io.WriteString(s, string(c))
	case 'v':
		_, _ = fmt.Fprintf(s, "%q", string(c))
	}
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
	if n == nil {
		return errors.Newf("nil %T value to unmarshal to", n)
	}
	p := fastjson.Parser{}
	val, err := p.ParseBytes(data)
	if err != nil {
		// try our luck if data contains an unquoted string
		return n.Append(NilLangRef, unescape(data))
	}
	switch val.Type() {
	case fastjson.TypeObject:
		ob, _ := val.Object()
		ob.Visit(func(key []byte, v *fastjson.Value) {
			if dat := v.GetStringBytes(); len(dat) > 0 {
				_ = n.Append(MakeRef(key), unescape(dat))
			}
		})
	case fastjson.TypeString:
		if dat := val.GetStringBytes(); len(dat) > 0 {
			_ = n.Append(NilLangRef, unescape(dat))
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
		n.Append(NilLangRef, Content(data[1:len(data)-1]))
	}
	return nil
}

func (n NaturalLanguageValues) GobEncode() ([]byte, error) {
	if len(n) == 0 {
		return []byte{}, nil
	}
	b := new(bytes.Buffer)
	gg := gob.NewEncoder(b)
	mm := make([]kv, 0, len(n))
	for ref, val := range n {
		mm = append(mm, kv{K: []byte(ref.String()), V: val})
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
		_ = n.Append(MakeRef(m.K), m.V)
	}
	return nil
}

// Equals
func (n NaturalLanguageValues) Equals(with NaturalLanguageValues) bool {
	if n.Count() != with.Count() {
		return false
	}

	for l, wv := range with {
		nv, ok := n[l]
		if !ok {
			return false
		}
		if nv.Equals(wv) {
			continue
		}
		return false
	}
	return true
}
