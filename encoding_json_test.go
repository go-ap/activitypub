package activitypub

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func Test_JSONWrite(t *testing.T) {
	type args struct {
		b *[]byte
		c []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_JSONWriteActivity(t *testing.T) {
	type args struct {
		b *[]byte
		a Activity
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteActivityValue(tt.args.b, tt.args.a); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteActivityValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteBoolProp(t *testing.T) {
	type args struct {
		b *[]byte
		n string
		t bool
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteBoolProp(tt.args.b, tt.args.n, tt.args.t); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteBoolProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteComma(t *testing.T) {
	type args struct {
		b *[]byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_JSONWriteDurationProp(t *testing.T) {
	type args struct {
		b *[]byte
		n string
		d time.Duration
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteDurationProp(tt.args.b, tt.args.n, tt.args.d); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteDurationProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteFloatProp(t *testing.T) {
	type args struct {
		b *[]byte
		n string
		f float64
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteFloatProp(tt.args.b, tt.args.n, tt.args.f); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteFloatProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteIRIProp(t *testing.T) {
	type args struct {
		b *[]byte
		n string
		i LinkOrIRI
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteIRIProp(tt.args.b, tt.args.n, tt.args.i); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteIRIProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteIntProp(t *testing.T) {
	type args struct {
		b *[]byte
		n string
		d int64
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteIntProp(tt.args.b, tt.args.n, tt.args.d); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteIntProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteIntransitiveActivity(t *testing.T) {
	type args struct {
		b *[]byte
		i IntransitiveActivity
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteIntransitiveActivityValue(tt.args.b, tt.args.i); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteIntransitiveActivityValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteItemCollection(t *testing.T) {
	type args struct {
		b   *[]byte
		col ItemCollection
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteItemCollectionValue(tt.args.b, tt.args.col, true); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteItemCollectionValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteItemCollectionProp(t *testing.T) {
	type args struct {
		b   *[]byte
		n   string
		col ItemCollection
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteItemCollectionProp(tt.args.b, tt.args.n, tt.args.col, true); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteItemCollectionProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteItemProp(t *testing.T) {
	type args struct {
		b *[]byte
		n string
		i Item
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteItemProp(tt.args.b, tt.args.n, tt.args.i); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteItemProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteNaturalLanguageProp(t *testing.T) {
	type args struct {
		b  *[]byte
		n  string
		nl NaturalLanguageValues
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteNaturalLanguageProp(tt.args.b, tt.args.n, tt.args.nl); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteNaturalLanguageProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteObjectValue(t *testing.T) {
	type args struct {
		b *[]byte
		o Object
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteObjectValue(tt.args.b, tt.args.o); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteObjectValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteProp(t *testing.T) {
	type args struct {
		b    *[]byte
		name string
		val  []byte
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteProp(tt.args.b, tt.args.name, tt.args.val); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWritePropName(t *testing.T) {
	type args struct {
		b *[]byte
		s string
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWritePropName(tt.args.b, tt.args.s); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWritePropName() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteQuestionValue(t *testing.T) {
	type args struct {
		b *[]byte
		q Question
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteQuestionValue(tt.args.b, tt.args.q); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteQuestionValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteStringValue(t *testing.T) {
	type args struct {
		b *[]byte
		s string
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteStringValue(tt.args.b, tt.args.s); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteStringValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteStringProp(t *testing.T) {
	type args struct {
		b *[]byte
		n string
		s string
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteStringProp(tt.args.b, tt.args.n, tt.args.s); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteStringProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteTimeProp(t *testing.T) {
	type args struct {
		b *[]byte
		n string
		t time.Time
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteTimeProp(tt.args.b, tt.args.n, tt.args.t); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteTimeProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_JSONWriteValue(t *testing.T) {
	type args struct {
		b *[]byte
		s []byte
	}
	tests := []struct {
		name         string
		args         args
		wantNotEmpty bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotEmpty := JSONWriteValue(tt.args.b, tt.args.s); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func mockOb(id IRI, typ ActivityVocabularyType) LinkOrIRI {
	ob := ObjectNew(typ)
	ob.ID = id
	return ob
}

func TestMarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		arg     LinkOrIRI
		want    []byte
		wantErr bool
	}{
		{
			name: "empty",
			arg:  nil,
			want: []byte("null"),
		},
		{
			name: "Link to example.com",
			arg:  LinkNew("https://example.com", MentionType),
			want: []byte(`{"id":"https://example.com","type":"Mention"}`),
		},
		{
			name: "Note",
			arg:  mockOb("https://example.com", NoteType),
			want: []byte(`{"id":"https://example.com","type":"Note"}`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MarshalJSON(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONWriteValue(t *testing.T) {
	buff := func(s int) *[]byte {
		b := make([]byte, 0, s)
		return &b
	}
	type args struct {
		b *[]byte
		s []byte
	}
	tests := []struct {
		name         string
		args         args
		want         []byte
		wantNotEmpty bool
	}{
		{
			name: "empty",
			args: args{
				b: buff(0),
			},
			wantNotEmpty: false,
		},
		{
			name: "do not escape quotes at start-end",
			args: args{
				b: buff(20),
				s: []byte(`"https://example.com"`),
			},
			want:         []byte(`"https://example.com"`),
			wantNotEmpty: true,
		},
		{
			name: "escape quotes inside string",
			args: args{
				b: buff(80),
				s: []byte(`"application/ld+json; profile="https://www.w3.org/ns/activitystreams""`),
			},
			want:         []byte(`"application/ld+json; profile="https://www.w3.org/ns/activitystreams""`),
			wantNotEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNotEmpty := JSONWriteValue(tt.args.b, tt.args.s)
			if gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteStringValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
			if tt.wantNotEmpty && !bytes.Equal(*tt.args.b, tt.want) {
				t.Errorf("JSONWriteValue() = %s, want %s", *tt.args.b, tt.want)
			}
		})
	}
}

func buff(l int) *[]byte {
	b := make([]byte, 0, l)
	return &b
}

func TestJSONWriteStringValue(t *testing.T) {
	type args struct {
		b *[]byte
		s string
	}
	tests := []struct {
		name         string
		args         args
		want         string
		wantNotEmpty bool
	}{
		{
			name:         "empty",
			args:         args{},
			want:         "",
			wantNotEmpty: false,
		},
		{
			name: "escaped quote",
			args: args{
				b: buff(10),
				s: `ana"are`,
			},
			want:         `"ana\"are"`,
			wantNotEmpty: true,
		},
		{
			name: "already escaped quote",
			args: args{
				b: buff(10),
				s: `ana\"are`,
			},
			want:         `"ana\"are"`,
			wantNotEmpty: true,
		},
		{
			name: "already escaped quote and multiple other quotes",
			args: args{
				b: buff(10),
				s: `ana\"""are`,
			},
			want:         `"ana\"\"\"are"`,
			wantNotEmpty: true,
		},
		{
			name: "quote at the end",
			args: args{
				b: buff(10),
				s: `anaare"`,
			},
			want:         `"anaare\""`,
			wantNotEmpty: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNotEmpty := JSONWriteStringValue(tt.args.b, tt.args.s)
			if gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("JSONWriteStringValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
			if tt.wantNotEmpty && tt.want != string(*tt.args.b) {
				t.Errorf("JSONWriteStringValue() = %s, want %s", *tt.args.b, tt.want)
			}
		})
	}
}
