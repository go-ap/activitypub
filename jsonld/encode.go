package jsonld

import (
	"encoding/json"
	"github.com/fatih/structs"
)

type payloadWithContext struct {
	Context Context `json:"@context"`
	Obj     *interface{}
}

func (this *payloadWithContext) MarshalJSON() ([]byte, error) {
	a := make(map[string]interface{})
	a["@context"] = this.Context.Ref()

	objMap := structs.Map(*this.Obj)
	for k, v := range objMap {
		a[k] = v
	}
	return json.Marshal(a)
}

func (this *payloadWithContext) UnmarshalJSON() {}

type Encoder struct{}

func Marshal(v interface{}, c *Context) ([]byte, error) {
	if c != nil {
		p := payloadWithContext{*c, &v}
		return p.MarshalJSON()
	}
	return json.Marshal(v)
}
