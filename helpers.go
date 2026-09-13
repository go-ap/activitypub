package activitypub

import (
	"fmt"

	"github.com/go-ap/errors"
)

func To[T Objects | Links](it LinkOrIRI) (*T, error) {
	if ob, ok := it.(T); ok {
		return &ob, nil
	}
	return nil, fmt.Errorf("invalid cast for object %T", it)
}

// On handles in a generic way the call to fn(*T) if the "it" Item can be asserted to one of the Objects type.
// It also covers the case where "it" is a collection of items that match the assertion.
func On[T Objects | Links](it Item, fn func(*T) error) error {
	return callOnItemCollection[T, func(iri LinkOrIRI) (*T, error)](it, To[T], fn)
}

func callOnItemCollection[T Objects | Links, F func(iri LinkOrIRI) (*T, error)](it LinkOrIRI, fn F, callFn func(*T) error) error {
	call := func(it LinkOrIRI) error {
		if tt, err := fn(it); err != nil {
			return err
		} else {
			return callFn(tt)
		}
	}
	if !IsItemCollection(it) {
		return call(it)
	}
	return OnItemCollection(it, func(col *ItemCollection) error {
		errs := make([]error, 0, len(*col))
		for _, ob := range *col {
			if err := call(ob); err != nil {
				errs = append(errs, err)
			}
		}
		return errors.Join(errs...)
	})
}
