package activitypub

import (
	"fmt"
	"reflect"
	"testing"
)

func assertObjectWithTesting(fn canErrorFunc, expected Item) WithObjectFn {
	return func(p *Object) error {
		if !assertDeepEquals(fn, p, expected) {
			return fmt.Errorf("not equal")
		}
		return nil
	}
}

func TestOnObject(t *testing.T) {
	testObject := Object{
		ID: "https://example.com",
	}
	type args struct {
		it Item
		fn func(canErrorFunc, Item) WithObjectFn
	}
	tests := []struct {
		name     string
		args     args
		expected Item
		wantErr  bool
	}{
		{
			name:     "single",
			args:     args{testObject, assertObjectWithTesting},
			expected: &testObject,
			wantErr:  false,
		},
		{
			name:     "single fails",
			args:     args{Object{ID: "https://not-equals"}, assertObjectWithTesting},
			expected: &testObject,
			wantErr:  true,
		},
		{
			name:     "collectionOfObjects",
			args:     args{ItemCollection{testObject, testObject}, assertObjectWithTesting},
			expected: &testObject,
			wantErr:  false,
		},
		{
			name:     "collectionOfObjects fails",
			args:     args{ItemCollection{testObject, Object{ID: "https://not-equals"}}, assertObjectWithTesting},
			expected: &testObject,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		var logFn canErrorFunc
		if tt.wantErr {
			logFn = t.Logf
		} else {
			logFn = t.Errorf
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := OnObject(tt.args.it, tt.args.fn(logFn, tt.expected)); (err != nil) != tt.wantErr {
				t.Errorf("OnObject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func assertActivityWithTesting(fn canErrorFunc, expected Item) WithActivityFn {
	return func(p *Activity) error {
		if !assertDeepEquals(fn, p, expected) {
			return fmt.Errorf("not equal")
		}
		return nil
	}
}

func TestOnActivity(t *testing.T) {
	testActivity := Activity{
		ID: "https://example.com",
	}
	type args struct {
		it Item
		fn func(canErrorFunc, Item) WithActivityFn
	}
	tests := []struct {
		name     string
		args     args
		expected Item
		wantErr  bool
	}{
		{
			name:     "single",
			args:     args{testActivity, assertActivityWithTesting},
			expected: &testActivity,
			wantErr:  false,
		},
		{
			name:     "single fails",
			args:     args{Activity{ID: "https://not-equals"}, assertActivityWithTesting},
			expected: &testActivity,
			wantErr:  true,
		},
		{
			name:     "collectionOfActivitys",
			args:     args{ItemCollection{testActivity, testActivity}, assertActivityWithTesting},
			expected: &testActivity,
			wantErr:  false,
		},
		{
			name:     "collectionOfActivitys fails",
			args:     args{ItemCollection{testActivity, Activity{ID: "https://not-equals"}}, assertActivityWithTesting},
			expected: &testActivity,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		var logFn canErrorFunc
		if tt.wantErr {
			logFn = t.Logf
		} else {
			logFn = t.Errorf
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := OnActivity(tt.args.it, tt.args.fn(logFn, tt.expected)); (err != nil) != tt.wantErr {
				t.Errorf("OnActivity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func assertIntransitiveActivityWithTesting(fn canErrorFunc, expected Item) WithIntransitiveActivityFn {
	return func(p *IntransitiveActivity) error {
		if !assertDeepEquals(fn, p, expected) {
			return fmt.Errorf("not equal")
		}
		return nil
	}
}

func TestOnIntransitiveActivity(t *testing.T) {
	testIntransitiveActivity := IntransitiveActivity{
		ID: "https://example.com",
	}
	type args struct {
		it Item
		fn func(canErrorFunc, Item) WithIntransitiveActivityFn
	}
	tests := []struct {
		name     string
		args     args
		expected Item
		wantErr  bool
	}{
		{
			name:     "single",
			args:     args{testIntransitiveActivity, assertIntransitiveActivityWithTesting},
			expected: &testIntransitiveActivity,
			wantErr:  false,
		},
		{
			name:     "single fails",
			args:     args{IntransitiveActivity{ID: "https://not-equals"}, assertIntransitiveActivityWithTesting},
			expected: &testIntransitiveActivity,
			wantErr:  true,
		},
		{
			name:     "collectionOfIntransitiveActivitys",
			args:     args{ItemCollection{testIntransitiveActivity, testIntransitiveActivity}, assertIntransitiveActivityWithTesting},
			expected: &testIntransitiveActivity,
			wantErr:  false,
		},
		{
			name:     "collectionOfIntransitiveActivitys fails",
			args:     args{ItemCollection{testIntransitiveActivity, IntransitiveActivity{ID: "https://not-equals"}}, assertIntransitiveActivityWithTesting},
			expected: &testIntransitiveActivity,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		var logFn canErrorFunc
		if tt.wantErr {
			logFn = t.Logf
		} else {
			logFn = t.Errorf
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := OnIntransitiveActivity(tt.args.it, tt.args.fn(logFn, tt.expected)); (err != nil) != tt.wantErr {
				t.Errorf("OnIntransitiveActivity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func assertQuestionWithTesting(fn canErrorFunc, expected Item) WithQuestionFn {
	return func(p *Question) error {
		if !assertDeepEquals(fn, p, expected) {
			return fmt.Errorf("not equal")
		}
		return nil
	}
}

func TestOnQuestion(t *testing.T) {
	testQuestion := Question{
		ID: "https://example.com",
	}
	type args struct {
		it Item
		fn func(canErrorFunc, Item) WithQuestionFn
	}
	tests := []struct {
		name     string
		args     args
		expected Item
		wantErr  bool
	}{
		{
			name:     "single",
			args:     args{testQuestion, assertQuestionWithTesting},
			expected: &testQuestion,
			wantErr:  false,
		},
		{
			name:     "single fails",
			args:     args{Question{ID: "https://not-equals"}, assertQuestionWithTesting},
			expected: &testQuestion,
			wantErr:  true,
		},
		{
			name:     "collectionOfQuestions",
			args:     args{ItemCollection{testQuestion, testQuestion}, assertQuestionWithTesting},
			expected: &testQuestion,
			wantErr:  false,
		},
		{
			name:     "collectionOfQuestions fails",
			args:     args{ItemCollection{testQuestion, Question{ID: "https://not-equals"}}, assertQuestionWithTesting},
			expected: &testQuestion,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		var logFn canErrorFunc
		if tt.wantErr {
			logFn = t.Logf
		} else {
			logFn = t.Errorf
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := OnQuestion(tt.args.it, tt.args.fn(logFn, tt.expected)); (err != nil) != tt.wantErr {
				t.Errorf("OnQuestion() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOnCollection(t *testing.T) {
	t.Skipf("TODO")
}

func TestOnCollectionPage(t *testing.T) {
	t.Skipf("TODO")
}

func TestOnOrderedCollectionPage(t *testing.T) {
	t.Skipf("TODO")
}

type args[T Objects] struct {
	it T
	fn func(fn canErrorFunc, expected T) func(*T) error
}

type testPair[T Objects] struct {
	name     string
	args     args[T]
	expected T
	wantErr  bool
}

func assert[T Objects](fn canErrorFunc, expected T) func(*T) error {
	return func(p *T) error {
		if !assertDeepEquals(fn, *p, expected) {
			return fmt.Errorf("not equal")
		}
		return nil
	}
}

func TestOn(t *testing.T) {
	var tests = []testPair[Object]{
		{
			name:     "single object",
			args:     args[Object]{Object{ID: "https://example.com"}, assert[Object]},
			expected: Object{ID: "https://example.com"},
			wantErr:  false,
		},
		{
			name:     "single image",
			args:     args[Image]{Image{ID: "http://example.com"}, assert[Image]},
			expected: Image{ID: "http://example.com"},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		var logFn canErrorFunc
		if tt.wantErr {
			logFn = t.Logf
		} else {
			logFn = t.Errorf
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := On(tt.args.it, tt.args.fn(logFn, tt.expected)); (err != nil) != tt.wantErr {
				t.Errorf("On[%T]() error = %v, wantErr %v", tt.args.it, err, tt.wantErr)
			}
		})
	}
}

var (
	emptyPrintFn = func(string, ...any) {}

	fnPrintObj = func(printFn func(string, ...any)) func(_ *Object) error {
		return func(o *Object) error {
			printFn("%v", o)
			return nil
		}
	}

	fnObj = func(_ *Object) error { return nil }
	fnAct = func(_ *Actor) error { return nil }
	fnA   = func(_ *Activity) error { return nil }
	fnIA  = func(_ *IntransitiveActivity) error { return nil }

	maybeObject               Item = new(Object)
	notObject                 Item = new(Activity)
	maybeActor                Item = new(Actor)
	maybeActivity             Item = new(Activity)
	notIntransitiveActivity   Item = new(Activity)
	maybeIntransitiveActivity Item = new(IntransitiveActivity)
	colOfObjects              Item = ItemCollection{Object{ID: "unum"}, Object{ID: "duo"}}
	colOfNotObjects           Item = ItemCollection{Activity{ID: "unum"}, Activity{ID: "duo"}}
)

func Benchmark_ToObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToObject(maybeObject)
	}
}

func Benchmark_To_T_Object(b *testing.B) {
	for i := 0; i < b.N; i++ {
		To[Object](maybeObject)
	}
}

func Benchmark_ToActor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToActor(maybeActor)
	}
}

func Benchmark_To_T_Actor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		To[Actor](maybeActor)
	}
}

func Benchmark_ToActivity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToActivity(maybeActivity)
	}
}

func Benchmark_To_T_Activity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		To[Activity](maybeActivity)
	}
}

func Benchmark_ToIntransitiveActivityHappy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToIntransitiveActivity(maybeIntransitiveActivity)
	}
}

func Benchmark_To_T_IntransitiveActivityHappy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		To[IntransitiveActivity](maybeIntransitiveActivity)
	}
}

func Benchmark_ToIntransitiveActivityNotHappy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToIntransitiveActivity(notIntransitiveActivity)
	}
}

func Benchmark_To_T_IntransitiveActivityNotHappy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		To[IntransitiveActivity](notIntransitiveActivity)
	}
}

func Benchmark_OnObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OnObject(maybeObject, fnObj)
	}
}

func Benchmark_On_T_Object(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On[Object](maybeObject, fnObj)
	}
}
func Benchmark_OnActor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OnActor(maybeObject, fnAct)
	}
}
func Benchmark_On_T_Actor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On[Actor](maybeObject, fnAct)
	}
}
func Benchmark_OnActivity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OnActivity(maybeObject, fnA)
	}
}

func Benchmark_On_T_Activity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On[Activity](maybeObject, fnA)
	}
}
func Benchmark_OnIntransitiveActivity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OnIntransitiveActivity(maybeObject, fnIA)
	}
}

func Benchmark_On_T_IntransitiveActivity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On[IntransitiveActivity](maybeObject, fnIA)
	}
}
func Benchmark_OnObjectNotHappy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OnObject(notObject, fnObj)
	}
}

func Benchmark_On_T_ObjectNotHappy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On[Object](notObject, fnObj)
	}
}

func Benchmark_OnObjectHappyCol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OnObject(colOfObjects, fnObj)
	}
}

func Benchmark_On_T_ObjectHappyCol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On[Object](colOfObjects, fnObj)
	}
}

func Benchmark_OnObjectNotHappyCol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OnObject(colOfNotObjects, fnObj)
	}
}

func Benchmark_On_T_ObjectNotHappyCol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		On[Object](colOfNotObjects, fnObj)
	}
}

func TestDerefItem(t *testing.T) {
	tests := []struct {
		name string
		arg  Item
		want ItemCollection
	}{
		{
			name: "empty",
		},
		{
			name: "simple object",
			arg:  &Object{ID: "https://example.com"},
			want: ItemCollection{&Object{ID: "https://example.com"}},
		},
		{
			name: "simple IRI",
			arg:  IRI("https://example.com"),
			want: ItemCollection{IRI("https://example.com")},
		},
		{
			name: "IRI collection",
			arg:  IRIs{IRI("https://example.com"), IRI("https://example.com/~jdoe")},
			want: ItemCollection{IRI("https://example.com"), IRI("https://example.com/~jdoe")},
		},
		{
			name: "Item collection",
			arg: ItemCollection{
				&Object{ID: "https://example.com"},
				&Actor{ID: "https://example.com/~jdoe"},
			},
			want: ItemCollection{
				&Object{ID: "https://example.com"},
				&Actor{ID: "https://example.com/~jdoe"},
			},
		},
		{
			name: "mixed item collection",
			arg: ItemCollection{
				&Object{ID: "https://example.com"},
				IRI("https://example.com/666"),
				&Actor{ID: "https://example.com/~jdoe"},
			},
			want: ItemCollection{
				&Object{ID: "https://example.com"},
				IRI("https://example.com/666"),
				&Actor{ID: "https://example.com/~jdoe"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DerefItem(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DerefItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
