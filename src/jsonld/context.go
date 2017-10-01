package jsonld

import (
	"encoding/json"
	"activitypub"
)

type Ref string

type Context struct {
	URL      Ref                                 `json:"_"`
	Language activitypub.NaturalLanguageValue    `json:"@language"`
}

func (c *Context) Ref() Ref {
	return Ref(c.URL)
}
func (r *Ref) MarshalText() ([]byte, error) {
	return []byte(*r), nil
}

func (c *Context) MarshalJSON() ([]byte, error) {
	var a map[string]interface{}
	a = getMap(c)
	return json.Marshal(a)
}
