package events

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"

	"hive-honeypot-network/internal/database"
	"strconv"
)

// Store describes the database operation the handler needs
type Store interface {
	CreateEvent(
		ctx context.Context,
		source string,
		eventType string,
		occurredAt time.Time,
		eventData []byte,
	) (int64, time.Time, error)

	ListEvents(
		ctx context.Context,
		limit int,
		offset int,
	) ([]database.Event, error)
}

// EventInput represents the JSON sent by a client.
type EventInput struct {
	Source     string          `json:"source"`
	EventType  string          `json:"event_type"`
	OccurredAt time.Time       `json:"occurred_at"`
	EventData  json.RawMessage `json:"event_data"`
}

// EventHandler handles incoming event requests.
type EventHandler struct {
	Store Store
}

// writeJSON sends a JSON response.
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// Create handles POST /api/events.
func (h EventHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		writeJSON(w, http.StatusMethodNotAllowed,
			map[string]string{"error": "Method not allowed"})
		return
	}

	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || mediaType != "application/json" {
		writeJSON(w, http.StatusUnsupportedMediaType,
			map[string]string{"error": "Content-Type must be application/json"})
		return
	}

	// Limit request bodies to 1 MB.
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var input EventInput
	if err := decoder.Decode(&input); err != nil {
		var maxBytesErr *http.MaxBytesError
		if errors.As(err, &maxBytesErr) {
			writeJSON(w, http.StatusRequestEntityTooLarge,
				map[string]string{"error": "Request body is too large"})
			return
		}

		writeJSON(w, http.StatusBadRequest,
			map[string]string{"error": "Invalid JSON request"})
		return
	}

	// Reject a second JSON value after the first one.
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		writeJSON(w, http.StatusBadRequest,
			map[string]string{"error": "Request must contain one JSON object"})
		return
	}

	input.Source = strings.TrimSpace(input.Source)
	input.EventType = strings.TrimSpace(input.EventType)

	if input.Source == "" || utf8.RuneCountInString(input.Source) > 100 {
		writeJSON(w, http.StatusBadRequest,
			map[string]string{"error": "source must contain 1 to 100 characters"})
		return
	}

	if input.EventType == "" || utf8.RuneCountInString(input.EventType) > 100 {
		writeJSON(w, http.StatusBadRequest,
			map[string]string{"error": "event_type must contain 1 to 100 characters"})
		return
	}

	if input.OccurredAt.IsZero() {
		writeJSON(w, http.StatusBadRequest,
			map[string]string{"error": "occurred_at is required"})
		return
	}

	eventData := strings.TrimSpace(string(input.EventData))
	if eventData == "" || !json.Valid([]byte(eventData)) ||
		!strings.HasPrefix(eventData, "{") {
		writeJSON(w, http.StatusBadRequest,
			map[string]string{"error": "event_data must be a JSON object"})
		return
	}

	if h.Store == nil {
		writeJSON(w, http.StatusInternalServerError,
			map[string]string{"error": "Event storage is unavailable"})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	id, receivedAt, err := h.Store.CreateEvent(
		ctx,
		input.Source,
		input.EventType,
		input.OccurredAt,
		[]byte(eventData),
	)
	if err != nil {
		// Do not expose internal database details to the client.
		writeJSON(w, http.StatusInternalServerError,
			map[string]string{"error": "Could not store event"})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]any{
		"id":          id,
		"received_at": receivedAt,
	})
}

// List handles GET /api/events with pagination
func (h EventHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		writeJSON(w, http.StatusMethodNotAllowed,
			map[string]string{"error": "Method not allowed"})
		return
	}

	page := 1
	limit := 20

	query := r.URL.Query()

	if values, exists := query["page"]; exists {
		if len(values) != 1 {
			writeJSON(w, http.StatusBadRequest,
				map[string]string{"error": "page must be a positive integer"})
			return
		}

		value, err := strconv.Atoi(values[0])
		if err != nil || value < 1 || value > 1000000 {
			writeJSON(w, http.StatusBadRequest,
				map[string]string{"error": "page must be between 1 and 1000000"})
			return
		}
		page = value
	}

	if values, exists := query["limit"]; exists {
		if len(values) != 1 {
			writeJSON(w, http.StatusBadRequest,
				map[string]string{"error": "limit must be between 1 and 100"})
			return
		}

		value, err := strconv.Atoi(values[0])
		if err != nil || value < 1 || value > 100 {
			writeJSON(w, http.StatusBadRequest,
				map[string]string{"error": "limit must be between 1 and 100"})
			return
		}
		limit = value
	}

	offset := (page - 1) * limit

	if h.Store == nil {
		writeJSON(w, http.StatusInternalServerError,
			map[string]string{"error": "Event storage is unavailable"})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Fetch one extra event to determine whether another page exists.
	events, err := h.Store.ListEvents(ctx, limit+1, offset)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError,
			map[string]string{"error": "Could not retrieve events"})
		return
	}

	hasMore := len(events) > limit
	if hasMore {
		events = events[:limit]
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"events":   events,
		"page":     page,
		"limit":    limit,
		"has_more": hasMore,
	})
}
