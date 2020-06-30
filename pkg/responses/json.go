package responses

import (
	"encoding/json"
	"net/http"
)

// Response represents a response of the Application REST API.
type Response struct {
	Status int         `json:"-"`
	Data   interface{} `json:"data,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
}

// Write writes a ApplicationResposne to the given response writer encoded as JSON.
func (r *Response) Write(w http.ResponseWriter) error {
	b, err := json.Marshal(r)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	_, err = w.Write(b)

	return err
}
