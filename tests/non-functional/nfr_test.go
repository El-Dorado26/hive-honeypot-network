
package nfrtests

import (
	"encoding/json"
	"testing"
	"time"
)

type Event struct {
	ID        string `json:"id"`
	Source    string `json:"source"`
	Timestamp string `json:"timestamp"`
	EventType string `json:"event_type"`
}

// Reference validator for the planned HIVE event contract.
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

type validationError string

func (e validationError) Error() string {
	return string(e)
}

const errInvalidEvent validationError = "missing required event field"

func TestMalformedJSONIsRejected(t *testing.T) {
	err := validateEvent([]byte(`{"id":`))

	if err == nil {
		t.Fatal("expected malformed JSON to be rejected")
	}
}

func TestMissingRequiredFieldIsRejected(t *testing.T) {
	input := []byte(`{
		"id": "evt-001",
		"source": "ssh"
	}`)

	if err := validateEvent(input); err == nil {
		t.Fatal("expected incomplete event to be rejected")
	}
}

func TestValidEventIsAccepted(t *testing.T) {
	input := []byte(`{
		"id": "evt-001",
		"source": "ssh",
		"timestamp": "2026-10-11T10:00:00Z",
		"event_type": "authentication_failure"
	}`)

	if err := validateEvent(input); err != nil {
		t.Fatalf("expected valid event to be accepted: %v", err)
	}
}

func TestValidatorRecoversAfterInvalidInput(t *testing.T) {
	for i := 0; i < 100; i++ {
		_ = validateEvent([]byte(`{"id":`))
	}

	valid := []byte(`{
		"id": "evt-002",
		"source": "http",
		"timestamp": "2026-10-11T10:00:00Z",
		"event_type": "request"
	}`)

	if err := validateEvent(valid); err != nil {
		t.Fatalf("valid event failed after invalid inputs: %v", err)
	}
}

func TestValidationPerformance(t *testing.T) {
	input := []byte(`{
		"id": "evt-003",
		"source": "ssh",
		"timestamp": "2026-10-11T10:00:00Z",
		"event_type": "authentication_failure"
	}`)

	const iterations = 100
	var durations []time.Duration

	for i := 0; i < iterations; i++ {
		start := time.Now()
		err := validateEvent(input)
		elapsed := time.Since(start)

		if err != nil {
			t.Fatalf("valid event rejected: %v", err)
		}

		durations = append(durations, elapsed)
	}

	// Sort the 100 durations to obtain the nearest-rank
	// 95th percentile: the 95th observation.
	for i := 0; i < len(durations); i++ {
		for j := i + 1; j < len(durations); j++ {
			if durations[j] < durations[i] {
				durations[i], durations[j] =
					durations[j], durations[i]
			}
		}
	}

	p95 := durations[94]
	t.Logf("Validator-only p95: %v", p95)

	if p95 > 500*time.Millisecond {
		t.Fatalf("validator p95 %v exceeds 500 ms", p95)
	}
}
