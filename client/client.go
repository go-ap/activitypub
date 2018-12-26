package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	as "github.com/mariusor/activitypub.go/activitystreams"
)

type signFn func(*http.Request) error
type logFn func(...interface{})

// UserAgent value that the client uses when performing requests
var UserAgent = "activitypub-go-http-client"

// ErrorLogger
var ErrorLogger logFn = func(el ...interface{}) {}

// InfoLogger
var InfoLogger logFn = func(el ...interface{}) {}

// Sign is the default function to use when signing requests
// Usually this is done using HTTP-Signatures
// See https://github.com/spacemonkeygo/httpsig
//    var key *rsa.PrivateKey = ...
//    signer := httpsig.NewSigner("foo", key, httpsig.RSASHA256, nil)
//    client.Sign = signer.Sign
var Sign signFn = func(r *http.Request) error { return nil }

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

type client struct{}

func NewClient() client {
	return client{}
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
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	//req.Header.Set("Cache-Control", "no-cache")
	if err = Sign(req); err != nil {
		err := errorf(as.IRI(req.URL.String()), "Unable to sign request (method %q, previous error: %s)", req.Method, err)
		return req, err
	}
	return req, nil
}

// Head
func (c client) Head(url string) (resp *http.Response, err error) {
	req, err := c.req(http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

// Get wrapper over the functionality offered by the default http.Client object
func (c client) Get(url string) (resp *http.Response, err error) {
	req, err := c.req(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

// Post wrapper over the functionality offered by the default http.Client object
func (c *client) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := c.req(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return http.DefaultClient.Do(req)
}

// Put wrapper over the functionality offered by the default http.Client object
func (c client) Put(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := c.req(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return http.DefaultClient.Do(req)
}

// Delete wrapper over the functionality offered by the default http.Client object
func (c client) Delete(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := c.req(http.MethodDelete, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return http.DefaultClient.Do(req)
}
