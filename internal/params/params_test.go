package params_test

import (
	"testing"

	"github.com/chatbotkit/go-sdk/internal/params"
)

// MockOrder is a test type that mimics the generated order types.
type MockOrder string

const (
	MockOrderAsc  MockOrder = "asc"
	MockOrderDesc MockOrder = "desc"
)

func TestBuildListQuery(t *testing.T) {
	tests := []struct {
		name   string
		cursor *string
		order  *MockOrder
		take   *int64
		meta   map[string]string
		want   map[string]string
	}{
		{
			name: "all nil values",
			want: map[string]string{},
		},
		{
			name:   "cursor only",
			cursor: ptr("abc123"),
			want:   map[string]string{"cursor": "abc123"},
		},
		{
			name:  "order only",
			order: ptrOrder(MockOrderAsc),
			want:  map[string]string{"order": "asc"},
		},
		{
			name: "take only",
			take: ptr(int64(50)),
			want: map[string]string{"take": "50"},
		},
		{
			name: "meta only",
			meta: map[string]string{"key1": "value1", "key2": "value2"},
			want: map[string]string{"meta[key1]": "value1", "meta[key2]": "value2"},
		},
		{
			name:   "all values",
			cursor: ptr("cursor123"),
			order:  ptrOrder(MockOrderDesc),
			take:   ptr(int64(100)),
			meta:   map[string]string{"foo": "bar"},
			want:   map[string]string{"cursor": "cursor123", "order": "desc", "take": "100", "meta[foo]": "bar"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := params.BuildListQuery(tt.cursor, tt.order, tt.take, tt.meta)

			// Check that all expected values are present
			for k, v := range tt.want {
				if got.Get(k) != v {
					t.Errorf("BuildListQuery() key %q = %q, want %q", k, got.Get(k), v)
				}
			}

			// Check that no unexpected values are present
			for k := range got {
				if _, ok := tt.want[k]; !ok {
					t.Errorf("BuildListQuery() unexpected key %q = %q", k, got.Get(k))
				}
			}
		})
	}
}

func TestSetCursor(t *testing.T) {
	tests := []struct {
		name   string
		cursor *string
		want   string
	}{
		{
			name:   "nil cursor",
			cursor: nil,
			want:   "",
		},
		{
			name:   "valid cursor",
			cursor: ptr("abc123"),
			want:   "abc123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := make(map[string][]string)
			params.SetCursor(q, tt.cursor)
			if got := q["cursor"]; len(got) > 0 && got[0] != tt.want {
				t.Errorf("SetCursor() = %q, want %q", got[0], tt.want)
			}
		})
	}
}

func TestSetOrder(t *testing.T) {
	tests := []struct {
		name  string
		order *MockOrder
		want  string
	}{
		{
			name:  "nil order",
			order: nil,
			want:  "",
		},
		{
			name:  "asc order",
			order: ptrOrder(MockOrderAsc),
			want:  "asc",
		},
		{
			name:  "desc order",
			order: ptrOrder(MockOrderDesc),
			want:  "desc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := make(map[string][]string)
			params.SetOrder(q, tt.order)
			if got := q["order"]; len(got) > 0 && got[0] != tt.want {
				t.Errorf("SetOrder() = %q, want %q", got[0], tt.want)
			}
		})
	}
}

func TestSetTake(t *testing.T) {
	tests := []struct {
		name string
		take *int64
		want string
	}{
		{
			name: "nil take",
			take: nil,
			want: "",
		},
		{
			name: "valid take",
			take: ptr(int64(50)),
			want: "50",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := make(map[string][]string)
			params.SetTake(q, tt.take)
			if got := q["take"]; len(got) > 0 && got[0] != tt.want {
				t.Errorf("SetTake() = %q, want %q", got[0], tt.want)
			}
		})
	}
}

func TestSetMeta(t *testing.T) {
	tests := []struct {
		name string
		meta map[string]string
		want map[string]string
	}{
		{
			name: "nil meta",
			meta: nil,
			want: map[string]string{},
		},
		{
			name: "empty meta",
			meta: map[string]string{},
			want: map[string]string{},
		},
		{
			name: "single meta",
			meta: map[string]string{"key": "value"},
			want: map[string]string{"meta[key]": "value"},
		},
		{
			name: "multiple meta",
			meta: map[string]string{"a": "1", "b": "2"},
			want: map[string]string{"meta[a]": "1", "meta[b]": "2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := make(map[string][]string)
			params.SetMeta(q, tt.meta)
			for k, v := range tt.want {
				if got := q[k]; len(got) > 0 && got[0] != v {
					t.Errorf("SetMeta() key %q = %q, want %q", k, got[0], v)
				}
			}
		})
	}
}

// Helper functions for creating pointers
func ptr[T any](v T) *T {
	return &v
}

func ptrOrder(v MockOrder) *MockOrder {
	return &v
}
