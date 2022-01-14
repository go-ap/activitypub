package activitypub

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestIRI_GetLink(t *testing.T) {
	val := "http://example.com"
	u := IRI(val)
	if u.GetLink() != IRI(val) {
		t.Errorf("IRI %q should equal %q", u, val)
	}
}

func TestIRI_String(t *testing.T) {
	val := "http://example.com"
	u := IRI(val)
	if u.String() != val {
		t.Errorf("IRI %q should equal %q", u, val)
	}
}

func TestIRI_GetID(t *testing.T) {
	i := IRI("http://example.com")
	if id := i.GetID(); !id.IsValid() || id != ID(i) {
		t.Errorf("ID %q (%T) should equal %q (%T)", id, id, i, ID(i))
	}
}

func TestIRI_GetType(t *testing.T) {
	i := IRI("http://example.com")
	if i.GetType() != IRIType {
		t.Errorf("Invalid type for %T object %s, expected %s", i, i.GetType(), IRIType)
	}
}

func TestIRI_IsLink(t *testing.T) {
	i := IRI("http://example.com")
	if i.IsLink() != true {
		t.Errorf("%T.IsLink() returned %t, expected %t", i, i.IsLink(), true)
	}
}

func TestIRI_IsObject(t *testing.T) {
	i := IRI("http://example.com")
	if i.IsObject() != false {
		t.Errorf("%T.IsObject() returned %t, expected %t", i, i.IsObject(), false)
	}
}

func TestIRI_UnmarshalJSON(t *testing.T) {
	val := "http://example.com"
	i := IRI("")

	err := i.UnmarshalJSON([]byte(val))
	if err != nil {
		t.Error(err)
	}
	if val != i.String() {
		t.Errorf("%T invalid value after Unmarshal %q, expected %q", i, i, val)
	}
}

func TestIRI_MarshalJSON(t *testing.T) {
	value := []byte("http://example.com")
	i := IRI(value)

	v, err := i.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	expected := fmt.Sprintf("%q", value)
	if expected != string(v) {
		t.Errorf("Invalid value after MarshalJSON: %s, expected %s", v, expected)
	}
}

func TestFlattenToIRI(t *testing.T) {
	t.Skipf("TODO")
}

func TestIRI_URL(t *testing.T) {
	t.Skipf("TODO")
}

func TestIRIs_Contains(t *testing.T) {
	t.Skipf("TODO")
}

func TestIRI_Equals(t *testing.T) {
	{
		i1 := IRI("http://example.com")
		i2 := IRI("http://example.com")
		// same host same scheme
		if !i1.Equals(i2, true) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example.com/ana/are/mere")
		i2 := IRI("http://example.com/ana/are/mere")
		// same host, same scheme and same path
		if !i1.Equals(i2, true) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("https://example.com")
		i2 := IRI("http://example.com")
		// same host different scheme
		if !i1.Equals(i2, false) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example.com/ana/are/mere")
		i2 := IRI("https://example.com/ana/are/mere")
		// same host, different scheme and same path
		if !i1.Equals(i2, false) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("https://example.com?ana=mere")
		i2 := IRI("http://example.com?ana=mere")
		// same host different scheme, same query
		if !i1.Equals(i2, false) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("https://example.com?ana=mere&foo=bar")
		i2 := IRI("http://example.com?foo=bar&ana=mere")
		// same host different scheme, same query - different order
		if !i1.Equals(i2, false) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example.com/ana/are/mere?foo=bar&ana=mere")
		i2 := IRI("https://example.com/ana/are/mere?ana=mere&foo=bar")
		// same host, different scheme and same path, same query different order
		if !i1.Equals(i2, false) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("https://example.com?ana=mere")
		i2 := IRI("http://example.com?ana=mere")
		// same host different scheme, same query
		if !i1.Equals(i2, false) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("https://example.com?ana=mere&foo=bar")
		i2 := IRI("http://example.com?foo=bar&ana=mere")
		// same host different scheme, same query - different order
		if !i1.Equals(i2, false) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example.com/ana/are/mere?foo=bar&ana=mere")
		i2 := IRI("https://example.com/ana/are/mere?ana=mere&foo=bar")
		// same host, different scheme and same path, same query different order
		if !i1.Equals(i2, false) {
			t.Errorf("%s should equal %s", i1, i2)
		}
	}
	///
	{
		i1 := IRI("http://example.com")
		i2 := IRI("https://example.com")
		// same host different scheme
		if i1.Equals(i2, true) {
			t.Errorf("%s should not equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example1.com")
		i2 := IRI("http://example.com")
		// different host same scheme
		if i1.Equals(i2, true) {
			t.Errorf("%s should not equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example.com/ana/1are/mere")
		i2 := IRI("http://example.com/ana/are/mere")
		// same host, same scheme and different path
		if i1.Equals(i2, true) {
			t.Errorf("%s should not equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example.com?ana1=mere")
		i2 := IRI("http://example.com?ana=mere")
		// same host same scheme, different query key
		if i1.Equals(i2, false) {
			t.Errorf("%s should not equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example.com?ana=mere")
		i2 := IRI("http://example.com?ana=mere1")
		// same host same scheme, different query value
		if i1.Equals(i2, false) {
			t.Errorf("%s should not equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("https://example.com?ana=mere&foo=bar")
		i2 := IRI("http://example.com?foo=bar1&ana=mere")
		// same host different scheme, different query value - different order
		if i1.Equals(i2, false) {
			t.Errorf("%s should not equal %s", i1, i2)
		}
	}
	{
		i1 := IRI("http://example.com/ana/are/mere?foo=bar&ana=mere")
		i2 := IRI("https://example.com/ana/are/mere?ana=mere&foo1=bar")
		// same host, different scheme and same path, different query key different order
		if i1.Equals(i2, false) {
			t.Errorf("%s should not equal %s", i1, i2)
		}
	}
}

func TestIRI_Contains(t *testing.T) {
	t.Skip("TODO")
}

func TestIRI_IsCollection(t *testing.T) {
	t.Skip("TODO")
}

func TestIRIs_UnmarshalJSON(t *testing.T) {
	type args struct {
		d []byte
	}
	tests := []struct {
		name string
		args args
		obj  IRIs
		want IRIs
		err  error
	}{
		{
			name: "empty",
			args: args{[]byte{'{', '}'}},
			want: nil,
			err:  nil,
		},
		{
			name: "IRI",
			args: args{[]byte("\"http://example.com\"")},
			want: IRIs{IRI("http://example.com")},
			err:  nil,
		},
		{
			name: "IRIs",
			args: args{[]byte(fmt.Sprintf("[%q, %q, %q]", "http://example.com", "http://example.net", "http://example.org"))},
			want: IRIs{
				IRI("http://example.com"),
				IRI("http://example.net"),
				IRI("http://example.org"),
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.obj.UnmarshalJSON(tt.args.d)
			if (err != nil && tt.err == nil) || (err == nil && tt.err != nil) {
				if !errors.Is(err, tt.err) {
					t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.err)
				}
				return
			}
			if !assertDeepEquals(t.Errorf, tt.obj, tt.want) {
				t.Errorf("UnmarshalJSON() got = %#v, want %#v", tt.obj, tt.want)
			}
		})
	}
}

func TestIRIs_MarshalJSON(t *testing.T) {
	value1 := []byte("http://example.com")
	value2 := []byte("http://example.net")
	value3 := []byte("http://example.org")
	i := IRIs{
		IRI(value1),
		IRI(value2),
		IRI(value3),
	}

	v, err := i.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	expected := fmt.Sprintf("[%q, %q, %q]", value1, value2, value3)
	if expected == string(v) {
		t.Errorf("Invalid value after MarshalJSON: %s, expected %s", v, expected)
	}
}

func TestIRI_AddPath(t *testing.T) {
	t.Skip("TODO")
}

func TestIRI_ItemMatches(t *testing.T) {
	t.Skip("TODO")
}

func TestIRI_GobDecode(t *testing.T) {
	tests := []struct {
		name    string
		i       IRI
		data    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			i:       "",
			data:    []byte{},
			wantErr: false,
		},
		{
			name:    "some iri",
			i:       "https://example.com",
			data:    gobValue([]byte("https://example.com")),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.GobDecode(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("GobDecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIRI_GobEncode(t *testing.T) {
	tests := []struct {
		name    string
		i       IRI
		want    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			i:       "",
			want:    []byte{},
			wantErr: false,
		},
		{
			name:    "some iri",
			i:       "https://example.com",
			want:    []byte("https://example.com"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.GobEncode()
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
