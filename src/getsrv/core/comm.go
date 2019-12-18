package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// This will send JSON response with content header assigned
func Send(items interface{}, w http.ResponseWriter) error {
	data, err := json.Marshal(items)
	if err != nil {
		return errors.New("Cannot serialize response")
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(data))
	return nil
}
