package client

import (
	"bytes"
	"fmt"
	"net/http"
)

// ErrorHandlerFunc is a data type for the default ErrorHandler function of the package
type ErrorHandlerFunc func(...error) http.HandlerFunc

// ErrorHandler is the error handler callback for the ActivityPub package.
// It can be overloaded from the packages that require it
var ErrorHandler ErrorHandlerFunc = func(errors ...error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		output := bytes.Buffer{}
		for i, e := range errors {
			output.WriteString(fmt.Sprintf("#%d %s\n", i, e.Error()))
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-Type", "text/plain")
		w.Write(output.Bytes())
	}
}

// Routes
// {actor}/ -> HandleActorGET (requires LoadActorMiddleware(r, w), HandleError(r, w))
// {actor}/{inbox} -> HandleActorInboxGET
// {actor}/{outbox} -> HandleActorOutboxGET

// S2S
// {actor}/inbox -> InboxRequest (requires LoadActorMIddleware(r, w))

// C2S
// {actor}/outbox -> OutboxRequest
