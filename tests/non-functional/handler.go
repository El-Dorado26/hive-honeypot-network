
package nfrtests

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Event struct {
	ID        string `json:"id"`
	Source    string `json:"source"`
	Timestamp string `json:"timestamp"`
	EventType string `json:"event_type"`
}

var errInvalidEvent = errors.New("missing required event field")

func validateEvent(input []byte) error {
	var event Event

	if err := json.Unmarshal(input, &event); err != nil {
		return err
	}

	if event.ID == "" ||
		event.Source == "" ||
		event.Timestamp == "" ||
		event.EventType == "" {
		return errInvalidEvent
	}

	return nil
}

// EventHandler is a prototype ingestion endpoint.
// It validates events but does not store them in a database.
func EventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := validateEvent(body); err != nil {
		http.Error(w, "invalid event", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
