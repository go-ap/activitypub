package build

import vocab "github.com/go-ap/activitypub"

func setLinkPropsFromObject(lnk *vocab.Link, ob vocab.Object) {
	if ob.ID != "" {
		lnk.ID = ob.ID
	}
	if ob.Type != nil {
		lnk.Type = ob.Type
	}
	if ob.Name != nil {
		lnk.Name = ob.Name
	}
	if ob.MediaType != "" {
		lnk.MediaType = ob.MediaType
	}
	if !vocab.IsNil(ob.Preview) {
		lnk.Preview = ob.Preview
	}
}

func WithHref(i vocab.IRI) vocab.WithLinkFn {
	return func(l *vocab.Link) error {
		l.Href = i
		return nil
	}
}

func Link(initFn ...InitFn) *vocab.Link {
	lnk := new(vocab.Link)
	_ob := vocab.Object{}
	for _, maybeFn := range initFn {
		switch fn := maybeFn.(type) {
		case vocab.WithObjectFn:
			_ = vocab.OnObject(&_ob, fn)
		case vocab.WithLinkFn:
			_ = vocab.OnLink(lnk, fn)
		}
	}
	setLinkPropsFromObject(lnk, _ob)
	return lnk
}

func Mention(initFn ...InitFn) *vocab.Mention {
	initFn = append([]InitFn{WithType(vocab.MentionType)}, initFn...)
	return Link(initFn...)
}
