package activitystreams

import (
	"encoding"
	"encoding/json"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/mariusor/activitypub.go/jsonld"
)

var (
	apUnmarshalerType   = reflect.TypeOf(new(Item)).Elem()
	unmarshalerType     = reflect.TypeOf(new(jsonld.Unmarshaler)).Elem()
	textUnmarshalerType = reflect.TypeOf(new(encoding.TextUnmarshaler)).Elem()
)

type mockObj map[string]json.RawMessage

func getType(j json.RawMessage) ActivityVocabularyType {
	mock := make(mockObj, 0)
	json.Unmarshal([]byte(j), &mock)

	for key, val := range mock {
		if strings.ToLower(key) == "type" {
			return ActivityVocabularyType(strings.Trim(string(val), "\""))
		}
	}
	return ""
}

func getAPObjectID(data []byte) ObjectID {
	i, err := jsonparser.GetString(data, "id")
	if err != nil {
		return ObjectID("")
	}
	return ObjectID(i)
}

func getAPType(data []byte) ActivityVocabularyType {
	t, err := jsonparser.GetString(data, "type")
	typ := ActivityVocabularyType(t)
	if err != nil {
		return ActivityVocabularyType("")
	}
	return typ
}

func getAPMimeType(data []byte) MimeType {
	t, err := jsonparser.GetString(data, "mediaType")
	if err != nil {
		return MimeType("")
	}
	return MimeType(t)
}
func getAPInt(data []byte, prop string) int64 {
	val, err := jsonparser.GetInt(data, prop)
	if err != nil {
	}
	return val
}

func getAPString(data []byte, prop string) string {
	val, err := jsonparser.GetString(data, prop)
	if err != nil {
	}
	return val
}

func getAPNaturalLanguageField(data []byte, prop string) NaturalLanguageValue {
	n := NaturalLanguageValue{}
	val, typ, _, err := jsonparser.Get(data, prop)
	if err != nil {
		return nil
	}
	switch typ {
	case jsonparser.Object:
		jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			if dataType == jsonparser.String {
				n.Append(LangRef(key), string(value))
			}
			return err
		})
	case jsonparser.String:
		n.Append(NilLangRef, string(val))
	}

	return n
}

func getAPTime(data []byte, prop string) time.Time {
	t := time.Time{}
	str, _ := jsonparser.GetUnsafeString(data, prop)
	t.UnmarshalText([]byte(str))
	return t
}

func unmarshalToAPObject(data []byte) Item {
	if _, err := url.ParseRequestURI(string(data)); err == nil {
		// try to see if it's an IRI
		return IRI(data)
	}

	i, err := getAPObjectByType(getAPType(data))
	if err != nil {
		return nil
	}
	p := reflect.PtrTo(reflect.TypeOf(i))
	if reflect.TypeOf(i).Implements(unmarshalerType) || p.Implements(unmarshalerType) {
		err = i.(jsonld.Unmarshaler).UnmarshalJSON(data)
	}
	if reflect.TypeOf(i).Implements(textUnmarshalerType) || p.Implements(textUnmarshalerType) {
		err = i.(encoding.TextUnmarshaler).UnmarshalText(data)
	}
	if err != nil {
		return nil
	}
	return i
}

func getAPItem(data []byte, prop string) Item {
	val, typ, _, err := jsonparser.Get(data, prop)
	if err != nil {
		return nil
	}
	switch typ {
	case jsonparser.String:
		if _, err = url.ParseRequestURI(string(val)); err == nil {
			// try to see if it's an IRI
			return IRI(val)
		}
	case jsonparser.Object:
		return unmarshalToAPObject(val)
	case jsonparser.Number:
		fallthrough
	case jsonparser.Array:
		fallthrough
	case jsonparser.Boolean:
		fallthrough
	case jsonparser.Null:
		fallthrough
	case jsonparser.Unknown:
		fallthrough
	default:
		return nil
	}
	return nil
}

func getAPItems(data []byte, prop string) ItemCollection {
	val, typ, _, err := jsonparser.Get(data, prop)
	if err != nil {
		return nil
	}

	var it ItemCollection
	switch typ {
	case jsonparser.Array:
		jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			var i Item
			switch dataType {
			case jsonparser.String:
				if _, err = url.ParseRequestURI(string(val)); err == nil {
					// try to see if it's an IRI
					i = IRI(value)
				}
			case jsonparser.Object:
				i = unmarshalToAPObject(value)
			case jsonparser.Number:
				fallthrough
			case jsonparser.Array:
				fallthrough
			case jsonparser.Boolean:
				fallthrough
			case jsonparser.Null:
				fallthrough
			case jsonparser.Unknown:
				fallthrough
			default:
				return
			}
			if err != nil {
				return
			}
			if i != nil {
				it.Append(i)
			}
		}, prop)
	case jsonparser.Object:
		jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			i := unmarshalToAPObject(value)
			if i != nil {
				it.Append(i)
			}
			return err
		}, prop)
	case jsonparser.String:
		s, _ := jsonparser.GetString(val)
		it.Append(URI(s))
	}
	return it
}
func getAPItemCollection(data []byte, prop string) ItemCollection {
	val, typ, _, err := jsonparser.Get(data, prop)
	if err != nil {
		return nil
	}

	var it ItemCollection
	switch typ {
	case jsonparser.Array:
		jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			i, err := getAPObjectByType(getAPType(value))
			if err != nil {
				return
			}
			err = i.(json.Unmarshaler).UnmarshalJSON(value)
			if err != nil {
				return
			}
			it.Append(i)
		}, prop)
	case jsonparser.Object:
		jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			i := Object{}
			err := i.UnmarshalJSON(val)
			it.Append(i)
			return err
		}, prop)
	case jsonparser.String:
		it.Append(IRI(val))
	}

	return it
}

func getURIField(data []byte, prop string) URI {
	val, err := jsonparser.GetString(data, prop)
	if err != nil {
		return URI("")
	}
	return URI(val)
}

func getAPLangRefField(data []byte, prop string) LangRef {
	val, err := jsonparser.GetString(data, prop)
	if err != nil {
		return LangRef("")
	}
	return LangRef(val)
}

/*
func unmarshal(data []byte, a interface{}) (interface{}, error) {
	ta := make(mockObj, 0)
	err := jsonld.Unmarshal(data, &ta)
	if err != nil {
		return nil, err
	}

	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		cField := typ.Field(i)
		cValue := val.Field(i)
		cTag := cField.Tag
		tag, _ := jsonld.LoadJSONLdTag(cTag)

		var vv reflect.Value
		for key, j := range ta {
			if j == nil {
				continue
			}
			if key == tag.Name {
				if cField.Type.Implements(textUnmarshalerType) {
					m, _ := cValue.Interface().(encoding.TextUnmarshaler)
					m.UnmarshalText(j)
					vv = reflect.ValueOf(m)
				}
				if cField.Type.Implements(unmarshalerType) {
					m, _ := cValue.Interface().(json.Unmarshaler)
					m.UnmarshalJSON(j)
					vv = reflect.ValueOf(m)
				}
				if cField.Type.Implements(apUnmarshalerType) {
					o := getAPObjectByType(getType(j))
					if o != nil {
						jsonld.Unmarshal([]byte(j), o)
						vv = reflect.ValueOf(o)
					}
				}
			}
			if vv.CanAddr() {
				cValue.Set(vv)
				fmt.Printf("\n\nReflected %q %q => %#v\n\n%#v\n", cField.Name, cField.Type, vv, tag.Name)
			}
		}
	}
	return a, nil
}
*/
