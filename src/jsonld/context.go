package jsonld

import "encoding/json"

type Ref string
type Context struct {
	URL string
}

func (c *Context) Ref() Ref {
	return Ref(c.URL)
}
func (r *Ref) MarshalText() ([]byte, error) {
	return []byte(*r), nil
}

func (c *Context) MarshalJSON() ([]byte, error) {
	return json.Marshal(getMap(c))
}
