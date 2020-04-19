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
			if gotNotEmpty := writeActivityValue(tt.args.b, tt.args.a); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeActivityValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeBoolProp(tt.args.b, tt.args.n, tt.args.t); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeBoolProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeDurationProp(tt.args.b, tt.args.n, tt.args.d); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeDurationProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeFloatProp(tt.args.b, tt.args.n, tt.args.f); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeFloatProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeIRIProp(tt.args.b, tt.args.n, tt.args.i); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeIRIProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeIntProp(tt.args.b, tt.args.n, tt.args.d); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeIntProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeIntransitiveActivityValue(tt.args.b, tt.args.i); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeIntransitiveActivityValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeItemCollectionValue(tt.args.b, tt.args.col); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeItemCollectionValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeItemCollectionProp(tt.args.b, tt.args.n, tt.args.col); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeItemCollectionProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeItemProp(tt.args.b, tt.args.n, tt.args.i); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeItemProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeNaturalLanguageProp(tt.args.b, tt.args.n, tt.args.nl); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeNaturalLanguageProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeObjectValue(tt.args.b, tt.args.o); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeObjectValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeProp(tt.args.b, tt.args.name, tt.args.val); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writePropName(tt.args.b, tt.args.s); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writePropName() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeQuestionValue(tt.args.b, tt.args.q); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeQuestionValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeStringValue(tt.args.b, tt.args.s); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeStringValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeStringProp(tt.args.b, tt.args.n, tt.args.s); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeStringProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeTimeProp(tt.args.b, tt.args.n, tt.args.t); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeTimeProp() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
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
			if gotNotEmpty := writeValue(tt.args.b, tt.args.s); gotNotEmpty != tt.wantNotEmpty {
				t.Errorf("writeValue() = %v, want %v", gotNotEmpty, tt.wantNotEmpty)
			}
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	t.Skip("TODO")
}
