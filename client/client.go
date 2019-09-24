package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	as "github.com/go-ap/activitystreams"
)

type RequestSignFn func(*http.Request) error
type LogFn func(...interface{})

type HttpClient interface {
	Client

	CanSign

	Head(string) (*http.Response, error)
	Get(string) (*http.Response, error)
	Post(string, string, io.Reader) (*http.Response, error)
	Put(string, string, io.Reader) (*http.Response, error)
	Delete(string, string, io.Reader) (*http.Response, error)
}

type CanSign interface {
	SignFn(fn RequestSignFn)
}

type Client interface {
	LoadIRI(as.IRI) (as.Item, error)
}

// UserAgent value that the client uses when performing requests
var UserAgent = "activitypub-go-http-client"
var ContentTypeJsonLD = `application/ld+json; profile="https://www.w3.org/ns/activitystreams"`
var ContentTypeActivityJson = `application/activity+json`

// ErrorLogger
var ErrorLogger LogFn = func(el ...interface{}) {}

// InfoLogger
var InfoLogger LogFn = func(el ...interface{}) {}

var defaultSign RequestSignFn = func(r *http.Request) error { return nil }

type err struct {
	msg string
	iri as.IRI
}

func errorf(i as.IRI, msg string, p ...interface{}) error {
	return &err{
		msg: fmt.Sprintf(msg, p...),
		iri: i,
	}
}

// Error returns the formatted error
func (e *err) Error() string {
	if len(e.iri) > 0 {
		return fmt.Sprintf("%s\nwhen loading: %s", e.msg, e.iri)
	} else {
		return fmt.Sprintf("%s", e.msg)
	}
}

type client struct {
	signFn RequestSignFn
}

func NewClient() *client {
	return &client{
		signFn: defaultSign,
	}
}

func (c *client) SignFn(fn RequestSignFn) {
	c.signFn = fn
}

// LoadIRI tries to dereference an IRI and load the full ActivityPub object it represents
func (c *client) LoadIRI(id as.IRI) (as.Item, error) {
	if len(id) == 0 {
		return nil, errorf(id, "Invalid IRI, nil value")
	}
	if _, err := url.ParseRequestURI(id.String()); err != nil {
		return nil, errorf(id, "Invalid IRI: %s", err)
	}
	var err error
	var obj as.Item

	var resp *http.Response
	if resp, err = c.Get(id.String()); err != nil {
		ErrorLogger(err.Error())
		return obj, err
	}
	if resp == nil {
		err := errorf(id, "Unable to load from the AP end point: nil response")
		ErrorLogger(err)
		return obj, err
	}
	if resp.StatusCode != http.StatusOK {
		err := errorf(id, "Unable to load from the AP end point: invalid status %d", resp.StatusCode)
		ErrorLogger(err)
		return obj, err
	}

	defer resp.Body.Close()
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		ErrorLogger(err)
		return obj, err
	}

	return as.UnmarshalJSON(body)
}

func (c *client) req(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}
	req.Header.Set("User-Agent", UserAgent)
	if method == http.MethodGet {
		req.Header.Add("Accept", ContentTypeJsonLD)
		req.Header.Add("Accept", ContentTypeActivityJson)
		req.Header.Add("Accept", "application/json")
	}
	if method == http.MethodPost {
		req.Header.Set("Content-Type", ContentTypeJsonLD)
	}
	if err = c.signFn(req); err != nil {
		err := errorf(as.IRI(req.URL.String()), "Unable to sign request (method %q, previous error: %s)", req.Method, err)
		return req, err
	}
	return req, nil
}

// Head
func (c client) Head(url string) (resp *http.Response, err error) {
	req, err := c.req(http.MethodHead, url, nil)
	var log LogFn
	if err != nil {
		log = ErrorLogger
	} else {
		log = InfoLogger
	}
	log(http.MethodHead, url)
	return http.DefaultClient.Do(req)
}

// Get wrapper over the functionality offered by the default http.Client object
func (c client) Get(url string) (resp *http.Response, err error) {
	req, err := c.req(http.MethodGet, url, nil)
	var log LogFn
	if err != nil {
		log = ErrorLogger
	} else {
		log = InfoLogger
	}
	log(http.MethodGet, url)
	return http.DefaultClient.Do(req)
}

// Post wrapper over the functionality offered by the default http.Client object
func (c *client) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := c.req(http.MethodPost, url, body)
	var log LogFn
	if err != nil {
		log = ErrorLogger
	} else {
		log = InfoLogger
	}
	log(http.MethodPost, url)
	req.Header.Set("Content-Type", contentType)
	return http.DefaultClient.Do(req)
}

// Put wrapper over the functionality offered by the default http.Client object
func (c client) Put(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := c.req(http.MethodPut, url, body)
	var log LogFn
	if err != nil {
		log = ErrorLogger
	} else {
		log = InfoLogger
	}
	log(http.MethodPut, url)
	req.Header.Set("Content-Type", contentType)
	return http.DefaultClient.Do(req)
}

// Delete wrapper over the functionality offered by the default http.Client object
func (c client) Delete(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := c.req(http.MethodDelete, url, body)
	var log LogFn
	if err != nil {
		log = ErrorLogger
	} else {
		log = InfoLogger
	}
	log(http.MethodDelete, url)
	req.Header.Set("Content-Type", contentType)
	return http.DefaultClient.Do(req)
}
