
package nfrtests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
	"time"
)

func sendEvent(body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/events",
		bytes.NewBufferString(body),
	)

	rec := httptest.NewRecorder()
	EventHandler(rec, req)
	return rec
}

const validEvent = `{
	"id": "evt-001",
	"source": "ssh",
	"timestamp": "2026-10-11T10:00:00Z",
	"event_type": "authentication_failure"
}`

func TestMalformedJSONIsRejected(t *testing.T) {
	rec := sendEvent(`{"id":`)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("got %d, want 400", rec.Code)
	}
}

func TestMissingRequiredFieldIsRejected(t *testing.T) {
	rec := sendEvent(`{
		"id": "evt-001",
		"source": "ssh"
	}`)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("got %d, want 400", rec.Code)
	}
}

func TestValidEventIsAccepted(t *testing.T) {
	rec := sendEvent(validEvent)

	if rec.Code != http.StatusAccepted {
		t.Fatalf("got %d, want 202", rec.Code)
	}
}

func TestNonPostMethodIsRejected(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/events", nil)
	rec := httptest.NewRecorder()

	EventHandler(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("got %d, want 405", rec.Code)
	}
}

func TestHandlerRecoversAfterInvalidRequests(t *testing.T) {
	for i := 0; i < 100; i++ {
		rec := sendEvent(`{"id":`)

		if rec.Code != http.StatusBadRequest {
			t.Fatalf("invalid request %d: got %d, want 400",
				i, rec.Code)
		}
	}

	rec := sendEvent(validEvent)

	if rec.Code != http.StatusAccepted {
		t.Fatalf("valid request after invalid requests: got %d, want 202",
			rec.Code)
	}
}

func TestEventHandlerResponseTime(t *testing.T) {
	const requests = 100
	durations := make([]time.Duration, 0, requests)

	for i := 0; i < requests; i++ {
		req := httptest.NewRequest(
			http.MethodPost,
			"/api/events",
			bytes.NewBufferString(validEvent),
		)
		rec := httptest.NewRecorder()

		start := time.Now()
		EventHandler(rec, req)
		elapsed := time.Since(start)

		if rec.Code != http.StatusAccepted {
			t.Fatalf("request %d: got %d, want 202",
				i, rec.Code)
		}

		durations = append(durations, elapsed)
	}

	sort.Slice(durations, func(i, j int) bool {
		return durations[i] < durations[j]
	})

	p95 := durations[94]
	t.Log(fmt.Sprintf("Prototype handler p95: %v", p95))

	if p95 > 500*time.Millisecond {
		t.Fatalf("p95 %v exceeds the preliminary 500 ms target", p95)
	}
}
