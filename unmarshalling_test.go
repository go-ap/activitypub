package activitystreams

import "testing"

func TestUnmarshalJSON(t *testing.T) {
	dataEmpty := []byte("{}")
	i, err := UnmarshalJSON(dataEmpty)
	if err != nil {
		t.Errorf("invalid unmarshalling %s", err)
	}

	o := *i.(*Object)
	validateEmptyObject(o, t)
}
