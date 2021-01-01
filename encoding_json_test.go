package activitypub

import (
	"testing"
	"time"
)

func Test_write(t *testing.T) {
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

func Test_writeActivity(t *testing.T) {
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
			if gotNotEmpty := writeActivityJSONValue(tt.args.b, tt.args.a); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeActivityJSONValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeBoolProp(t *testing.T) {
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
			if gotNotEmpty := writeBoolJSONProp(tt.args.b, tt.args.n, tt.args.t); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeBoolJSONProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeComma(t *testing.T) {
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

func Test_writeDurationProp(t *testing.T) {
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
			if gotNotEmpty := writeDurationJSONProp(tt.args.b, tt.args.n, tt.args.d); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeDurationJSONProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeFloatProp(t *testing.T) {
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
			if gotNotEmpty := writeFloatJSONProp(tt.args.b, tt.args.n, tt.args.f); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeFloatJSONProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeIRIProp(t *testing.T) {
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
			if gotNotEmpty := writeIRIJSONProp(tt.args.b, tt.args.n, tt.args.i); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeIRIJSONProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeIntProp(t *testing.T) {
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
			if gotNotEmpty := writeIntJSONProp(tt.args.b, tt.args.n, tt.args.d); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeIntJSONProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeIntransitiveActivity(t *testing.T) {
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
			if gotNotEmpty := writeIntransitiveActivityJSONValue(tt.args.b, tt.args.i); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeIntransitiveActivityJSONValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeItemCollection(t *testing.T) {
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
			if gotNotEmpty := writeItemCollectionJSONValue(tt.args.b, tt.args.col); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeItemCollectionJSONValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeItemCollectionProp(t *testing.T) {
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
			if gotNotEmpty := writeItemCollectionJSONProp(tt.args.b, tt.args.n, tt.args.col); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeItemCollectionJSONProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeItemProp(t *testing.T) {
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
			if gotNotEmpty := writeItemJSONProp(tt.args.b, tt.args.n, tt.args.i); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeItemJSONProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeNaturalLanguageProp(t *testing.T) {
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
			if gotNotEmpty := writeNaturalLanguageJSONProp(tt.args.b, tt.args.n, tt.args.nl); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeNaturalLanguageJSONProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeObject(t *testing.T) {
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
			if gotNotEmpty := writeObjectJSONValue(tt.args.b, tt.args.o); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeObjectJSONValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeProp(t *testing.T) {
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
			if gotNotEmpty := writeJSONProp(tt.args.b, tt.args.name, tt.args.val); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeJSONProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writePropName(t *testing.T) {
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
			if gotNotEmpty := writePropJSONName(tt.args.b, tt.args.s); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writePropJSONName() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeQuestion(t *testing.T) {
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
			if gotNotEmpty := writeQuestionJSONValue(tt.args.b, tt.args.q); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeQuestionJSONValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeS(t *testing.T) {
	type args struct {
		b *[]byte
		s string
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

func Test_writeString(t *testing.T) {
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
			if gotNotEmpty := writeStringJSONValue(tt.args.b, tt.args.s); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeStringJSONValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeStringProp(t *testing.T) {
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
			if gotNotEmpty := writeStringJSONProp(tt.args.b, tt.args.n, tt.args.s); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeStringJSONProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeTimeProp(t *testing.T) {
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
			if gotNotEmpty := writeTimeJSONProp(tt.args.b, tt.args.n, tt.args.t); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeTimeJSONProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func Test_writeValue(t *testing.T) {
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
			if gotNotEmpty := writeJSONValue(tt.args.b, tt.args.s); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeJSONValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	t.Skip("TODO")
}
