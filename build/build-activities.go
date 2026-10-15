package build

import vocab "github.com/go-ap/activitypub"

func Activity(initFn ...InitFn) *vocab.Activity {
	act := new(vocab.Activity)
	for _, maybeFn := range initFn {
		switch fn := maybeFn.(type) {
		case vocab.WithObjectFn:
			_ = vocab.OnObject(act, fn)
		case vocab.WithIntransitiveActivityFn:
			_ = vocab.OnIntransitiveActivity(act, fn)
		case vocab.WithActivityFn:
			_ = vocab.OnActivity(act, fn)
		}
	}
	return act
}

func IntransitiveActivity(initFn ...InitFn) *vocab.IntransitiveActivity {
	act := new(vocab.IntransitiveActivity)
	for _, maybeFn := range initFn {
		switch fn := maybeFn.(type) {
		case vocab.WithObjectFn:
			_ = vocab.OnObject(act, fn)
		case vocab.WithIntransitiveActivityFn:
			_ = vocab.OnIntransitiveActivity(act, fn)
		}
	}
	return act
}

func Create(initFn ...InitFn) *vocab.Activity {
	initFn = append([]InitFn{WithType(vocab.CreateType)}, initFn...)
	return Activity(initFn...)
}

func Update(initFn ...InitFn) *vocab.Activity {
	initFn = append([]InitFn{WithType(vocab.UpdateType)}, initFn...)
	return Activity(initFn...)
}

func Delete(initFn ...InitFn) *vocab.Activity {
	initFn = append([]InitFn{WithType(vocab.DeleteType)}, initFn...)
	return Activity(initFn...)
}

func Like(initFn ...InitFn) *vocab.Activity {
	initFn = append([]InitFn{WithType(vocab.LikeType)}, initFn...)
	return Activity(initFn...)
}

func Dislike(initFn ...InitFn) *vocab.Activity {
	initFn = append([]InitFn{WithType(vocab.DislikeType)}, initFn...)
	return Activity(initFn...)
}

func Announce(initFn ...InitFn) *vocab.Activity {
	initFn = append([]InitFn{WithType(vocab.AnnounceType)}, initFn...)
	return Activity(initFn...)
}

func Follow(initFn ...InitFn) *vocab.Activity {
	initFn = append([]InitFn{WithType(vocab.FollowType)}, initFn...)
	return Activity(initFn...)
}

func Accept(initFn ...InitFn) *vocab.Activity {
	initFn = append([]InitFn{WithType(vocab.AcceptType)}, initFn...)
	return Activity(initFn...)
}

func Reject(initFn ...InitFn) *vocab.Activity {
	initFn = append([]InitFn{WithType(vocab.RejectType)}, initFn...)
	return Activity(initFn...)
}

func Block(initFn ...InitFn) *vocab.Activity {
	initFn = append([]InitFn{WithType(vocab.BlockType)}, initFn...)
	return Activity(initFn...)
}

func Ignore(initFn ...InitFn) *vocab.Activity {
	initFn = append([]InitFn{WithType(vocab.IgnoreType)}, initFn...)
	return Activity(initFn...)
}

func Flag(initFn ...InitFn) *vocab.Activity {
	initFn = append([]InitFn{WithType(vocab.FlagType)}, initFn...)
	return Activity(initFn...)
}

func Undo(initFn ...InitFn) *vocab.Activity {
	initFn = append([]InitFn{WithType(vocab.UndoType)}, initFn...)
	return Activity(initFn...)
}

func Add(initFn ...InitFn) *vocab.Add {
	initFn = append([]InitFn{WithType(vocab.AddType)}, initFn...)
	return Activity(initFn...)
}

func Remove(initFn ...InitFn) *vocab.Activity {
	initFn = append([]InitFn{WithType(vocab.RemoveType)}, initFn...)
	return Activity(initFn...)
}

func Move(initFn ...InitFn) *vocab.Activity {
	initFn = append([]InitFn{WithType(vocab.MoveType)}, initFn...)
	return Activity(initFn...)
}

func WithObject(o ...vocab.Item) vocab.WithActivityFn {
	return func(act *vocab.Activity) error {
		act.Object = vocab.ItemCollection(o).Normalize()
		return nil
	}
}

func WithActor(a ...vocab.Item) vocab.WithIntransitiveActivityFn {
	return func(act *vocab.IntransitiveActivity) error {
		act.Actor = vocab.ItemCollection(a).Normalize()
		return nil
	}
}

func WithOrigin(o ...vocab.Item) vocab.WithIntransitiveActivityFn {
	return func(act *vocab.IntransitiveActivity) error {
		act.Origin = vocab.ItemCollection(o).Normalize()
		return nil
	}
}

func WithTarget(t ...vocab.Item) vocab.WithIntransitiveActivityFn {
	return func(act *vocab.IntransitiveActivity) error {
		act.Target = vocab.ItemCollection(t).Normalize()
		return nil
	}
}

func WithAnyOf(o ...vocab.Item) vocab.WithQuestionFn {
	return func(act *vocab.Question) error {
		act.AnyOf = vocab.ItemCollection(o).Normalize()
		return nil
	}
}

func WithOneOf(o ...vocab.Item) vocab.WithQuestionFn {
	return func(act *vocab.Question) error {
		act.OneOf = vocab.ItemCollection(o).Normalize()
		return nil
	}
}

func Question(initFn ...InitFn) *vocab.Question {
	initFn = append([]InitFn{WithType(vocab.QuestionType)}, initFn...)
	act := new(vocab.Question)
	for _, maybeFn := range initFn {
		switch fn := maybeFn.(type) {
		case vocab.WithObjectFn:
			_ = vocab.OnObject(act, fn)
		case vocab.WithIntransitiveActivityFn:
			_ = vocab.OnIntransitiveActivity(act, fn)
		case vocab.WithQuestionFn:
			_ = vocab.OnQuestion(act, fn)
		}
	}
	return act
}
