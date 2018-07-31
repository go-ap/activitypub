package jsonld

import (
	"encoding/json"
	"strings"
)

// From the JSON-LD spec 3.3
// https://www.w3.org/TR/json-ld/#dfn-keyword
const (
	// @context
	// Used to define the short-hand names that are used throughout a JSON-LD document.
	// These short-hand names are called terms and help developers to express specific identifiers in a compact manner.
	// The @context keyword is described in detail in section 5.1 The Context.
	ContextKw Term = "@context"
	// @id
	//Used to uniquely identify things that are being described in the document with IRIs or blank node identifiers.
	// This keyword is described in section 5.3 Node Identifiers.
	IdKw Term = "@id"
	// @value
	// Used to specify the data that is associated with a particular property in the graph.
	// This keyword is described in section 6.9 String Internationalization and section 6.4 Typed Values.
	ValueKw Term = "@value"
	// @language
	// Used to specify the language for a particular string value or the default language of a JSON-LD document.
	// This keyword is described in section 6.9 String Internationalization.
	LanguageKw Term = "@language"
	//@type
	//Used to set the data type of a node or typed value. This keyword is described in section 6.4 Typed Values.
	TypeKw Term = "@type"
	// @container
	// Used to set the default container type for a term. This keyword is described in section 6.11 Sets and Lists.
	ContainerKw Term = "@container"
	//@list
	//Used to express an ordered set of data. This keyword is described in section 6.11 Sets and Lists.
	ListKw Term = "@list"
	// @set
	// Used to express an unordered set of data and to ensure that values are always represented as arrays.
	// This keyword is described in section 6.11 Sets and Lists.
	SetKw Term = "@set"
	// @reverse
	// Used to express reverse properties. This keyword is described in section 6.12 Reverse Properties.
	ReverseKw Term = "@reverse"
	// @index
	// Used to specify that a container is used to index information and that processing should continue deeper
	// into a JSON data structure. This keyword is described in section 6.16 Data Indexing.
	IndexKw Term = "@index"
	// @base
	// Used to set the base IRI against which relative IRIs are resolved. T
	// his keyword is described in section 6.1 Base IRI.
	BaseKw Term = "@base"
	// @vocab
	// Used to expand properties and values in @type with a common prefix IRI.
	// This keyword is described in section 6.2 Default Vocabulary.
	VocabKw Term = "@vocab"
	// @graph
	// Used to express a graph. This keyword is described in section 6.13 Named Graphs.
	GraphKw Term = "@graph"
)

// Ref basic type
type (
	LangRef string
	Term    string
	IRI     string
	Terms   []Term
)

type IRILike interface {
	IsCompact() bool
	IsAbsolute() bool
	IsRelative() bool
}

func (i IRI) IsCompact() bool {
	return !i.IsAbsolute() && strings.Contains(string(i), ":")
}
func (i IRI) IsAbsolute() bool {
	return strings.Contains(string(i), "https://")
}
func (i IRI) IsRelative() bool {
	return !i.IsAbsolute()
}

var keywords = Terms{
	BaseKw,
	ContextKw,
	ContainerKw,
	GraphKw,
	IdKw,
	IndexKw,
	LanguageKw,
	ListKw,
	ReverseKw,
	SetKw,
	TypeKw,
	ValueKw,
	VocabKw,
}

type ContextObject struct {
	ID   interface{} `jsonld:"@id,omitempty,collapsible"`
	Type interface{} `jsonld:"@type,omitempty,collapsible"`
}

// Context is the basic JSON-LD element. It is used to map terms to IRIs or JSON objects.
// Terms are case sensitive and any valid string that is not a reserved JSON-LD
// keyword can be used as a term.
type Context map[Term]IRI

//type Context Collapsible

// Collapsible is an interface used by the JSON-LD marshaller to collapse a struct to one single value
type Collapsible interface {
	Collapse() interface{}
}

// Collapse returns the plain text collapsed value of the current Context object
func (c *Context) Collapse() interface{} {
	if len(*c) == 1 {
		for _, iri := range *c {
			return iri
		}
	}
	return c
}

// Collapse returns the plain text collapsed value of the current IRI string
func (i IRI) Collapse() interface{} {
	return i
}

// MarshalText basic stringify function
func (i IRI) MarshalText() ([]byte, error) {
	return []byte(i), nil
}

// MarshalJSON returns the JSON document represented by the current Context
func (c *Context) MarshalJSON() ([]byte, error) {
	a := reflectToJSONValue(c)
	if a.isScalar {
		return json.Marshal(a.scalar)
	}
	return json.Marshal(a.object)
}

// UnmarshalJSON tries to load the Context from the incoming json value
func (c *Context) UnmarshalJSON(data []byte) error {
	return nil
}
