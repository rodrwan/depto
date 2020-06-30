package errors

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// NewAPI ...
func NewAPI(internal error, public string, status int) *API {
	return &API{
		Internal: internal,
		Status:   status,
		ErrorMessage: Message{
			Public: public,
		},
	}
}

// Message ...
type Message struct {
	Public string `json:"message,omitempty"`
}

// API ...
type API struct {
	Internal     error   `json:"-"`
	ErrorMessage Message `json:"error"`
	Status       int     `json:"-"`
}

func (h *API) Write(log *logrus.Logger, w http.ResponseWriter) {
	b, err := json.Marshal(h)
	if err != nil {
		NewInternalServerError(err).Write(log, w)
		return
	}

	log.Error(h.Internal.Error())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(h.Status)
	_, err = w.Write(b)
	if err != nil {
		NewInternalServerError(err).Write(log, w)
	}

	return
}

func (h API) Error() string {
	return fmt.Sprintf("Internal: %s | Public: %s", h.Internal, h.ErrorMessage.Public)
}

// NewInternalServerError ...
func NewInternalServerError(err error) *API {
	return &API{
		ErrorMessage: Message{Public: "Internal Server Error"},
		Internal:     err,
		Status:       http.StatusInternalServerError,
	}
}

// NewNotFound ...
func NewNotFound(err error) *API {
	return &API{
		ErrorMessage: Message{Public: "Not Found"},
		Internal:     err,
		Status:       http.StatusNotFound,
	}
}

// NewBadRequest ...
func NewBadRequest(err error, msg string) *API {
	return &API{
		ErrorMessage: Message{Public: "Bad request"},
		Internal:     err,
		Status:       http.StatusBadRequest,
	}
}

// NewTimeout ...
func NewTimeout(err error, msg string) *API {
	return &API{
		ErrorMessage: Message{Public: "Request take too long"},
		Internal:     err,
		Status:       http.StatusRequestTimeout,
	}
}
