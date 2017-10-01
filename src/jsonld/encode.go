package jsonld

import (
	"encoding/json"
	"reflect"
)

type payloadWithContext struct {
	Context Context `json:"@context"`
	Obj     *interface{}
}

func IsEmpty(v reflect.Value) bool {
	var ret bool
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		ret = v.IsNil()
	case reflect.Slice, reflect.Map:
		ret = v.Len() == 0
	case reflect.Struct:
		ret = func(reflect.Value) bool {
			var ret bool = true
			for i := 0; i < v.NumField(); i++ {
				ret = ret && IsEmpty(v.Field(i))
			}
			return ret
		}(v)
	case reflect.String:
		ret = v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ret = v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ret = v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		ret = v.Float() == 0.0
	}
	return ret
}

func getMap(v interface{}) map[string]interface{} {
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
		if cField.Anonymous {
			for k, v := range getMap(cValue.Interface()) {
				a[k] = v
			}
		} else {
			if !IsEmpty(cValue) {
				a[cField.Name] = cValue.Interface()
			}
		}
	}

	return a
}

func (p *payloadWithContext) MarshalJSON() ([]byte, error) {
	a := make(map[string]interface{})
	a["@context"] = p.Context.Ref()

	for k, v := range getMap(*p.Obj) {
		a[k] = v
	}

	return json.Marshal(a)
}

func (p *payloadWithContext) UnmarshalJSON() {}

type Encoder struct{}

func Marshal(v interface{}, c *Context) ([]byte, error) {
	if c == nil {
		return json.Marshal(v)
	}
	p := payloadWithContext{*c, &v}
	return p.MarshalJSON()
}
