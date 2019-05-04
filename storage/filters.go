package storage

import (
	as "github.com/go-ap/activitystreams"
)

// Filterable
type Filterable interface {
	Types() []as.ActivityVocabularyType
	IRIs() []as.IRI
}

// Paginator
type Paginator interface {
	QueryString() string
	BasePage() Paginator
	CurrentPage() Paginator
	NextPage() Paginator
	PrevPage() Paginator
	FirstPage() Paginator
	CurrentIndex() int
}
