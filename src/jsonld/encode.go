package jsonld

import (
	"encoding/json"
	"reflect"
)

type payloadWithContext struct {
	Context Context `json:"@context"`
	Obj     *interface{}
}

func recurse(copy, original reflect.Value) {
	switch original.Kind() {
	// The first cases handle nested structures and translate them recursively

	// If it is a pointer we need to unwrap and call once again
	case reflect.Ptr:
		// To get the actual value of the original we have to call Elem()
		// At the same time this unwraps the pointer so we don't end up in
		// an infinite recursion
		originalValue := original.Elem()
		// Check if the pointer is nil
		if !originalValue.IsValid() {
			return
		}
		// Allocate a new object and set the pointer to it
		copy.Set(reflect.New(originalValue.Type()))
		// Unwrap the newly created pointer
		recurse(copy.Elem(), originalValue)

		// If it is an interface (which is very similar to a pointer), do basically the
		// same as for the pointer. Though a pointer is not the same as an interface so
		// note that we have to call Elem() after creating a new object because otherwise
		// we would end up with an actual pointer
	case reflect.Interface:
		// Get rid of the wrapping interface
		originalValue := original.Elem()
		// Create a new object. Now new gives us a pointer, but we want the value it
		// points to, so we have to call Elem() to unwrap it
		copyValue := reflect.New(originalValue.Type()).Elem()
		recurse(copyValue, originalValue)
		copy.Set(copyValue)

		// If it is a struct we translate each field
	case reflect.Struct:
		for i := 0; i < original.NumField(); i += 1 {
			recurse(copy.Field(i), original.Field(i))
		}

		// If it is a slice we create a new slice and translate each element
	case reflect.Slice:
		copy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i += 1 {
			recurse(copy.Index(i), original.Index(i))
		}

		// If it is a map we create a new map and translate each value
	case reflect.Map:
		copy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			// New gives us a pointer, but again we want the value
			copyValue := reflect.New(originalValue.Type()).Elem()
			recurse(copyValue, originalValue)
			copy.SetMapIndex(key, copyValue)
		}

		// Otherwise we cannot traverse anywhere so this finishes the the recursion

		// If it is a string translate it (yay finally we're doing what we came for)
	case reflect.String:
	default:
		copy.Set(original)
	}
}

func getMap(v interface{}) map[string]interface{} {
	a := make(map[string]interface{})
	typ := reflect.TypeOf(v)
	val := reflect.ValueOf(v)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return a
	}
	for i := 0; i < typ.NumField(); i++ {
		cField := typ.Field(i)
		cValue := val.Field(i)
		if cField.Anonymous {
			for k, v := range getMap(cValue.Interface()) {
				a[k] = v
			}
		} else {
			a[cField.Name] = cValue.Interface()
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
	if c != nil {
		p := payloadWithContext{*c, &v}
		return p.MarshalJSON()
	}
	return json.Marshal(v)
}
