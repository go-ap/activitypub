package jsonld

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

const (
	tagLabel       = "jsonld"
	tagOmitEmpty   = "omitempty"
	tagCollapsible = "collapsible"
)

func Marshal(v interface{}, c *Context) ([]byte, error) {
	p := payloadWithContext{c, &v}
	return p.MarshalJSON()
}

type payloadWithContext struct {
	Context *Context `jsonld:"@context,omitempty,collapsible"`
	Obj     *interface{}
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	case reflect.Struct:
		return func(reflect.Value) bool {
			var ret bool = true
			for i := 0; i < v.NumField(); i++ {
				ret = ret && isEmptyValue(v.Field(i))
			}
			return ret
		}(v)
	}
	return false
}

func reflectToJsonLdMap(v interface{}) map[string]interface{} {
	a := make(map[string]interface{})
	typ := reflect.TypeOf(v)
	val := reflect.ValueOf(v)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		cField := typ.Field(i)
		cValue := val.Field(i)
		cTag := cField.Tag

		jsonLdTag, ok := loadJsonLdTag(cTag)
		omitEmpty := ok && jsonLdTag.omitEmpty
		if jsonLdTag.ignore {
			continue
		}
		if cField.Anonymous {
			for k, v := range reflectToJsonLdMap(cValue.Interface()) {
				a[k] = v
			}
			continue
		}
		empty := isEmptyValue(cValue)
		if !empty || empty && !omitEmpty {
			a[jsonLdName(cField.Name, jsonLdTag)] = cValue.Interface()
		}
	}

	return a
}

func (p *payloadWithContext) MarshalJSON() ([]byte, error) {
	a := make(map[string]interface{})
	if p.Context != nil {
		typ := reflect.TypeOf(*p)
		cMirror, _ := typ.FieldByName("Context")
		jsonLdTag, ok := loadJsonLdTag(cMirror.Tag)
		omitEmpty := ok && jsonLdTag.omitEmpty
		collapsible := ok && jsonLdTag.collapsible

		con := reflectToJsonLdMap(p.Context)
		if len(con) > 0 || !omitEmpty {
			for _, v := range con {
				a[jsonLdName(cMirror.Name, jsonLdTag)] = v
				if len(con) == 1 && collapsible {
					break
				}
			}
		}
	}
	if *p.Obj != nil {
		oMap := reflectToJsonLdMap(*p.Obj)
		if len(oMap) == 0 {
			return nil, fmt.Errorf("invalid object to marshall")
		}
		for k, v := range oMap {
			a[k] = v
		}
	}
	return json.Marshal(a)
}

func (p *payloadWithContext) UnmarshalJSON() {}

type Encoder struct{}

type jsonLdTag struct {
	name        string
	ignore      bool
	omitEmpty   bool
	collapsible bool
}

func loadJsonLdTag(tag reflect.StructTag) (jsonLdTag, bool) {
	jlTag, ok := tag.Lookup(tagLabel)
	if !ok {
		return jsonLdTag{}, false
	}
	val := strings.Split(jlTag, ",")
	cont := func(arr []string, s string) bool {
		for _, v := range arr {
			if v == s {
				return true
			}
		}
		return false
	}
	t := jsonLdTag{
		omitEmpty:   cont(val, tagOmitEmpty),
		collapsible: cont(val, tagCollapsible),
	}
	t.name, t.ignore = func(v string) (string, bool) {
		if len(v) > 0 && v != "_" {
			return v, false
		} else {
			return "", true
		}
	}(val[0])

	return t, true
}

func jsonLdName(n string, tag jsonLdTag) string {
	if len(tag.name) > 0 {
		return tag.name
	}
	return n
}
