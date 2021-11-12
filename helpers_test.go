package activitypub

import (
	"fmt"
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
