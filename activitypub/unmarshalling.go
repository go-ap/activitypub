package activitypub

import (
	"encoding"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/buger/jsonparser"
)

var (
	apUnmarshalerType   = reflect.TypeOf(new(ObjectOrLink)).Elem()
	unmarshalerType     = reflect.TypeOf(new(json.Unmarshaler)).Elem()
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

func getAPNaturalLanguageField(data []byte, prop string) NaturalLanguageValue {
	n := NaturalLanguageValue{}
	val, typ, _, err := jsonparser.Get(data, prop)
	if err != nil {
		return NaturalLanguageValue(nil)
	}
	switch typ {
	case jsonparser.Object:
		jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			vv, err := jsonparser.GetString(val, string(key))
			n.Append(LangRef(key), vv)
			return err
		}, prop)
	case jsonparser.String:
		n.Append("-", string(val))
	}

	return n
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
