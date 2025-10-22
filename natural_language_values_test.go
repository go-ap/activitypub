package activitypub

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"strconv"
	"testing"

	json "github.com/go-ap/jsonld"
	"golang.org/x/text/language"
)

func TestNaturalLanguageValue_MarshalJSON(t *testing.T) {
	p := NaturalLanguageValues{
		English: Content("the test"),
		French:  Content("le test"),
	}
	js1 := `{"en":"the test","fr":"le test"}`
	js2 := `{"fr":"le test","en":"the test"}`
	out, err := p.MarshalJSON()
	if err != nil {
		t.Errorf("Error: '%s'", err)
	}
	if js1 != string(out) && js2 != string(out) {
		t.Errorf("Different marshal result '%s', instead of '%s' or '%s'", out, js1, js2)
	}
	p1 := NaturalLanguageValues{
		English: Content("the test"),
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
			Ref:   MakeRef([]byte("en")),
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
			Ref:   MakeRef([]byte("en")),
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
			Ref:   MakeRef([]byte("en")),
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
	a := NaturalLanguageValues{NilLangRef: testVal}
	if !a.Get(NilLangRef).Equals(testVal) {
		t.Errorf("Invalid Get result. Expected %s received %s", testVal, a.Get(NilLangRef))
	}
}

func TestNaturalLanguageValue_Append(t *testing.T) {
	a := make(NaturalLanguageValues)

	if len(a) != 0 {
		t.Errorf("Invalid initialization of %T. Size %d > 0 ", a, len(a))
	}
	langEn := English
	valEn := Content("random value")

	_ = a.Append(langEn, valEn)
	if len(a) != 1 {
		t.Errorf("Invalid append of one element to %T. Size %d != 1", a, len(a))
	}
	if !a.Get(langEn).Equals(valEn) {
		t.Errorf("Invalid append of one element to %T. Value of %q not equal to %q, but %q", a, langEn, valEn, a.Get(langEn))
	}
	langDe := MakeRef([]byte("de"))
	valDe := Content("randomisch")
	_ = a.Append(langDe, valDe)

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

func TestNaturalLanguageValue_UnmarshalFullObjectJSON(t *testing.T) {
	langEn := []byte(AmericanEnglish.String())
	valEn := Content("random")
	langDe := []byte(German.String())
	valDe := Content("zufällig\n")

	// m := make(map[string]string)
	// m[langEn] = valEn
	// m[langDe] = valDe

	rawJson := `{
		"` + string(langEn) + `": "` + valEn.String() + `",
		"` + string(langDe) + `": "` + valDe.String() + `"
	}`

	a := make(NaturalLanguageValues)
	_ = a.Append(AmericanEnglish, valEn)
	_ = a.Append(German, valDe)
	err := a.UnmarshalJSON([]byte(rawJson))
	if err != nil {
		t.Error(err)
	}
	for lang, val := range a {
		if lang != AmericanEnglish && lang != German {
			t.Errorf("Invalid json unmarshal for %T. Expected lang %q or %q, found %q", a, langEn, langDe, lang)
		}

		if lang == AmericanEnglish && !val.Equals(valEn) {
			t.Errorf("Invalid json unmarshal for %T. Expected value %q, found %q", a, valEn, val)
		}
		if lang == German && !val.Equals(valDe) {
			t.Errorf("Invalid json unmarshal for %T. Expected value %q, found %q", a, valDe, val)
		}
	}
}

func TestNaturalLanguageValue_UnmarshalJSON(t *testing.T) {
	l := NilLangRef
	dataEmpty := []byte("")

	_ = l.UnmarshalJSON(dataEmpty)
	if l != NilLangRef {
		t.Errorf("Unmarshaled object %T should be an empty string, received %q", l, l)
	}
}

func TestNaturalLanguageValue_UnmarshalText(t *testing.T) {
	l := NilLangRef
	dataEmpty := []byte("")

	_ = l.UnmarshalText(dataEmpty)
	if l != NilLangRef {
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
	val := Content("test")
	tst := NaturalLanguageValues{English: val}
	j, err := tst.MarshalText()
	if err != nil {
		t.Errorf("Error marshaling: %s", err)
	}
	if j == nil {
		t.Errorf("Error marshaling: nil value returned")
	}
	expected := fmt.Sprintf("%s[%s]", val, English)
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
			English: Content("test"),
			German:  Content("test"),
		}
		result, err := m.MarshalJSON()
		if err != nil {
			t.Errorf("Failed marshaling '%v'", err)
		}
		mRes1 := `{"en":"test","de":"test"}`
		mRes2 := `{"de":"test","en":"test"}`
		if string(result) != mRes1 && string(result) != mRes2 {
			t.Errorf("Different results '%v' vs. '%v' or '%v'", string(result), mRes1, mRes2)
		}
		// n := NaturalLanguageValuesNew()
		// result, err := n.MarshalJSON()

		s := make(map[LangRef]string)
		s[LangRef(language.English)] = "test"
		n1 := NaturalLanguageValues{
			English: Content("test"),
		}
		result1, err1 := n1.MarshalJSON()
		if err1 != nil {
			t.Errorf("Failed marshaling '%v'", err1)
		}
		mRes3 := `"test"`
		if string(result1) != mRes3 {
			t.Errorf("Different results '%v' vs. '%v'", string(result1), mRes3)
		}
	}
	{
		val := Content("test")
		tst := NaturalLanguageValues{NilLangRef: val}
		j, err := tst.MarshalJSON()
		if err != nil {
			t.Errorf("Error marshaling: %s", err)
		}
		if j == nil {
			t.Errorf("Error marshaling: nil value returned")
		}
		expected := fmt.Sprintf("\"%s\"", val)
		if string(j) != expected {
			t.Errorf("Wrong value: %s, expected %s", j, expected)
		}
	}
	{
		val := Content("test")
		tst := NaturalLanguageValues{English: val}
		j, err := tst.MarshalJSON()
		if err != nil {
			t.Errorf("Error marshaling: %s", err)
		}
		if j == nil {
			t.Errorf("Error marshaling: nil value returned")
		}
		expected := fmt.Sprintf("\"%s\"", val)
		if string(j) != expected {
			t.Errorf("Wrong value: %s, expected %s", j, expected)
		}
	}
	{
		valEn := Content("test")
		valFr := Content("teste")
		tst := NaturalLanguageValues{English: valEn, French: valFr}
		j, err := tst.MarshalJSON()
		if err != nil {
			t.Errorf("Error marshaling: %s", err)
		}
		if j == nil {
			t.Errorf("Error marshaling: nil value returned")
		}
		expected1 := fmt.Sprintf(`{"%s":"%s","%s":"%s"}`, English, valEn, French, valFr)
		expected2 := fmt.Sprintf(`{"%s":"%s","%s":"%s"}`, French, valFr, English, valEn)
		if string(j) != expected1 && string(j) != expected2 {
			t.Errorf("Wrong value: '%s', expected '%s' or '%s'", j, expected1, expected2)
		}
	}
	{
		valEn := Content("test\nwith new line")
		valFr := Content("teste\navec une ligne nouvelle")
		tst := NaturalLanguageValues{English: valEn, French: valFr}
		j, err := tst.MarshalJSON()
		if err != nil {
			t.Errorf("Error marshaling: %s", err)
		}
		if j == nil {
			t.Errorf("Error marshaling: nil value returned")
		}
		expected1 := fmt.Sprintf(`{"%s":%s,"%s":%s}`, English, strconv.Quote(valEn.String()), French, strconv.Quote(valFr.String()))
		expected2 := fmt.Sprintf(`{"%s":%s,"%s":%s}`, French, strconv.Quote(valFr.String()), English, strconv.Quote(valEn.String()))
		if string(j) != expected1 && string(j) != expected2 {
			t.Errorf("Wrong value: '%s', expected '%s' or '%s'", j, expected1, expected2)
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
		if !l.Equals(Content("ana are mere\n")) {
			t.Errorf("Invalid %T value %q, expected %q", l, l, "ana are mere\n")
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
		if !l.Equals(Content("ana are mere\n")) {
			t.Errorf("Invalid %T value %q, expected %q", l, l, "ana are mere\n")
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
			n: NaturalLanguageValues{
				English: Content("test123#"),
			},
			args: args{
				with: NaturalLanguageValues{
					English: Content("test123#"),
				},
			},
			want: true,
		},
		{
			name: "not-equal-key",
			n: NaturalLanguageValues{
				English: Content("test123#"),
			},
			args: args{
				with: NaturalLanguageValues{
					French: Content("test123#"),
				},
			},
			want: false,
		},
		{
			name: "not-equal-value",
			n: NaturalLanguageValues{
				English: Content("test123#"),
			},
			args: args{
				with: NaturalLanguageValues{
					English: Content("test123"),
				},
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
	_ = gg.Encode(a)
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
			l:       NilLangRef,
			data:    []byte{},
			wantErr: false,
		},
		{
			name:    "some text",
			l:       MakeRef([]byte("ana are")),
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
				Ref:   MakeRef([]byte("ana")),
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
			n: NaturalLanguageValues{
				Und: []byte("are mere"),
			},
			want:    gobValue([]kv{{K: []byte("und"), V: []byte("are mere")}}),
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
			n: NaturalLanguageValues{
				Und: []byte("are mere"),
			},
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

func TestLangRefValue_String(t *testing.T) {
	type fields struct {
		Ref   LangRef
		Value Content
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "empty",
			fields: fields{},
			want:   "",
		},
		{
			name: "empty LangRef, not empty Value",
			fields: fields{
				Value: Content("not empty"),
			},
			want: "not empty",
		},
		{
			name: "nil LangRef, empty Value",
			fields: fields{
				Ref:   NilLangRef,
				Value: Content(""),
			},
			want: "",
		},
		{
			name: "nil LangRef, non empty Value",
			fields: fields{
				Ref:   NilLangRef,
				Value: Content("test"),
			},
			want: "test",
		},
		{
			name: "ro-example",
			fields: fields{
				Ref:   MakeRef([]byte("ro")),
				Value: Content("example"),
			},
			want: "example[ro]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LangRefValue{
				Ref:   tt.fields.Ref,
				Value: tt.fields.Value,
			}
			if got := l.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLangRefValue_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		data    []byte
		want    LangRefValue
	}{
		{
			name: "empty",
			data: []byte(""),
			want: LangRefValue{
				Ref: NilLangRef,
			},
		},
		{
			name: "empty LangRef, not empty Value",
			want: LangRefValue{
				Ref:   NilLangRef,
				Value: Content("not empty"),
			},
			data: []byte("not empty"),
		},
		{
			name: "nil LangRef, non empty Value",
			want: LangRefValue{
				Ref:   NilLangRef,
				Value: Content("test"),
			},
			data: []byte("test"),
		},
		{
			name: "ro-example",
			data: []byte(`{"ro":"example"}`),
			want: LangRefValue{
				Ref:   MakeRef([]byte("ro")),
				Value: Content("example"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LangRefValue{}
			if _ = l.UnmarshalJSON(tt.data); !reflect.DeepEqual(l, tt.want) {
				t.Errorf("UnmarshalJSON() got = %+v, want %+v", l, tt.want)
			}
		})
	}
}
