package files_sdk

import (
	"encoding/json"
	"time"

	lib "github.com/Files-com/files-sdk-go/lib"
)

type ExternalEvent struct {
	EventType string    `json:"event_type,omitempty"`
	Status    string    `json:"status,omitempty"`
	Body      string    `json:"body,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type ExternalEventCollection []ExternalEvent

type ExternalEventListParams struct {
	Page       int             `url:"page,omitempty"`
	PerPage    int             `url:"per_page,omitempty"`
	Action     string          `url:"action,omitempty"`
	Cursor     string          `url:"cursor,omitempty"`
	SortBy     json.RawMessage `url:"sort_by,omitempty"`
	Filter     json.RawMessage `url:"filter,omitempty"`
	FilterGt   json.RawMessage `url:"filter_gt,omitempty"`
	FilterGteq json.RawMessage `url:"filter_gteq,omitempty"`
	FilterLike json.RawMessage `url:"filter_like,omitempty"`
	FilterLt   json.RawMessage `url:"filter_lt,omitempty"`
	FilterLteq json.RawMessage `url:"filter_lteq,omitempty"`
	lib.ListParams
}

func (e *ExternalEvent) UnmarshalJSON(data []byte) error {
	type externalEvent ExternalEvent
	var v externalEvent
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*e = ExternalEvent(v)
	return nil
}

func (e *ExternalEventCollection) UnmarshalJSON(data []byte) error {
	type externalEvents []ExternalEvent
	var v externalEvents
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*e = ExternalEventCollection(v)
	return nil
}
