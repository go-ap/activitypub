package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"strconv"
	"testing"

	json "github.com/go-ap/jsonld"
)

func TestNaturalLanguageValue_MarshalJSON(t *testing.T) {
	p := NaturalLanguageValues{
		{
			"en", Content("the test"),
		},
		{
			"fr", Content("le test"),
		},
	}
	js := "{\"en\":\"the test\",\"fr\":\"le test\"}"
	out, err := p.MarshalJSON()
	if err != nil {
		t.Errorf("Error: '%s'", err)
	}
	if js != string(out) {
		t.Errorf("Different marshal result '%s', instead of '%s'", out, js)
	}
	p1 := NaturalLanguageValues{
		{
			"en", Content("the test"),
		},
	}

	out1, err1 := p1.MarshalJSON()

	if err1 != nil {
		t.Errorf("Error: '%s'", err1)
	}
	txt := `"the test"`
	if txt != string(out1) {
		t.Errorf("Different marshal result '%s', instead of '%s'", out1, txt)
	}
}

func TestLangRefValue_MarshalJSON(t *testing.T) {
	{
		tst := LangRefValue{
			Ref:   NilLangRef,
			Value: Content("test"),
		}
		j, err := tst.MarshalJSON()
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		expected := `"test"`
		if string(j) != expected {
			t.Errorf("Different marshal result '%s', expected '%s'", j, expected)
		}
	}
	{
		tst := LangRefValue{
			Ref:   "en",
			Value: Content("test"),
		}
		j, err := tst.MarshalJSON()
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		expected := `"en":"test"`
		if string(j) != expected {
			t.Errorf("Different marshal result '%s', expected '%s'", j, expected)
		}
	}
	{
		tst := LangRefValue{
			Ref:   "en",
			Value: Content("test\nwith characters\tneeding escaping\r\n"),
		}
		j, err := tst.MarshalJSON()
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		expected := `"en":"test\nwith characters\tneeding escaping\r\n"`
		if string(j) != expected {
			t.Errorf("Different marshal result '%s', expected '%s'", j, expected)
		}
	}
}

func TestLangRefValue_MarshalText(t *testing.T) {
	{
		tst := LangRefValue{
			Ref:   NilLangRef,
			Value: Content("test"),
		}
		j, err := tst.MarshalText()
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		expected := "test"
		if string(j) != expected {
			t.Errorf("Different marshal result '%s', expected '%s'", j, expected)
		}
	}
	{
		tst := LangRefValue{
			Ref:   "en",
			Value: Content("test"),
		}
		j, err := tst.MarshalText()
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		expected := "test[en]"
		if string(j) != expected {
			t.Errorf("Different marshal result '%s', expected '%s'", j, expected)
		}
	}
}

func TestNaturalLanguageValue_Get(t *testing.T) {
	testVal := Content("test")
	a := NaturalLanguageValues{{NilLangRef, testVal}}
	if !a.Get(NilLangRef).Equals(testVal) {
		t.Errorf("Invalid Get result. Expected %s received %s", testVal, a.Get(NilLangRef))
	}
}

func TestNaturalLanguageValue_Set(t *testing.T) {
	testVal := Content("test")
	a := NaturalLanguageValues{{NilLangRef, Content("ana are mere")}}
	err := a.Set(LangRef("en"), testVal)
	if err != nil {
		t.Errorf("Received error when doing Set %s", err)
	}
}

func TestNaturalLanguageValue_Append(t *testing.T) {
	var a NaturalLanguageValues

	if len(a) != 0 {
		t.Errorf("Invalid initialization of %T. Size %d > 0 ", a, len(a))
	}
	langEn := LangRef("en")
	valEn := Content("random value")

	a.Append(langEn, valEn)
	if len(a) != 1 {
		t.Errorf("Invalid append of one element to %T. Size %d != 1", a, len(a))
	}
	if !a.Get(langEn).Equals(valEn) {
		t.Errorf("Invalid append of one element to %T. Value of %q not equal to %q, but %q", a, langEn, valEn, a.Get(langEn))
	}
	langDe := LangRef("de")
	valDe := Content("randomisch")
	a.Append(langDe, valDe)

	if len(a) != 2 {
		t.Errorf("Invalid append of one element to %T. Size %d != 2", a, len(a))
	}
	if !a.Get(langEn).Equals(valEn) {
		t.Errorf("Invalid append of one element to %T. Value of %q not equal to %q, but %q", a, langEn, valEn, a.Get(langEn))
	}
	if !a.Get(langDe).Equals(valDe) {
		t.Errorf("Invalid append of one element to %T. Value of %q not equal to %q, but %q", a, langDe, valDe, a.Get(langDe))
	}
}

func TestLangRef_UnmarshalJSON(t *testing.T) {
	lang := "en-US"
	json := `"` + lang + `"`

	var a LangRef
	a.UnmarshalJSON([]byte(json))

	if string(a) != lang {
		t.Errorf("Invalid json unmarshal for %T. Expected %q, found %q", lang, lang, string(a))
	}
}

func TestNaturalLanguageValue_UnmarshalFullObjectJSON(t *testing.T) {
	langEn := "en-US"
	valEn := Content("random")
	langDe := "de-DE"
	valDe := Content("zufällig\n")

	// m := make(map[string]string)
	// m[langEn] = valEn
	// m[langDe] = valDe

	json := `{
		"` + langEn + `": "` + valEn.String() + `",
		"` + langDe + `": "` + valDe.String() + `"
	}`

	a := make(NaturalLanguageValues, 0)
	_ = a.Append(LangRef(langEn), valEn)
	_ = a.Append(LangRef(langDe), valDe)
	err := a.UnmarshalJSON([]byte(json))
	if err != nil {
		t.Error(err)
	}
	for lang, val := range a {
		if val.Ref != LangRef(langEn) && val.Ref != LangRef(langDe) {
			t.Errorf("Invalid json unmarshal for %T. Expected lang %q or %q, found %q", a, langEn, langDe, lang)
		}

		if val.Ref == LangRef(langEn) && !val.Value.Equals(valEn) {
			t.Errorf("Invalid json unmarshal for %T. Expected value %q, found %q", a, valEn, val.Value)
		}
		if val.Ref == LangRef(langDe) && !val.Value.Equals(valDe) {
			t.Errorf("Invalid json unmarshal for %T. Expected value %q, found %q", a, valDe, val.Value)
		}
	}
}

func TestNaturalLanguageValue_UnmarshalJSON(t *testing.T) {
	l := LangRef("")
	dataEmpty := []byte("")

	l.UnmarshalJSON(dataEmpty)
	if l != "" {
		t.Errorf("Unmarshaled object %T should be an empty string, received %q", l, l)
	}
}

func TestNaturalLanguageValue_UnmarshalText(t *testing.T) {
	l := LangRef("")
	dataEmpty := []byte("")

	l.UnmarshalText(dataEmpty)
	if l != "" {
		t.Errorf("Unmarshaled object %T should be an empty string, received %q", l, l)
	}
}

func TestNaturalLanguageValue_First(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValueNew(t *testing.T) {
	n := NaturalLanguageValuesNew()

	if len(n) != 0 {
		t.Errorf("Initial %T should have length 0, received %d", n, len(n))
	}
}

func TestNaturalLanguageValue_MarshalText(t *testing.T) {
	nlv := LangRefValue{
		Ref:   "en",
		Value: Content("test"),
	}
	tst := NaturalLanguageValues{nlv}
	j, err := tst.MarshalText()
	if err != nil {
		t.Errorf("Error marshaling: %s", err)
	}
	if j == nil {
		t.Errorf("Error marshaling: nil value returned")
	}
	expected := fmt.Sprintf("%s[%s]", nlv.Value, nlv.Ref)
	if string(j) != expected {
		t.Errorf("Wrong value: %s, expected %s", j, expected)
	}
}

func TestNaturalLanguageValues_Append(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_First(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_Get(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_MarshalJSON(t *testing.T) {
	{
		m := NaturalLanguageValues{
			{
				"en", Content("test"),
			},
			{
				"de", Content("test"),
			},
		}
		result, err := m.MarshalJSON()
		if err != nil {
			t.Errorf("Failed marshaling '%v'", err)
		}
		mRes := "{\"en\":\"test\",\"de\":\"test\"}"
		if string(result) != mRes {
			t.Errorf("Different results '%v' vs. '%v'", string(result), mRes)
		}
		// n := NaturalLanguageValuesNew()
		// result, err := n.MarshalJSON()

		s := make(map[LangRef]string)
		s["en"] = "test"
		n1 := NaturalLanguageValues{{
			"en", Content("test"),
		}}
		result1, err1 := n1.MarshalJSON()
		if err1 != nil {
			t.Errorf("Failed marshaling '%v'", err1)
		}
		mRes1 := `"test"`
		if string(result1) != mRes1 {
			t.Errorf("Different results '%v' vs. '%v'", string(result1), mRes1)
		}
	}
	{
		nlv := LangRefValue{
			Ref:   NilLangRef,
			Value: Content("test"),
		}
		tst := NaturalLanguageValues{nlv}
		j, err := tst.MarshalJSON()
		if err != nil {
			t.Errorf("Error marshaling: %s", err)
		}
		if j == nil {
			t.Errorf("Error marshaling: nil value returned")
		}
		expected := fmt.Sprintf("\"%s\"", nlv.Value)
		if string(j) != expected {
			t.Errorf("Wrong value: %s, expected %s", j, expected)
		}
	}
	{
		nlv := LangRefValue{
			Ref:   "en",
			Value: Content("test"),
		}
		tst := NaturalLanguageValues{nlv}
		j, err := tst.MarshalJSON()
		if err != nil {
			t.Errorf("Error marshaling: %s", err)
		}
		if j == nil {
			t.Errorf("Error marshaling: nil value returned")
		}
		expected := fmt.Sprintf("\"%s\"", nlv.Value)
		if string(j) != expected {
			t.Errorf("Wrong value: %s, expected %s", j, expected)
		}
	}
	{
		nlvEn := LangRefValue{
			Ref:   "en",
			Value: Content("test"),
		}
		nlvFr := LangRefValue{
			Ref:   "fr",
			Value: Content("teste"),
		}
		tst := NaturalLanguageValues{nlvEn, nlvFr}
		j, err := tst.MarshalJSON()
		if err != nil {
			t.Errorf("Error marshaling: %s", err)
		}
		if j == nil {
			t.Errorf("Error marshaling: nil value returned")
		}
		expected := fmt.Sprintf("{\"%s\":\"%s\",\"%s\":\"%s\"}", nlvEn.Ref, nlvEn.Value, nlvFr.Ref, nlvFr.Value)
		if string(j) != expected {
			t.Errorf("Wrong value: %s, expected %s", j, expected)
		}
	}
	{
		nlvEn := LangRefValue{
			Ref:   "en",
			Value: Content("test\nwith new line"),
		}
		nlvFr := LangRefValue{
			Ref:   "fr",
			Value: Content("teste\navec une ligne nouvelle"),
		}
		tst := NaturalLanguageValues{nlvEn, nlvFr}
		j, err := tst.MarshalJSON()
		if err != nil {
			t.Errorf("Error marshaling: %s", err)
		}
		if j == nil {
			t.Errorf("Error marshaling: nil value returned")
		}
		expected := fmt.Sprintf("{\"%s\":%s,\"%s\":%s}", nlvEn.Ref, strconv.Quote(nlvEn.Value.String()), nlvFr.Ref, strconv.Quote(nlvFr.Value.String()))
		if string(j) != expected {
			t.Errorf("Wrong value: %s, expected %s", j, expected)
		}
	}
}

func TestNaturalLanguageValues_MarshalText(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_Set(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_UnmarshalJSON(t *testing.T) {
	{
		lang := []byte{'e', 'n'}
		val := []byte{'a', 'n', 'a', ' ', 'a', 'r', 'e', ' ', 'm', 'e', 'r', 'e', '\n'}
		js := fmt.Sprintf(`[{"%s": "%s"}]`, lang, val)
		n := NaturalLanguageValues{}
		err := n.UnmarshalJSON([]byte(js))
		if err != nil {
			t.Errorf("Unexpected error when unmarshaling %T: %s", n, err)
		}

		if n.Count() != 1 {
			t.Errorf("Invalid number of elements %d, expected %d", n.Count(), 1)
		}
		l := n.First()
		if !l.Value.Equals(Content("ana are mere\n")) {
			t.Errorf("Invalid %T value %q, expected %q", l, l.Value, "ana are mere\n")
		}
		if l.Ref != "en" {
			t.Errorf("Invalid %T ref %q, expected %q", l, l.Ref, "en")
		}
	}
	{
		ob := make(map[string]string)
		ob["en"] = "ana are mere\n"
		js, err := json.Marshal(ob)
		if err != nil {
			t.Errorf("Unexpected error when marshaling %T: %s", ob, err)
		}
		n := NaturalLanguageValues{}
		err = n.UnmarshalJSON(js)
		if err != nil {
			t.Errorf("Unexpected error when unmarshaling %T: %s", n, err)
		}

		if n.Count() != 1 {
			t.Errorf("Invalid number of elements %d, expected %d", n.Count(), 1)
		}
		l := n.First()
		if !l.Value.Equals(Content("ana are mere\n")) {
			t.Errorf("Invalid %T value %q, expected %q", l, l.Value, "ana are mere\n")
		}
		if l.Ref != "en" {
			t.Errorf("Invalid %T ref %q, expected %q", l, l.Ref, "en")
		}
	}
}

func TestNaturalLanguageValues_UnmarshalText(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValuesNew(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_String(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_Count(t *testing.T) {
	t.Skipf("TODO")
}

func TestNaturalLanguageValues_Equals(t *testing.T) {
	type args struct {
		with NaturalLanguageValues
	}
	tests := []struct {
		name string
		n    NaturalLanguageValues
		args args
		want bool
	}{
		{
			name: "equal-key-value",
			n: NaturalLanguageValues{LangRefValue{
				Ref:   "en",
				Value: Content("test123#"),
			}},
			args: args{
				with: NaturalLanguageValues{LangRefValue{
					Ref:   "en",
					Value: Content("test123#"),
				}},
			},
			want: true,
		},
		{
			name: "not-equal-key",
			n: NaturalLanguageValues{LangRefValue{
				Ref:   "en",
				Value: Content("test123#"),
			}},
			args: args{
				with: NaturalLanguageValues{LangRefValue{
					Ref:   "fr",
					Value: Content("test123#"),
				}},
			},
			want: false,
		},
		{
			name: "not-equal-value",
			n: NaturalLanguageValues{LangRefValue{
				Ref:   "en",
				Value: Content("test123#"),
			}},
			args: args{
				with: NaturalLanguageValues{LangRefValue{
					Ref:   "en",
					Value: Content("test123"),
				}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.Equals(tt.args.with); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContent_String(t *testing.T) {
	t.Skip("TODO")
}

func TestContent_UnmarshalJSON(t *testing.T) {
	t.Skip("TODO")
}

func TestContent_UnmarshalText(t *testing.T) {
	t.Skip("TODO")
}

func gobValue(a interface{}) []byte {
	b := bytes.Buffer{}
	gg := gob.NewEncoder(&b)
	gg.Encode(a)
	return b.Bytes()
}

func TestContent_GobEncode(t *testing.T) {
	tests := []struct {
		name    string
		c       Content
		want    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			c:       Content{},
			want:    []byte{},
			wantErr: false,
		},
		{
			name:    "empty value",
			c:       Content{'0'},
			want:    gobValue([]byte{'0'}),
			wantErr: false,
		},
		{
			name:    "some text",
			c:       Content{'a', 'n', 'a', ' ', 'a', 'r', 'e'},
			want:    gobValue([]byte{'a', 'n', 'a', ' ', 'a', 'r', 'e'}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GobEncode()
			if (err != nil) != tt.wantErr {
				t.Errorf("GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GobEncode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContent_GobDecode(t *testing.T) {
	tests := []struct {
		name    string
		c       Content
		data    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			c:       Content{},
			data:    []byte{},
			wantErr: false,
		},
		{
			name:    "empty value",
			c:       Content{'0'},
			data:    gobValue([]byte{'0'}),
			wantErr: false,
		},
		{
			name:    "some text",
			c:       Content{'a', 'n', 'a', ' ', 'a', 'r', 'e'},
			data:    gobValue([]byte{'a', 'n', 'a', ' ', 'a', 'r', 'e'}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.GobDecode(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("GobDecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLangRef_GobDecode(t *testing.T) {
	tests := []struct {
		name    string
		l       LangRef
		data    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			l:       "",
			data:    []byte{},
			wantErr: false,
		},
		{
			name:    "some text",
			l:       LangRef("ana are"),
			data:    gobValue([]byte("ana are")),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.l.GobDecode(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("GobDecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLangRef_GobEncode(t *testing.T) {
	tests := []struct {
		name    string
		l       LangRef
		want    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			l:       "",
			want:    []byte{},
			wantErr: false,
		},
		{
			name:    "some text",
			l:       LangRef("ana are"),
			want:    gobValue([]byte{'a', 'n', 'a', ' ', 'a', 'r', 'e'}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.l.GobEncode()
			if (err != nil) != tt.wantErr {
				t.Errorf("GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GobEncode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLangRefValue_GobEncode(t *testing.T) {
	type fields struct {
		Ref   LangRef
		Value Content
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			fields:  fields{},
			want:    []byte{},
			wantErr: false,
		},
		{
			name: "some values",
			fields: fields{
				Ref:   "ana",
				Value: Content("are mere"),
			},
			want:    gobValue(kv{K: []byte("ana"), V: []byte("are mere")}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LangRefValue{
				Ref:   tt.fields.Ref,
				Value: tt.fields.Value,
			}
			got, err := l.GobEncode()
			if (err != nil) != tt.wantErr {
				t.Errorf("GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GobEncode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLangRefValue_GobDecode(t *testing.T) {
	type fields struct {
		Ref   LangRef
		Value Content
	}
	tests := []struct {
		name    string
		fields  fields
		data    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			fields:  fields{},
			data:    gobValue(kv{}),
			wantErr: false,
		},
		{
			name: "some values",
			fields: fields{
				Ref:   "ana",
				Value: Content("are mere"),
			},
			data:    gobValue(kv{K: []byte("ana"), V: []byte("are mere")}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LangRefValue{
				Ref:   tt.fields.Ref,
				Value: tt.fields.Value,
			}
			if err := l.GobDecode(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("GobDecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNaturalLanguageValues_GobEncode(t *testing.T) {
	tests := []struct {
		name    string
		n       NaturalLanguageValues
		want    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			n:       NaturalLanguageValues{},
			want:    []byte{},
			wantErr: false,
		},
		{
			name: "some values",
			n: NaturalLanguageValues{{
				Ref:   "ana",
				Value: []byte("are mere"),
			}},
			want:    gobValue([]kv{{K: []byte("ana"), V: []byte("are mere")}}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.GobEncode()
			if (err != nil) != tt.wantErr {
				t.Errorf("GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GobEncode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNaturalLanguageValues_GobDecode(t *testing.T) {
	tests := []struct {
		name    string
		n       NaturalLanguageValues
		data    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			n:       NaturalLanguageValues{},
			data:    []byte{},
			wantErr: false,
		},
		{
			name: "some values",
			n: NaturalLanguageValues{{
				Ref:   "ana",
				Value: []byte("are mere"),
			}},
			data:    gobValue([]kv{{K: []byte("ana"), V: []byte("are mere")}}),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.n.GobDecode(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("GobDecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
