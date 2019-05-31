package handler

import (
	"errors"
	"fmt"
	"github.com/go-ap/storage"
	as "github.com/go-ap/activitystreams"
	j "github.com/go-ap/jsonld"
	"net/http"
)

// ActivityHandlerFn is the type that we're using to represent handlers that process requests containing
// an ActivityStreams Activity. It needs to implement the http.Handler interface.
//
// It is considered that following the execution of the handler, we return a pair formed of a HTTP status together with
//  an IRI representing a new Object - in the case of transitive activities that had a side effect, or
//  an error.
// In the case of intransitive activities the iri will always be empty.
type ActivityHandlerFn func(CollectionType, *http.Request, storage.ActivitySaver) (as.IRI, int, error)

func (a ActivityHandlerFn) Storage(r *http.Request) (storage.ActivitySaver, error) {
	ctxVal := r.Context().Value(RepositoryKey)
	st, ok := ctxVal.(storage.ActivitySaver)
	if !ok {
		return nil, errors.New("unable to load storage from context")
	}
	return st, nil
}

// ValidMethod validates if the current handler can process the current request
func (a ActivityHandlerFn) ValidMethod(r *http.Request) bool {
	return r.Method != http.MethodPost
}

// ValidateRequest validates if the current handler can process the current request
func (a ActivityHandlerFn) ValidateRequest(r *http.Request) (int, error) {
	if !a.ValidMethod(r) {
		return http.StatusNotAcceptable, fmt.Errorf("invalid HTTP method %s", r.Method)
	}
	return http.StatusOK, nil
}

// ServeHTTP implements the http.Handler interface for the ActivityHandlerFn type
func (a ActivityHandlerFn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var dat []byte
	var iri as.IRI
	var err error
	var status = http.StatusInternalServerError

	writeResponse := func(status int, dat []byte) {
		w.WriteHeader(status)
		if r.Method == http.MethodGet {
			w.Write(dat)
		}
	}

	if status, err = a.ValidateRequest(r); err != nil {
		dat = []byte(err.Error())
		writeResponse(status, dat)
		return
	}

	st, err := a.Storage(r)
	if err != nil {
		dat = []byte(err.Error())
		writeResponse(status, dat)
		return
	}

	if iri, status, err = a(Typer.Type(r), r, st); err != nil {
		dat = []byte(err.Error())
		writeResponse(status, dat)
		return
	}

	w.WriteHeader(status)
	switch status {
	case http.StatusCreated:
		dat = []byte("CREATED")
		w.Header().Set("Location", iri.String())
	case http.StatusGone:
		dat = []byte("DELETED")
	default:
		dat = []byte("OK")
	}
	w.Write(dat)
}

// MethodValidator is the interface need to be implemented to specify if an HTTP request's method
// is supported by the implementor object
type MethodValidator interface {
	ValidMethod(r *http.Request) bool
}

// RequestValidator is the interface need to be implemented to specify if the whole HTTP request
// is valid in the context of the implementor object
type RequestValidator interface {
	ValidateRequest(r *http.Request) (int, error)
}

// CollectionHandlerFn is the type that we're using to represent handlers that will return ActivityStreams
// Collection or OrderedCollection objects. It needs to implement the http.Handler interface.
type CollectionHandlerFn func(CollectionType, *http.Request, storage.CollectionLoader) (as.CollectionInterface, error)

// ValidMethod validates if the current handler can process the current request
func (c CollectionHandlerFn) ValidMethod(r *http.Request) bool {
	return r.Method == http.MethodGet || r.Method == http.MethodHead
}

// ValidateRequest validates if the current handler can process the current request
func (c CollectionHandlerFn) ValidateRequest(r *http.Request) (int, error) {
	if !c.ValidMethod(r) {
		return http.StatusMethodNotAllowed, fmt.Errorf("invalid HTTP method %s", r.Method)
	}
	return http.StatusOK, nil
}

// CtxtKey type alias for the key under which we're storing the Collection Storage in the Request's context
type CtxtKey string

var RepositoryKey = CtxtKey("__repo")

func (c CollectionHandlerFn) Storage(r *http.Request) (storage.CollectionLoader, error) {
	ctxVal := r.Context().Value(RepositoryKey)
	repo, ok := ctxVal.(storage.CollectionLoader)
	if !ok {
		return nil, errors.New("unable to load storage from context")
	}
	return repo, nil
}

// ServeHTTP implements the http.Handler interface for the CollectionHandlerFn type
func (c CollectionHandlerFn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var dat []byte

	var status = http.StatusInternalServerError
	var err error

	writeResponse := func(status int, dat []byte) {
		w.WriteHeader(status)
		if r.Method == http.MethodGet {
			w.Write(dat)
		}
	}

	status, err = c.ValidateRequest(r)
	if err != nil {
		dat = []byte(err.Error())
		writeResponse(status, dat)
		return
	}

	st, err := c.Storage(r)
	if err != nil {
		dat = []byte(err.Error())
		writeResponse(status, dat)
		return
	}

	col, err := c(Typer.Type(r), r, st)
	if err != nil {
		dat = []byte(err.Error())
		writeResponse(status, dat)
		return
	}
	if dat, err = j.WithContext(j.IRI(as.ActivityBaseURI)).Marshal(col); err != nil {
		dat = []byte(err.Error())
	}

	status = http.StatusOK
	writeResponse(status, dat)
}

// ItemHandlerFn is the type that we're using to represent handlers that return ActivityStreams
// objects. It needs to implement the http.Handler interface
type ItemHandlerFn func(*http.Request, storage.ObjectSaver) (as.Item, error)

func (i ItemHandlerFn) Storage(r *http.Request) (storage.ObjectSaver, error) {
	ctxVal := r.Context().Value(RepositoryKey)
	st, ok := ctxVal.(storage.ObjectSaver)
	if !ok {
		return nil, errors.New("unable to load storage from context")
	}
	return st, nil
}

// ValidMethod validates if the current handler can process the current request
func (i ItemHandlerFn) ValidMethod(r *http.Request) bool {
	return r.Method != http.MethodGet && r.Method != http.MethodHead
}

// ValidateRequest validates if the current handler can process the current request
func (i ItemHandlerFn) ValidateRequest(r *http.Request) (int, error) {
	if !i.ValidMethod(r) {
		return http.StatusMethodNotAllowed, fmt.Errorf("invalid HTTP method %s", r.Method)
	}
	return http.StatusOK, nil
}

// ServeHTTP implements the http.Handler interface for the ItemHandlerFn type
func (i ItemHandlerFn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var dat []byte
	var err error
	status := http.StatusInternalServerError

	writeResponse := func(status int, dat []byte) {
		w.WriteHeader(status)
		if r.Method == http.MethodGet {
			w.Write(dat)
		}
	}

	status, err = i.ValidateRequest(r)
	if err != nil {
		dat = []byte(err.Error())
		writeResponse(status, dat)
		return
	}

	st, err := i.Storage(r)
	if err != nil {
		dat = []byte(err.Error())
		writeResponse(status, dat)
		return
	}

	it, err := i(r, st)
	if err != nil {
		dat = []byte(err.Error())
		writeResponse(status, dat)
		return
	}
	if dat, err = j.WithContext(j.IRI(as.ActivityBaseURI)).Marshal(it); err != nil {
		dat = []byte(err.Error())
		writeResponse(status, dat)
		return
	}

	status = http.StatusOK
	writeResponse(status, dat)
}
