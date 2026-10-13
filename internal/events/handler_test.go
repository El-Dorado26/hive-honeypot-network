package events

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// fakeStore imitates database storage without PostgreSQL.
type fakeStore struct {
	called     bool
	fail       bool
	source     string
	eventType  string
	occurredAt time.Time
	eventData  []byte
}

func (f *fakeStore) CreateEvent(
	ctx context.Context,
	source string,
	eventType string,
	occurredAt time.Time,
	eventData []byte,
) (int64, time.Time, error) {
	f.called = true
	f.source = source
	f.eventType = eventType
	f.occurredAt = occurredAt
	f.eventData = eventData

	if f.fail {
		return 0, time.Time{}, errors.New("simulated database error")
	}

	return 42, time.Date(2026, 10, 13, 10, 0, 0, 0, time.UTC), nil
}

func sendRequest(
	store *fakeStore,
	method string,
	contentType string,
	body string,
) *httptest.ResponseRecorder {
	handler := EventHandler{Store: store}

	req := httptest.NewRequest(method, "/api/events", strings.NewReader(body))
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	recorder := httptest.NewRecorder()
	handler.Create(recorder, req)
	return recorder
}

const validEvent = `{
	"source": "ssh-honeypot",
	"event_type": "login_attempt",
	"occurred_at": "2026-10-13T09:00:00Z",
	"event_data": {
		"username": "admin",
		"result": "failed"
	}
}`

func TestCreateValidEvent(t *testing.T) {
	store := &fakeStore{}
	recorder := sendRequest(store, http.MethodPost, "application/json", validEvent)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d: %s", recorder.Code, recorder.Body.String())
	}

	if !store.called {
		t.Fatal("expected store to be called")
	}

	if store.source != "ssh-honeypot" {
		t.Errorf("unexpected source: %s", store.source)
	}

	var response map[string]any
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatalf("response is not valid JSON: %v", err)
	}

	if response["id"] != float64(42) {
		t.Errorf("expected event ID 42, got %v", response["id"])
	}
}

func TestCreateRejectsInvalidRequests(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		contentType string
		body        string
		wantStatus  int
	}{
		{
			name:        "malformed JSON",
			method:      http.MethodPost,
			contentType: "application/json",
			body:        `{"source":`,
			wantStatus:  http.StatusBadRequest,
		},
		{
			name:        "missing source",
			method:      http.MethodPost,
			contentType: "application/json",
			body:        `{"event_type":"login_attempt","occurred_at":"2026-10-13T09:00:00Z","event_data":{}}`,
			wantStatus:  http.StatusBadRequest,
		},
		{
			name:        "unknown field",
			method:      http.MethodPost,
			contentType: "application/json",
			body:        `{"source":"ssh","event_type":"login","occurred_at":"2026-10-13T09:00:00Z","event_data":{},"extra":"unexpected"}`,
			wantStatus:  http.StatusBadRequest,
		},
		{
			name:        "event data is not an object",
			method:      http.MethodPost,
			contentType: "application/json",
			body:        `{"source":"ssh","event_type":"login","occurred_at":"2026-10-13T09:00:00Z","event_data":[]}`,
			wantStatus:  http.StatusBadRequest,
		},
		{
			name:        "wrong method",
			method:      http.MethodGet,
			contentType: "application/json",
			body:        validEvent,
			wantStatus:  http.StatusMethodNotAllowed,
		},
		{
			name:        "wrong content type",
			method:      http.MethodPost,
			contentType: "text/plain",
			body:        validEvent,
			wantStatus:  http.StatusUnsupportedMediaType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &fakeStore{}
			recorder := sendRequest(store, tt.method, tt.contentType, tt.body)

			if recorder.Code != tt.wantStatus {
				t.Errorf("expected status %d, got %d: %s",
					tt.wantStatus, recorder.Code, recorder.Body.String())
			}

			if store.called {
				t.Error("store should not be called for an invalid request")
			}
		})
	}
}

func TestCreateRejectsOversizedRequest(t *testing.T) {
	store := &fakeStore{}
	body := `{"source":"ssh","event_type":"login","occurred_at":"2026-10-13T09:00:00Z","event_data":{"padding":"` +
		strings.Repeat("a", 1<<20) + `"}}`

	recorder := sendRequest(store, http.MethodPost, "application/json", body)

	if recorder.Code != http.StatusRequestEntityTooLarge {
		t.Errorf("expected status 413, got %d", recorder.Code)
	}

	if store.called {
		t.Error("store should not be called for an oversized request")
	}
}

func TestCreateHandlesDatabaseFailure(t *testing.T) {
	store := &fakeStore{fail: true}
	recorder := sendRequest(store, http.MethodPost, "application/json", validEvent)

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500, got %d", recorder.Code)
	}

	if !store.called {
		t.Error("expected store to be called")
	}

	if strings.Contains(recorder.Body.String(), "simulated database error") {
		t.Error("internal database error must not be exposed to clients")
	}
}
