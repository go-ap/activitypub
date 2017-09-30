package jsonld

import (
	"encoding/json"
)

type Ref string
type Context struct {
	URL string
}

func (this *Context) Ref() Ref {
	return Ref(this.URL)
}
func (this *Ref) MarshalText() ([]byte, error) {
	return []byte(*this), nil
}
func (this *Context) MarshalJSON() ([]byte, error) {
	return json.Marshal(this)
}
